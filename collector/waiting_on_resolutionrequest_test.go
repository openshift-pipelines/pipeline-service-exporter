package collector

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/tektoncd/pipeline/pkg/apis/resolution/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestResolutionRequestsStats(t *testing.T) {
	objs := []client.Object{}
	scheme := runtime.NewScheme()
	_ = v1beta1.AddToScheme(scheme)
	c := fake.NewClientBuilder().WithScheme(scheme).WithObjects(objs...).Build()

	mockResolutionRequests := []*v1beta1.ResolutionRequest{
		// should bump the counter
		{
			ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace", Name: "test-1"},
			Status:     v1beta1.ResolutionRequestStatus{},
		},
		// should bump the counter
		{
			ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace", Name: "test-2"},
			Status:     v1beta1.ResolutionRequestStatus{},
		},
		// should not bump the counter
		{
			ObjectMeta: metav1.ObjectMeta{Namespace: "test-namespace", Name: "test-3"},
			Status: v1beta1.ResolutionRequestStatus{
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
	}

	ctx := context.TODO()
	for _, rr := range mockResolutionRequests {
		err := c.Create(ctx, rr)
		assert.NoError(t, err)
	}

	reconciler := buildReconciler(c, nil, nil)
	// initiate first scan
	reconciler.resetResoultionRequestsStats(ctx)
	// initiate second scan (where repeats mean bump the metric)
	reconciler.resetResoultionRequestsStats(ctx)
	label := prometheus.Labels{NS_LABEL: "test-namespace"}
	validateGaugeVec(t, reconciler.waitRRCollector.waitingResolutionRequest, label, float64(2))
	// third pass should reset and still be one
	reconciler.resetResoultionRequestsStats(ctx)
	validateGaugeVec(t, reconciler.waitRRCollector.waitingResolutionRequest, label, float64(2))
	// deletion, then another pass, should now be zero
	err := c.Delete(ctx, mockResolutionRequests[0])
	assert.NoError(t, err)
	// third pass should reset and still be one
	reconciler.resetResoultionRequestsStats(ctx)
	validateGaugeVec(t, reconciler.waitRRCollector.waitingResolutionRequest, label, float64(1))

}
