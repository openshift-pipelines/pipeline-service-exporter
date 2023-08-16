package collector

import (
	"context"
	"fmt"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type ReconcileTaskRun struct {
	client        client.Client
	scheme        *runtime.Scheme
	eventRecorder record.EventRecorder
	trCollector   *TaskRunCollector
}

type trStartTimeEventFilter struct {
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
			return true
		}
	}
	return false
}

func SetupTaskRunController(mgr ctrl.Manager) error {
	reconciler := &ReconcileTaskRun{
		client:        mgr.GetClient(),
		scheme:        mgr.GetScheme(),
		eventRecorder: mgr.GetEventRecorderFor("MetricExporterTaskRuns"),
		trCollector:   NewTaskRunCollector(),
	}
	return ctrl.NewControllerManagedBy(mgr).For(&v1beta1.TaskRun{}).WithEventFilter(&trStartTimeEventFilter{}).Complete(reconciler)
}

func (r *ReconcileTaskRun) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()
	log := log.FromContext(ctx)

	//TODO remember, keep track of when pipeline-service and RHTAP starts moving from v1beta1 to v1
	tr := &v1beta1.TaskRun{}
	err := r.client.Get(ctx, types.NamespacedName{Namespace: request.Namespace, Name: request.Name}, tr)
	if err != nil && !errors.IsNotFound(err) {
		return reconcile.Result{}, err
	}
	if err != nil {
		log.V(4).Info(fmt.Sprintf("ignoring deleted taskrun %q", request.NamespacedName))
		return reconcile.Result{}, nil
	}

	// based on our WithEventFilter we should only be getting called with the start time is set
	log.V(4).Info(fmt.Sprintf("recording schedule duration for %q", request.NamespacedName))
	r.trCollector.bumpScheduledDuration(tr, calculateScheduledDurationTaskRun(tr))
	return reconcile.Result{}, nil
}
