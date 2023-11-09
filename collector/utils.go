/*
 Copyright 2023 The Pipeline Service Authors.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package collector

import (
	"context"
	"fmt"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	"k8s.io/apimachinery/pkg/types"
	"knative.dev/pkg/apis"
	"os"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	FILTER_THRESHOLD                    = "FILTER_THRESHOLD"
	DEFAULT_THRESHOLD                   = float64(300)
	ENABLE_GAP_METRIC_ADDITIONAL_LABELS = "ENABLE_GAP_METRIC_ADDITIONAL_LABELS"
	NS_LABEL                            = "namespace"
	PIPELINE_NAME_LABEL                 = "pipelinename"
	TASK_NAME_LABEL                     = "taskname"
	COMPLETED_LABEL                     = "completed"
	UPCOMING_LABEL                      = "upcoming"
	STATUS_LABEL                        = "status"
	SUCCEEDED                           = "succeded"
	FAILED                              = "failed"
)

func pipelineRunPipelineRef(pr *v1.PipelineRun) string {
	val := ""
	ref := pr.Spec.PipelineRef
	if ref != nil {
		val = ref.Name
		if len(val) == 0 {
			for _, p := range ref.Params {
				if strings.TrimSpace(p.Name) == "name" {
					return p.Value.StringVal
				}
			}
		}
	} else {
		if len(pr.GenerateName) > 0 {
			return pr.GenerateName
		}
		// at this point, the pipelinerun name should not have any random aspects, and is
		// essentially a constant, with a minimal cardinality impact
		val = pr.Name
	}
	return val
}

func taskRef(labels map[string]string) string {
	task, _ := labels[pipeline.TaskLabelKey]
	pipelineTask, _ := labels[pipeline.PipelineTaskLabelKey]
	clusterTask, _ := labels[pipeline.ClusterTaskLabelKey]
	taskRun, _ := labels[pipeline.TaskRunLabelKey]
	switch {
	case len(task) > 0:
		return task
	case len(pipelineTask) > 0:
		return pipelineTask
	case len(clusterTask) > 0:
		return clusterTask
	case len(taskRun) > 0:
		return taskRun
	}
	return ""
}

func optionalMetricEnabled(envVarName string) bool {
	env := os.Getenv(envVarName)
	enabled := len(env) > 0
	// any random setting means true
	if enabled {
		// allow for users to turn off by setting to false
		bv, err := strconv.ParseBool(env)
		if err == nil && !bv {
			return false
		}
		return true
	}
	return false
}

func calculateScheduledDuration(created, started time.Time) float64 {
	if created.IsZero() || started.IsZero() {
		return 0
	}
	return float64(started.Sub(created).Milliseconds())
}

func prNotDoneOrHasNoKids(pr *v1.PipelineRun) bool {
	if len(pr.Status.ChildReferences) < 1 {
		return true
	}
	// in case there are gaps between a pipelinerun being marked done but the complete timestamp is not set, with the
	// understanding that the complete timestamp is not processed before any completed taskrun complete timestamps have been processed
	if pr.Status.CompletionTime == nil {
		return true
	}
	return false
}

func sortTaskRunsForGapCalculations(pr *v1.PipelineRun, oc client.Client, ctx context.Context) ([]*v1.TaskRun, []*v1.TaskRun) {
	sortedTaskRunsByCreateTimes := []*v1.TaskRun{}
	reverseOrderSortedTaskRunsByCompletionTimes := []*v1.TaskRun{}
	// prior testing in staging proved that with enough concurrency, this array is minimally not sorted based on when
	// the task runs were created, so we explicitly sort for that; also, this sorting will allow us to effectively
	// address parallel taskruns vs. taskrun dependencies and ordering (where tekton does not create a taskrun until its dependencies
	// have completed).
	for _, kidRef := range pr.Status.ChildReferences {
		if kidRef.Kind != "TaskRun" {
			continue
		}
		kid := &v1.TaskRun{}
		err := oc.Get(ctx, types.NamespacedName{Namespace: pr.Namespace, Name: kidRef.Name}, kid)
		if err != nil {
			ctrl.Log.Info(fmt.Sprintf("could not calculate gap for taskrun %s:%s: %s", pr.Namespace, kidRef.Name, err.Error()))
			continue
		}

		sortedTaskRunsByCreateTimes = append(sortedTaskRunsByCreateTimes, kid)
		// don't add taskruns that did not complete i.e. presumably timed out of failed; any taskruns that dependended
		// on should not have even been created
		if kid.Status.CompletionTime != nil {
			reverseOrderSortedTaskRunsByCompletionTimes = append(reverseOrderSortedTaskRunsByCompletionTimes, kid)

		}
	}
	sort.SliceStable(sortedTaskRunsByCreateTimes, func(i, j int) bool {
		return sortedTaskRunsByCreateTimes[i].CreationTimestamp.Time.Before(sortedTaskRunsByCreateTimes[j].CreationTimestamp.Time)
	})
	sort.SliceStable(reverseOrderSortedTaskRunsByCompletionTimes, func(i, j int) bool {
		return reverseOrderSortedTaskRunsByCompletionTimes[i].Status.CompletionTime.Time.After(reverseOrderSortedTaskRunsByCompletionTimes[j].Status.CompletionTime.Time)
	})
	return sortedTaskRunsByCreateTimes, reverseOrderSortedTaskRunsByCompletionTimes
}

type GapEntry struct {
	status    string
	pipeline  string
	completed string
	upcoming  string
	gap       float64
}

func calculateGaps(ctx context.Context, pr *v1.PipelineRun, oc client.Client, sortedTaskRunsByCreateTimes []*v1.TaskRun, reverseOrderSortedTaskRunsByCompletionTimes []*v1.TaskRun) []GapEntry {
	gapEntries := []GapEntry{}
	prRef := pipelineRunPipelineRef(pr)
	for index, tr := range sortedTaskRunsByCreateTimes {
		succeedCondition := pr.Status.GetCondition(apis.ConditionSucceeded)
		if succeedCondition == nil {
			ctrl.Log.Info(fmt.Sprintf("WARNING: pipielinerun %s:%s marked done but has nil succeed condition", pr.Namespace, pr.Name))
			continue
		}
		if succeedCondition.IsUnknown() {
			ctrl.Log.Info(fmt.Sprintf("WARNING: pipielinerun %s:%s marked done but has unknown succeed condition", pr.Namespace, pr.Name))
			continue
		}
		gapEntry := GapEntry{}
		status := SUCCEEDED
		if succeedCondition.IsFalse() {
			status = FAILED
		}
		gapEntry.status = status
		gapEntry.pipeline = prRef

		if index == 0 {
			ctrl.Log.V(4).Info(fmt.Sprintf("first task %s for pipeline %s", taskRef(tr.Labels), prRef))
			// our first task is simple, just work off of the pipelinerun
			gapEntry.gap = float64(tr.CreationTimestamp.Time.Sub(pr.CreationTimestamp.Time).Milliseconds())
			gapEntry.completed = prRef
			gapEntry.upcoming = taskRef(tr.Labels)
			gapEntries = append(gapEntries, gapEntry)
			continue
		}

		firstKid := sortedTaskRunsByCreateTimes[0]

		// so using the first taskrun completion time addresses sequential / chaining dependencies;
		// for parallel, if the first taskrun's completion time is not after this taskrun's create time,
		// that means parallel taskruns, and we work off of the pipelinerun; NOTE: this focuses on "top level" parallel task runs
		// with absolutely no dependencies.  Once any sort of dependency is established, there are no more top level parallel taskruns.
		if firstKid.Status.CompletionTime != nil && firstKid.Status.CompletionTime.Time.After(tr.CreationTimestamp.Time) {
			ctrl.Log.V(4).Info(fmt.Sprintf("task %s considered parallel for pipeline %s", taskRef(tr.Labels), prRef))
			gapEntry.gap = float64(tr.CreationTimestamp.Time.Sub(pr.CreationTimestamp.Time).Milliseconds())
			gapEntry.completed = prRef
			gapEntry.upcoming = taskRef(tr.Labels)
			gapEntries = append(gapEntries, gapEntry)
			continue
		}

		// Conversely, task run chains can run in parallel, and a taskrun can depend on multiple chains or threads of taskruns. We want to find the chain
		// that finished last, but before we are created.  We traverse through our reverse sorted on completion time list to determine that.  But yes, we don't reproduce the DAG
		// graph (there is no clean dependency import path in tekton for that) to confirm the edges.  This approximation is sufficient.

		// get whatever completed first
		timeToCalculateWith := time.Time{}
		trToCalculateWith := &v1.TaskRun{}
		if len(reverseOrderSortedTaskRunsByCompletionTimes) > 0 {
			trToCalculateWith = reverseOrderSortedTaskRunsByCompletionTimes[len(reverseOrderSortedTaskRunsByCompletionTimes)-1]
			timeToCalculateWith = trToCalculateWith.Status.CompletionTime.Time
		} else {
			// if no taskruns completed, that means any taskruns created were created as part of the initial pipelinerun creation,
			// so use the pipelinerun creation time
			timeToCalculateWith = pr.CreationTimestamp.Time
		}
		for _, tr2 := range reverseOrderSortedTaskRunsByCompletionTimes {
			if tr2.Name == tr.Name {
				continue
			}
			ctrl.Log.V(4).Info(fmt.Sprintf("comparing candidate %s to current task %s", taskRef(tr2.Labels), taskRef(tr.Labels)))
			if !tr2.Status.CompletionTime.Time.After(tr.CreationTimestamp.Time) {
				ctrl.Log.V(4).Info(fmt.Sprintf("%s did not complete after so use it to compute gap for current task %s", taskRef(tr2.Labels), taskRef(tr.Labels)))
				trToCalculateWith = tr2
				timeToCalculateWith = tr2.Status.CompletionTime.Time
				break
			}
			ctrl.Log.V(4).Info(fmt.Sprintf("skipping %s as a gap candidate for current task %s is OK", taskRef(tr2.Labels), taskRef(tr.Labels)))
		}
		gapEntry.gap = float64(tr.CreationTimestamp.Time.Sub(timeToCalculateWith).Milliseconds())
		gapEntry.completed = taskRef(trToCalculateWith.Labels)
		gapEntry.upcoming = taskRef(tr.Labels)
		gapEntries = append(gapEntries, gapEntry)
	}
	return gapEntries
}

func filter(numerator, denominator float64) bool {
	threshold := DEFAULT_THRESHOLD
	thresholdStr := os.Getenv(FILTER_THRESHOLD)
	if len(thresholdStr) > 0 {
		thresholdOverride, err := strconv.ParseFloat(thresholdStr, 64)
		if err != nil {
			ctrl.Log.V(6).Info(fmt.Sprintf("error parsing %s env of %s: %s", FILTER_THRESHOLD, thresholdStr, err.Error()))
		} else {
			threshold = thresholdOverride
		}
	}
	// if overhead is non-zero, but total duration is less that 40 seconds,
	// this is a simpler, most likely user defined pipeline which does not fall
	// under our image building based overhead concerns
	if numerator > 0 && denominator < threshold {
		return true
	}
	//TODO we don't have a sense for it yet, but at some point we may get an idea
	// of what is unacceptable overhead regardless of total duration, where we don't
	// try to mitigate the tekton controller and user pipelineruns sharing the same
	// cluster resources
	return false
}
