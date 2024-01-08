package collector

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"testing"
)

func TestReconcileOverhead_Reconcile(t *testing.T) {
	// rather than using golang mocks, grabbed actual RHTAP pipelinerun/taskruns from staging
	// to drive the gap metric, given its trickiness
	objs := []client.Object{}
	scheme := runtime.NewScheme()
	_ = v1.AddToScheme(scheme)
	_ = v1beta1.AddToScheme(scheme)
	c := fake.NewClientBuilder().WithScheme(scheme).WithObjects(objs...).Build()
	overheadReconciler := &ReconcileOverhead{
		client:    c,
		collector: NewOverheadCollector(),
	}
	var err error
	// first we test with the samples pulled from actual RHTAP yaml to best capture the parallel task executions
	prs := []v1beta1.PipelineRun{}
	trs := []v1beta1.TaskRun{}
	prs, err = pipelineRunFromActualRHTAPYaml()
	if err != nil {
		t.Fatalf(fmt.Sprintf("%s", err.Error()))
	}
	trs, err = taskRunsFromActualRHTAPYaml()
	if err != nil {
		t.Fatalf(fmt.Sprintf("%s", err.Error()))
	}

	ctx := context.TODO()
	for _, trv1beta1 := range trs {
		// mimic what the tekton conversion webhook will do
		tr := &v1.TaskRun{}
		err = trv1beta1.ConvertTo(ctx, tr)
		assert.NoError(t, err)
		err = c.Create(ctx, tr)
		assert.NoError(t, err)
	}
	for index, prv1beta1 := range prs {
		// mimic what the tekton conversion webhook will do
		pr := &v1.PipelineRun{}
		err = prv1beta1.ConvertTo(ctx, pr)
		assert.NoError(t, err)
		err = c.Create(ctx, pr)
		assert.NoError(t, err)
		request := reconcile.Request{
			NamespacedName: types.NamespacedName{
				Namespace: pr.Namespace,
				Name:      pr.Name,
			},
		}
		_, err = overheadReconciler.Reconcile(ctx, request)
		label := prometheus.Labels{NS_LABEL: pr.Namespace, STATUS_LABEL: SUCCEEDED}
		// with our actual RHTAP samples the first entry had 0 scheduling overhead so we created a metric,
		// but the rest was filtered
		var observer prometheus.Observer
		var histogram prometheus.Histogram
		var metric *dto.Metric
		if index == 0 {
			validateHistogramVec(t, overheadReconciler.collector.scheduling, label, true)
		} else {
			observer, err = overheadReconciler.collector.scheduling.GetMetricWith(label)
			assert.NoError(t, err)
			assert.NotNil(t, observer)
			histogram = observer.(prometheus.Histogram)
			metric = &dto.Metric{}
			histogram.Write(metric)
			assert.NotNil(t, metric.Histogram)
			assert.NotNil(t, metric.Histogram.SampleCount)
			assert.Equal(t, *metric.Histogram.SampleCount, uint64(0))
		}
		observer, err = overheadReconciler.collector.execution.GetMetricWith(label)
		assert.NoError(t, err)
		assert.NotNil(t, observer)
		histogram = observer.(prometheus.Histogram)
		metric = &dto.Metric{}
		histogram.Write(metric)
		assert.NotNil(t, metric.Histogram)
		assert.NotNil(t, metric.Histogram.SampleCount)
		assert.Equal(t, *metric.Histogram.SampleCount, uint64(0))

	}
	metrics.Registry.Unregister(overheadReconciler.collector.execution)
	metrics.Registry.Unregister(overheadReconciler.collector.scheduling)

}

func TestReconcileOverhead_Reconcile_MissingTaskRuns(t *testing.T) {
	// rather than using golang mocks, grabbed actual RHTAP pipelinerun/taskruns from staging
	// to drive the gap metric, given its trickiness
	objs := []client.Object{}
	scheme := runtime.NewScheme()
	_ = v1.AddToScheme(scheme)
	_ = v1beta1.AddToScheme(scheme)
	c := fake.NewClientBuilder().WithScheme(scheme).WithObjects(objs...).Build()
	overheadReconciler := &ReconcileOverhead{
		client:    c,
		collector: NewOverheadCollector(),
	}
	var err error
	// first we test with the samples pulled from actual RHTAP yaml to best capture the parallel task executions
	prs := []v1beta1.PipelineRun{}
	prs, err = pipelineRunFromActualRHTAPYaml()
	if err != nil {
		t.Fatalf(fmt.Sprintf("%s", err.Error()))
	}
	// but in this test we make sure no stats are generated if the taskruns are missing

	ctx := context.TODO()
	for _, prv1beta1 := range prs {
		// mimic what the tekton conversion webhook will do
		pr := &v1.PipelineRun{}
		err = prv1beta1.ConvertTo(ctx, pr)
		assert.NoError(t, err)
		err = c.Create(ctx, pr)
		assert.NoError(t, err)
		request := reconcile.Request{
			NamespacedName: types.NamespacedName{
				Namespace: pr.Namespace,
				Name:      pr.Name,
			},
		}
		_, err = overheadReconciler.Reconcile(ctx, request)
		label := prometheus.Labels{NS_LABEL: pr.Namespace, STATUS_LABEL: SUCCEEDED}
		validateHistogramVecZeroCount(t, overheadReconciler.collector.execution, label)
	}
	metrics.Registry.Unregister(overheadReconciler.collector.execution)
	metrics.Registry.Unregister(overheadReconciler.collector.scheduling)

}
