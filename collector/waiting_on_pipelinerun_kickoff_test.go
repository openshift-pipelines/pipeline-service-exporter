package collector

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"testing"
	"time"
)

var testTimeout = &metav1.Duration{time.Nanosecond}

func TestPipelineRunKickoffStats(t *testing.T) {
	objs := []client.Object{}
	scheme := runtime.NewScheme()
	_ = v1.AddToScheme(scheme)
	c := fake.NewClientBuilder().WithScheme(scheme).WithObjects(objs...).Build()

	os.Setenv(PipelineRunKickoffFilterEnvName, "test-namespace-2")

	mockPipelineRuns := []*v1.PipelineRun{
		{
			ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace", Name: "test-1"},
			Spec: v1.PipelineRunSpec{
				Status: v1.PipelineRunSpecStatusCancelled,
			},
			Status: v1.PipelineRunStatus{
				PipelineRunStatusFields: v1.PipelineRunStatusFields{},
			},
		},
		// should bump the counter
		{
			ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace", Name: "test-2"},
			Status:     v1.PipelineRunStatus{},
		},
		// should bump the counter
		{
			ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace", Name: "test-3"},
			Status: v1.PipelineRunStatus{
				Status: duckv1.Status{
					Conditions: duckv1.Conditions{
						apis.Condition{
							Type:   apis.ConditionSucceeded,
							Status: corev1.ConditionUnknown,
						},
					},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace", Name: "test-4"},
			Status: v1.PipelineRunStatus{
				Status: duckv1.Status{
					Conditions: duckv1.Conditions{
						apis.Condition{
							Type:    apis.ConditionSucceeded,
							Status:  corev1.ConditionUnknown,
							Reason:  "foo",
							Message: "bar",
						},
					},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace", Name: "test-5"},
			Status: v1.PipelineRunStatus{
				Status: duckv1.Status{
					Conditions: duckv1.Conditions{
						apis.Condition{
							Type:   apis.ConditionSucceeded,
							Status: corev1.ConditionFalse,
						},
					},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace", Name: "test-6"},
			Status: v1.PipelineRunStatus{
				Status: duckv1.Status{
					Conditions: duckv1.Conditions{
						apis.Condition{
							Type:   apis.ConditionSucceeded,
							Status: corev1.ConditionTrue,
						},
					},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace-2", Name: "test-3"},
			Status: v1.PipelineRunStatus{
				Status: duckv1.Status{
					Conditions: duckv1.Conditions{
						apis.Condition{
							Type:   apis.ConditionSucceeded,
							Status: corev1.ConditionUnknown,
						},
					},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace", Name: "test-20"},
			Spec: v1.PipelineRunSpec{
				Status: v1.PipelineRunSpecStatusCancelledRunFinally,
			},
			Status: v1.PipelineRunStatus{
				PipelineRunStatusFields: v1.PipelineRunStatusFields{},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace", Name: "test-21"},
			Spec: v1.PipelineRunSpec{
				Status: v1.PipelineRunSpecStatusCancelledRunFinally,
			},
			Status: v1.PipelineRunStatus{
				PipelineRunStatusFields: v1.PipelineRunStatusFields{},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace", Name: "test-22"},
			Spec: v1.PipelineRunSpec{
				Status: v1.PipelineRunSpecStatusStoppedRunFinally,
			},
			Status: v1.PipelineRunStatus{
				PipelineRunStatusFields: v1.PipelineRunStatusFields{},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace", Name: "test-23"},
			Spec: v1.PipelineRunSpec{
				Status: v1.PipelineRunSpecStatusPending,
			},
			Status: v1.PipelineRunStatus{
				PipelineRunStatusFields: v1.PipelineRunStatusFields{},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace", Name: "test-24"},
			Status: v1.PipelineRunStatus{
				PipelineRunStatusFields: v1.PipelineRunStatusFields{
					StartTime: &metav1.Time{time.Now()},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace", Name: "test-25"},
			Spec: v1.PipelineRunSpec{
				Timeouts: &v1.TimeoutFields{Pipeline: testTimeout},
			},
			Status: v1.PipelineRunStatus{
				PipelineRunStatusFields: v1.PipelineRunStatusFields{
					StartTime: &metav1.Time{time.Now().Add(5 * time.Second)},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace", Name: "test-26"},
			Status: v1.PipelineRunStatus{
				PipelineRunStatusFields: v1.PipelineRunStatusFields{
					ChildReferences: []v1.ChildStatusReference{
						{
							Name: "child-1",
						},
					},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace", Name: "test-27"},
			Status: v1.PipelineRunStatus{
				PipelineRunStatusFields: v1.PipelineRunStatusFields{
					SkippedTasks: []v1.SkippedTask{
						{
							Name: "child-1",
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

	reconciler := buildReconciler(c, nil, nil)
	// initiate first scan
	reconciler.resetPipelineRunKickoffStats(ctx)
	// initiate second scan (where repeats mean bump the metric)
	reconciler.resetPipelineRunKickoffStats(ctx)
	label := prometheus.Labels{NS_LABEL: "test-namespace"}
	validateGaugeVec(t, reconciler.waitPRKickoffCollector.waitPipelineRunKickoff, label, float64(1))
	// third pass should reset and still be one
	reconciler.resetPipelineRunKickoffStats(ctx)
	validateGaugeVec(t, reconciler.waitPRKickoffCollector.waitPipelineRunKickoff, label, float64(1))
	// deletion, then another pass, should now be zero
	err := c.Delete(ctx, mockPipelineRuns[1])
	assert.NoError(t, err)
	reconciler.resetPipelineRunKickoffStats(ctx)
	validateGaugeVec(t, reconciler.waitPRKickoffCollector.waitPipelineRunKickoff, label, float64(0))
	// change the last remaining one so it passes, should still be zero
	mockPipelineRuns[2].Status.StartTime = &metav1.Time{time.Now()}
	err = c.Update(ctx, mockPipelineRuns[2])
	assert.NoError(t, err)
	reconciler.resetPipelineRunKickoffStats(ctx)
	validateGaugeVec(t, reconciler.waitPRKickoffCollector.waitPipelineRunKickoff, label, float64(0))
	unregisterStats(reconciler)
}
