package collector

import (
	"context"
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
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"time"
)

var (
	controllerLog = ctrl.Log.WithName("controller")
)

func NewManager(cfg *rest.Config, options ctrl.Options) (ctrl.Manager, error) {
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

	err = SetupPipelineRunScheduleDurationController(mgr)
	if err != nil {
		return nil, err
	}

	err = SetupPipelineRunTaskRunGapController(mgr)
	if err != nil {
		return nil, err
	}

	err = SetupTaskRunScheduleDurationController(mgr)
	if err != nil {
		return nil, err
	}

	err = SetupPVCThrottledController(mgr)
	if err != nil {
		return nil, err
	}

	err = SetupPodCreateToKubeletDurationController(mgr)
	if err != nil {
		return nil, err
	}

	err = SetupPodKubeletToContainerStartDurationController(mgr)
	if err != nil {
		return nil, err
	}

	err = SetupTaskReferenceWaitTimeController(mgr)
	if err != nil {
		return nil, err
	}

	err = SetupPipelineReferenceWaitTimeController(mgr)
	if err != nil {
		return nil, err
	}

	err = SetupPodCreateToCompleteTimeController(mgr)
	if err != nil {
		return nil, err
	}

	err = SetupOverheadController(mgr)
	if err != nil {
		return nil, err
	}

	return mgr, nil
}
