package collector

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"testing"
	"time"
)

func TestTaskRefWaitTimeFilter_Update(t *testing.T) {
	filter := &taskRefWaitTimeFilter{waitDuration: NewTaskReferenceWaitTimeMetric()}
	now := time.Now()
	for _, tc := range []struct {
		name                  string
		oldTR                 *v1.TaskRun
		newTR                 *v1.TaskRun
		expectedRC            bool
		expectedNonZeroMetric bool
	}{
		{
			name:  "not started",
			oldTR: &v1.TaskRun{},
			newTR: &v1.TaskRun{},
		},
		{
			name:  "done",
			oldTR: &v1.TaskRun{},
			newTR: &v1.TaskRun{
				Status: v1.TaskRunStatus{
					Status: duckv1.Status{Conditions: duckv1.Conditions{
						{
							Type:   apis.ConditionSucceeded,
							Status: corev1.ConditionTrue,
						},
					}},
				},
			},
		},
		{
			name: "both running, same transition time",
			oldTR: &v1.TaskRun{
				Status: v1.TaskRunStatus{
					Status: duckv1.Status{Conditions: duckv1.Conditions{
						{
							Type:               apis.ConditionSucceeded,
							Status:             corev1.ConditionUnknown,
							Reason:             v1.TaskRunReasonResolvingTaskRef,
							LastTransitionTime: apis.VolatileTime{Inner: metav1.NewTime(now)},
						},
					}},
				},
			},
			newTR: &v1.TaskRun{
				Status: v1.TaskRunStatus{
					Status: duckv1.Status{Conditions: duckv1.Conditions{
						{
							Type:               apis.ConditionSucceeded,
							Status:             corev1.ConditionUnknown,
							Reason:             v1.TaskRunReasonResolvingTaskRef,
							LastTransitionTime: apis.VolatileTime{Inner: metav1.NewTime(now)},
						},
					}},
				},
			},
		},
		{
			name: "both running, diff transition time",
			oldTR: &v1.TaskRun{
				Status: v1.TaskRunStatus{
					Status: duckv1.Status{Conditions: duckv1.Conditions{
						{
							Type:               apis.ConditionSucceeded,
							Status:             corev1.ConditionUnknown,
							Reason:             v1.TaskRunReasonResolvingTaskRef,
							LastTransitionTime: apis.VolatileTime{Inner: metav1.NewTime(now)},
						},
					}},
				},
			},
			newTR: &v1.TaskRun{
				Status: v1.TaskRunStatus{
					Status: duckv1.Status{Conditions: duckv1.Conditions{
						{
							Type:               apis.ConditionSucceeded,
							Status:             corev1.ConditionUnknown,
							Reason:             v1.TaskRunReasonResolvingTaskRef,
							LastTransitionTime: apis.VolatileTime{Inner: metav1.NewTime(now.Add(1 * time.Second))},
						},
					}},
				},
			},
		},
		{
			name:                  "wait over",
			expectedNonZeroMetric: true,
			oldTR: &v1.TaskRun{
				Status: v1.TaskRunStatus{
					Status: duckv1.Status{Conditions: duckv1.Conditions{
						{
							Type:               apis.ConditionSucceeded,
							Status:             corev1.ConditionUnknown,
							Reason:             v1.TaskRunReasonResolvingTaskRef,
							LastTransitionTime: apis.VolatileTime{Inner: metav1.NewTime(now)},
						},
					}},
				},
			},
			newTR: &v1.TaskRun{
				Status: v1.TaskRunStatus{
					Status: duckv1.Status{Conditions: duckv1.Conditions{
						{
							Type:               apis.ConditionSucceeded,
							Status:             corev1.ConditionUnknown,
							Reason:             v1.TaskRunReasonRunning.String(),
							LastTransitionTime: apis.VolatileTime{Inner: metav1.NewTime(now.Add(1 * time.Second))},
						},
					}},
				},
			},
		},
	} {
		ev := event.UpdateEvent{
			ObjectOld: tc.oldTR,
			ObjectNew: tc.newTR,
		}
		rc := filter.Update(ev)
		if rc != tc.expectedRC {
			t.Errorf(fmt.Sprintf("tc %s expected %v but got %v", tc.name, tc.expectedRC, rc))
		}
		labels := prometheus.Labels{NS_LABEL: tc.newTR.Namespace}
		if tc.expectedNonZeroMetric {
			validateHistogramVec(t, filter.waitDuration, labels, false)
		} else {
			observer, err := filter.waitDuration.GetMetricWith(labels)
			assert.NoError(t, err)
			assert.NotNil(t, observer)
			histogram := observer.(prometheus.Histogram)
			metric := &dto.Metric{}
			histogram.Write(metric)
			assert.NotNil(t, metric.Histogram)
			assert.NotNil(t, metric.Histogram.SampleCount)
			if tc.newTR.IsDone() {
				assert.NotZero(t, metric.Histogram.SampleCount)
			}
		}
	}
}
