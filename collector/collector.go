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
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"k8s.io/apimachinery/pkg/types"
	"os"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	ENABLE_PIPELINERUN_SCHEDULED_DURATION_PIPELINENAME_LABEL = "ENABLE_PIPELINERUN_SCHEDULED_DURATION_PIPELINENAME_LABEL"
	ENABLE_TASKRUN_SCHEDULED_DURATION_TASKNAME_LABEL         = "ENABLE_TASKRUN_SCHEDULED_DURATION_TASKNAME_LABEL"
	ENABLE_GAP_METRIC_ADDITIONAL_LABELS                      = "ENABLE_GAP_METRIC_ADDITIONAL_LABELS"
	NS_LABEL                                                 = "namespace"
	PIPELINE_NAME_LABEL                                      = "pipelinename"
	TASK_NAME_LABEL                                          = "taskname"
	COMPLETED_LABEL                                          = "completed"
	UPCOMING_LABEL                                           = "upcoming"
)

type PipelineRunScheduledCollector struct {
	durationScheduled *prometheus.HistogramVec
	prSchedNameLabel  bool
}

type PipelineRunTaskRunGapCollector struct {
	trGaps           *prometheus.HistogramVec
	additionalLabels bool
}

type ThrottledByPVCQuotaCollector struct {
	pvcThrottle *prometheus.GaugeVec
}

type TaskRunCollector struct {
	durationScheduled *prometheus.HistogramVec
	trSchedNameLabel  bool
}

func optionalLabelEnabled(labelName string) bool {
	env := os.Getenv(labelName)
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

func pipelineRunPipelineRef(pr *v1beta1.PipelineRun) string {
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

func taskRunTaskRef(tr *v1beta1.TaskRun, pr *v1beta1.PipelineRun) string {
	val := ""
	ref := tr.Spec.TaskRef
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
		if len(tr.GenerateName) > 0 {
			return tr.GenerateName
		}
		if pr != nil {
			if pr.Spec.PipelineSpec != nil {
				ps := pr.Spec.PipelineSpec
				if ps != nil {
					for _, t := range ps.Tasks {
						name := t.Name
						suffix := fmt.Sprintf("-%s", name)
						if len(suffix) < 2 {
							for _, p := range t.Params {
								if p.Name == "name" {
									name = p.Value.StringVal
									suffix = fmt.Sprintf("-%s", name)
									break
								}
							}
						}
						if strings.HasSuffix(tr.Name, suffix) {
							return name
						}

					}
				}
			}
			for _, t := range pr.Spec.TaskRunSpecs {
				suffix := fmt.Sprintf("-%s", t.PipelineTaskName)
				if strings.HasSuffix(tr.Name, suffix) {
					return t.PipelineTaskName
				}
			}
		}
		// at this point, the taskrun name should not have any random aspects, and is
		// essentially a constant, with a minimal cardinality impact
		val = tr.Name
	}
	return val
}

func NewPipelineRunScheduledCollector() *PipelineRunScheduledCollector {
	labelNames := []string{NS_LABEL}
	prSchedStatNameEnabled := optionalLabelEnabled(ENABLE_PIPELINERUN_SCHEDULED_DURATION_PIPELINENAME_LABEL)
	if prSchedStatNameEnabled {
		labelNames = append(labelNames, PIPELINE_NAME_LABEL)
	}
	durationScheduled := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "pipelinerun_duration_scheduled_seconds",
		Help: "Duration in seconds for a PipelineRun to be scheduled.",
		// reminder: exponential buckets need a start value greater than 0
		// the results in buckets of 0.1, 0.5, 2.5, 12.5, 62.5, 312.5 seconds
		Buckets: prometheus.ExponentialBuckets(0.1, 5, 6),
	}, labelNames)

	pipelineRunScheduledCollector := &PipelineRunScheduledCollector{
		durationScheduled: durationScheduled,
		prSchedNameLabel:  prSchedStatNameEnabled,
	}
	metrics.Registry.MustRegister(durationScheduled)

	return pipelineRunScheduledCollector
}

func (c *PipelineRunScheduledCollector) bumpScheduledDuration(pr *v1beta1.PipelineRun, scheduleDuration float64) {
	labels := map[string]string{NS_LABEL: pr.Namespace}
	if c.prSchedNameLabel {
		labels[PIPELINE_NAME_LABEL] = pipelineRunPipelineRef(pr)
	}
	c.durationScheduled.With(labels).Observe(scheduleDuration)
}

func NewPVCThrottledCollector() *ThrottledByPVCQuotaCollector {
	labelNames := []string{NS_LABEL}
	pvcThrottled := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pipelinerun_failed_by_pvc_quota_count",
		Help: "Number of PipelineRuns who were marked failed because PVC Resource Quotas prevented the creation of required PVCs",
	}, labelNames)
	pvcThrottledCollector := &ThrottledByPVCQuotaCollector{
		pvcThrottle: pvcThrottled,
	}
	metrics.Registry.MustRegister(pvcThrottled)
	return pvcThrottledCollector
}

func (c *ThrottledByPVCQuotaCollector) incPVCThrottle(pr *v1beta1.PipelineRun) {
	labels := map[string]string{NS_LABEL: pr.Namespace}
	c.pvcThrottle.With(labels).Inc()
}

func (c *ThrottledByPVCQuotaCollector) zeroPVCThrottle(ns string) {
	labels := map[string]string{NS_LABEL: ns}
	c.pvcThrottle.With(labels).Set(float64(0))
}

func NewPipelineRunTaskRunGapCollector() *PipelineRunTaskRunGapCollector {
	labelNames := []string{NS_LABEL}
	additionalLabels := optionalMetricEnabled(ENABLE_GAP_METRIC_ADDITIONAL_LABELS)
	if additionalLabels {
		labelNames = append(labelNames, PIPELINE_NAME_LABEL, COMPLETED_LABEL, UPCOMING_LABEL)
	}
	trGaps := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "pipelinerun_gap_between_taskruns_milliseconds",
		Help: "Duration in milliseconds between a taskrun completing and the next taskrun being created within a pipelinerun.  If a pipelinerun only has one taskrun, use pipelinerun_duration_scheduled_seconds.",
		// reminder: exponential buckets need a start value greater than 0
		// the results in buckets of 0.1, 0.5, 2.5, 12.5, 62.5, 312.5 milliseconds
		Buckets: prometheus.ExponentialBuckets(0.1, 5, 6),
	}, labelNames)

	pipelineRunTaskRunGapCollector := &PipelineRunTaskRunGapCollector{
		trGaps:           trGaps,
		additionalLabels: additionalLabels,
	}
	metrics.Registry.MustRegister(trGaps)

	return pipelineRunTaskRunGapCollector
}

func (c *PipelineRunTaskRunGapCollector) bumpGapDuration(pr *v1beta1.PipelineRun, oc client.Client, ctx context.Context) {
	if len(pr.Status.ChildReferences) < 1 {
		return
	}
	// in case there are gaps between a pipelinerun being marked done but the complete timestamp is not set, with the
	// understanding that the complete timestamp is not processed before any completed taskrun complete timestamps have been processed
	if pr.Status.CompletionTime == nil {
		return
	}

	sortedTaskRunsByCreateTimes := []*v1beta1.TaskRun{}
	reverseOrderSortedTaskRunsByCompletionTimes := []*v1beta1.TaskRun{}
	// prior testing in staging proved that with enough concurrency, this array is minimally not sorted based on when
	// the task runs were created, so we explicitly sort for that; also, this sorting will allow us to effectively
	// address parallel taskruns vs. taskrun dependencies and ordering (where tekton does not create a taskrun until its dependencies
	// have completed).
	for _, kidRef := range pr.Status.ChildReferences {
		if kidRef.Kind != "TaskRun" {
			continue
		}
		kid := &v1beta1.TaskRun{}
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
	prRef := pipelineRunPipelineRef(pr)
	for index, tr := range sortedTaskRunsByCreateTimes {
		labels := map[string]string{NS_LABEL: pr.Namespace}
		if c.additionalLabels {
			labels[PIPELINE_NAME_LABEL] = prRef
		}

		if index == 0 {
			ctrl.Log.V(4).Info(fmt.Sprintf("first task %s for pipeline %s", taskRunTaskRef(tr, pr), prRef))
			// our first task is simple, just work off of the pipelinerun
			gap := tr.CreationTimestamp.Time.Sub(pr.CreationTimestamp.Time).Milliseconds()
			if c.additionalLabels {
				labels[COMPLETED_LABEL] = prRef
				labels[UPCOMING_LABEL] = taskRunTaskRef(tr, pr)
			}
			c.trGaps.With(labels).Observe(float64(gap))
			continue
		}

		firstKid := sortedTaskRunsByCreateTimes[0]

		// so using the first taskrun completion time addresses sequential / chaining dependencies;
		// for parallel, if the first taskrun's completion time is not after this taskrun's create time,
		// that means parallel taskruns, and we work off of the pipelinerun; NOTE: this focuses on "top level" parallel task runs
		// with absolutely no dependencies.  Once any sort of dependency is established, there are no more top level parallel taskruns.
		if firstKid.Status.CompletionTime != nil && firstKid.Status.CompletionTime.Time.After(tr.CreationTimestamp.Time) {
			ctrl.Log.V(4).Info(fmt.Sprintf("task %s considered parallel for pipeline %s", taskRunTaskRef(tr, pr), prRef))
			gap := tr.CreationTimestamp.Time.Sub(pr.CreationTimestamp.Time).Milliseconds()
			if c.additionalLabels {
				labels[COMPLETED_LABEL] = prRef
				labels[UPCOMING_LABEL] = taskRunTaskRef(tr, pr)
			}
			c.trGaps.With(labels).Observe(float64(gap))
			continue
		}

		// Conversely, task run chains can run in parallel, and a taskrun can depend on multiple chains or threads of taskruns. We want to find the chain
		// that finished last, but before we are created.  We traverse through our reverse sorted on completion time list to determine that.  But yes, we don't reproduce the DAG
		// graph (there is no clean dependency import path in tekton for that) to confirm the edges.  This approximation is sufficient.

		// get whatever completed first
		timeToCalculateWith := time.Time{}
		trToCalculateWith := &v1beta1.TaskRun{}
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
			ctrl.Log.V(4).Info(fmt.Sprintf("comparing candidate %s to current task %s", taskRunTaskRef(tr2, pr), taskRunTaskRef(tr, pr)))
			if !tr2.Status.CompletionTime.Time.After(tr.CreationTimestamp.Time) {
				ctrl.Log.V(4).Info(fmt.Sprintf("%s did not complete after so use it to compute gap for current task %s", taskRunTaskRef(tr2, pr), taskRunTaskRef(tr, pr)))
				trToCalculateWith = tr2
				timeToCalculateWith = tr2.Status.CompletionTime.Time
				break
			}
			ctrl.Log.V(4).Info(fmt.Sprintf("skipping %s as a gap candidate for current task %s is OK", taskRunTaskRef(tr2, pr), taskRunTaskRef(tr, pr)))
		}
		gap := tr.CreationTimestamp.Time.Sub(timeToCalculateWith).Milliseconds()
		if c.additionalLabels {
			labels[COMPLETED_LABEL] = taskRunTaskRef(trToCalculateWith, pr)
			labels[UPCOMING_LABEL] = taskRunTaskRef(tr, pr)
		}
		c.trGaps.With(labels).Observe(float64(gap))
	}

	return
}

func NewTaskRunCollector() *TaskRunCollector {
	labelNames := []string{NS_LABEL}
	trSchedStatNameEnabled := optionalLabelEnabled(ENABLE_TASKRUN_SCHEDULED_DURATION_TASKNAME_LABEL)
	if trSchedStatNameEnabled {
		labelNames = append(labelNames, TASK_NAME_LABEL)
	}
	durationScheduled := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "taskrun_duration_scheduled_seconds",
		Help: "Duration in seconds for a TaskRun to be scheduled.",
		// reminder: exponential buckets need a start value greater than 0
		// the results in buckets of 0.1, 0.5, 2.5, 12.5, 62.5, 312.5 seconds
		Buckets: prometheus.ExponentialBuckets(0.1, 5, 6),
	}, labelNames)

	taskRunCollector := &TaskRunCollector{
		durationScheduled: durationScheduled,
		trSchedNameLabel:  trSchedStatNameEnabled,
	}
	metrics.Registry.MustRegister(durationScheduled)

	return taskRunCollector

}

func (c *TaskRunCollector) bumpScheduledDuration(tr *v1beta1.TaskRun, scheduleDuration float64) {
	labels := map[string]string{NS_LABEL: tr.Namespace}
	switch {
	case c.trSchedNameLabel:
		val := ""
		if tr.Spec.TaskRef != nil {
			val = tr.Spec.TaskRef.Name
		} else {
			val = tr.Name
		}
		labels[TASK_NAME_LABEL] = val

	}
	c.durationScheduled.With(labels).Observe(scheduleDuration)
}
