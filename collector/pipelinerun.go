package collector

import (
	"context"
	"github.com/go-kit/log/level"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

func SetupPipelineRunCachingClient(mgr ctrl.Manager) {
	// we do not set up a reconciler to get a caching client, but we make one list
	// call up top to get the ball rolling; confirmed via break points in a debugger that
	// k8s methods for creating a shared informer and watch are called
	prList := v1beta1.PipelineRunList{}
	opts := &client.ListOptions{Namespace: metav1.NamespaceAll}
	mgr.GetClient().List(context.Background(), &prList, opts)
	psCollector, err := NewCollector(logger, mgr.GetClient())
	if err != nil {
		level.Error(logger).Log("msg", "Couldn't create collector", "error", err)
		os.Exit(1)
	}
	metrics.Registry.MustRegister(psCollector)
}
