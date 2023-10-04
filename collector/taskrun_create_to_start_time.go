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

/*
  Observed sequence from testing and tekton/k8s code examination:

- taskrun created by external user or pipelinerun reconciler, and k8s sets the taskrun create timestap
- on first reconcile event loop, taskrun ConditionSucceeded condition is intialized to pending and the startime is set
- "prepare" will see if resolutionrequest is needed, and if so, error is returned to controller fw and requeue happens
- otherwise, or on next reconcile event loop, volume claim templates / workspaces processed, params processed, and if no errors, pod create attempted
- if error requeue
- otherwise pod status converted to taskrun status, with conditions now updated and set running, controller then puts back on queue with requeueAfter(timeDuration)

Now concurrent event streams from both Pod updates or TaskRun updates, or requeuesAfter(timeDuration) calling reconcile again, can update task run status and conditions

On pods
- create time stems from tekton creating
- pod start time means kubelet has "accepted" per godoc for scheduling, but no image pulls have occurred, where
using of "latest" for the tag means always pull per godoc, but use of a specific SHA means pull if not on local CRI-O / "node cache"

So,
- no real diff of worth between taskrun startime and its conditions
- pod create vs. pod start time captures how long the external factor of the Kubelet agreeing the schedule the pod takes
- pod start time vs. first container start captures how long to pull images and schedule the container
- where as the upstream latency metric of the last transition time of the `corev1.PodScheduled` condition minus the pod create time
is "perhaps" the sum of those two
*/

func SetupTaskRunScheduleDurationController(mgr ctrl.Manager) error {
	filter := &trStartTimeEventFilter{metric: NewTaskRunScheduledMetric()}
	reconciler := &ReconcileTaskRunScheduled{
		client:        mgr.GetClient(),
		scheme:        mgr.GetScheme(),
		eventRecorder: mgr.GetEventRecorderFor("MetricExporterScheduledTaskRuns"),
	}
	return ctrl.NewControllerManagedBy(mgr).For(&v1beta1.TaskRun{}).WithEventFilter(filter).Complete(reconciler)
}

func NewTaskRunScheduledMetric() *prometheus.HistogramVec {
	labelNames := []string{NS_LABEL, TASK_NAME_LABEL}
	durationScheduled := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "taskrun_duration_scheduled_seconds",
		Help: "Duration in seconds for a TaskRun to be 'scheduled', meaning it has been received by the Tekton controller.  This is an indication of how quickly create events from the API server are arriving to the Tekton controller.",
		// reminder: exponential buckets need a start value greater than 0
		// the results in buckets of 0.1, 0.5, 2.5, 12.5, 62.5, 312.5 seconds
		Buckets: prometheus.ExponentialBuckets(0.1, 5, 6),
	}, labelNames)

	metrics.Registry.MustRegister(durationScheduled)

	return durationScheduled

}

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

func (r *ReconcileTaskRunScheduled) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	return reconcile.Result{}, nil
}

func bumpTaskRunScheduledDuration(scheduleDuration float64, tr *v1beta1.TaskRun, metric *prometheus.HistogramVec) {
	labels := map[string]string{NS_LABEL: tr.Namespace, TASK_NAME_LABEL: taskRef(tr.Labels)}
	metric.With(labels).Observe(scheduleDuration)
}

func calculateScheduledDurationTaskRun(taskrun *v1beta1.TaskRun) float64 {
	return calcuateScheduledDuration(taskrun.CreationTimestamp.Time, taskrun.Status.StartTime.Time)
}
