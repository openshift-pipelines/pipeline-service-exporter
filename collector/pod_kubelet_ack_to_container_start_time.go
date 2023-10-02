package collector

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

type ReconcilePodKubeletToContainerLatency struct {
	client        client.Client
	scheme        *runtime.Scheme
	eventRecorder record.EventRecorder
}

func (r *ReconcilePodKubeletToContainerLatency) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	return reconcile.Result{}, nil
}

func NewPodKubeletToContainerStartDurationMetric() *prometheus.HistogramVec {
	labelNames := []string{NS_LABEL, TASK_NAME_LABEL}
	metric := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "taskrun_pod_duration_kubelet_to_container_start_milliseconds",
		Help:    "Duration in milliseconds between the pod start time and the first container to start. This should include any overhead to pull container images, plus any kubelet to linux scheduling overhead.",
		Buckets: prometheus.ExponentialBuckets(float64(100), float64(5), 6),
	}, labelNames)
	metrics.Registry.MustRegister(metric)
	return metric
}

func SetupPodKubeletToContainerStartDurationController(mgr ctrl.Manager) error {
	filter := &kubeletContainerLatencyFilter{metric: NewPodKubeletToContainerStartDurationMetric()}
	reconciler := &ReconcilePodKubeletToContainerLatency{
		client:        mgr.GetClient(),
		scheme:        mgr.GetScheme(),
		eventRecorder: mgr.GetEventRecorderFor("MetricExporterStageTwoPods"),
	}
	return ctrl.NewControllerManagedBy(mgr).For(&corev1.Pod{}).WithEventFilter(filter).Complete(reconciler)
}

type PodKubeletToContainerLatencyCollector struct {
	podStartFirstContainerStartCollector *prometheus.HistogramVec
}

type kubeletContainerLatencyFilter struct {
	metric *prometheus.HistogramVec
}

func (f *kubeletContainerLatencyFilter) Create(event.CreateEvent) bool {
	return false
}

func (f *kubeletContainerLatencyFilter) Generic(event.GenericEvent) bool {
	return false
}

func (f *kubeletContainerLatencyFilter) Delete(event.DeleteEvent) bool {
	return false
}

func (f *kubeletContainerLatencyFilter) Update(e event.UpdateEvent) bool {
	oldpod, okold := e.ObjectOld.(*corev1.Pod)
	newpod, oknew := e.ObjectNew.(*corev1.Pod)
	if okold && oknew {
		labels := map[string]string{NS_LABEL: newpod.Namespace, TASK_NAME_LABEL: taskRef(newpod.Labels)}

		if oldpod.Status.StartTime == nil && newpod.Status.StartTime == nil {
			return false
		}
		if len(oldpod.Status.ContainerStatuses) == 0 && len(newpod.Status.ContainerStatuses) != 0 {
			// see if any of the new container statuses have a start time
			for _, cs := range newpod.Status.ContainerStatuses {
				if cs.State.Terminated != nil && !cs.State.Terminated.StartedAt.IsZero() {
					f.metric.With(labels).Observe(calculateTaskRunPodStartedToFirstContainerStartedDuration(newpod))
					return false
				}
				if cs.State.Running != nil && !cs.State.Running.StartedAt.IsZero() {
					f.metric.With(labels).Observe(calculateTaskRunPodStartedToFirstContainerStartedDuration(newpod))
					return false
				}
			}
		}
		if len(oldpod.Status.ContainerStatuses) != 0 && len(newpod.Status.ContainerStatuses) != 0 {
			for _, cs := range oldpod.Status.ContainerStatuses {
				// if old already has a container start, then quit
				if cs.State.Terminated != nil && !cs.State.Terminated.StartedAt.IsZero() {
					return false
				}
				// if old already has a container start, then quit
				if cs.State.Running != nil && !cs.State.Running.StartedAt.IsZero() {
					return false
				}
			}
			// if old had container statuses but no start times, then let's see if new has start times
			for _, cs := range newpod.Status.ContainerStatuses {
				if cs.State.Terminated != nil && !cs.State.Terminated.StartedAt.IsZero() {
					f.metric.With(labels).Observe(calculateTaskRunPodStartedToFirstContainerStartedDuration(newpod))
					return false
				}
				if cs.State.Running != nil && !cs.State.Running.StartedAt.IsZero() {
					f.metric.With(labels).Observe(calculateTaskRunPodStartedToFirstContainerStartedDuration(newpod))
					return false
				}
			}
		}
	}
	return false
}

// this minimally captures any time the kubelet spends pulling container images for the pod
func calculateTaskRunPodStartedToFirstContainerStartedDuration(pod *corev1.Pod) float64 {
	if pod.Status.StartTime == nil || pod.Status.StartTime.IsZero() {
		return 0
	}
	if len(pod.Status.ContainerStatuses) == 0 {
		return 0
	}
	var firstTime *metav1.Time
	for _, cs := range pod.Status.ContainerStatuses {
		if cs.State.Running != nil && !cs.State.Running.StartedAt.IsZero() {
			if firstTime == nil {
				firstTime = &cs.State.Running.StartedAt
				continue
			}
			if cs.State.Running.StartedAt.Before(firstTime) {
				firstTime = &cs.State.Running.StartedAt
				continue
			}
		}
		if cs.State.Terminated != nil && !cs.State.Terminated.StartedAt.IsZero() {
			if firstTime == nil {
				firstTime = &cs.State.Terminated.StartedAt
				continue
			}
			if cs.State.Terminated.StartedAt.Before(firstTime) {
				firstTime = &cs.State.Terminated.StartedAt
				continue
			}
		}
	}
	if firstTime == nil {
		return 0
	}
	return float64(firstTime.Time.Sub(pod.Status.StartTime.Time).Milliseconds())
}
