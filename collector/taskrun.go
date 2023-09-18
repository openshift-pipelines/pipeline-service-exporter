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
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type ReconcileTaskRunScheduled struct {
	client        client.Client
	scheme        *runtime.Scheme
	eventRecorder record.EventRecorder
}

type trStartTimeEventFilter struct {
	metric *prometheus.HistogramVec
}

func (f *trStartTimeEventFilter) Create(event.CreateEvent) bool {
	return false
}

func (f *trStartTimeEventFilter) Generic(event.GenericEvent) bool {
	return false
}

func (f *trStartTimeEventFilter) Delete(event.DeleteEvent) bool {
	return false
}

func (f *trStartTimeEventFilter) Update(e event.UpdateEvent) bool {
	//TODO remember, keep track of when pipeline-service and RHTAP starts moving from v1beta1 to v1
	oldTR, okold := e.ObjectOld.(*v1beta1.TaskRun)
	newTR, oknew := e.ObjectNew.(*v1beta1.TaskRun)
	if okold && oknew {
		if oldTR.Status.StartTime == nil && newTR.Status.StartTime != nil {
			bumpTaskRunScheduledDuration(calculateScheduledDurationTaskRun(newTR), newTR, f.metric)
			return false
		}
	}
	return false
}

func SetupTaskRunScheduleDurationController(mgr ctrl.Manager) error {
	filter := &trStartTimeEventFilter{metric: NewTaskRunScheduledMetric()}
	reconciler := &ReconcileTaskRunScheduled{
		client:        mgr.GetClient(),
		scheme:        mgr.GetScheme(),
		eventRecorder: mgr.GetEventRecorderFor("MetricExporterScheduledTaskRuns"),
	}
	return ctrl.NewControllerManagedBy(mgr).For(&v1beta1.TaskRun{}).WithEventFilter(filter).Complete(reconciler)
}

func (r *ReconcileTaskRunScheduled) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	return reconcile.Result{}, nil
}
