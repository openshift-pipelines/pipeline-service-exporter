package collector

import (
	"context"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func SetupPipelineRunCachingClient(mgr ctrl.Manager) {
	// we do not set up a reconciler to get a caching client, but we make one list
	// call up top to get the ball rolling; confirmed via break points in a debugger that
	// k8s methods for creating a shared informer and watch are called
	prList := v1beta1.PipelineRunList{}
	opts := &client.ListOptions{Namespace: metav1.NamespaceAll}
	mgr.GetClient().List(context.Background(), &prList, opts)
}
