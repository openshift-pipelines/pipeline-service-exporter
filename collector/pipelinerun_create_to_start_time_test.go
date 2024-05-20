package collector

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
	"testing"
	"time"
)

func TestPipelineRunStartTimeEventFilter_Update(t *testing.T) {
	filter := &startTimeEventFilter{}
	for _, tc := range []struct {
		name       string
		oldPR      *v1.PipelineRun
		newPR      *v1.PipelineRun
		expectedRC bool
	}{
		{
			name:  "not started",
			oldPR: &v1.PipelineRun{},
			newPR: &v1.PipelineRun{},
		},
		{
			name:  "just started",
			oldPR: &v1.PipelineRun{},
			newPR: &v1.PipelineRun{
				Status: v1.PipelineRunStatus{
					PipelineRunStatusFields: v1.PipelineRunStatusFields{StartTime: &metav1.Time{}},
				},
			},
		},
		{
			name: "udpate after started",
			oldPR: &v1.PipelineRun{
				Status: v1.PipelineRunStatus{
					PipelineRunStatusFields: v1.PipelineRunStatusFields{StartTime: &metav1.Time{}},
				},
			},
			newPR: &v1.PipelineRun{
				Status: v1.PipelineRunStatus{
					PipelineRunStatusFields: v1.PipelineRunStatusFields{StartTime: &metav1.Time{}},
				},
			},
		},
		{
			name: "completed",
			oldPR: &v1.PipelineRun{
				Status: v1.PipelineRunStatus{
					PipelineRunStatusFields: v1.PipelineRunStatusFields{StartTime: &metav1.Time{}},
				},
			},
			newPR: &v1.PipelineRun{
				Status: v1.PipelineRunStatus{
					Status: duckv1.Status{
						ObservedGeneration: 0,
						Conditions: duckv1.Conditions{{
							Type:   "Succeeded",
							Status: corev1.ConditionTrue,
						}},
						Annotations: nil,
					},
					PipelineRunStatusFields: v1.PipelineRunStatusFields{StartTime: &metav1.Time{}},
				},
			},
		},
	} {
		ev := event.UpdateEvent{
			ObjectOld: tc.oldPR,
			ObjectNew: tc.newPR,
		}
		filter.metric = NewPipelineRunScheduledMetric()
		rc := filter.Update(ev)
		if rc != tc.expectedRC {
			t.Errorf(fmt.Sprintf("tc %s expected %v but got %v", tc.name, tc.expectedRC, rc))
		}
		metrics.Registry.Unregister(filter.metric)
	}

}

func TestPipelineRunScheduleCollection(t *testing.T) {
	mockPipelineRuns := []*v1.PipelineRun{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:              "test-pipelinerun-1",
				Namespace:         "test-namespace",
				UID:               types.UID("test-pipelinerun-1"),
				CreationTimestamp: metav1.NewTime(time.Now().UTC()),
			},
			Status: v1.PipelineRunStatus{
				Status: duckv1.Status{
					ObservedGeneration: 0,
					Conditions: duckv1.Conditions{{
						Type:   "Succeeded",
						Status: corev1.ConditionTrue,
					}},
					Annotations: nil,
				},
				PipelineRunStatusFields: v1.PipelineRunStatusFields{
					StartTime:      &metav1.Time{Time: time.Now().UTC().Add(5 * time.Second)},
					CompletionTime: &metav1.Time{Time: time.Now().UTC().Add(10 * time.Second)},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:              "test-pipelinerun-2",
				Namespace:         "test-namespace",
				UID:               types.UID("test-pipelinerun-2"),
				CreationTimestamp: metav1.NewTime(time.Now().UTC()),
			},
			Status: v1.PipelineRunStatus{
				Status: duckv1.Status{
					ObservedGeneration: 0,
					Conditions: duckv1.Conditions{{
						Type:   "Succeeded",
						Status: corev1.ConditionTrue,
					}},
					Annotations: nil,
				},
				PipelineRunStatusFields: v1.PipelineRunStatusFields{
					StartTime:      &metav1.Time{Time: time.Now().UTC().Add(5 * time.Second)},
					CompletionTime: &metav1.Time{Time: time.Now().UTC().Add(10 * time.Second)},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:              "test-pipelinerun-3",
				Namespace:         "test-namespace",
				UID:               types.UID("test-pipelinerun-3"),
				CreationTimestamp: metav1.NewTime(time.Now().UTC()),
			},
			Status: v1.PipelineRunStatus{
				Status: duckv1.Status{
					ObservedGeneration: 0,
					Conditions: duckv1.Conditions{{
						Type:   "Succeeded",
						Status: corev1.ConditionTrue,
					}},
					Annotations: nil,
				},
				PipelineRunStatusFields: v1.PipelineRunStatusFields{
					ChildReferences: []v1.ChildStatusReference{
						{
							TypeMeta: runtime.TypeMeta{
								Kind: "TaskRun",
							},
							Name: "test-taskrun-1",
						},
						{
							TypeMeta: runtime.TypeMeta{
								Kind: "TaskRun",
							},
							Name: "test-taskrun-2",
						},
					},
					StartTime:      &metav1.Time{Time: time.Now().UTC().Add(5 * time.Second)},
					CompletionTime: &metav1.Time{Time: time.Now().UTC().Add(10 * time.Second)},
				},
			},
		},
	}
	for _, pr := range mockPipelineRuns {
		metric := NewPipelineRunScheduledMetric()
		label := prometheus.Labels{NS_LABEL: "test-namespace", STATUS_LABEL: SUCCEEDED}
		bumpPipelineRunScheduledDuration(calculateScheduledDurationPipelineRun(pr), pr, metric)
		validateHistogramVec(t, metric, label, false)
		metrics.Registry.Unregister(metric)
	}

}

func TestCalculatePipelineRunScheduledDuration(t *testing.T) {
	now := time.Now()
	for _, tc := range []struct {
		expectedAmt float64
		pr          *v1.PipelineRun
	}{
		{
			expectedAmt: 5,
			pr: &v1.PipelineRun{
				ObjectMeta: metav1.ObjectMeta{
					Name:              "test-pipelinerun-1",
					Namespace:         "test-namespace",
					CreationTimestamp: metav1.NewTime(now),
				},
				Status: v1.PipelineRunStatus{
					Status: duckv1.Status{
						ObservedGeneration: 0,
						Conditions: duckv1.Conditions{{
							Type:   "Succeeded",
							Status: corev1.ConditionTrue,
						}},
						Annotations: nil,
					},
					PipelineRunStatusFields: v1.PipelineRunStatusFields{
						StartTime:      &metav1.Time{Time: now.Add(5 * time.Second)},
						CompletionTime: &metav1.Time{Time: now.Add(10 * time.Second)},
					},
				},
			},
		},
		{
			expectedAmt: 5,
			pr: &v1.PipelineRun{
				ObjectMeta: metav1.ObjectMeta{
					Name:              "test-pipelinerun-1",
					Namespace:         "test-namespace",
					CreationTimestamp: metav1.NewTime(now),
				},
				Status: v1.PipelineRunStatus{
					Status: duckv1.Status{
						ObservedGeneration: 0,
						Conditions: duckv1.Conditions{{
							Type:   "Failed",
							Status: corev1.ConditionTrue,
						}},
						Annotations: nil,
					},
					PipelineRunStatusFields: v1.PipelineRunStatusFields{
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
