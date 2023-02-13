/*
 Copyright 2023 The Pipeline Service Authors.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package collector

import (
	"context"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	tektonclient "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
)

// logger initialized from go-kit/log
var logger log.Logger

// getPipelineRuns returns the list of PipelineRuns
func getPipelineRuns() (*v1beta1.PipelineRunList, error) {
	prs := &v1beta1.PipelineRunList{
		Items: []v1beta1.PipelineRun{},
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

	prs, err = tektonClient.TektonV1beta1().PipelineRuns("").List(context.Background(), metav1.ListOptions{})
	if len(prs.Items) == 0 {
		return prs, err
	}
	return prs, nil
}

func calculateScheduledDuration(pipelineRun v1beta1.PipelineRun) (float64, error) {
	var durationScheduled float64

	// Fetch the creation and scheduled times
	createdTime := pipelineRun.ObjectMeta.CreationTimestamp.Time
	startTime := pipelineRun.Status.StartTime.Time

	// Check if either one of these values is zero
	if createdTime.IsZero() || startTime.IsZero() {
		return 0, level.Error(logger).Log("msg", "could not calculate scheduled duration, as creation time or scheduled time is not set.")
	}

	durationScheduled = startTime.Sub(createdTime).Seconds()
	return durationScheduled, nil
}

func calculateCompletedDuration(pipelineRun v1beta1.PipelineRun) (float64, error) {
	var timeCompleted float64

	// Fetch the scheduled and completion times
	startTime := pipelineRun.Status.StartTime.Time
	completionTime := pipelineRun.Status.CompletionTime.Time

	// Check if either one of these values is zero
	if completionTime.IsZero() {
		return 0, level.Error(logger).Log("msg", "pipelineRun is not completed yet.")
	}

	timeCompleted = completionTime.Sub(startTime).Seconds()
	return timeCompleted, nil
}
