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
	"sigs.k8s.io/controller-runtime/pkg/metrics"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func SetupPodCreateToCompleteTimeController(mgr ctrl.Manager) error {
	reconciler := &ReconcilePodCreateToCompleteTime{
		client:        mgr.GetClient(),
		scheme:        mgr.GetScheme(),
		eventRecorder: mgr.GetEventRecorderFor("MetricExporterPodCreateToCompleteTime"),
	}
	filter := NewPodCreateToCompleteFilter()
	metrics.Registry.MustRegister(filter.duration)
	return ctrl.NewControllerManagedBy(mgr).For(&corev1.Pod{}).WithEventFilter(filter).Complete(reconciler)
}

func NewPodCreateToCompleteMetric() *prometheus.HistogramVec {
	labelNames := []string{NS_LABEL, PIPELINE_NAME_LABEL, TASK_NAME_LABEL}
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "tekton_pods_create_to_complete_seconds",
		Help: "Since tekton's duration are only from start time to completion, we provide a create time to completion for comparisons and potential alerting",
		// reminder: exponential buckets need a start value greater than 0
		// the results in buckets of 0.1, 0.5, 2.5, 12.5, 62.5, 312.5 seconds
		Buckets: prometheus.ExponentialBuckets(0.1, 5, 6),
	}, labelNames)
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
			labels := map[string]string{NS_LABEL: newpod.Namespace, PIPELINE_NAME_LABEL: pipelineRef(newpod.Labels), TASK_NAME_LABEL: taskRef(newpod.Labels)}
			f.duration.With(labels).Observe(newTerminatedState.FinishedAt.Time.Sub(newpod.CreationTimestamp.Time).Seconds())
		}
	}
	return false
}

type ReconcilePodCreateToCompleteTime struct {
	client        client.Client
	scheme        *runtime.Scheme
	eventRecorder record.EventRecorder
}

func (r *ReconcilePodCreateToCompleteTime) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	return reconcile.Result{}, nil
}
