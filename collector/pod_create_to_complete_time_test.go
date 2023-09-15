package collector

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"testing"
	"time"
)

func TestPodCreateToCompleteFilter_Update(t *testing.T) {
	filter := NewPodCreateToCompleteFilter()
	now := time.Now()
	for _, tc := range []struct {
		name                  string
		old                   *corev1.Pod
		new                   *corev1.Pod
		expectedRC            bool
		expectedNonZeroMetric bool
	}{
		{
			name: "not started",
			old:  &corev1.Pod{},
			new:  &corev1.Pod{},
		},
		{
			name: "running",
			old: &corev1.Pod{
				Status: corev1.PodStatus{
					ContainerStatuses: []corev1.ContainerStatus{
						{
							State: corev1.ContainerState{
								Running:    &corev1.ContainerStateRunning{},
								Terminated: nil,
							},
						},
						{
							State: corev1.ContainerState{
								Running:    &corev1.ContainerStateRunning{},
								Terminated: nil,
							},
						},
					},
				},
			},
			new: &corev1.Pod{
				Status: corev1.PodStatus{
					ContainerStatuses: []corev1.ContainerStatus{
						{
							State: corev1.ContainerState{
								Running:    &corev1.ContainerStateRunning{},
								Terminated: nil,
							},
						},
						{
							State: corev1.ContainerState{
								Running:    &corev1.ContainerStateRunning{},
								Terminated: nil,
							},
						},
					},
				},
			},
		},
		{
			name: "not all done 1",
			old: &corev1.Pod{
				Status: corev1.PodStatus{
					ContainerStatuses: []corev1.ContainerStatus{
						{
							State: corev1.ContainerState{
								Running:    &corev1.ContainerStateRunning{},
								Terminated: nil,
							},
						},
						{
							State: corev1.ContainerState{
								Running:    &corev1.ContainerStateRunning{},
								Terminated: nil,
							},
						},
					},
				},
			},
			new: &corev1.Pod{
				Status: corev1.PodStatus{
					ContainerStatuses: []corev1.ContainerStatus{
						{
							State: corev1.ContainerState{
								Running: nil,
								Terminated: &corev1.ContainerStateTerminated{
									FinishedAt: metav1.NewTime(now.Add(3 * time.Second)),
								},
							},
						},
						{
							State: corev1.ContainerState{
								Running:    &corev1.ContainerStateRunning{},
								Terminated: nil,
							},
						},
					},
				},
			},
		},
		{
			name: "not all done 2",
			old: &corev1.Pod{
				Status: corev1.PodStatus{
					ContainerStatuses: []corev1.ContainerStatus{
						{
							State: corev1.ContainerState{
								Running: &corev1.ContainerStateRunning{},
								Terminated: &corev1.ContainerStateTerminated{
									FinishedAt: metav1.NewTime(now.Add(3 * time.Second)),
								},
							},
						},
						{
							State: corev1.ContainerState{
								Running:    &corev1.ContainerStateRunning{},
								Terminated: nil,
							},
						},
					},
				},
			},
			new: &corev1.Pod{
				Status: corev1.PodStatus{
					ContainerStatuses: []corev1.ContainerStatus{
						{
							State: corev1.ContainerState{
								Running: nil,
								Terminated: &corev1.ContainerStateTerminated{
									FinishedAt: metav1.NewTime(now.Add(3 * time.Second)),
								},
							},
						},
						{
							State: corev1.ContainerState{
								Running:    &corev1.ContainerStateRunning{},
								Terminated: nil,
							},
						},
					},
				},
			},
		},
		{
			name:                  "done",
			expectedNonZeroMetric: true,
			old: &corev1.Pod{
				Status: corev1.PodStatus{
					ContainerStatuses: []corev1.ContainerStatus{
						{
							State: corev1.ContainerState{
								Running: &corev1.ContainerStateRunning{},
								Terminated: &corev1.ContainerStateTerminated{
									FinishedAt: metav1.NewTime(now.Add(3 * time.Second)),
								},
							},
						},
						{
							State: corev1.ContainerState{
								Running:    &corev1.ContainerStateRunning{},
								Terminated: nil,
							},
						},
					},
				},
			},
			new: &corev1.Pod{
				Status: corev1.PodStatus{
					ContainerStatuses: []corev1.ContainerStatus{
						{
							State: corev1.ContainerState{
								Running: nil,
								Terminated: &corev1.ContainerStateTerminated{
									FinishedAt: metav1.NewTime(now.Add(3 * time.Second)),
								},
							},
						},
						{
							State: corev1.ContainerState{
								Running: &corev1.ContainerStateRunning{},
								Terminated: &corev1.ContainerStateTerminated{
									FinishedAt: metav1.NewTime(now.Add(3 * time.Second)),
								},
							},
						},
					},
				},
			},
		},
		{
			name: "event after done",
			old: &corev1.Pod{
				Status: corev1.PodStatus{
					ContainerStatuses: []corev1.ContainerStatus{
						{
							State: corev1.ContainerState{
								Running: &corev1.ContainerStateRunning{},
								Terminated: &corev1.ContainerStateTerminated{
									FinishedAt: metav1.NewTime(now.Add(3 * time.Second)),
								},
							},
						},
						{
							State: corev1.ContainerState{
								Running: &corev1.ContainerStateRunning{},
								Terminated: &corev1.ContainerStateTerminated{
									FinishedAt: metav1.NewTime(now.Add(3 * time.Second)),
								},
							},
						},
					},
				},
			},
			new: &corev1.Pod{
				Status: corev1.PodStatus{
					ContainerStatuses: []corev1.ContainerStatus{
						{
							State: corev1.ContainerState{
								Running: nil,
								Terminated: &corev1.ContainerStateTerminated{
									FinishedAt: metav1.NewTime(now.Add(3 * time.Second)),
								},
							},
						},
						{
							State: corev1.ContainerState{
								Running: &corev1.ContainerStateRunning{},
								Terminated: &corev1.ContainerStateTerminated{
									FinishedAt: metav1.NewTime(now.Add(3 * time.Second)),
								},
							},
						},
					},
				},
			},
		},
	} {
		filter.duration.Reset()
		ev := event.UpdateEvent{
			ObjectOld: tc.old,
			ObjectNew: tc.new,
		}
		rc := filter.Update(ev)
		if rc != tc.expectedRC {
			t.Errorf(fmt.Sprintf("tc %s expected %v but got %v", tc.name, tc.expectedRC, rc))
		}
		labels := prometheus.Labels{NS_LABEL: tc.new.Namespace, PIPELINE_NAME_LABEL: pipelineRef(tc.new.Labels), TASK_NAME_LABEL: taskRef(tc.new.Labels)}
		if tc.expectedNonZeroMetric {
			validateHistogramVec(t, filter.duration, labels, false)
		} else {
			observer, err := filter.duration.GetMetricWith(labels)
			assert.NoError(t, err)
			assert.NotNil(t, observer)
			histogram := observer.(prometheus.Histogram)
			metric := &dto.Metric{}
			histogram.Write(metric)
			assert.NotNil(t, metric.Histogram)
			assert.Equal(t, uint64(0), *metric.Histogram.SampleCount)
		}
	}
}
