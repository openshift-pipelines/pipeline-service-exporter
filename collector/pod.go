package collector

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type ReconcilePodCreateToKubeletLatency struct {
	client        client.Client
	scheme        *runtime.Scheme
	eventRecorder record.EventRecorder
}

func (r *ReconcilePodCreateToKubeletLatency) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	return reconcile.Result{}, nil
}

func SetupPodCreateToKubeletDurationController(mgr ctrl.Manager) error {
	filter := &createKubeletLatencyFilter{metric: NewPodCreateToKubeletDurationMetric()}
	reconciler := &ReconcilePodCreateToKubeletLatency{
		client:        mgr.GetClient(),
		scheme:        mgr.GetScheme(),
		eventRecorder: mgr.GetEventRecorderFor("MetricExporterStageOnePods"),
	}
	return ctrl.NewControllerManagedBy(mgr).For(&corev1.Pod{}).WithEventFilter(filter).Complete(reconciler)
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

type ReconcilePodKubeletToContainerLatency struct {
	client        client.Client
	scheme        *runtime.Scheme
	eventRecorder record.EventRecorder
}

func (r *ReconcilePodKubeletToContainerLatency) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	return reconcile.Result{}, nil
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
