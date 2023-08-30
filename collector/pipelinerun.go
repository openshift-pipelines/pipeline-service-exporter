package collector

import (
	"context"
	"fmt"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"strconv"

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

type ReconcilePipelineRunScheduled struct {
	client        client.Client
	scheme        *runtime.Scheme
	eventRecorder record.EventRecorder
	prCollector   *PipelineRunScheduledCollector
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
	reconciler := &ReconcilePipelineRunScheduled{
		client:        mgr.GetClient(),
		scheme:        mgr.GetScheme(),
		eventRecorder: mgr.GetEventRecorderFor("MetricExporterPipelineRunsScheduled"),
		prCollector:   NewPipelineRunScheduledCollector(),
	}
	return ctrl.NewControllerManagedBy(mgr).For(&v1beta1.PipelineRun{}).WithEventFilter(&startTimeEventFilter{}).Complete(reconciler)
}

func (r *ReconcilePipelineRunScheduled) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
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
	r.prCollector.bumpScheduledDuration(pr, calculateScheduledDurationPipelineRun(pr))
	return reconcile.Result{}, nil
}

type ReconcilePipelineRunTaskRunGap struct {
	client        client.Client
	scheme        *runtime.Scheme
	eventRecorder record.EventRecorder
	prCollector   *PipelineRunTaskRunGapCollector
}

type taskRunGapEventFilter struct {
}

func (f *taskRunGapEventFilter) Create(event.CreateEvent) bool {
	return false
}

func (f *taskRunGapEventFilter) Delete(event.DeleteEvent) bool {
	return false
}

func (f *taskRunGapEventFilter) Update(e event.UpdateEvent) bool {

	//TODO remember, keep track of when pipeline-service and RHTAP starts moving from v1beta1 to v1
	oldPR, okold := e.ObjectOld.(*v1beta1.PipelineRun)
	newPR, oknew := e.ObjectNew.(*v1beta1.PipelineRun)
	// the real-time filtering involes retrieving the taskruns that are childs of this pipelinerun, so we only
	// calculate when the pipelinerun transtions to done, and then compare the kinds; note - do not need to check for cancel,
	// as eventually those PRs will be marked done once any running TRs are done
	if okold && oknew {
		// NOTE: confirmed that the succeeded condition is marked done and the completion timestamp is set at the same time
		if !oldPR.IsDone() && newPR.IsDone() {
			return true
		}
	}
	return false
}

func (f *taskRunGapEventFilter) Generic(event.GenericEvent) bool {
	return false
}

func optionalMetricEnabled(envVarName string) bool {
	env := os.Getenv(envVarName)
	enabled := len(env) > 0
	// any random setting means true
	if enabled {
		// allow for users to turn off by setting to false
		bv, err := strconv.ParseBool(env)
		if err == nil && !bv {
			return false
		}
		return true
	}
	return false
}

func SetupPipelineRunTaskRunGapController(mgr ctrl.Manager) error {
	reconciler := &ReconcilePipelineRunTaskRunGap{
		client:        mgr.GetClient(),
		scheme:        mgr.GetScheme(),
		eventRecorder: mgr.GetEventRecorderFor("MetricExporterPipelineRunsTaskRunGap"),
		prCollector:   NewPipelineRunTaskRunGapCollector(),
	}
	return ctrl.NewControllerManagedBy(mgr).For(&v1beta1.PipelineRun{}).WithEventFilter(&taskRunGapEventFilter{}).Complete(reconciler)
}

func (r *ReconcilePipelineRunTaskRunGap) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
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
	log.V(4).Info(fmt.Sprintf("recording taskrun gap for %q", request.NamespacedName))
	r.prCollector.bumpGapDuration(pr, r.client, ctx)
	return reconcile.Result{}, nil
}
