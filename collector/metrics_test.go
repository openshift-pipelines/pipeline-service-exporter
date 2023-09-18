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
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"testing"
	"time"
)

func TestCalculatePipelineRunScheduledDuration(t *testing.T) {
	now := time.Now()
	for _, tc := range []struct {
		expectedAmt float64
		pr          *v1beta1.PipelineRun
	}{
		{
			expectedAmt: 5,
			pr: &v1beta1.PipelineRun{
				ObjectMeta: metav1.ObjectMeta{
					Name:              "test-pipelinerun-1",
					Namespace:         "test-namespace",
					CreationTimestamp: metav1.NewTime(now),
				},
				Status: v1beta1.PipelineRunStatus{
					Status: duckv1.Status{
						ObservedGeneration: 0,
						Conditions: duckv1.Conditions{{
							Type:   "Succeeded",
							Status: corev1.ConditionTrue,
						}},
						Annotations: nil,
					},
					PipelineRunStatusFields: v1beta1.PipelineRunStatusFields{
						StartTime:      &metav1.Time{Time: now.Add(5 * time.Second)},
						CompletionTime: &metav1.Time{Time: now.Add(10 * time.Second)},
					},
				},
			},
		},
		{
			expectedAmt: 5,
			pr: &v1beta1.PipelineRun{
				ObjectMeta: metav1.ObjectMeta{
					Name:              "test-pipelinerun-1",
					Namespace:         "test-namespace",
					CreationTimestamp: metav1.NewTime(now),
				},
				Status: v1beta1.PipelineRunStatus{
					Status: duckv1.Status{
						ObservedGeneration: 0,
						Conditions: duckv1.Conditions{{
							Type:   "Failed",
							Status: corev1.ConditionTrue,
						}},
						Annotations: nil,
					},
					PipelineRunStatusFields: v1beta1.PipelineRunStatusFields{
						StartTime:      &metav1.Time{Time: now.Add(5 * time.Second)},
						CompletionTime: &metav1.Time{Time: now.Add(10 * time.Second)},
					},
				},
			},
		},
	} {
		got := calculateScheduledDurationPipelineRun(tc.pr)
		if got != tc.expectedAmt {
			t.Errorf("Scheduled Duration is not as expected. Got %v, expected %v", got, tc.expectedAmt)
		}
	}
}

func TestCalculateTaskRunScheduledDuration(t *testing.T) {
	now := time.Now()
	for _, tc := range []struct {
		expectedAmt float64
		tr          *v1beta1.TaskRun
	}{
		{
			expectedAmt: 5,
			tr: &v1beta1.TaskRun{
				ObjectMeta: metav1.ObjectMeta{
					Name:              "test-taskrun-1",
					Namespace:         "test-namespace",
					CreationTimestamp: metav1.NewTime(now),
				},
				Status: v1beta1.TaskRunStatus{
					Status: duckv1.Status{
						ObservedGeneration: 0,
						Conditions: duckv1.Conditions{{
							Type:   "Succeeded",
							Status: corev1.ConditionTrue,
						}},
						Annotations: nil,
					},
					TaskRunStatusFields: v1beta1.TaskRunStatusFields{
						StartTime:      &metav1.Time{Time: now.Add(5 * time.Second)},
						CompletionTime: &metav1.Time{Time: now.Add(10 * time.Second)},
					},
				},
			},
		},
		{
			expectedAmt: 5,
			tr: &v1beta1.TaskRun{
				ObjectMeta: metav1.ObjectMeta{
					Name:              "test-taskrun-1",
					Namespace:         "test-namespace",
					CreationTimestamp: metav1.NewTime(now),
				},
				Status: v1beta1.TaskRunStatus{
					Status: duckv1.Status{
						ObservedGeneration: 0,
						Conditions: duckv1.Conditions{{
							Type:   "Failed",
							Status: corev1.ConditionTrue,
						}},
						Annotations: nil,
					},
					TaskRunStatusFields: v1beta1.TaskRunStatusFields{
						StartTime:      &metav1.Time{Time: now.Add(5 * time.Second)},
						CompletionTime: &metav1.Time{Time: now.Add(10 * time.Second)},
					},
				},
			},
		},
	} {
		got := calculateScheduledDurationTaskRun(tc.tr)
		if got != tc.expectedAmt {
			t.Errorf("Scheduled Duration is not as expected. Got %v, expected %v", got, tc.expectedAmt)
		}
	}
}

func TestCalculateTaskRunPodStartedToFirstContainerStartedDuration(t *testing.T) {
	now := time.Now()
	for _, tc := range []struct {
		expectedAmt float64
		pod         *corev1.Pod
	}{
		{
			expectedAmt: 0,
			pod: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:              "test-taskrun-1",
					Namespace:         "test-namespace",
					CreationTimestamp: metav1.NewTime(now),
				},
			},
		},
		{
			expectedAmt: 0,
			pod: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:              "test-taskrun-1",
					Namespace:         "test-namespace",
					CreationTimestamp: metav1.NewTime(now),
				},
				Status: corev1.PodStatus{StartTime: &metav1.Time{Time: now.Add(3 * time.Second)}},
			},
		},
		{
			expectedAmt: 0,
			pod: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:              "test-taskrun-1",
					Namespace:         "test-namespace",
					CreationTimestamp: metav1.NewTime(now),
				},
				Status: corev1.PodStatus{StartTime: &metav1.Time{Time: now.Add(3 * time.Second)}},
			},
		},
		{
			expectedAmt: 2000,
			pod: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:              "test-taskrun-1",
					Namespace:         "test-namespace",
					CreationTimestamp: metav1.NewTime(now),
				},
				Status: corev1.PodStatus{
					StartTime: &metav1.Time{Time: now.Add(3 * time.Second)},
					ContainerStatuses: []corev1.ContainerStatus{
						{
							State: corev1.ContainerState{
								Running: &corev1.ContainerStateRunning{StartedAt: metav1.Time{Time: now.Add(5 * time.Second)}},
							},
						},
					},
				},
			},
		},
		{
			expectedAmt: 2000,
			pod: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:              "test-taskrun-1",
					Namespace:         "test-namespace",
					CreationTimestamp: metav1.NewTime(now),
				},
				Status: corev1.PodStatus{
					StartTime: &metav1.Time{Time: now.Add(3 * time.Second)},
					ContainerStatuses: []corev1.ContainerStatus{
						{
							State: corev1.ContainerState{
								Terminated: &corev1.ContainerStateTerminated{StartedAt: metav1.Time{Time: now.Add(5 * time.Second)}},
							},
						},
					},
				},
			},
		},
		{
			expectedAmt: 2000,
			pod: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:              "test-taskrun-1",
					Namespace:         "test-namespace",
					CreationTimestamp: metav1.NewTime(now),
				},
				Status: corev1.PodStatus{
					StartTime: &metav1.Time{Time: now.Add(3 * time.Second)},
					ContainerStatuses: []corev1.ContainerStatus{
						{
							State: corev1.ContainerState{
								Running: &corev1.ContainerStateRunning{StartedAt: metav1.Time{Time: now.Add(6 * time.Second)}},
							},
						},
						{
							State: corev1.ContainerState{
								Terminated: &corev1.ContainerStateTerminated{StartedAt: metav1.Time{Time: now.Add(5 * time.Second)}},
							},
						},
					},
				},
			},
		},
		{
			expectedAmt: 2000,
			pod: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:              "test-taskrun-1",
					Namespace:         "test-namespace",
					CreationTimestamp: metav1.NewTime(now),
				},
				Status: corev1.PodStatus{
					StartTime: &metav1.Time{Time: now.Add(3 * time.Second)},
					ContainerStatuses: []corev1.ContainerStatus{
						{
							State: corev1.ContainerState{
								Running: &corev1.ContainerStateRunning{StartedAt: metav1.Time{Time: now.Add(5 * time.Second)}},
							},
						},
						{
							State: corev1.ContainerState{
								Terminated: &corev1.ContainerStateTerminated{StartedAt: metav1.Time{Time: now.Add(6 * time.Second)}},
							},
						},
					},
				},
			},
		},
	} {
		got := calculateTaskRunPodStartedToFirstContainerStartedDuration(tc.pod)
		if got != tc.expectedAmt {
			t.Errorf("expected %v but got %v", tc.expectedAmt, got)
		}
	}
}

func TestCalculateTaskRunPodCreatedToKubeletAcceptsAndStartTimeSetDuration(t *testing.T) {
	now := time.Now()
	for _, tc := range []struct {
		expectedAmt float64
		pod         *corev1.Pod
	}{
		{
			expectedAmt: 3000,
			pod: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:              "test-taskrun-1",
					Namespace:         "test-namespace",
					CreationTimestamp: metav1.NewTime(now),
				},
				Status: corev1.PodStatus{StartTime: &metav1.Time{Time: now.Add(3 * time.Second)}},
			},
		},
		{
			expectedAmt: 0,
			pod: &corev1.Pod{
				ObjectMeta: metav1.ObjectMeta{
					Name:              "test-taskrun-1",
					Namespace:         "test-namespace",
					CreationTimestamp: metav1.NewTime(now),
				},
			},
		},
	} {
		got := calculateTaskRunPodCreatedToKubeletAcceptsAndStartTimeSetDuration(tc.pod)
		if got != tc.expectedAmt {
			t.Errorf("expected %v but got %v", tc.expectedAmt, got)
		}
	}
}
