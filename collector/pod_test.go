package collector

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"testing"
	"time"
)

func TestCreateKubeletLatencyFilter_Update(t *testing.T) {
	filter := &createKubeletLatencyFilter{
		metric: NewPodCreateToKubeletDurationMetric(),
	}
	for _, tc := range []struct {
		name           string
		oldPod         *corev1.Pod
		newPod         *corev1.Pod
		expectedRC     bool
		expectedMetric bool
	}{
		{
			name:   "both start times nil",
			oldPod: &corev1.Pod{},
			newPod: &corev1.Pod{},
		},
		{
			name:           "only new start time set",
			expectedMetric: true,
			oldPod:         &corev1.Pod{},
			newPod: &corev1.Pod{
				Status: corev1.PodStatus{StartTime: &metav1.Time{Time: time.Now()}},
			},
		},
		{
			name: "both start times set",
			oldPod: &corev1.Pod{
				Status: corev1.PodStatus{StartTime: &metav1.Time{Time: time.Now()}},
			},
			newPod: &corev1.Pod{
				Status: corev1.PodStatus{StartTime: &metav1.Time{Time: time.Now()}},
			},
		},
	} {
		ev := event.UpdateEvent{
			ObjectOld: tc.oldPod,
			ObjectNew: tc.newPod,
		}
		rc := filter.Update(ev)
		if rc != tc.expectedRC {
			t.Errorf(fmt.Sprintf("tc %s expected %v but got %v", tc.name, tc.expectedRC, rc))
		}
		if tc.expectedMetric {
			validateHistogramVec(t, filter.metric, prometheus.Labels{NS_LABEL: tc.newPod.Namespace, TASK_NAME_LABEL: taskRef(tc.newPod.Labels)}, false)
		}
	}
}

func TestKubeletContainerLatencyFilter_Update(t *testing.T) {
	filter := &kubeletContainerLatencyFilter{
		metric: NewPodKubeletToContainerStartDurationMetric(),
	}
	for _, tc := range []struct {
		name           string
		oldPod         *corev1.Pod
		newPod         *corev1.Pod
		expectedRC     bool
		expectedMetric bool
	}{
		{
			name:   "both start times nil",
			oldPod: &corev1.Pod{},
			newPod: &corev1.Pod{},
		},
		{
			name: "both start times set, no container status",
			oldPod: &corev1.Pod{
				Status: corev1.PodStatus{StartTime: &metav1.Time{Time: time.Now()}},
			},
			newPod: &corev1.Pod{
				Status: corev1.PodStatus{StartTime: &metav1.Time{Time: time.Now()}},
			},
		},
		{
			name: "both start times set, no container status",
			oldPod: &corev1.Pod{
				Status: corev1.PodStatus{StartTime: &metav1.Time{Time: time.Now()}},
			},
			newPod: &corev1.Pod{
				Status: corev1.PodStatus{StartTime: &metav1.Time{Time: time.Now()}},
			},
		},
		{
			name: "both start times set, new only container running state",
			oldPod: &corev1.Pod{
				Status: corev1.PodStatus{StartTime: &metav1.Time{Time: time.Now()}},
			},
			newPod: &corev1.Pod{
				Status: corev1.PodStatus{
					StartTime: &metav1.Time{Time: time.Now()},
					ContainerStatuses: []corev1.ContainerStatus{
						{
							State: corev1.ContainerState{
								Running: &corev1.ContainerStateRunning{
									StartedAt: metav1.NewTime(time.Now()),
								},
							},
						},
					},
				},
			},
		},
		{
			name:           "both start times set, new only container terminated state",
			expectedMetric: true,
			oldPod: &corev1.Pod{
				Status: corev1.PodStatus{StartTime: &metav1.Time{Time: time.Now()}},
			},
			newPod: &corev1.Pod{
				Status: corev1.PodStatus{
					StartTime: &metav1.Time{Time: time.Now()},
					ContainerStatuses: []corev1.ContainerStatus{
						{
							State: corev1.ContainerState{
								Terminated: &corev1.ContainerStateTerminated{
									StartedAt: metav1.NewTime(time.Now()),
								},
							},
						},
					},
				},
			},
		},
		{
			name:           "both start times set, old container array set, not status, new container terminated state",
			expectedMetric: true,
			oldPod: &corev1.Pod{
				Status: corev1.PodStatus{
					StartTime:         &metav1.Time{Time: time.Now()},
					ContainerStatuses: []corev1.ContainerStatus{},
				},
			},
			newPod: &corev1.Pod{
				Status: corev1.PodStatus{
					StartTime: &metav1.Time{Time: time.Now()},
					ContainerStatuses: []corev1.ContainerStatus{
						{
							State: corev1.ContainerState{
								Terminated: &corev1.ContainerStateTerminated{
									StartedAt: metav1.NewTime(time.Now()),
								},
							},
						},
					},
				},
			},
		},
		{
			name:           "both start times set, old container array set, not status, new container running state",
			expectedMetric: true,
			oldPod: &corev1.Pod{
				Status: corev1.PodStatus{
					StartTime:         &metav1.Time{Time: time.Now()},
					ContainerStatuses: []corev1.ContainerStatus{},
				},
			},
			newPod: &corev1.Pod{
				Status: corev1.PodStatus{
					StartTime: &metav1.Time{Time: time.Now()},
					ContainerStatuses: []corev1.ContainerStatus{
						{
							State: corev1.ContainerState{
								Running: &corev1.ContainerStateRunning{
									StartedAt: metav1.NewTime(time.Now()),
								},
							},
						},
					},
				},
			},
		},
	} {
		ev := event.UpdateEvent{
			ObjectOld: tc.oldPod,
			ObjectNew: tc.newPod,
		}
		rc := filter.Update(ev)
		if rc != tc.expectedRC {
			t.Errorf(fmt.Sprintf("tc %s expected %v but got %v", tc.name, tc.expectedRC, rc))
		}
		if tc.expectedMetric {
			validateHistogramVec(t, filter.metric, prometheus.Labels{NS_LABEL: tc.newPod.Namespace, TASK_NAME_LABEL: taskRef(tc.newPod.Labels)}, false)
		}
	}
}
