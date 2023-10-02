package collector

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func SetupPipelineRunScheduleDurationController(mgr ctrl.Manager) error {
	filter := &startTimeEventFilter{
		metric: NewPipelineRunScheduledMetric(),
	}
	reconciler := &ReconcilePipelineRunScheduled{
		client:        mgr.GetClient(),
		scheme:        mgr.GetScheme(),
		eventRecorder: mgr.GetEventRecorderFor("MetricExporterPipelineRunsScheduled"),
	}
	return ctrl.NewControllerManagedBy(mgr).For(&v1beta1.PipelineRun{}).WithEventFilter(filter).Complete(reconciler)
}

type PipelineRunScheduledCollector struct {
	durationScheduled *prometheus.HistogramVec
	prSchedNameLabel  bool
}

func NewPipelineRunScheduledMetric() *prometheus.HistogramVec {
	labelNames := []string{NS_LABEL, PIPELINE_NAME_LABEL}
	durationScheduled := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "pipelinerun_duration_scheduled_seconds",
		Help: "Duration in seconds for a PipelineRun to be 'scheduled', meaning it has been received by the Tekton controller.  This is an indication of how quickly create events from the API server are arriving to the Tekton controller.",
		// reminder: exponential buckets need a start value greater than 0
		// the results in buckets of 0.1, 0.5, 2.5, 12.5, 62.5, 312.5 seconds
		Buckets: prometheus.ExponentialBuckets(0.1, 5, 6),
	}, labelNames)

	metrics.Registry.MustRegister(durationScheduled)

	return durationScheduled
}

func bumpPipelineRunScheduledDuration(scheduleDuration float64, pr *v1beta1.PipelineRun, metric *prometheus.HistogramVec) {
	labels := map[string]string{NS_LABEL: pr.Namespace, PIPELINE_NAME_LABEL: pipelineRunPipelineRef(pr)}
	metric.With(labels).Observe(scheduleDuration)
}

func calculateScheduledDurationPipelineRun(pipelineRun *v1beta1.PipelineRun) float64 {
	return calcuateScheduledDuration(pipelineRun.CreationTimestamp.Time, pipelineRun.Status.StartTime.Time)
}

type ReconcilePipelineRunScheduled struct {
	client        client.Client
	scheme        *runtime.Scheme
	eventRecorder record.EventRecorder
	prCollector   *PipelineRunScheduledCollector
}

type startTimeEventFilter struct {
	metric *prometheus.HistogramVec
}

func (f *startTimeEventFilter) Create(event.CreateEvent) bool {
	return false
}

func (f *startTimeEventFilter) Delete(event.DeleteEvent) bool {
	return false
}

func (f *startTimeEventFilter) Update(e event.UpdateEvent) bool {

	//TODO remember, keep track of when pipeline-service and RHTAP starts moving from v1beta1 to v1
	oldPR, okold := e.ObjectOld.(*v1beta1.PipelineRun)
	newPR, oknew := e.ObjectNew.(*v1beta1.PipelineRun)
	if okold && oknew {
		if oldPR.Status.StartTime == nil && newPR.Status.StartTime != nil {
			bumpPipelineRunScheduledDuration(calculateScheduledDurationPipelineRun(newPR), newPR, f.metric)
			return false
		}
	}
	return false
}

func (f *startTimeEventFilter) Generic(event.GenericEvent) bool {
	return false
}

func (r *ReconcilePipelineRunScheduled) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	return reconcile.Result{}, nil
}
