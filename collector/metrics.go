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
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"time"
)

func calcuateScheduledDuration(created, started time.Time) float64 {
	if created.IsZero() || started.IsZero() {
		return 0
	}
	return started.Sub(created).Seconds()
}

func calculateScheduledDurationPipelineRun(pipelineRun *v1beta1.PipelineRun) float64 {
	return calcuateScheduledDuration(pipelineRun.CreationTimestamp.Time, pipelineRun.Status.StartTime.Time)
}

func calculateScheduledDurationTaskRun(taskrun *v1beta1.TaskRun) float64 {
	return calcuateScheduledDuration(taskrun.CreationTimestamp.Time, taskrun.Status.StartTime.Time)
}
