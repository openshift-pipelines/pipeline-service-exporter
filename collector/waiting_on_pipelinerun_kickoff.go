package collector

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	v1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	"knative.dev/pkg/apis"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

type WaitingOnPipelineRunKickoffCollector struct {
	waitPipelineRunKickoff *prometheus.GaugeVec
}

func NewWaitingOnPipelineRunKickoffCollector() *WaitingOnPipelineRunKickoffCollector {
	labelNames := []string{NS_LABEL}
	waitPipelineRunKickoff := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pipelinerun_kickoff_not_attempted_count",
		Help: "Number of PipelineRuns where the Tekton Controller has yet to attempt to process its correctly defined Task specifications for multiple scan iterations",
	}, labelNames)
	waitPipelineRunKickoffCollector := &WaitingOnPipelineRunKickoffCollector{
		waitPipelineRunKickoff: waitPipelineRunKickoff,
	}
	metrics.Registry.Register(waitPipelineRunKickoff)
	return waitPipelineRunKickoffCollector
}

func (c *WaitingOnPipelineRunKickoffCollector) IncCollector(ns string) {
	labels := map[string]string{NS_LABEL: ns}
	c.waitPipelineRunKickoff.With(labels).Inc()
}

func (c *WaitingOnPipelineRunKickoffCollector) ZeroCollector(ns string) {
	labels := map[string]string{NS_LABEL: ns}
	c.waitPipelineRunKickoff.With(labels).Set(float64(0))
}

func (r *ExporterReconcile) resetPipelineRunKickoffStats(ctx context.Context) {
	cacheCopy := buildLastScanCopy(r.waitPRKickoffCollector, r.waitPRKickoffCache)

	// however, we'll clear out cache to avoid long term accumulation, memory leak, as things like dynamically created test namespaces
	// accumulate
	r.waitPRKickoffCache = map[string]map[string]struct{}{}

	prList := &v1.PipelineRunList{}
	err := r.client.List(ctx, prList)
	deadlockTracker := &DeadlockTracker{
		collector:         r.waitPRKickoffCollector,
		filter:            r.pipelineRunKickoffNamespaceFilter,
		flaggedNamespaces: map[string]struct{}{},
		lastScan:          cacheCopy,
		currentScan:       r.waitPRKickoffCache,
	}
	if err == nil {
		for _, pr := range prList.Items {
			deadlockTracker.deadlocked = func() bool {
				if pr.IsDone() || pr.IsCancelled() || pr.IsGracefullyCancelled() || pr.IsGracefullyStopped() || pr.IsPending() {
					return false
				}

				throttled, _, _ := isPipelineRunThrottled(&pr, r.client, context.Background())
				if throttled {
					return false
				}

				if len(pr.Status.ChildReferences) > 0 || len(pr.Status.SkippedTasks) > 0 {
					return false
				}

				c := pr.GetStatusCondition().GetCondition(apis.ConditionSucceeded)
				if c != nil && (len(c.Reason) > 0 || len(c.Message) > 0) {
					return false
				}

				// FYI this is set before the Task DAG is built
				if pr.HasStarted() {
					return false
				}

				// any namespace throttling could defer taskrun/pod creation; we wait until all throttled items have been processed
				_, nsThrotled := deadlockTracker.flaggedNamespaces[pr.Namespace]
				if nsThrotled {
					return false
				}

				controllerLog.Info(fmt.Sprintf("no pipelinerun kickoff yet for pipelinerun %s:%s, %s", pr.Namespace, pr.Name, createJSONFormattedString(pr)))
				return true

			}
			deadlockTracker.PerformDeadlockDetection(pr.Name, pr.Namespace)
		}
	} else {
		controllerLog.Error(err, "pipeline run query for kickoff attempts failed with an error")
	}

	// if a namespace is in the cache, but not our most recent scan, zero it out too, as the namespace is either
	// deleted or has all its PipelineRuns pruned.
	zeroOutPriorHitNamespacesThatAreNowEmpty(r.waitPRKickoffCollector, cacheCopy, r.waitPRKickoffCache)
}
