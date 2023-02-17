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
	"github.com/go-kit/log"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
)

// logger initialized from go-kit/log
var logger log.Logger

func calculateScheduledDuration(pipelineRun v1beta1.PipelineRun) float64 {
	var durationScheduled float64

	// Fetch the creation and scheduled times
	createdTime := pipelineRun.ObjectMeta.CreationTimestamp.Time
	startTime := pipelineRun.Status.StartTime.Time

	// Check if either one of these values is zero
	if createdTime.IsZero() || startTime.IsZero() {
		return 0
	}

	durationScheduled = startTime.Sub(createdTime).Seconds()
	return durationScheduled
}

func calculateCompletedDuration(pipelineRun v1beta1.PipelineRun) float64 {
	var timeCompleted float64

	// Fetch the scheduled and completion times
	startTime := pipelineRun.Status.StartTime.Time
	completionTime := pipelineRun.Status.CompletionTime.Time

	// Check if completionTime is zero
	if completionTime.IsZero() {
		return 0
	}

	timeCompleted = completionTime.Sub(startTime).Seconds()
	return timeCompleted
}
