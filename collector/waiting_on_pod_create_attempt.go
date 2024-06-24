package collector

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	v1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	"github.com/tektoncd/pipeline/pkg/pod"
	corev1 "k8s.io/api/core/v1"
	"knative.dev/pkg/apis"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

type WaitingOnPodCreateAttemptCollector struct {
	waitPodCreate *prometheus.GaugeVec
}

func NewWaitingOnPodCreateAttemptCollector() *WaitingOnPodCreateAttemptCollector {
	labelNames := []string{NS_LABEL}
	waitPodCreate := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "taskrun_pod_create_not_attempted_or_pending_count",
		Help: "Number of TaskRuns where the Tekton Controller has yet to attempt to create its underlying Pod, or the TaskRun is still in Pending state for multiple scan iterations",
	}, labelNames)
	waitPodCreateCollector := &WaitingOnPodCreateAttemptCollector{
		waitPodCreate: waitPodCreate,
	}
	metrics.Registry.Register(waitPodCreate)
	return waitPodCreateCollector
}

func (c *WaitingOnPodCreateAttemptCollector) IncCollector(ns string) {
	labels := map[string]string{NS_LABEL: ns}
	c.waitPodCreate.With(labels).Inc()
}

func (c *WaitingOnPodCreateAttemptCollector) ZeroCollector(ns string) {
	labels := map[string]string{NS_LABEL: ns}
	c.waitPodCreate.With(labels).Set(float64(0))
}

func (r *ExporterReconcile) resetPodCreateAttemptedStats(ctx context.Context) {
	cacheCopy := buildLastScanCopy(r.pvcCollector, r.waitPodNSCache)

	// however, we'll clear out cache to avoid long term accumulation, memory leak, as things like dynamically created test namespaces
	// accumulate
	r.waitPodNSCache = map[string]map[string]struct{}{}

	trList := &v1.TaskRunList{}
	err := r.client.List(ctx, trList)
	deadlockTracker := &DeadlockTracker{
		collector:         r.waitPodCollector,
		filter:            r.podCreateNamespaceFilter,
		flaggedNamespaces: map[string]struct{}{},
		lastScan:          cacheCopy,
		currentScan:       r.waitPodNSCache,
	}
	if err == nil {
		for _, tr := range trList.Items {
			deadlockTracker.deadlocked = func() bool {
				if len(tr.Status.PodName) > 0 {
					return false
				}
				// see if the success condition has been handled by either TaskRunStatus.MarkResourceFailed
				// or TaskRunStatus.MarkReasonOngoing, which are the two methods called in the TaskRun reconciler
				// when handling pod create errors
				// failure will result in IsDone being
				if tr.IsDone() || tr.IsCancelled() {
					return false
				}

				if isTaskRunThrottled(&tr) {
					return false
				}

				// reason and message will be set when MarkReasonOngoing called, which should occur on a create pod error;
				// also see updateIncompleteTaskRunStatus for how Pod conditions are handled while in progress
				con := tr.Status.GetCondition(apis.ConditionSucceeded)
				if con != nil &&
					con.Status == corev1.ConditionUnknown &&
					(len(con.Reason) > 0 || len(con.Message) > 0) {
					return false
				}

				// if some pre-processing of the TaskRun aborted the reconciliation path before attempting to create the pod,
				// the TaskRun controller via finishReconcileUpdateEmitEvents and then updateLabelsAndAnnotations will
				// attempt to set the "pipeline.tekton.dev/release" annotation; I see that in my samples from konflux prod
				_, ok := tr.Annotations[pod.ReleaseAnnotation]
				if ok {
					return false
				}

				// any namespace throttling could defer pod creation; we wait until all throttled items have been processed
				_, nsThrotled := deadlockTracker.flaggedNamespaces[tr.Namespace]
				if nsThrotled {
					return false
				}

				controllerLog.Info(fmt.Sprintf("no pod creation yet for taskrun %s:%s, %s", tr.Namespace, tr.Name, createJSONFormattedString(tr)))
				return true
			}
			deadlockTracker.PerformDeadlockDetection(tr.Name, tr.Namespace)
		}
	} else {
		controllerLog.Error(err, "task run query for pod create attempts failed with an error")
	}

	// if a namespace is in the cache, but not our most recent scan, zero it out too, as the namespace is either
	// deleted or has all its TaskRuns pruned.
	zeroOutPriorHitNamespacesThatAreNowEmpty(r.waitPodCollector, cacheCopy, r.waitPodNSCache)
}
