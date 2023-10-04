package collector

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"github.com/tektoncd/pipeline/pkg/reconciler/volumeclaim"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"knative.dev/pkg/apis"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"strings"
	"sync"
	"time"
)

func SetupPVCThrottledController(mgr ctrl.Manager) error {
	reconciler := &ReconcilePVCThrottled{
		client:        mgr.GetClient(),
		scheme:        mgr.GetScheme(),
		eventRecorder: mgr.GetEventRecorderFor("MetricExporterPVCThrottled"),
		prCollector:   NewPVCThrottledCollector(),
		nsCache:       map[string]struct{}{},
		listMutex:     sync.RWMutex{},
	}
	err := ctrl.NewControllerManagedBy(mgr).For(&v1beta1.PipelineRun{}).WithEventFilter(&pvcThrottledFilter{}).Complete(reconciler)
	if err == nil {
		mgr.Add(reconciler)
	}
	return err
}

type ThrottledByPVCQuotaCollector struct {
	pvcThrottle *prometheus.GaugeVec
}

type ReconcilePVCThrottled struct {
	client        client.Client
	scheme        *runtime.Scheme
	eventRecorder record.EventRecorder
	prCollector   *ThrottledByPVCQuotaCollector
	nsCache       map[string]struct{}
	listMutex     sync.RWMutex
}

func (r *ReconcilePVCThrottled) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()
	log := log.FromContext(ctx)

	// so we don't collide with the periodic relist/reset in Start; we should still be able to process events concurrently
	r.listMutex.RLock()
	defer r.listMutex.RUnlock()

	//TODO remember, keep track of when pipeline-service and RHTAP starts moving from v1beta1 to v1
	pr := &v1beta1.PipelineRun{}
	err := r.client.Get(ctx, types.NamespacedName{Namespace: request.Namespace, Name: request.Name}, pr)
	if err != nil && !errors.IsNotFound(err) {
		return reconcile.Result{}, err
	}
	if err != nil {
		// given the various complexities around deletion processing and controllers, we do not decrement our
		// metric in real time, but rather reset the metrics in our background poller attuned to the operator's pruner
		// interval.
		log.V(4).Info(fmt.Sprintf("processing deleted pipelinerun %q", request.NamespacedName))
		return reconcile.Result{}, nil
	}

	// based on our WithEventFilter we should only be getting called if this got throttled by PVC
	log.V(4).Info(fmt.Sprintf("recording pvc throttling for %q", request.NamespacedName))
	r.prCollector.incPVCThrottle(pr)
	return reconcile.Result{}, nil
}

// Start - we do a long running runnable to reset the metric in case we miss delete events, as controller relist does not duplicate
// delete events like it can create/update events
func (r *ReconcilePVCThrottled) Start(ctx context.Context) error {
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
			eventTicker.Stop()
			return nil
		}
	}
	return nil
}

func failedBecauseOfPVCQuota(pr *v1beta1.PipelineRun) bool {
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

func (r *ReconcilePVCThrottled) resetPVCStats(ctx context.Context) {
	r.listMutex.Lock()
	defer r.listMutex.Unlock()
	// originally considered using pvcThrottle.Reset() but wanted to allow for history based searches from metrics
	// console, so we are trying keeping track of namespaces; for now, not worried about history across exporter restart
	for ns := range r.nsCache {
		r.prCollector.zeroPVCThrottle(ns)
	}
	prList := &v1beta1.PipelineRunList{}
	err := r.client.List(ctx, prList)
	nsWithPVCThrottle := map[string]struct{}{}
	if err == nil {
		for _, pr := range prList.Items {
			r.nsCache[pr.Namespace] = struct{}{}
			if failedBecauseOfPVCQuota(&pr) {
				r.prCollector.incPVCThrottle(&pr)
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
			r.prCollector.zeroPVCThrottle(pr.Namespace)
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

	//TODO remember, keep track of when pipeline-service and RHTAP starts moving from v1beta1 to v1
	oldPR, okold := e.ObjectOld.(*v1beta1.PipelineRun)
	newPR, oknew := e.ObjectNew.(*v1beta1.PipelineRun)
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

func (c *ThrottledByPVCQuotaCollector) incPVCThrottle(pr *v1beta1.PipelineRun) {
	labels := map[string]string{NS_LABEL: pr.Namespace}
	c.pvcThrottle.With(labels).Inc()
}

func (c *ThrottledByPVCQuotaCollector) zeroPVCThrottle(ns string) {
	labels := map[string]string{NS_LABEL: ns}
	c.pvcThrottle.With(labels).Set(float64(0))
}
