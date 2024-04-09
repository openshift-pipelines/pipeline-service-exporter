package collector

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	v1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
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
		Name: "taskrun_pod_create_not_attempted_count",
		Help: "Number of TaskRuns where the Tekton Controller has yet to attempt to create its underlying Pod",
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

func attemptedToCreatePod(tr *v1.TaskRun) bool {
	if len(tr.Status.PodName) > 0 {
		return true
	}
	// see if the success condition has been handled by either TaskRunStatus.MarkResourceFailed
	// or TaskRunStatus.MarkReasonOngoing, which are the two methods called in the TaskRun reconciler
	// when handling pod create errors

	// failure will result in IsDone being
	if tr.IsDone() {
		controllerLog.Info("GGM is done")
		return true
	}

	// reason and message will be set when MarkReasonOngoing called
	con := tr.Status.GetCondition(apis.ConditionSucceeded)
	if con != nil &&
		con.Status == corev1.ConditionUnknown &&
		len(con.Reason) > 0 &&
		len(con.Message) > 0 {
		return true
	}
	controllerLog.Info(fmt.Sprintf("no pod creation yet for %s:%s", tr.Namespace, tr.Name))
	return false
}

func (r *ExporterReconcile) resetPodCreateAttemptedStats(ctx context.Context) {
	innerReset(r.pvcCollector, r.waitPodNSCache)
	// however, we'll clear out cache to avoid long term accumulation, memory leak, as things like dynamically created test clusters
	// accumulate; as long as we maintain history for permanent, active tenant namespaces, that is OK
	r.waitPodNSCache = map[string]struct{}{}

	trList := &v1.TaskRunList{}
	err := r.client.List(ctx, trList)
	nsWithWaitOnPod := map[string]struct{}{}
	if err == nil {
		for _, tr := range trList.Items {
			r.waitPodNSCache[tr.Namespace] = struct{}{}
			if !attemptedToCreatePod(&tr) {
				r.waitPodCollector.IncCollector(tr.Namespace)
				nsWithWaitOnPod[tr.Namespace] = struct{}{}
				continue
			}

			_, ok := nsWithWaitOnPod[tr.Namespace]
			if ok {
				continue
			}
			r.waitPodCollector.ZeroCollector(tr.Namespace)
		}
	}
}
