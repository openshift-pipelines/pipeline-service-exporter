package collector

import (
	"context"
	"github.com/go-kit/log/level"
	v1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	tektonclient "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

func (c *PipelineServiceCollector) getPipelineRuns() ([]*v1beta1.PipelineRun, error) {
	var pipelineRuns []*v1beta1.PipelineRun
	limit := int64(100)
	listOptions := metav1.ListOptions{
		Limit: limit,
	}

	config, err := rest.InClusterConfig()
	if err != nil {
		level.Error(logger).Log("msg", "error reading InClusterConfig", "error", err)
		return nil, err
	}

	tektonClient, err := tektonclient.NewForConfig(config)
	if err != nil {
		level.Error(logger).Log("msg", "error creating a tektonClient", "error", err)
		return nil, err
	}

	for {
		prs, err := tektonClient.TektonV1beta1().PipelineRuns("").List(context.Background(), listOptions)
		if err != nil {
			return nil, err
		}

		pipelineRunItems := prs.Items

		for i := range pipelineRunItems {
			pipelineRuns = append(pipelineRuns, &pipelineRunItems[i])
		}

		if prs.Continue == "" {
			break
		}
		listOptions.Continue = prs.Continue
	}

	return pipelineRuns, nil
}
