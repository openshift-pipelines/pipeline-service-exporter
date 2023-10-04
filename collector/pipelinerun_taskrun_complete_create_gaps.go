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
	"knative.dev/pkg/apis"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sort"
	"time"
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
	if len(pr.Status.ChildReferences) < 1 {
		return
	}
	// in case there are gaps between a pipelinerun being marked done but the complete timestamp is not set, with the
	// understanding that the complete timestamp is not processed before any completed taskrun complete timestamps have been processed
	if pr.Status.CompletionTime == nil {
		return
	}

	sortedTaskRunsByCreateTimes := []*v1.TaskRun{}
	reverseOrderSortedTaskRunsByCompletionTimes := []*v1.TaskRun{}
	// prior testing in staging proved that with enough concurrency, this array is minimally not sorted based on when
	// the task runs were created, so we explicitly sort for that; also, this sorting will allow us to effectively
	// address parallel taskruns vs. taskrun dependencies and ordering (where tekton does not create a taskrun until its dependencies
	// have completed).
	for _, kidRef := range pr.Status.ChildReferences {
		if kidRef.Kind != "TaskRun" {
			continue
		}
		kid := &v1.TaskRun{}
		err := oc.Get(ctx, types.NamespacedName{Namespace: pr.Namespace, Name: kidRef.Name}, kid)
		if err != nil {
			ctrl.Log.Info(fmt.Sprintf("could not calculate gap for taskrun %s:%s: %s", pr.Namespace, kidRef.Name, err.Error()))
			continue
		}

		sortedTaskRunsByCreateTimes = append(sortedTaskRunsByCreateTimes, kid)
		// don't add taskruns that did not complete i.e. presumably timed out of failed; any taskruns that dependended
		// on should not have even been created
		if kid.Status.CompletionTime != nil {
			reverseOrderSortedTaskRunsByCompletionTimes = append(reverseOrderSortedTaskRunsByCompletionTimes, kid)

		}
	}
	sort.SliceStable(sortedTaskRunsByCreateTimes, func(i, j int) bool {
		return sortedTaskRunsByCreateTimes[i].CreationTimestamp.Time.Before(sortedTaskRunsByCreateTimes[j].CreationTimestamp.Time)
	})
	sort.SliceStable(reverseOrderSortedTaskRunsByCompletionTimes, func(i, j int) bool {
		return reverseOrderSortedTaskRunsByCompletionTimes[i].Status.CompletionTime.Time.After(reverseOrderSortedTaskRunsByCompletionTimes[j].Status.CompletionTime.Time)
	})
	prRef := pipelineRunPipelineRef(pr)
	for index, tr := range sortedTaskRunsByCreateTimes {
		succeedCondition := pr.Status.GetCondition(apis.ConditionSucceeded)
		if succeedCondition == nil {
			ctrl.Log.Info(fmt.Sprintf("WARNING: pipielinerun %s:%s marked done but has nil succeed condition", pr.Namespace, pr.Name))
			continue
		}
		if succeedCondition.IsUnknown() {
			ctrl.Log.Info(fmt.Sprintf("WARNING: pipielinerun %s:%s marked done but has unknown succeed condition", pr.Namespace, pr.Name))
			continue
		}
		status := SUCCEEDED
		if succeedCondition.IsFalse() {
			status = FAILED
		}
		labels := map[string]string{NS_LABEL: pr.Namespace, STATUS_LABEL: status}
		if c.additionalLabels {
			labels[PIPELINE_NAME_LABEL] = prRef
		}

		if index == 0 {
			ctrl.Log.V(4).Info(fmt.Sprintf("first task %s for pipeline %s", taskRef(tr.Labels), prRef))
			// our first task is simple, just work off of the pipelinerun
			gap := tr.CreationTimestamp.Time.Sub(pr.CreationTimestamp.Time).Milliseconds()
			if c.additionalLabels {
				labels[COMPLETED_LABEL] = prRef
				labels[UPCOMING_LABEL] = taskRef(tr.Labels)
			}
			c.trGaps.With(labels).Observe(float64(gap))
			continue
		}

		firstKid := sortedTaskRunsByCreateTimes[0]

		// so using the first taskrun completion time addresses sequential / chaining dependencies;
		// for parallel, if the first taskrun's completion time is not after this taskrun's create time,
		// that means parallel taskruns, and we work off of the pipelinerun; NOTE: this focuses on "top level" parallel task runs
		// with absolutely no dependencies.  Once any sort of dependency is established, there are no more top level parallel taskruns.
		if firstKid.Status.CompletionTime != nil && firstKid.Status.CompletionTime.Time.After(tr.CreationTimestamp.Time) {
			ctrl.Log.V(4).Info(fmt.Sprintf("task %s considered parallel for pipeline %s", taskRef(tr.Labels), prRef))
			gap := tr.CreationTimestamp.Time.Sub(pr.CreationTimestamp.Time).Milliseconds()
			if c.additionalLabels {
				labels[COMPLETED_LABEL] = prRef
				labels[UPCOMING_LABEL] = taskRef(tr.Labels)
			}
			c.trGaps.With(labels).Observe(float64(gap))
			continue
		}

		// Conversely, task run chains can run in parallel, and a taskrun can depend on multiple chains or threads of taskruns. We want to find the chain
		// that finished last, but before we are created.  We traverse through our reverse sorted on completion time list to determine that.  But yes, we don't reproduce the DAG
		// graph (there is no clean dependency import path in tekton for that) to confirm the edges.  This approximation is sufficient.

		// get whatever completed first
		timeToCalculateWith := time.Time{}
		trToCalculateWith := &v1.TaskRun{}
		if len(reverseOrderSortedTaskRunsByCompletionTimes) > 0 {
			trToCalculateWith = reverseOrderSortedTaskRunsByCompletionTimes[len(reverseOrderSortedTaskRunsByCompletionTimes)-1]
			timeToCalculateWith = trToCalculateWith.Status.CompletionTime.Time
		} else {
			// if no taskruns completed, that means any taskruns created were created as part of the initial pipelinerun creation,
			// so use the pipelinerun creation time
			timeToCalculateWith = pr.CreationTimestamp.Time
		}
		for _, tr2 := range reverseOrderSortedTaskRunsByCompletionTimes {
			if tr2.Name == tr.Name {
				continue
			}
			ctrl.Log.V(4).Info(fmt.Sprintf("comparing candidate %s to current task %s", taskRef(tr2.Labels), taskRef(tr.Labels)))
			if !tr2.Status.CompletionTime.Time.After(tr.CreationTimestamp.Time) {
				ctrl.Log.V(4).Info(fmt.Sprintf("%s did not complete after so use it to compute gap for current task %s", taskRef(tr2.Labels), taskRef(tr.Labels)))
				trToCalculateWith = tr2
				timeToCalculateWith = tr2.Status.CompletionTime.Time
				break
			}
			ctrl.Log.V(4).Info(fmt.Sprintf("skipping %s as a gap candidate for current task %s is OK", taskRef(tr2.Labels), taskRef(tr.Labels)))
		}
		gap := tr.CreationTimestamp.Time.Sub(timeToCalculateWith).Milliseconds()
		if c.additionalLabels {
			labels[COMPLETED_LABEL] = taskRef(trToCalculateWith.Labels)
			labels[UPCOMING_LABEL] = taskRef(tr.Labels)
		}
		c.trGaps.With(labels).Observe(float64(gap))
	}

	return
}
