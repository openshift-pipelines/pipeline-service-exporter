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

// processedScheduledPipelineRuns contains the list of PipelineRuns which are already processed for calculating the schedule time
//var processedScheduledPipelineRuns = make(map[types.UID]bool)

func calculateScheduledDuration(pipelineRun v1beta1.PipelineRun) (float64, error) {
	fmt.Println("\n\n Inside calculateScheduledDuration")
	var durationScheduled float64

	// Fetch the creation and scheduled times
	createdTime := pipelineRun.ObjectMeta.CreationTimestamp.Time
	scheduledTime := pipelineRun.Status.StartTime.Time

	fmt.Println("createdTime: ", createdTime)
	fmt.Println("scheduledTime", scheduledTime)

	// Check if either one of these values is zero
	if createdTime.IsZero() || scheduledTime.IsZero() {
		return 0, fmt.Errorf("could not calculate scheduled duration, as creation time or scheduled time is not set")
	}

	//for k, v := range processedScheduledPipelineRuns {
	//	fmt.Printf("UID: %v, Processed: %v\n", k, v)
	//}

	// Check if we have already processed this pipeline run
	//if _, ok := processedScheduledPipelineRuns[pipelineRun.UID]; ok {
	//	return 0, nil
	//}

	// Mark this pipeline run as processed
	//processedScheduledPipelineRuns[pipelineRun.UID] = true
	durationScheduled = scheduledTime.Sub(createdTime).Seconds()

	return durationScheduled, nil
}

// processedCompletedPipelineRuns contains the list of PipelineRuns which are already processed for calculating the completion time
//var processedCompletedPipelineRuns = make(map[types.UID]bool)

func calculateCompletedDuration(pipelineRun v1beta1.PipelineRun) (float64, error) {
	fmt.Println("\n\n Inside calculateCompletedDuration")
	var durationCompleted float64

	// Fetch the scheduled and completion times
	scheduledTime := pipelineRun.Status.StartTime.Time
	completionTime := pipelineRun.Status.CompletionTime.Time

	fmt.Println("scheduledTime: ", scheduledTime)
	fmt.Println("completionTime: ", completionTime)

	// Check if either one of these values is zero
	if completionTime.IsZero() {
		return 0, fmt.Errorf("PipelineRun has not completed yet")
	}

	//for k, v := range processedCompletedPipelineRuns {
	//	fmt.Printf("UID: %v, Processed: %v\n", k, v)
	//}
	//fmt.Println("UID of the current PipelineRun: ", pipelineRun.UID)
	// Check if we have already processed this pipeline run
	//if _, ok := processedCompletedPipelineRuns[pipelineRun.UID]; ok {
	//	return 0, nil
	//}

	// Mark this pipeline run as processed
	//processedCompletedPipelineRuns[pipelineRun.UID] = true
	durationCompleted = completionTime.Sub(scheduledTime).Seconds()

	return durationCompleted, nil
}
