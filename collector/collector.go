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
	"os"
	ctrl "sigs.k8s.io/controller-runtime"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

const (
	ENABLE_PIPELINERUN_SCHEDULED_DURATION_PIPELINENAME_LABEL = "ENABLE_PIPELINERUN_SCHEDULED_DURATION_PIPELINENAME_LABEL"
	ENABLE_TASKRUN_SCHEDULED_DURATION_TASKNAME_LABEL         = "ENABLE_TASKRUN_SCHEDULED_DURATION_TASKNAME_LABEL"
	ENABLE_GAP_METRIC                                        = "ENABLE_GAP_METRIC"
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
	trGaps *prometheus.HistogramVec
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
	switch {
	case c.prSchedNameLabel:
		val := ""
		if pr.Spec.PipelineRef != nil {
			val = pr.Spec.PipelineRef.Name
		} else {
			val = pr.Name
		}
		labels[PIPELINE_NAME_LABEL] = val

	}
	c.durationScheduled.With(labels).Observe(scheduleDuration)
}

func NewPipelineRunTaskRunGapCollector() *PipelineRunTaskRunGapCollector {
	labelNames := []string{NS_LABEL}
	trGaps := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "pipelinerun_duration_between_taskruns_milliseconds",
		Help: "Duration in milliseconds between a taskrun completing and the next taskrun being created within a pipelinerun.  If a pipelinerun only has one taskrun, use pipelinerun_duration_scheduled_seconds.",
		// reminder: exponential buckets need a start value greater than 0
		// the results in buckets of 0.1, 0.5, 2.5, 12.5, 62.5, 312.5 milliseconds
		Buckets: prometheus.ExponentialBuckets(0.1, 5, 6),
	}, append(labelNames, PIPELINE_NAME_LABEL, COMPLETED_LABEL, UPCOMING_LABEL))

	pipelineRunTaskRunGapCollector := &PipelineRunTaskRunGapCollector{
		trGaps: trGaps,
	}
	metrics.Registry.MustRegister(trGaps)

	return pipelineRunTaskRunGapCollector
}

func (c *PipelineRunTaskRunGapCollector) bumpGapDuration(pr *v1beta1.PipelineRun, oc client.Client, ctx context.Context) {
	labels := map[string]string{NS_LABEL: pr.Namespace}
	val := ""
	if pr.Spec.PipelineRef != nil {
		val = pr.Spec.PipelineRef.Name
	} else {
		val = pr.Name
	}
	labels[PIPELINE_NAME_LABEL] = val

	if len(pr.Status.ChildReferences) < 2 {
		return
	}

	for index, kidRef := range pr.Status.ChildReferences {
		if kidRef.Kind != "TaskRun" {
			continue
		}
		if index == 0 {
			kid := &v1beta1.TaskRun{}
			err := oc.Get(ctx, types.NamespacedName{Namespace: pr.Namespace, Name: kidRef.Name}, kid)
			if err != nil {
				ctrl.Log.Info(fmt.Sprintf("could not calcuate gap for taskrun %s:%s: %s", pr.Namespace, kidRef.Name, err.Error()))
				continue
			}
			gap := kid.CreationTimestamp.Time.Sub(pr.CreationTimestamp.Time)
			labels[COMPLETED_LABEL] = pr.Name
			labels[UPCOMING_LABEL] = kid.Name
			c.trGaps.With(labels).Observe(float64(gap.Milliseconds()))
			continue
		}

		prevKidRef := pr.Status.ChildReferences[index-1]
		if prevKidRef.Kind != "TaskRun" {
			continue
		}
		kid := &v1beta1.TaskRun{}
		err := oc.Get(ctx, types.NamespacedName{Namespace: pr.Namespace, Name: kidRef.Name}, kid)
		if err != nil {
			ctrl.Log.Info(fmt.Sprintf("could not calcuate gap for taskrun %s:%s: %s", pr.Namespace, kidRef.Name, err.Error()))
			continue
		}
		prevKid := &v1beta1.TaskRun{}
		err = oc.Get(ctx, types.NamespacedName{Namespace: pr.Namespace, Name: prevKidRef.Name}, prevKid)
		if err != nil {
			ctrl.Log.Info(fmt.Sprintf("could not calcuate gap for taskrun %s:%s: %s", pr.Namespace, prevKidRef.Name, err.Error()))
			continue
		}

		gap := kid.CreationTimestamp.Time.Sub(prevKid.Status.CompletionTime.Time).Milliseconds()
		labels[COMPLETED_LABEL] = prevKidRef.Name
		labels[UPCOMING_LABEL] = kidRef.Name
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
