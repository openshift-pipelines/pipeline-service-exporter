package collector

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	"github.com/tektoncd/pipeline/pkg/pod"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"testing"
)

func TestResetPodCreateAttemptedStats(t *testing.T) {
	objs := []client.Object{}
	scheme := runtime.NewScheme()
	_ = v1.AddToScheme(scheme)
	c := fake.NewClientBuilder().WithScheme(scheme).WithObjects(objs...).Build()

	os.Setenv(PodCreateFilterEnvName, "test-namespace-2")

	mockTaskRuns := []*v1.TaskRun{
		{
			ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace", Name: "test-1"},
			Status: v1.TaskRunStatus{
				TaskRunStatusFields: v1.TaskRunStatusFields{
					PodName: "test-1-pod",
				},
			},
		},
		// should bump the counter
		{
			ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace", Name: "test-2"},
			Status:     v1.TaskRunStatus{},
		},
		// should bump the counter
		{
			ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace", Name: "test-3"},
			Status: v1.TaskRunStatus{
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
			Status: v1.TaskRunStatus{
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
			Status: v1.TaskRunStatus{
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
			Status: v1.TaskRunStatus{
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
		// should not bump the counter because of the filter
		{
			ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace", Name: "test-7", Annotations: map[string]string{pod.ReleaseAnnotation: "foo"}},
			Status:     v1.TaskRunStatus{},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace-2", Name: "test-3"},
			Status: v1.TaskRunStatus{
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
	}
	ctx := context.TODO()
	for _, pr := range mockTaskRuns {
		err := c.Create(ctx, pr)
		assert.NoError(t, err)
	}

	reconciler := buildReconciler(c, nil, nil)
	// initiate first scan
	reconciler.resetPodCreateAttemptedStats(ctx)
	// initiate second scan (where repeats mean bump the metric)
	reconciler.resetPodCreateAttemptedStats(ctx)
	label := prometheus.Labels{NS_LABEL: "test-namespace"}
	validateGaugeVec(t, reconciler.waitPodCollector.waitPodCreate, label, float64(2))
	// third pass should reset and still be two
	reconciler.resetPodCreateAttemptedStats(ctx)
	validateGaugeVec(t, reconciler.waitPodCollector.waitPodCreate, label, float64(2))
	// deletion, then another pass, should now be one
	err := c.Delete(ctx, mockTaskRuns[1])
	assert.NoError(t, err)
	reconciler.resetPodCreateAttemptedStats(ctx)
	validateGaugeVec(t, reconciler.waitPodCollector.waitPodCreate, label, float64(1))
	// change the last remaining one so it passes, should now be zero
	mockTaskRuns[2].ObjectMeta.Annotations = map[string]string{pod.ReleaseAnnotation: "foo"}
	err = c.Update(ctx, mockTaskRuns[2])
	assert.NoError(t, err)
	reconciler.resetPodCreateAttemptedStats(ctx)
	validateGaugeVec(t, reconciler.waitPodCollector.waitPodCreate, label, float64(0))
	unregisterStats(reconciler)
}
