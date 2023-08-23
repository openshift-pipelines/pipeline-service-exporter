package collector

import (
	"fmt"
	"testing"

	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

func TestStartTimeEventFilter_Update(t *testing.T) {
	filter := &startTimeEventFilter{}
	for _, tc := range []struct {
		name       string
		oldPR      *v1beta1.PipelineRun
		newPR      *v1beta1.PipelineRun
		expectedRC bool
	}{
		{
			name:  "not started",
			oldPR: &v1beta1.PipelineRun{},
			newPR: &v1beta1.PipelineRun{},
		},
		{
			name:  "just started",
			oldPR: &v1beta1.PipelineRun{},
			newPR: &v1beta1.PipelineRun{
				Status: v1beta1.PipelineRunStatus{
					PipelineRunStatusFields: v1beta1.PipelineRunStatusFields{StartTime: &metav1.Time{}},
				},
			},
			expectedRC: true,
		},
		{
			name: "udpate after started",
			oldPR: &v1beta1.PipelineRun{
				Status: v1beta1.PipelineRunStatus{
					PipelineRunStatusFields: v1beta1.PipelineRunStatusFields{StartTime: &metav1.Time{}},
				},
			},
			newPR: &v1beta1.PipelineRun{
				Status: v1beta1.PipelineRunStatus{
					PipelineRunStatusFields: v1beta1.PipelineRunStatusFields{StartTime: &metav1.Time{}},
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
	}

}

func TestTaskRunGapEventFilter_Update(t *testing.T) {
	filter := &taskRunGapEventFilter{}
	for _, tc := range []struct {
		name       string
		oldPR      *v1beta1.PipelineRun
		newPR      *v1beta1.PipelineRun
		expectedRC bool
	}{
		{
			name:  "not done no status",
			oldPR: &v1beta1.PipelineRun{},
			newPR: &v1beta1.PipelineRun{},
		},
		{
			name:  "not done status unknown",
			oldPR: &v1beta1.PipelineRun{},
			newPR: &v1beta1.PipelineRun{
				Status: v1beta1.PipelineRunStatus{
					Status: duckv1.Status{
						Conditions: []apis.Condition{
							{
								Type:   apis.ConditionSucceeded,
								Status: corev1.ConditionUnknown,
							},
						},
					},
				},
			},
		},
		{
			name:  "just done succeed",
			oldPR: &v1beta1.PipelineRun{},
			newPR: &v1beta1.PipelineRun{
				Status: v1beta1.PipelineRunStatus{
					Status: duckv1.Status{
						Conditions: []apis.Condition{
							{
								Type:   apis.ConditionSucceeded,
								Status: corev1.ConditionTrue,
							},
						},
					},
				},
			},
			expectedRC: true,
		},
		{
			name:  "just done failed",
			oldPR: &v1beta1.PipelineRun{},
			newPR: &v1beta1.PipelineRun{
				Status: v1beta1.PipelineRunStatus{
					Status: duckv1.Status{
						Conditions: []apis.Condition{
							{
								Type:   apis.ConditionSucceeded,
								Status: corev1.ConditionFalse,
							},
						},
					},
				},
			},
			expectedRC: true,
		},
		{
			name: "udpate after done",
			oldPR: &v1beta1.PipelineRun{
				Status: v1beta1.PipelineRunStatus{
					Status: duckv1.Status{
						Conditions: []apis.Condition{
							{
								Type:   apis.ConditionSucceeded,
								Status: corev1.ConditionFalse,
							},
						},
					},
				},
			},
			newPR: &v1beta1.PipelineRun{
				Status: v1beta1.PipelineRunStatus{
					Status: duckv1.Status{
						Conditions: []apis.Condition{
							{
								Type:   apis.ConditionSucceeded,
								Status: corev1.ConditionFalse,
							},
						},
					},
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
	}
}
