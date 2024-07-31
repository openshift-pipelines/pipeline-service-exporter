package collector

import (
	"context"
	"fmt"
	"knative.dev/pkg/apis"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/tektoncd/pipeline/pkg/apis/resolution/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

type WaitingOnResolutionRequestCollector struct {
	waitingResolutionRequest *prometheus.GaugeVec
}

func NewWaitingOnResolutionRequestCollector() *WaitingOnResolutionRequestCollector {
	labelNames := []string{NS_LABEL}
	waitingResolutionRequest := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pending_resolutionrequest_count",
		Help: "Number of uncompleted ResolutionRequests for multiple scan iterations",
	}, labelNames)
	waitingOnResolutionRequestCollector := &WaitingOnResolutionRequestCollector{
		waitingResolutionRequest: waitingResolutionRequest,
	}
	metrics.Registry.Register(waitingResolutionRequest)
	return waitingOnResolutionRequestCollector
}

func (c *WaitingOnResolutionRequestCollector) IncCollector(ns string) {
	labels := map[string]string{NS_LABEL: ns}
	c.waitingResolutionRequest.With(labels).Inc()
}

func (c *WaitingOnResolutionRequestCollector) ZeroCollector(ns string) {
	labels := map[string]string{NS_LABEL: ns}
	c.waitingResolutionRequest.With(labels).Set(float64(0))
}

func (r *ExporterReconcile) resetResoultionRequestsStats(ctx context.Context) {
	cacheCopy := buildLastScanCopy(r.waitRRCollector, r.waitRRCache)

	// however, we'll clear out cache to avoid long term accumulation, memory leak, as things like dynamically created test namespaces
	// accumulate
	r.waitRRCache = map[string]map[string]struct{}{}

	rrList := &v1beta1.ResolutionRequestList{}
	err := r.client.List(ctx, rrList)
	deadlockTracker := &DeadlockTracker{
		collector:         r.waitRRCollector,
		filter:            r.pipelineRunKickoffNamespaceFilter,
		flaggedNamespaces: map[string]struct{}{},
		lastScan:          cacheCopy,
		currentScan:       r.waitRRCache,
	}
	if err == nil {
		for _, rr := range rrList.Items {
			deadlockTracker.deadlocked = func() bool {
				if rr.IsDone() {
					return false
				}

				if rr.Status.GetCondition(apis.ConditionSucceeded) != nil {
					return false
				}

				controllerLog.Info(fmt.Sprintf("resoultionrequest not started %s:%s, %s", rr.Namespace, rr.Name, createJSONFormattedString(rr)))
				return true

			}
			deadlockTracker.PerformDeadlockDetection(rr.Name, rr.Namespace)
		}
	} else {
		controllerLog.Error(err, "resolutionrequest query for kickoff attempts failed with an error")
	}

	// if a namespace is in the cache, but not our most recent scan, zero it out too, as the namespace is either
	// deleted or has all its PipelineRuns and owned objects like ResoulutionRequests pruned.
	zeroOutPriorHitNamespacesThatAreNowEmpty(r.waitRRCollector, cacheCopy, r.waitRRCache)
}
