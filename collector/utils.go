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
	"github.com/tektoncd/pipeline/pkg/apis/pipeline"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	ENABLE_GAP_METRIC_ADDITIONAL_LABELS = "ENABLE_GAP_METRIC_ADDITIONAL_LABELS"
	NS_LABEL                            = "namespace"
	PIPELINE_NAME_LABEL                 = "pipelinename"
	TASK_NAME_LABEL                     = "taskname"
	COMPLETED_LABEL                     = "completed"
	UPCOMING_LABEL                      = "upcoming"
	STATUS_LABEL                        = "status"
	SUCCEEDED                           = "succeded"
	FAILED                              = "failed"
)

func pipelineRunPipelineRef(pr *v1beta1.PipelineRun) string {
	val := ""
	ref := pr.Spec.PipelineRef
	if ref != nil {
		val = ref.Name
		if len(val) == 0 {
			for _, p := range ref.Params {
				if strings.TrimSpace(p.Name) == "name" {
					return p.Value.StringVal
				}
			}
		}
	} else {
		if len(pr.GenerateName) > 0 {
			return pr.GenerateName
		}
		// at this point, the pipelinerun name should not have any random aspects, and is
		// essentially a constant, with a minimal cardinality impact
		val = pr.Name
	}
	return val
}

func taskRef(labels map[string]string) string {
	task, _ := labels[pipeline.TaskLabelKey]
	pipelineTask, _ := labels[pipeline.PipelineTaskLabelKey]
	clusterTask, _ := labels[pipeline.ClusterTaskLabelKey]
	taskRun, _ := labels[pipeline.TaskRunLabelKey]
	switch {
	case len(task) > 0:
		return task
	case len(pipelineTask) > 0:
		return pipelineTask
	case len(clusterTask) > 0:
		return clusterTask
	case len(taskRun) > 0:
		return taskRun
	}
	return ""
}

func optionalMetricEnabled(envVarName string) bool {
	env := os.Getenv(envVarName)
	enabled := len(env) > 0
	// any random setting means true
	if enabled {
		// allow for users to turn off by setting to false
		bv, err := strconv.ParseBool(env)
		if err == nil && !bv {
			return false
		}
		return true
	}
	return false
}

func calcuateScheduledDuration(created, started time.Time) float64 {
	if created.IsZero() || started.IsZero() {
		return 0
	}
	return started.Sub(created).Seconds()
}
