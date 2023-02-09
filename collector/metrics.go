package collector

import (
	"context"
	"fmt"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	tektonclient "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

// getPipelineRuns returns the list of PipelineRuns
func getPipelineRuns() (*v1beta1.PipelineRunList, error) {
	prs := &v1beta1.PipelineRunList{
		Items: []v1beta1.PipelineRun{},
	}

	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	//config, err := clientcmd.BuildConfigFromFlags("", "")

	tektonClient, err := tektonclient.NewForConfig(config)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	prs, err = tektonClient.TektonV1beta1().PipelineRuns("default").List(context.Background(), metav1.ListOptions{})
	if len(prs.Items) == 0 {
		return prs, err
	}
	return prs, nil
}

func calculateScheduledDuration(pipelineRun v1beta1.PipelineRun) (float64, error) {
	fmt.Println("\n\n Inside calculateScheduledDuration")
	var durationScheduled float64

	// Fetch the creation and scheduled times
	createdTime := pipelineRun.ObjectMeta.CreationTimestamp.Time
	startTime := pipelineRun.Status.StartTime.Time

	fmt.Println("createdTime: ", createdTime)
	fmt.Println("startTime: ", startTime)

	// Check if either one of these values is zero
	if createdTime.IsZero() || startTime.IsZero() {
		return 0, fmt.Errorf("could not calculate scheduled duration, as creation time or scheduled time is not set")
	}

	durationScheduled = startTime.Sub(createdTime).Seconds()

	return durationScheduled, nil
}

func calculateCompletedDuration(pipelineRun v1beta1.PipelineRun) (float64, error) {
	fmt.Println("\n\n Inside calculateCompletedDuration")
	var timeCompleted float64

	// Fetch the scheduled and completion times
	startTime := pipelineRun.Status.StartTime.Time
	completionTime := pipelineRun.Status.CompletionTime.Time

	fmt.Println("startTime: ", startTime)
	fmt.Println("completionTime: ", completionTime)

	// Check if either one of these values is zero
	if completionTime.IsZero() {
		return 0, fmt.Errorf("PipelineRun has not completed yet")
	}

	timeCompleted = completionTime.Sub(startTime).Seconds()

	return timeCompleted, nil
}
