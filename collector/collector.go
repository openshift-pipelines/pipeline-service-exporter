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
	"github.com/tektoncd/pipeline/pkg/apis/pipeline"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"k8s.io/apimachinery/pkg/types"
	"knative.dev/pkg/apis"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
	"sort"
	"strings"
	"time"
)

const (
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

type TaskRunScheduledCollector struct {
	durationScheduled *prometheus.HistogramVec
	trSchedNameLabel  bool
}

type TaskRunRunningCollector struct {
	durationRunning *prometheus.HistogramVec
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

func bumpTaskRunScheduledDuration(scheduleDuration float64, tr *v1beta1.TaskRun, metric *prometheus.HistogramVec) {
	labels := map[string]string{NS_LABEL: tr.Namespace, TASK_NAME_LABEL: taskRef(tr.Labels)}
	metric.With(labels).Observe(scheduleDuration)
}

func NewPipelineRunScheduledMetric() *prometheus.HistogramVec {
	labelNames := []string{NS_LABEL, PIPELINE_NAME_LABEL}
	durationScheduled := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "pipelinerun_duration_scheduled_seconds",
		Help: "Duration in seconds for a PipelineRun to be 'scheduled', meaning it has been received by the Tekton controller.  This is an indication of how quickly create events from the API server are arriving to the Tekton controller.",
		// reminder: exponential buckets need a start value greater than 0
		// the results in buckets of 0.1, 0.5, 2.5, 12.5, 62.5, 312.5 seconds
		Buckets: prometheus.ExponentialBuckets(0.1, 5, 6),
	}, labelNames)

	metrics.Registry.MustRegister(durationScheduled)

	return durationScheduled
}

func bumpPipelineRunScheduledDuration(scheduleDuration float64, pr *v1beta1.PipelineRun, metric *prometheus.HistogramVec) {
	labels := map[string]string{NS_LABEL: pr.Namespace, PIPELINE_NAME_LABEL: pipelineRunPipelineRef(pr)}
	metric.With(labels).Observe(scheduleDuration)
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
	labelNames := []string{NS_LABEL, STATUS_LABEL}
	additionalLabels := optionalMetricEnabled(ENABLE_GAP_METRIC_ADDITIONAL_LABELS)
	if additionalLabels {
		labelNames = append(labelNames, PIPELINE_NAME_LABEL, COMPLETED_LABEL, UPCOMING_LABEL)
	}
	trGaps := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "pipelinerun_gap_between_taskruns_milliseconds",
		Help: "Duration in milliseconds between a taskrun completing and the next taskrun being created within a pipelinerun.  For a pipelinerun's first taskrun, the duration is the time between that taskrun's creation and the pipelinerun's creation.",
		// reminder: exponential buckets need a start value greater than 0
		// the results in buckets of 100, 500, 2500, 12500, 62500, 312500 milliseconds
		Buckets: prometheus.ExponentialBuckets(float64(100), float64(5), 6),
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
		succeedCondition := pr.Status.GetCondition(apis.ConditionSucceeded)
		if succeedCondition == nil {
			ctrl.Log.Info(fmt.Sprintf("WARNING: pipielinerun %s:%s marked done but has nil succeed condition", pr.Namespace, pr.Name))
			continue
		}
		if succeedCondition.IsUnknown() {
			ctrl.Log.Info(fmt.Sprintf("WARNING: pipielinerun %s:%s marked done but has unknown succeed condition", pr.Namespace, pr.Name))
			continue
		}
		status := SUCCEEDED
		if succeedCondition.IsFalse() {
			status = FAILED
		}
		labels := map[string]string{NS_LABEL: pr.Namespace, STATUS_LABEL: status}
		if c.additionalLabels {
			labels[PIPELINE_NAME_LABEL] = prRef
		}

		if index == 0 {
			ctrl.Log.V(4).Info(fmt.Sprintf("first task %s for pipeline %s", taskRef(tr.Labels), prRef))
			// our first task is simple, just work off of the pipelinerun
			gap := tr.CreationTimestamp.Time.Sub(pr.CreationTimestamp.Time).Milliseconds()
			if c.additionalLabels {
				labels[COMPLETED_LABEL] = prRef
				labels[UPCOMING_LABEL] = taskRef(tr.Labels)
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
			ctrl.Log.V(4).Info(fmt.Sprintf("task %s considered parallel for pipeline %s", taskRef(tr.Labels), prRef))
			gap := tr.CreationTimestamp.Time.Sub(pr.CreationTimestamp.Time).Milliseconds()
			if c.additionalLabels {
				labels[COMPLETED_LABEL] = prRef
				labels[UPCOMING_LABEL] = taskRef(tr.Labels)
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
			ctrl.Log.V(4).Info(fmt.Sprintf("comparing candidate %s to current task %s", taskRef(tr2.Labels), taskRef(tr.Labels)))
			if !tr2.Status.CompletionTime.Time.After(tr.CreationTimestamp.Time) {
				ctrl.Log.V(4).Info(fmt.Sprintf("%s did not complete after so use it to compute gap for current task %s", taskRef(tr2.Labels), taskRef(tr.Labels)))
				trToCalculateWith = tr2
				timeToCalculateWith = tr2.Status.CompletionTime.Time
				break
			}
			ctrl.Log.V(4).Info(fmt.Sprintf("skipping %s as a gap candidate for current task %s is OK", taskRef(tr2.Labels), taskRef(tr.Labels)))
		}
		gap := tr.CreationTimestamp.Time.Sub(timeToCalculateWith).Milliseconds()
		if c.additionalLabels {
			labels[COMPLETED_LABEL] = taskRef(trToCalculateWith.Labels)
			labels[UPCOMING_LABEL] = taskRef(tr.Labels)
		}
		c.trGaps.With(labels).Observe(float64(gap))
	}

	return
}

func NewTaskRunScheduledMetric() *prometheus.HistogramVec {
	labelNames := []string{NS_LABEL, TASK_NAME_LABEL}
	durationScheduled := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "taskrun_duration_scheduled_seconds",
		Help: "Duration in seconds for a TaskRun to be 'scheduled', meaning it has been received by the Tekton controller.  This is an indication of how quickly create events from the API server are arriving to the Tekton controller.",
		// reminder: exponential buckets need a start value greater than 0
		// the results in buckets of 0.1, 0.5, 2.5, 12.5, 62.5, 312.5 seconds
		Buckets: prometheus.ExponentialBuckets(0.1, 5, 6),
	}, labelNames)

	metrics.Registry.MustRegister(durationScheduled)

	return durationScheduled

}

func NewPodCreateToKubeletDurationMetric() *prometheus.HistogramVec {
	labelNames := []string{NS_LABEL, TASK_NAME_LABEL}
	metric := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "taskrun_pod_duration_kubelet_acknowledged_milliseconds",
		Help:    "Duration in milliseconds between the pod creation time and pod start time, where the pod start time is set once the kubelet has acknowledged the pod, but has not yet pulled its images.",
		Buckets: prometheus.ExponentialBuckets(float64(100), float64(5), 6),
	}, labelNames)
	metrics.Registry.MustRegister(metric)
	return metric
}

func NewPodKubeletToContainerStartDurationMetric() *prometheus.HistogramVec {
	labelNames := []string{NS_LABEL, TASK_NAME_LABEL}
	metric := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "taskrun_pod_duration_kubelet_to_container_start_milliseconds",
		Help:    "Duration in milliseconds between the pod start time and the first container to start. This should include any overhead to pull container images, plus any kubelet to linux scheduling overhead.",
		Buckets: prometheus.ExponentialBuckets(float64(100), float64(5), 6),
	}, labelNames)
	metrics.Registry.MustRegister(metric)
	return metric
}
