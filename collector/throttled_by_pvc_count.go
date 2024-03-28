package collector

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	"github.com/tektoncd/pipeline/pkg/reconciler/volumeclaim"
	"knative.dev/pkg/apis"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
	"strings"
	"time"
)

type ThrottledByPVCQuotaCollector struct {
	pvcThrottle *prometheus.GaugeVec
}

// Start - we do a long running runnable to reset the pvc metric in case we miss delete events, as controller relist does not duplicate
// delete events like it can create/update events
func (r *ExporterReconcile) Start(ctx context.Context) error {
	//this matches the scheduling interval for pruner in the operator's TektonConfig object
	//for now we are going to refrain from reading in the TektonConfig, bringing in a 3rd party
	// golang cronjob schedule parser, and pulling the value; but if we end up changing it with
	// some frequency, we'll start doing that
	eventTicker := time.NewTicker(2 * time.Minute)
	for {
		select {
		case <-eventTicker.C:
			r.resetPVCStats(ctx)
		case <-ctx.Done():
			controllerLog.Info("ReconcilePVCThrottled Runnable context is marked as done, exiting")
			eventTicker.Stop()
			return nil
		}
	}
	return nil
}

func failedBecauseOfPVCQuota(pr *v1.PipelineRun) bool {
	c := pr.GetStatusCondition().GetCondition(apis.ConditionSucceeded)
	if c == nil {
		return false
	}
	if !c.IsFalse() {
		return false
	}
	if c.Reason != volumeclaim.ReasonCouldntCreateWorkspacePVC || !strings.Contains(c.Message, "exceeded quota") {
		return false
	}
	return true
}

func (r *ExporterReconcile) resetPVCStats(ctx context.Context) {
	// originally considered using pvcThrottle.Reset() but wanted to allow for history based searches from metrics
	// console, so we are trying keeping track of namespaces; for now, not worried about history across exporter restart
	for ns := range r.nsCache {
		r.pvcCollector.zeroPVCThrottle(ns)
	}
	// however, we'll clear out cache to avoid long term accumulation, memory leak, as things like dynamically created test clusters
	// accumulate; as long as we maintain history for permanent, active tenant namespaces, that is OK
	r.nsCache = map[string]struct{}{}

	prList := &v1.PipelineRunList{}
	err := r.client.List(ctx, prList)
	nsWithPVCThrottle := map[string]struct{}{}
	if err == nil {
		for _, pr := range prList.Items {
			r.nsCache[pr.Namespace] = struct{}{}
			if failedBecauseOfPVCQuota(&pr) {
				r.pvcCollector.incPVCThrottle(&pr)
				nsWithPVCThrottle[pr.Namespace] = struct{}{}
				continue
			}
			// in case this is a namespace we did not see in prior invocations of resetPVCStats,
			// we want to get explicit 0 counts if there is not any PVC throttling for a namespace,
			// but we make sure we did not increment it previously in this loop (that is easier/cheaper then getting the metric and then
			// hydrating the value like we do in our unit tests), so we set to 0
			_, ok := nsWithPVCThrottle[pr.Namespace]
			if ok {
				continue
			}
			r.pvcCollector.zeroPVCThrottle(pr.Namespace)
		}
	}
}

type pvcThrottledFilter struct {
}

func (f *pvcThrottledFilter) Create(event.CreateEvent) bool {
	return false
}

func (f *pvcThrottledFilter) Delete(e event.DeleteEvent) bool {
	return false
}

func (f *pvcThrottledFilter) Update(e event.UpdateEvent) bool {

	oldPR, okold := e.ObjectOld.(*v1.PipelineRun)
	newPR, oknew := e.ObjectNew.(*v1.PipelineRun)
	if okold && oknew {
		if !failedBecauseOfPVCQuota(oldPR) && failedBecauseOfPVCQuota(newPR) {
			return true
		}
	}
	return false
}

func (f *pvcThrottledFilter) Generic(event.GenericEvent) bool {
	return false
}

func NewPVCThrottledCollector() *ThrottledByPVCQuotaCollector {
	labelNames := []string{NS_LABEL}
	pvcThrottled := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pipelinerun_failed_by_pvc_quota_count",
		Help: "Number of PipelineRuns who were marked failed because PVC Resource Quotas prevented the creation of required PVCs",
	}, labelNames)
	pvcThrottledCollector := &ThrottledByPVCQuotaCollector{
		pvcThrottle: pvcThrottled,
	}
	metrics.Registry.MustRegister(pvcThrottled)
	return pvcThrottledCollector
}

func (c *ThrottledByPVCQuotaCollector) incPVCThrottle(pr *v1.PipelineRun) {
	labels := map[string]string{NS_LABEL: pr.Namespace}
	c.pvcThrottle.With(labels).Inc()
}

func (c *ThrottledByPVCQuotaCollector) zeroPVCThrottle(ns string) {
	labels := map[string]string{NS_LABEL: ns}
	c.pvcThrottle.With(labels).Set(float64(0))
}
