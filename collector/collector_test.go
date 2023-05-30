package collector

import (
	"context"
	"testing"
	"time"

	"github.com/go-kit/log"
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
)

func TestNewCollector(t *testing.T) {
	logger := log.NewNopLogger()
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
	for _, pr := range mockPipelineRuns {
		err := c.Create(context.TODO(), pr)
		assert.NoError(t, err)
	}

	collector, err := NewCollector(logger, c)
	assert.NoError(t, err)
	assert.NotNil(t, collector)
	assert.NotNil(t, collector.durationScheduled)
	assert.NotNil(t, collector.durationCompleted)
	metricReceived := false
	ch := make(chan prometheus.Metric)
	go func() {
		for {
			select {
			case m := <-ch:
				if m != nil {
					t.Logf("metric received: %#v\n", m)
					metricReceived = true
				}
			}

		}
	}()
	collector.collect(ch)
	label := prometheus.Labels{"name": "test-pipelinerun-1", "uid": "test-pipelinerun-1"}
	g, e := collector.durationScheduled.GetMetricWith(label)
	assert.NoError(t, e)
	assert.NotNil(t, g)
	label = prometheus.Labels{"name": "test-pipelinerun-2", "uid": "test-pipelinerun-2"}
	g, e = collector.durationCompleted.GetMetricWith(label)
	assert.NoError(t, e)
	assert.NotNil(t, g)
	assert.True(t, metricReceived)
	close(ch)
}
