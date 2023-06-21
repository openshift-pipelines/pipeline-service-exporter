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
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type ReconcilePipelineRun struct {
	client        client.Client
	scheme        *runtime.Scheme
	eventRecorder record.EventRecorder
	psCollector   *PipelineServiceCollector
}

func SetupPipelineRunCachingClient(mgr ctrl.Manager) error {
	reconciler := &ReconcilePipelineRun{
		client:        mgr.GetClient(),
		scheme:        mgr.GetScheme(),
		eventRecorder: mgr.GetEventRecorderFor("MetricExporterPipelineRuns"),
		psCollector:   NewCollector(),
	}
	return ctrl.NewControllerManagedBy(mgr).For(&v1beta1.PipelineRun{}).Complete(reconciler)
}

func (r *ReconcilePipelineRun) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()
	log := ctrl.Log.WithName("metricsexporter").WithValues("namespace", request.Namespace, "resource", request.Name)

	pr := &v1beta1.PipelineRun{}
	err := r.client.Get(ctx, types.NamespacedName{Namespace: request.Namespace, Name: request.Name}, pr)
	if err != nil && !errors.IsNotFound(err) {
		return reconcile.Result{}, err
	}
	if err != nil {
		log.V(4).Info(fmt.Sprintf("ignoring deleted pipelinerun %q", request.NamespacedName))
		return reconcile.Result{}, nil
	}

	// we wait until the pipelinerun has reached a terminal state before recording schedule time to minimize the amount of samples recorded;
	// at this time we want to avoid updating the pipelinerun to indicate that we have recorded the metric
	if pr.IsDone() || pr.IsCancelled() {
		log.V(4).Info(fmt.Sprintf("recording schedule duration for %q", request.NamespacedName))
		r.psCollector.bumpScheduledDuration(pr.Namespace, calculateScheduledDuration(pr))
	}
	return reconcile.Result{}, nil
}
