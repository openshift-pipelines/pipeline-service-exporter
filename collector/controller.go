package collector

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline"
	pipelinev1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	pipelinev1client "github.com/tektoncd/pipeline/pkg/client/clientset/versioned/typed/pipeline/v1"
	corev1 "k8s.io/api/core/v1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/apimachinery/pkg/util/wait"
	k8sscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/record"
	"net/http"
	_ "net/http/pprof"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"time"
)

var (
	controllerLog = ctrl.Log.WithName("controller")
)

type PollCollector interface {
	IncCollector(ns string)
	ZeroCollector(ns string)
}

func NewManager(cfg *rest.Config, options ctrl.Options, pprofPort string) (ctrl.Manager, error) {
	// we have seen in testing that this path can get invoked prior to the PipelineRun CRD getting generated,
	// and controller-runtime does not retry on missing CRDs.
	// so we are going to wait on the CRDs existing before moving forward.
	apiextensionsClient := apiextensionsclient.NewForConfigOrDie(cfg)
	pipelineClient := pipelinev1client.NewForConfigOrDie(cfg)
	if err := wait.PollImmediate(time.Second*5, time.Minute*5, func() (done bool, err error) {
		_, err = apiextensionsClient.ApiextensionsV1().CustomResourceDefinitions().Get(context.TODO(), "pipelineruns.tekton.dev", metav1.GetOptions{})
		if err != nil {
			controllerLog.Error(err, "get of pipelinerun CRD failed")
			return false, nil
		}
		controllerLog.Info("get of pipelinerun CRD returned successfully")
		// in addition to the CRD check we've got in several controller-runtime based RHTAP controllers, metrics-exporter
		// recently saw some intermittent issues even after this when setting up of watches or lists timed out as tekton
		// was still ramping up, and controller runtime would exit out of initialization.  For example:
		// "Failed to watch *v1.TaskRun: the server is currently unable to handle the request (get taskruns.tekton.dev)"
		// "Failed to watch *v1.PipelineRun: the server is currently unable to handle the request (get pipelineruns.tekton.dev)"
		// "failed to list *v1.TaskRun: the server was unable to return a response in the time allotted, but may still be processing the request (get taskruns.tekton.dev)"
		// "Failed to watch *v1.TaskRun: failed to list *v1.TaskRun: the server was unable to return a response in the time allotted, but may still be processing the request (get taskruns.tekton.dev)"
		//
		// So we now try to see a list return successfully before we move on to controller-runtime initialization
		_, err = pipelineClient.PipelineRuns("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			controllerLog.Error(err, "list of pipelineruns failed")
			return false, nil
		}
		controllerLog.Info("list of pipelineruns returned successfully")
		return true, nil
	}); err != nil {
		controllerLog.Error(err, "waiting for pipelinerun CRD to be created")
		return nil, err
	}

	options.Scheme = runtime.NewScheme()
	if err := k8sscheme.AddToScheme(options.Scheme); err != nil {
		return nil, err
	}
	if err := pipelinev1.AddToScheme(options.Scheme); err != nil {
		return nil, err
	}

	var mgr ctrl.Manager
	var err error
	var labelReq *labels.Requirement
	// only get/watch/cache pods with the tekton pipeline label
	labelReq, err = labels.NewRequirement(pipeline.PipelineLabelKey, selection.Exists, []string{})
	if err != nil {
		return nil, err
	}
	podSelector := labels.NewSelector().Add(*labelReq)
	selectors := cache.SelectorsByObject{
		&pipelinev1.PipelineRun{}: {},
		&pipelinev1.TaskRun{}:     {},
		&corev1.Pod{}: cache.ObjectSelector{
			Label: podSelector,
		},
	}
	cacheOptions := cache.Options{SelectorsByObject: selectors}
	options.NewCache = cache.BuilderWithOptions(cacheOptions)

	mgr, err = ctrl.NewManager(cfg, options)
	if err != nil {
		return nil, err
	}

	err = SetupController(mgr, pprofPort)

	return mgr, nil
}

type pprof struct {
	port string
}

func (p *pprof) Start(ctx context.Context) error {
	srv := &http.Server{Addr: ":" + p.port}
	controllerLog.Info(fmt.Sprintf("starting ppprof on %s", p.port))
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			controllerLog.Info(fmt.Sprintf("pprof server err: %s", err.Error()))
		}
	}()
	<-ctx.Done()
	controllerLog.Info("Shutting down pprof")
	srv.Shutdown(ctx)
	return nil
}

func SetupController(mgr ctrl.Manager, pprofPort string) error {
	exportFilter := &ExporterFilter{
		noReconcile:  []predicate.Predicate{},
		yesReconcile: []predicate.Predicate{},
	}

	// yesReconcile are metrics with non-empty Reconcile methods
	exportFilter.yesReconcile = append(exportFilter.yesReconcile, &overheadGapEventFilter{client: mgr.GetClient()})
	exportFilter.yesReconcile = append(exportFilter.yesReconcile, &taskRunGapEventFilter{})

	// noReconcile are metrics with empty Reconcile methods
	exportFilter.noReconcile = append(exportFilter.noReconcile, &pipelineRefWaitTimeFilter{waitDuration: NewPipelineReferenceWaitTimeMetric()})
	exportFilter.noReconcile = append(exportFilter.noReconcile, &startTimeEventFilter{metric: NewPipelineRunScheduledMetric()})
	exportFilter.noReconcile = append(exportFilter.noReconcile, NewPodCreateToCompleteFilter())
	exportFilter.noReconcile = append(exportFilter.noReconcile, &createKubeletLatencyFilter{metric: NewPodCreateToKubeletDurationMetric()})
	exportFilter.noReconcile = append(exportFilter.noReconcile, &kubeletContainerLatencyFilter{metric: NewPodKubeletToContainerStartDurationMetric()})
	exportFilter.noReconcile = append(exportFilter.noReconcile, &taskRefWaitTimeFilter{waitDuration: NewTaskReferenceWaitTimeMetric()})
	exportFilter.noReconcile = append(exportFilter.noReconcile, &trStartTimeEventFilter{metric: NewTaskRunScheduledMetric()})

	r := buildReconciler(mgr.GetClient(), mgr.GetScheme(), mgr.GetEventRecorderFor("MetricsExporter"))

	err := ctrl.NewControllerManagedBy(mgr).For(&pipelinev1.PipelineRun{}).
		WithOptions(controller.Options{MaxConcurrentReconciles: 32}).
		WithEventFilter(exportFilter).
		Complete(r)

	if err != nil {
		return err
	}

	err = mgr.Add(r)
	if err != nil {
		return err
	}
	if len(pprofPort) > 0 {
		pp := &pprof{port: pprofPort}
		err = mgr.Add(pp)
		if err != nil {
			return err
		}
	}

	err = ctrl.NewControllerManagedBy(mgr).For(&pipelinev1.TaskRun{}).
		WithOptions(controller.Options{MaxConcurrentReconciles: 32}).
		WithEventFilter(exportFilter).
		Complete(r)

	if err != nil {
		return err
	}

	err = ctrl.NewControllerManagedBy(mgr).For(&corev1.Pod{}).
		WithOptions(controller.Options{MaxConcurrentReconciles: 32}).
		WithEventFilter(exportFilter).
		Complete(r)

	if err != nil {
		return err
	}
	return nil
}

type ExporterFilter struct {
	noReconcile  []predicate.Predicate
	yesReconcile []predicate.Predicate
	pvcFilter    predicate.Predicate
}

func (f *ExporterFilter) Generic(event.GenericEvent) bool {
	return false
}

func (f *ExporterFilter) Create(event.CreateEvent) bool {
	return false
}

func (f *ExporterFilter) Delete(event.DeleteEvent) bool {
	return false
}

func (f ExporterFilter) Update(e event.UpdateEvent) bool {
	// yes, distinguishing between metrics which need to Reconcile and those which can fully be handled
	// in the filter's Update event is not required to determine if this Update implementation should return
	// 'true' and have controller runtime call Reconcile; however, I have found being explicit on this detail
	// is helpful and informative when working on this component.

	for _, p := range f.noReconcile {
		p.Update(e)
	}
	callReconcile := false
	for _, p := range f.yesReconcile {
		callReconcile = p.Update(e) || callReconcile
	}
	return callReconcile
}

type ExporterReconcile struct {
	client                            client.Client
	scheme                            *runtime.Scheme
	eventRecorder                     record.EventRecorder
	overheadCollector                 *OverheadCollector
	gapAdditionalLabels               bool
	prGapCollector                    *PipelineRunTaskRunGapCollector
	trGaps                            *prometheus.HistogramVec
	pvcNSCache                        map[string]struct{}
	waitPodNSCache                    map[string]map[string]struct{}
	waitPRKickoffCache                map[string]map[string]struct{}
	pvcCollector                      *ThrottledByPVCQuotaCollector
	waitPodCollector                  *WaitingOnPodCreateAttemptCollector
	waitPRKickoffCollector            *WaitingOnPipelineRunKickoffCollector
	podCreateNamespaceFilter          map[string]struct{}
	pipelineRunKickoffNamespaceFilter map[string]struct{}
}

func buildReconciler(client client.Client, scheme *runtime.Scheme, eventRecorder record.EventRecorder) *ExporterReconcile {
	prTrGapCollector := NewPipelineRunTaskRunGapCollector()
	r := &ExporterReconcile{
		client:                   client,
		scheme:                   scheme,
		eventRecorder:            eventRecorder,
		overheadCollector:        NewOverheadCollector(),
		gapAdditionalLabels:      prTrGapCollector.additionalLabels,
		prGapCollector:           prTrGapCollector,
		trGaps:                   prTrGapCollector.trGaps,
		pvcNSCache:               map[string]struct{}{},
		waitPodNSCache:           map[string]map[string]struct{}{},
		waitPRKickoffCache:       map[string]map[string]struct{}{},
		pvcCollector:             NewPVCThrottledCollector(),
		waitPodCollector:         NewWaitingOnPodCreateAttemptCollector(),
		waitPRKickoffCollector:   NewWaitingOnPipelineRunKickoffCollector(),
		podCreateNamespaceFilter: podCreateNameSpaceFilter(),
	}
	return r
}

func innerReset(collector PollCollector, nsCache []string) {
	// originally considered using pvcThrottle.Reset() but wanted to allow for history based searches from metrics
	// console, so we are trying keeping track of namespaces; for now, not worried about history across exporter restart
	for _, ns := range nsCache {
		collector.ZeroCollector(ns)
	}

}

func (r *ExporterReconcile) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	// replace with golang errors.Join(errs ...error) when we go to golang 1.20
	errorMsg := ""
	//TODO if we start providing something other than the empty Result object, we'll need to build a list and handle non-standard results
	result, err := r.ReconcileOverhead(ctx, request)
	if err != nil {
		errorMsg = fmt.Sprintf("%s\n%s", errorMsg, err.Error())
	}
	result, err = r.ReconcilePipelineRunTaskRunGap(ctx, request)
	if err != nil {
		errorMsg = fmt.Sprintf("%s\n%s", errorMsg, err.Error())
	}
	if len(errorMsg) > 0 {
		return result, fmt.Errorf("%s", errorMsg)
	}
	return result, nil
}

// Start - we do a long running runnable to reset the pvc metric in case we miss delete events, as controller relist does not duplicate
// delete events like it can create/update events
func (r *ExporterReconcile) Start(ctx context.Context) error {
	//this matches the scheduling interval for pruner in the operator's TektonConfig object
	//for now we are going to refrain from reading in the TektonConfig, bringing in a 3rd party
	// golang cronjob schedule parser, and pulling the value; but if we end up changing it with
	// some frequency, we'll start doing that
	// side note: the wait interval for the polling style metrics in core tekton is 30 seconds at last check
	eventTicker := time.NewTicker(2 * time.Minute)
	for {
		select {
		case <-eventTicker.C:
			r.resetPVCStats(ctx)
			r.resetPodCreateAttemptedStats(ctx)
			r.resetPipelineRunKickoffStats(ctx)
		case <-ctx.Done():
			controllerLog.Info("ReconcilePVCThrottled Runnable context is marked as done, exiting")
			eventTicker.Stop()
			return nil
		}
	}
	return nil
}
