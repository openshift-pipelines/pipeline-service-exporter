package collector

import (
	"context"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func TestMetricCollection(t *testing.T) {
	objs := []client.Object{}
	scheme := runtime.NewScheme()
	_ = v1beta1.AddToScheme(scheme)
	c := fake.NewClientBuilder().WithScheme(scheme).WithObjects(objs...).Build()

	mockPipelineRuns := []*v1beta1.PipelineRun{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:              "test-pipelinerun-1",
				Namespace:         "test-namespace",
				UID:               types.UID("test-pipelinerun-1"),
				CreationTimestamp: metav1.NewTime(time.Now().UTC()),
			},
			Status: v1beta1.PipelineRunStatus{
				Status: duckv1.Status{
					ObservedGeneration: 0,
					Conditions: duckv1.Conditions{{
						Type:   "Succeeded",
						Status: corev1.ConditionTrue,
					}},
					Annotations: nil,
				},
				PipelineRunStatusFields: v1beta1.PipelineRunStatusFields{
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
			Status: v1beta1.PipelineRunStatus{
				Status: duckv1.Status{
					ObservedGeneration: 0,
					Conditions: duckv1.Conditions{{
						Type:   "Succeeded",
						Status: corev1.ConditionTrue,
					}},
					Annotations: nil,
				},
				PipelineRunStatusFields: v1beta1.PipelineRunStatusFields{
					StartTime:      &metav1.Time{Time: time.Now().UTC().Add(5 * time.Second)},
					CompletionTime: &metav1.Time{Time: time.Now().UTC().Add(10 * time.Second)},
				},
			},
		},
	}
	reconciler := &ReconcilePipelineRun{client: c, psCollector: NewCollector()}
	ctx := context.TODO()
	for _, pr := range mockPipelineRuns {
		err := c.Create(context.TODO(), pr)
		assert.NoError(t, err)
		request := reconcile.Request{
			NamespacedName: types.NamespacedName{
				Namespace: pr.Namespace,
				Name:      pr.Name,
			},
		}
		_, err = reconciler.Reconcile(ctx, request)
	}

	label := prometheus.Labels{"namespace": "test-namespace"}
	g, e := reconciler.psCollector.durationScheduled.GetMetricWith(label)
	assert.NoError(t, e)
	assert.NotNil(t, g)
}
