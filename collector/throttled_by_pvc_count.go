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
)

type ThrottledByPVCQuotaCollector struct {
	pvcThrottle *prometheus.GaugeVec
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
	innerReset(r.pvcCollector, r.pvcNSCache)
	// however, we'll clear out cache to avoid long term accumulation, memory leak, as things like dynamically created test clusters
	// accumulate; as long as we maintain history for permanent, active tenant namespaces, that is OK
	r.pvcNSCache = map[string]struct{}{}

	prList := &v1.PipelineRunList{}
	err := r.client.List(ctx, prList)
	nsWithPVCThrottle := map[string]struct{}{}
	if err == nil {
		for _, pr := range prList.Items {
			r.pvcNSCache[pr.Namespace] = struct{}{}
			if failedBecauseOfPVCQuota(&pr) {
				r.pvcCollector.IncCollector(pr.Namespace)
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
			r.pvcCollector.ZeroCollector(pr.Namespace)
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

func (c *ThrottledByPVCQuotaCollector) IncCollector(ns string) {
	labels := map[string]string{NS_LABEL: ns}
	c.pvcThrottle.With(labels).Inc()
}

func (c *ThrottledByPVCQuotaCollector) ZeroCollector(ns string) {
	labels := map[string]string{NS_LABEL: ns}
	c.pvcThrottle.With(labels).Set(float64(0))
}
