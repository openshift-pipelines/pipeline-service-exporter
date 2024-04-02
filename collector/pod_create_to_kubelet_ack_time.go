package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
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

func NewPodCreateToKubeletDurationMetric() *prometheus.HistogramVec {
	labelNames := []string{NS_LABEL, TASK_NAME_LABEL}
	metric := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "taskrun_pod_duration_kubelet_acknowledged_milliseconds",
		Help:    "Duration in milliseconds between the pod creation time and pod start time, where the pod start time is set once the kubelet has acknowledged the pod, but has not yet pulled its images.",
		Buckets: prometheus.ExponentialBuckets(float64(100), float64(5), 6),
	}, labelNames)
	metrics.Registry.MustRegister(metric)
	return metric
}

type createKubeletLatencyFilter struct {
	metric *prometheus.HistogramVec
}

func (f *createKubeletLatencyFilter) Create(event.CreateEvent) bool {
	return false
}

func (f *createKubeletLatencyFilter) Generic(event.GenericEvent) bool {
	return false
}

func (f *createKubeletLatencyFilter) Delete(event.DeleteEvent) bool {
	return false
}

func (f *createKubeletLatencyFilter) Update(e event.UpdateEvent) bool {

	oldpod, okold := e.ObjectOld.(*corev1.Pod)
	newpod, oknew := e.ObjectNew.(*corev1.Pod)
	if okold && oknew {
		if oldpod.Status.StartTime == nil && newpod.Status.StartTime != nil {
			labels := map[string]string{NS_LABEL: newpod.Namespace, TASK_NAME_LABEL: taskRef(newpod.Labels)}
			f.metric.With(labels).Observe(calculateTaskRunPodCreatedToKubeletAcceptsAndStartTimeSetDuration(newpod))
			return false
		}
	}
	return false
}

// this captures how long it takes for the kubelet to accept the pod after the pod is created
func calculateTaskRunPodCreatedToKubeletAcceptsAndStartTimeSetDuration(pod *corev1.Pod) float64 {
	if pod.Status.StartTime == nil || pod.Status.StartTime.IsZero() {
		return 0
	}
	return float64(pod.Status.StartTime.Time.Sub(pod.CreationTimestamp.Time).Milliseconds())
}
