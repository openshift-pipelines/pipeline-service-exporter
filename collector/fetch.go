package collector

import (
	"context"

	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (c *PipelineServiceCollector) getPipelineRuns() ([]v1beta1.PipelineRun, error) {
	//NOTE:  there is no need to filter queries based on counts, worried about pagination, like with the old non-caching, polling client implementation;
	// if over time we think we need to filter at all, it should be based on labels ...appstudio based labels for example:
	//"appstudio.openshift.io/pac-provision"
	//"appstudio.openshift.io/application"
	//"appstudio.openshift.io/component"
	// like is seen with the build-service; we can also keep an eye on how the core tekton pipelines controller is tuned
	prList := v1beta1.PipelineRunList{}
	opts := &client.ListOptions{Namespace: metav1.NamespaceAll}
	err := c.client.List(context.Background(), &prList, opts)
	return prList.Items, err
}
