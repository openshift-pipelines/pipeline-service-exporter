package collector

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	v1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"knative.dev/pkg/apis"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func SetupOverheadController(mgr ctrl.Manager) error {
	reconciler := &ReconcileOverhead{
		client:        mgr.GetClient(),
		scheme:        mgr.GetScheme(),
		eventRecorder: mgr.GetEventRecorderFor("MetricExporterExecutionOverhead"),
		collector:     NewOverheadCollector(),
	}
	return ctrl.NewControllerManagedBy(mgr).For(&v1.PipelineRun{}).WithEventFilter(&taskRunGapEventFilter{}).Complete(reconciler)
}

type OverheadCollector struct {
	execution  *prometheus.HistogramVec
	scheduling *prometheus.HistogramVec
}

type ReconcileOverhead struct {
	client        client.Client
	scheme        *runtime.Scheme
	eventRecorder record.EventRecorder
	collector     *OverheadCollector
}

func NewOverheadCollector() *OverheadCollector {
	labelNames := []string{NS_LABEL, STATUS_LABEL}
	executionMetric := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "pipeline_service_execution_overhead_percentage",
		Help:    "Proportion of time elapsed between the completion of a TaskRun and the start of the next TaskRun within a PipelineRun to the total duration of successful PipelineRuns",
		Buckets: prometheus.DefBuckets,
	}, labelNames)
	schedulingMetric := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "pipeline_service_schedule_overhead_percentage",
		Help:    "Proportion of time elapsed waiting for the pipeline controller to receive create events compared to the total duration of successful PipelineRuns",
		Buckets: prometheus.DefBuckets,
	}, labelNames)
	collector := &OverheadCollector{execution: executionMetric, scheduling: schedulingMetric}
	metrics.Registry.MustRegister(executionMetric, schedulingMetric)
	return collector
}

func (r *ReconcileOverhead) accumulateGaps(pr *v1.PipelineRun, oc client.Client, ctx context.Context) (float64, bool) {
	if prNotDoneOrHasNoKids(pr) {
		return float64(0), false
	}
	gapTotal := float64(0)

	sortedTaskRunsByCreateTimes, reverseOrderSortedTaskRunsByCompletionTimes := sortTaskRunsForGapCalculations(pr, oc, ctx)

	gapEntries := calculateGaps(ctx, pr, oc, sortedTaskRunsByCreateTimes, reverseOrderSortedTaskRunsByCompletionTimes)
	for _, gapEntry := range gapEntries {
		gapTotal = gapTotal + gapEntry.gap
	}

	return gapTotal, true
}

func (r *ReconcileOverhead) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()
	log := log.FromContext(ctx)

	pr := &v1.PipelineRun{}
	err := r.client.Get(ctx, types.NamespacedName{Namespace: request.Namespace, Name: request.Name}, pr)
	if err != nil && !errors.IsNotFound(err) {
		return reconcile.Result{}, err
	}
	if err != nil {
		log.V(4).Info(fmt.Sprintf("ignoring deleted pipelinerun %q", request.NamespacedName))
		return reconcile.Result{}, nil
	}
	succeedCondition := pr.Status.GetCondition(apis.ConditionSucceeded)
	if succeedCondition != nil && !succeedCondition.IsUnknown() {
		gapTotal, foundGaps := r.accumulateGaps(pr, r.client, ctx)
		status := SUCCEEDED
		if succeedCondition.IsFalse() {
			status = FAILED
		}
		labels := map[string]string{NS_LABEL: pr.Namespace, STATUS_LABEL: status}
		totalDuration := float64(pr.Status.CompletionTime.Time.Sub(pr.Status.StartTime.Time).Milliseconds())
		if foundGaps {
			if !filter(gapTotal, totalDuration) {
				overhead := gapTotal / totalDuration
				r.collector.execution.With(labels).Observe(overhead)
			} else {
				log.V(4).Info(fmt.Sprintf("filtering execution metric for %s with gap %v and total %v",
					request.NamespacedName.String(), gapTotal, totalDuration))
			}
		}
		scheduleDuration := calculateScheduledDuration(pr.CreationTimestamp.Time, pr.Status.StartTime.Time)
		if !filter(scheduleDuration, totalDuration) {
			overhead := scheduleDuration / totalDuration
			r.collector.scheduling.With(labels).Observe(overhead)
		} else {
			log.V(4).Info(fmt.Sprintf("filtering scheduling metric for %s with gap %v and total %v",
				request.NamespacedName.String(), scheduleDuration, totalDuration))
		}
	}
	return reconcile.Result{}, nil

}
