package collector

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func SetupPipelineRunTaskRunGapController(mgr ctrl.Manager) error {
	reconciler := &ReconcilePipelineRunTaskRunGap{
		client:        mgr.GetClient(),
		scheme:        mgr.GetScheme(),
		eventRecorder: mgr.GetEventRecorderFor("MetricExporterPipelineRunsTaskRunGap"),
		prCollector:   NewPipelineRunTaskRunGapCollector(),
	}
	return ctrl.NewControllerManagedBy(mgr).For(&v1.PipelineRun{}).WithEventFilter(&taskRunGapEventFilter{}).Complete(reconciler)
}

type PipelineRunTaskRunGapCollector struct {
	trGaps           *prometheus.HistogramVec
	additionalLabels bool
}

func NewPipelineRunTaskRunGapCollector() *PipelineRunTaskRunGapCollector {
	labelNames := []string{NS_LABEL, STATUS_LABEL}
	additionalLabels := optionalMetricEnabled(ENABLE_GAP_METRIC_ADDITIONAL_LABELS)
	if additionalLabels {
		labelNames = append(labelNames, PIPELINE_NAME_LABEL, COMPLETED_LABEL, UPCOMING_LABEL)
	}
	trGaps := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "pipelinerun_gap_between_taskruns_milliseconds",
		Help: "Duration in milliseconds between a taskrun completing and the next taskrun being created within a pipelinerun.  For a pipelinerun's first taskrun, the duration is the time between that taskrun's creation and the pipelinerun's creation.",
		// reminder: exponential buckets need a start value greater than 0
		// the results in buckets of 100, 500, 2500, 12500, 62500, 312500 milliseconds
		Buckets: prometheus.ExponentialBuckets(float64(100), float64(5), 6),
	}, labelNames)

	pipelineRunTaskRunGapCollector := &PipelineRunTaskRunGapCollector{
		trGaps:           trGaps,
		additionalLabels: additionalLabels,
	}
	metrics.Registry.MustRegister(trGaps)

	return pipelineRunTaskRunGapCollector
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

	oldPR, okold := e.ObjectOld.(*v1.PipelineRun)
	newPR, oknew := e.ObjectNew.(*v1.PipelineRun)
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

func (r *ReconcilePipelineRunTaskRunGap) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
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

	// based on our WithEventFilter we should only be getting called with the start time is set
	log.V(4).Info(fmt.Sprintf("recording taskrun gap for %q", request.NamespacedName))
	r.prCollector.bumpGapDuration(pr, r.client, ctx)
	return reconcile.Result{}, nil
}

func (c *PipelineRunTaskRunGapCollector) bumpGapDuration(pr *v1.PipelineRun, oc client.Client, ctx context.Context) {
	if skipPipelineRun(pr) {
		return
	}

	sortedTaskRunsByCreateTimes, reverseOrderSortedTaskRunsByCompletionTimes, abort := sortTaskRunsForGapCalculations(pr, oc, ctx)

	if abort {
		return
	}

	prRef := pipelineRunPipelineRef(pr)
	gapEntries := calculateGaps(ctx, pr, oc, sortedTaskRunsByCreateTimes, reverseOrderSortedTaskRunsByCompletionTimes)
	for _, gapEntry := range gapEntries {
		labels := map[string]string{
			NS_LABEL:     pr.Namespace,
			STATUS_LABEL: gapEntry.status,
		}
		if c.additionalLabels {
			labels[PIPELINE_NAME_LABEL] = prRef
			labels[COMPLETED_LABEL] = gapEntry.completed
			labels[UPCOMING_LABEL] = gapEntry.upcoming
		}
		c.trGaps.With(labels).Observe(gapEntry.gap)
	}

	return
}
