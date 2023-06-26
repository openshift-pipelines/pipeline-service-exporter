package collector

import (
	"context"
	"time"

	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	k8sscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/cache"

	pipelinev1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
)

var (
	controllerLog = ctrl.Log.WithName("controller")
)

func NewManager(cfg *rest.Config, options ctrl.Options) (ctrl.Manager, error) {
	// we have seen in testing that this path can get invoked prior to the PipelineRun CRD getting generated,
	// and controller-runtime does not retry on missing CRDs.
	// so we are going to wait on the CRDs existing before moving forward.
	apiextensionsClient := apiextensionsclient.NewForConfigOrDie(cfg)
	if err := wait.PollImmediate(time.Second*5, time.Minute*5, func() (done bool, err error) {
		_, err = apiextensionsClient.ApiextensionsV1().CustomResourceDefinitions().Get(context.TODO(), "pipelineruns.tekton.dev", metav1.GetOptions{})
		if err != nil {
			controllerLog.Error(err, "get of pipelinerun CRD failed")
			return false, nil
		}
		controllerLog.Info("get of pipelinerun CRD returned successfully")
		return true, nil
	}); err != nil {
		controllerLog.Error(err, "waiting for pipelinerun CRD to be created")
		return nil, err
	}

	options.Scheme = runtime.NewScheme()
	if err := k8sscheme.AddToScheme(options.Scheme); err != nil {
		return nil, err
	}
	//TODO v1 tekton API is coming soon
	if err := pipelinev1beta1.AddToScheme(options.Scheme); err != nil {
		return nil, err
	}
	options.NewCache = cache.BuilderWithOptions(cache.Options{
		SelectorsByObject: cache.SelectorsByObject{
			&pipelinev1beta1.PipelineRun{}: {},
		}})

	var mgr ctrl.Manager
	var err error
	mgr, err = ctrl.NewManager(cfg, options)
	if err != nil {
		return nil, err
	}

	err = SetupPipelineRunScheduleDurationController(mgr)
	if err != nil {
		return nil, err
	}

	return mgr, nil
}
