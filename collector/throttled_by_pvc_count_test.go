package collector

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"github.com/tektoncd/pipeline/pkg/reconciler/volumeclaim"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sync"
	"testing"
)

func TestPvcThrottledFilter_Update(t *testing.T) {
	filter := &pvcThrottledFilter{}
	for _, tc := range []struct {
		name       string
		oldPR      *v1beta1.PipelineRun
		newPR      *v1beta1.PipelineRun
		expectedRC bool
	}{
		{
			name:  "not failed",
			oldPR: &v1beta1.PipelineRun{},
			newPR: &v1beta1.PipelineRun{},
		},
		{
			name:  "failed with right reason and message",
			oldPR: &v1beta1.PipelineRun{},
			newPR: &v1beta1.PipelineRun{
				Status: v1beta1.PipelineRunStatus{
					Status: duckv1.Status{
						Conditions: duckv1.Conditions{
							apis.Condition{
								Type:    apis.ConditionSucceeded,
								Status:  corev1.ConditionFalse,
								Reason:  volumeclaim.ReasonCouldntCreateWorkspacePVC,
								Message: "exceeded quota",
							},
						},
					},
				},
			},
			expectedRC: true,
		},
		{
			name:  "failed with right reason but wrong message",
			oldPR: &v1beta1.PipelineRun{},
			newPR: &v1beta1.PipelineRun{
				Status: v1beta1.PipelineRunStatus{
					Status: duckv1.Status{
						Conditions: duckv1.Conditions{
							apis.Condition{
								Type:    apis.ConditionSucceeded,
								Status:  corev1.ConditionFalse,
								Reason:  volumeclaim.ReasonCouldntCreateWorkspacePVC,
								Message: "api server unavailable",
							},
						},
					},
				},
			},
		},
		{
			name:  "failed with right message but wrong reason",
			oldPR: &v1beta1.PipelineRun{},
			newPR: &v1beta1.PipelineRun{
				Status: v1beta1.PipelineRunStatus{
					Status: duckv1.Status{
						Conditions: duckv1.Conditions{
							apis.Condition{
								Type:    apis.ConditionSucceeded,
								Status:  corev1.ConditionFalse,
								Reason:  corev1.PodReasonUnschedulable,
								Message: "exceeded quota",
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

func TestResetPVCStats(t *testing.T) {
	objs := []client.Object{}
	scheme := runtime.NewScheme()
	_ = v1beta1.AddToScheme(scheme)
	c := fake.NewClientBuilder().WithScheme(scheme).WithObjects(objs...).Build()

	mockPipelineRuns := []*v1beta1.PipelineRun{
		{
			ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace", Name: "test-1"},
			Status: v1beta1.PipelineRunStatus{
				Status: duckv1.Status{
					Conditions: duckv1.Conditions{
						apis.Condition{
							Type:    apis.ConditionSucceeded,
							Status:  corev1.ConditionFalse,
							Reason:  volumeclaim.ReasonCouldntCreateWorkspacePVC,
							Message: "exceeded quota",
						},
					},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace", Name: "test-2"},
			Status: v1beta1.PipelineRunStatus{
				Status: duckv1.Status{
					Conditions: duckv1.Conditions{
						apis.Condition{
							Type:    apis.ConditionSucceeded,
							Status:  corev1.ConditionFalse,
							Reason:  volumeclaim.ReasonCouldntCreateWorkspacePVC,
							Message: "exceeded quota",
						},
					},
				},
			},
		},
	}
	ctx := context.TODO()
	for _, pr := range mockPipelineRuns {
		err := c.Create(ctx, pr)
		assert.NoError(t, err)
	}

	pvcReconciler := &ReconcilePVCThrottled{client: c, prCollector: NewPVCThrottledCollector(), listMutex: sync.RWMutex{}, nsCache: map[string]struct{}{}}
	pvcReconciler.resetPVCStats(ctx)
	label := prometheus.Labels{NS_LABEL: "test-namespace"}
	validateGaugeVec(t, pvcReconciler.prCollector.pvcThrottle, label, float64(2))
	// second pass should reset and still be two
	pvcReconciler.resetPVCStats(ctx)
	validateGaugeVec(t, pvcReconciler.prCollector.pvcThrottle, label, float64(2))
	// deletion, then another pass, should now be one
	err := c.Delete(ctx, mockPipelineRuns[0])
	assert.NoError(t, err)
	pvcReconciler.resetPVCStats(ctx)
	validateGaugeVec(t, pvcReconciler.prCollector.pvcThrottle, label, float64(1))
}
