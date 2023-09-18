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
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

/*
  Observed sequence from testing and tekton/k8s code examination:

- taskrun created by external user or pipelinerun reconciler, and k8s sets the taskrun create timestap
- on first reconcile event loop, taskrun ConditionSucceeded condition is intialized to pending and the startime is set
- "prepare" will see if resolutionrequest is needed, and if so, error is returned to controller fw and requeue happens
- otherwise, or on next reconcile event loop, volume claim templates / workspaces processed, params processed, and if no errors, pod create attempted
- if error requeue
- otherwise pod status converted to taskrun status, with conditions now updated and set running, controller then puts back on queue with requeueAfter(timeDuration)

Now concurrent event streams from both Pod updates or TaskRun updates, or requeuesAfter(timeDuration) calling reconcile again, can update task run status and conditions

On pods
- create time stems from tekton creating
- pod start time means kubelet has "accepted" per godoc for scheduling, but no image pulls have occurred, where
using of "latest" for the tag means always pull per godoc, but use of a specific SHA means pull if not on local CRI-O / "node cache"

So,
- no real diff of worth between taskrun startime and its conditions
- pod create vs. pod start time captures how long the external factor of the Kubelet agreeing the schedule the pod takes
- pod start time vs. first container start captures how long to pull images and schedule the container
- where as the upstream latency metric of the last transition time of the `corev1.PodScheduled` condition minus the pod create time
is "perhaps" the sum of those two
*/

func calcuateScheduledDuration(created, started time.Time) float64 {
	if created.IsZero() || started.IsZero() {
		return 0
	}
	return started.Sub(created).Seconds()
}

func calculateScheduledDurationPipelineRun(pipelineRun *v1beta1.PipelineRun) float64 {
	return calcuateScheduledDuration(pipelineRun.CreationTimestamp.Time, pipelineRun.Status.StartTime.Time)
}

func calculateScheduledDurationTaskRun(taskrun *v1beta1.TaskRun) float64 {
	return calcuateScheduledDuration(taskrun.CreationTimestamp.Time, taskrun.Status.StartTime.Time)
}

// this minimally captures any time the kubelet spends pulling container images for the pod
func calculateTaskRunPodStartedToFirstContainerStartedDuration(pod *corev1.Pod) float64 {
	if pod.Status.StartTime == nil || pod.Status.StartTime.IsZero() {
		return 0
	}
	if len(pod.Status.ContainerStatuses) == 0 {
		return 0
	}
	var firstTime *metav1.Time
	for _, cs := range pod.Status.ContainerStatuses {
		if cs.State.Running != nil && !cs.State.Running.StartedAt.IsZero() {
			if firstTime == nil {
				firstTime = &cs.State.Running.StartedAt
				continue
			}
			if cs.State.Running.StartedAt.Before(firstTime) {
				firstTime = &cs.State.Running.StartedAt
				continue
			}
		}
		if cs.State.Terminated != nil && !cs.State.Terminated.StartedAt.IsZero() {
			if firstTime == nil {
				firstTime = &cs.State.Terminated.StartedAt
				continue
			}
			if cs.State.Terminated.StartedAt.Before(firstTime) {
				firstTime = &cs.State.Terminated.StartedAt
				continue
			}
		}
	}
	if firstTime == nil {
		return 0
	}
	return float64(firstTime.Time.Sub(pod.Status.StartTime.Time).Milliseconds())
}

// this captures how long it takes for the kubelet to accept the pod after the pod is created
func calculateTaskRunPodCreatedToKubeletAcceptsAndStartTimeSetDuration(pod *corev1.Pod) float64 {
	if pod.Status.StartTime == nil || pod.Status.StartTime.IsZero() {
		return 0
	}
	return float64(pod.Status.StartTime.Time.Sub(pod.CreationTimestamp.Time).Milliseconds())
}
