package collector

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
	"testing"
	"time"
)

func TestTaskRunStartTimeEventFilter_Update(t *testing.T) {
	filter := &trStartTimeEventFilter{}
	for _, tc := range []struct {
		name       string
		oldTR      *v1.TaskRun
		newTR      *v1.TaskRun
		expectedRC bool
	}{
		{
			name:  "not started",
			oldTR: &v1.TaskRun{},
			newTR: &v1.TaskRun{},
		},
		{
			name:  "just started",
			oldTR: &v1.TaskRun{},
			newTR: &v1.TaskRun{
				Status: v1.TaskRunStatus{
					TaskRunStatusFields: v1.TaskRunStatusFields{StartTime: &metav1.Time{}},
				},
			},
		},
		{
			name: "udpate after started",
			oldTR: &v1.TaskRun{
				Status: v1.TaskRunStatus{
					TaskRunStatusFields: v1.TaskRunStatusFields{StartTime: &metav1.Time{}},
				},
			},
			newTR: &v1.TaskRun{
				Status: v1.TaskRunStatus{
					TaskRunStatusFields: v1.TaskRunStatusFields{StartTime: &metav1.Time{}},
				},
			},
		},
		{
			name: "completed",
			oldTR: &v1.TaskRun{
				Status: v1.TaskRunStatus{
					TaskRunStatusFields: v1.TaskRunStatusFields{StartTime: &metav1.Time{}},
				},
			},
			newTR: &v1.TaskRun{
				Status: v1.TaskRunStatus{
					Status: duckv1.Status{
						ObservedGeneration: 0,
						Conditions: duckv1.Conditions{{
							Type:   "Succeeded",
							Status: corev1.ConditionTrue,
						}},
						Annotations: nil,
					},
					TaskRunStatusFields: v1.TaskRunStatusFields{StartTime: &metav1.Time{}},
				},
			},
		},
	} {
		ev := event.UpdateEvent{
			ObjectOld: tc.oldTR,
			ObjectNew: tc.newTR,
		}
		rc := filter.Update(ev)
		filter.metric = NewTaskRunScheduledMetric()

		if rc != tc.expectedRC {
			t.Errorf(fmt.Sprintf("tc %s expected %v but got %v", tc.name, tc.expectedRC, rc))
		}
		metrics.Registry.Unregister(filter.metric)
	}

}

func TestTaskRunScheduledCollection(t *testing.T) {
	mockTaskRuns := []*v1.TaskRun{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:              "test-taskrun-1",
				Namespace:         "test-namespace",
				UID:               types.UID("test-taskrun-1"),
				CreationTimestamp: metav1.NewTime(time.Now().UTC()),
			},
			Status: v1.TaskRunStatus{
				Status: duckv1.Status{
					ObservedGeneration: 0,
					Conditions: duckv1.Conditions{{
						Type:   "Succeeded",
						Status: corev1.ConditionTrue,
					}},
					Annotations: nil,
				},
				TaskRunStatusFields: v1.TaskRunStatusFields{
					StartTime:      &metav1.Time{Time: time.Now().UTC().Add(5 * time.Second)},
					CompletionTime: &metav1.Time{Time: time.Now().UTC().Add(10 * time.Second)},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:              "test-taskrun-2",
				Namespace:         "test-namespace",
				UID:               types.UID("test-taskrun-2"),
				CreationTimestamp: metav1.NewTime(time.Now().UTC()),
			},
			Status: v1.TaskRunStatus{
				Status: duckv1.Status{
					ObservedGeneration: 0,
					Conditions: duckv1.Conditions{{
						Type:   "Succeeded",
						Status: corev1.ConditionUnknown,
					}},
					Annotations: nil,
				},
				TaskRunStatusFields: v1.TaskRunStatusFields{
					StartTime: &metav1.Time{Time: time.Now().UTC().Add(5 * time.Second)},
				},
			},
		},
	}

	for _, tr := range mockTaskRuns {
		metric := NewTaskRunScheduledMetric()
		label := prometheus.Labels{NS_LABEL: "test-namespace", TASK_NAME_LABEL: taskRef(tr.Labels), STATUS_LABEL: SUCCEEDED}
		bumpTaskRunScheduledDuration(calculateScheduledDurationTaskRun(tr), tr, metric)
		validateHistogramVec(t, metric, label, false)
		metrics.Registry.Unregister(metric)
	}
}

func TestCalculateTaskRunScheduledDuration(t *testing.T) {
	now := time.Now()
	for _, tc := range []struct {
		expectedAmt float64
		tr          *v1.TaskRun
	}{
		{
			expectedAmt: 5,
			tr: &v1.TaskRun{
				ObjectMeta: metav1.ObjectMeta{
					Name:              "test-taskrun-1",
					Namespace:         "test-namespace",
					CreationTimestamp: metav1.NewTime(now),
				},
				Status: v1.TaskRunStatus{
					Status: duckv1.Status{
						ObservedGeneration: 0,
						Conditions: duckv1.Conditions{{
							Type:   "Succeeded",
							Status: corev1.ConditionTrue,
						}},
						Annotations: nil,
					},
					TaskRunStatusFields: v1.TaskRunStatusFields{
						StartTime:      &metav1.Time{Time: now.Add(5 * time.Second)},
						CompletionTime: &metav1.Time{Time: now.Add(10 * time.Second)},
					},
				},
			},
		},
		{
			expectedAmt: 5,
			tr: &v1.TaskRun{
				ObjectMeta: metav1.ObjectMeta{
					Name:              "test-taskrun-1",
					Namespace:         "test-namespace",
					CreationTimestamp: metav1.NewTime(now),
				},
				Status: v1.TaskRunStatus{
					Status: duckv1.Status{
						ObservedGeneration: 0,
						Conditions: duckv1.Conditions{{
							Type:   "Failed",
							Status: corev1.ConditionTrue,
						}},
						Annotations: nil,
					},
					TaskRunStatusFields: v1.TaskRunStatusFields{
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
