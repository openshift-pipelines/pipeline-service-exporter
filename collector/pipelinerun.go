package collector

import (
	"context"
	"fmt"
	"sigs.k8s.io/controller-runtime/pkg/event"

	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type ReconcilePipelineRun struct {
	client        client.Client
	scheme        *runtime.Scheme
	eventRecorder record.EventRecorder
	psCollector   *PipelineServiceCollector
}

type startTimeEventFilter struct {
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
			return true
		}
	}
	return false
}

func (f *startTimeEventFilter) Generic(event.GenericEvent) bool {
	return false
}

func SetupPipelineRunScheduleDurationController(mgr ctrl.Manager) error {
	reconciler := &ReconcilePipelineRun{
		client:        mgr.GetClient(),
		scheme:        mgr.GetScheme(),
		eventRecorder: mgr.GetEventRecorderFor("MetricExporterPipelineRuns"),
		psCollector:   NewCollector(),
	}
	return ctrl.NewControllerManagedBy(mgr).For(&v1beta1.PipelineRun{}).WithEventFilter(&startTimeEventFilter{}).Complete(reconciler)
}

func (r *ReconcilePipelineRun) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()
	log := log.FromContext(ctx)

	//TODO remember, keep track of when pipeline-service and RHTAP starts moving from v1beta1 to v1
	pr := &v1beta1.PipelineRun{}
	err := r.client.Get(ctx, types.NamespacedName{Namespace: request.Namespace, Name: request.Name}, pr)
	if err != nil && !errors.IsNotFound(err) {
		return reconcile.Result{}, err
	}
	if err != nil {
		log.V(4).Info(fmt.Sprintf("ignoring deleted pipelinerun %q", request.NamespacedName))
		return reconcile.Result{}, nil
	}

	// based on our WithEventFilter we should only be getting called with the start time is set
	log.V(4).Info(fmt.Sprintf("recording schedule duration for %q", request.NamespacedName))
	r.psCollector.bumpScheduledDuration(pr.Namespace, calculateScheduledDuration(pr))
	return reconcile.Result{}, nil
}
