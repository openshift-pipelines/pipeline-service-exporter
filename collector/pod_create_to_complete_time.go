package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

func NewPodCreateToCompleteMetric() *prometheus.HistogramVec {
	labelNames := []string{NS_LABEL}
	c2cMetric := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "tekton_pods_create_to_complete_seconds",
		Help: "Since tekton's duration are only from start time to completion, we provide a create time to completion for comparisons and potential alerting",
		// reminder: exponential buckets need a start value greater than 0
		// the results in buckets of 0.1, 0.5, 2.5, 12.5, 62.5, 312.5 seconds
		Buckets: prometheus.ExponentialBuckets(0.1, 5, 6),
	}, labelNames)
	metrics.Registry.MustRegister(c2cMetric)
	return c2cMetric
}

func NewPodCreateToCompleteFilter() *podCreateToCompleteFilter {
	return &podCreateToCompleteFilter{
		duration: NewPodCreateToCompleteMetric(),
	}
}

type podCreateToCompleteFilter struct {
	duration *prometheus.HistogramVec
}

func (f *podCreateToCompleteFilter) Create(event.CreateEvent) bool {
	return false
}

func (f *podCreateToCompleteFilter) Generic(event.GenericEvent) bool {
	return false
}

func (f *podCreateToCompleteFilter) Delete(e event.DeleteEvent) bool {
	return false
}

func (f *podCreateToCompleteFilter) Update(e event.UpdateEvent) bool {

	oldpod, okold := e.ObjectOld.(*corev1.Pod)
	newpod, oknew := e.ObjectNew.(*corev1.Pod)
	if okold && oknew {
		var oldTerminatedState *corev1.ContainerStateTerminated
		for _, status := range oldpod.Status.ContainerStatuses {
			oldTerminatedState = status.State.Terminated
			// if at least one container not terminated, keep set at nil and break
			if oldTerminatedState == nil {
				break
			}
		}
		var newTerminatedState *corev1.ContainerStateTerminated
		for _, status := range newpod.Status.ContainerStatuses {
			newTerminatedState = status.State.Terminated
			// if at least one container not terminated, keep set at nil and break
			if newTerminatedState == nil {
				break
			}
		}

		// if first transition when old pod still had non-terminated containers, but the new pod does not, process
		if oldTerminatedState == nil && newTerminatedState != nil {
			labels := map[string]string{NS_LABEL: newpod.Namespace}
			// we've seen in staging, especially with errors and short durations, and corroborated by comments I see in tekton,
			// where it is conceivable node times are not synchronized, when controller has been scheduled to other nodes than the pods, weird timestamps, etc.
			// so we check
			if newTerminatedState.FinishedAt.Time.Before(newpod.CreationTimestamp.Time) {
				f.duration.With(labels).Observe(float64(0))
				return false
			}
			f.duration.With(labels).Observe(newTerminatedState.FinishedAt.Time.Sub(newpod.CreationTimestamp.Time).Seconds())
		}
	}
	return false
}
