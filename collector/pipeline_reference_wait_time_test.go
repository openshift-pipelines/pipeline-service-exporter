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

func TestPipelineRefWaitTimeFilter_Update(t *testing.T) {
	filter := &pipelineRefWaitTimeFilter{waitDuration: NewPipelineReferenceWaitTimeMetric()}
	now := time.Now()
	for _, tc := range []struct {
		name                  string
		oldPR                 *v1.PipelineRun
		newPR                 *v1.PipelineRun
		expectedRC            bool
		expectedNonZeroMetric bool
	}{
		{
			name:  "not started",
			oldPR: &v1.PipelineRun{},
			newPR: &v1.PipelineRun{},
		},
		{
			name:  "done",
			oldPR: &v1.PipelineRun{},
			newPR: &v1.PipelineRun{
				Status: v1.PipelineRunStatus{
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
			oldPR: &v1.PipelineRun{
				Status: v1.PipelineRunStatus{
					Status: duckv1.Status{Conditions: duckv1.Conditions{
						{
							Type:               apis.ConditionSucceeded,
							Status:             corev1.ConditionUnknown,
							Reason:             "ResolvingPipelineRef", // waiting for tag/release with constant moved to the api package
							LastTransitionTime: apis.VolatileTime{Inner: metav1.NewTime(now)},
						},
					}},
				},
			},
			newPR: &v1.PipelineRun{
				Status: v1.PipelineRunStatus{
					Status: duckv1.Status{Conditions: duckv1.Conditions{
						{
							Type:               apis.ConditionSucceeded,
							Status:             corev1.ConditionUnknown,
							Reason:             "ResolvingPipelineRef", // waiting for tag/release with constant moved to the api package
							LastTransitionTime: apis.VolatileTime{Inner: metav1.NewTime(now)},
						},
					}},
				},
			},
		},
		{
			name: "both running, diff transition time",
			oldPR: &v1.PipelineRun{
				Status: v1.PipelineRunStatus{
					Status: duckv1.Status{Conditions: duckv1.Conditions{
						{
							Type:               apis.ConditionSucceeded,
							Status:             corev1.ConditionUnknown,
							Reason:             "ResolvingPipelineRef", // waiting for tag/release with constant moved to the api package
							LastTransitionTime: apis.VolatileTime{Inner: metav1.NewTime(now)},
						},
					}},
				},
			},
			newPR: &v1.PipelineRun{
				Status: v1.PipelineRunStatus{
					Status: duckv1.Status{Conditions: duckv1.Conditions{
						{
							Type:               apis.ConditionSucceeded,
							Status:             corev1.ConditionUnknown,
							Reason:             "ResolvingPipelineRef", // waiting for tag/release with constant moved to the api package
							LastTransitionTime: apis.VolatileTime{Inner: metav1.NewTime(now.Add(1 * time.Second))},
						},
					}},
				},
			},
		},
		{
			name:                  "wait over",
			expectedNonZeroMetric: true,
			oldPR: &v1.PipelineRun{
				Status: v1.PipelineRunStatus{
					Status: duckv1.Status{Conditions: duckv1.Conditions{
						{
							Type:               apis.ConditionSucceeded,
							Status:             corev1.ConditionUnknown,
							Reason:             "ResolvingPipelineRef", // waiting for tag/release with constant moved to the api package
							LastTransitionTime: apis.VolatileTime{Inner: metav1.NewTime(now)},
						},
					}},
				},
			},
			newPR: &v1.PipelineRun{
				Status: v1.PipelineRunStatus{
					Status: duckv1.Status{Conditions: duckv1.Conditions{
						{
							Type:               apis.ConditionSucceeded,
							Status:             corev1.ConditionUnknown,
							Reason:             "ResolvingPipelineRef", // waiting for tag/release with constant moved to the api package
							LastTransitionTime: apis.VolatileTime{Inner: metav1.NewTime(now.Add(1 * time.Second))},
						},
					}},
				},
			},
		},
	} {
		ev := event.UpdateEvent{
			ObjectOld: tc.oldPR,
			ObjectNew: tc.newPR,
		}
		rc := filter.Update(ev)
		if rc != tc.expectedRC {
			t.Errorf(fmt.Sprintf("tc %s expected %v but got %v", tc.name, tc.expectedRC, rc))
		}
		labels := prometheus.Labels{NS_LABEL: tc.newPR.Namespace, PIPELINE_NAME_LABEL: pipelineRef(tc.newPR.Labels)}
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
			if tc.newPR.IsDone() {
				assert.NotZero(t, metric.Histogram.SampleCount)
			}
		}
	}
}
