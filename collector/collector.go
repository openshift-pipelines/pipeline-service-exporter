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
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sync"
	"time"
)

// PipelineServiceCollector struct
type PipelineServiceCollector struct {
	logger            log.Logger
	mutex             sync.Mutex
	durationScheduled *prometheus.GaugeVec
	durationCompleted *prometheus.GaugeVec
	pipelineRuns      map[string]*v1beta1.PipelineRun
}

var completedDuration float64
var scheduledDuration float64

func NewCollector(logger log.Logger) (*PipelineServiceCollector, error) {
	durationScheduled := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pipelinerun_duration_scheduled_seconds",
		Help: "Duration in seconds for a PipelineRun to be scheduled.",
	}, []string{"name", "uid"})

	durationCompleted := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pipelinerun_duration_completed_seconds",
		Help: "Duration in seconds for a PipelineRun to complete.",
	}, []string{"name", "uid"})

	pipelineRuns := make(map[string]*v1beta1.PipelineRun)

	pipelineServiceCollector := &PipelineServiceCollector{
		logger:            logger,
		durationScheduled: durationScheduled,
		durationCompleted: durationCompleted,
		pipelineRuns:      pipelineRuns,
	}

	// Start a goroutine to fetch the PipelineRuns in the background
	go pipelineServiceCollector.fetchPipelineRuns()

	return pipelineServiceCollector, nil
}

func (c *PipelineServiceCollector) fetchPipelineRuns() {
	for {
		prs, err := c.getPipelineRuns()
		if err != nil {
			level.Error(c.logger).Log("msg", "Error fetching PipelineRuns:", "error", err)

		} else {
			// Update the pipelineRuns field with the new set of PipelineRuns
			c.mutex.Lock()

			for i := range prs {
				pr := &prs[i]
				// Check if a PipelineRun with the same name already exists in the map
				if existingPR, ok := c.pipelineRuns[(*pr).Name]; ok {
					// If it exists, check if the existing one is newer
					if existingPR.CreationTimestamp.Time.After((*pr).CreationTimestamp.Time) {
						continue
					}
				}

				// Below code block helps in keeping the length of processing PipelineRuns at 500
				// to keep the memory fixed.
				// Remove the oldest item if the number of PipelineRuns exceeds the limit
				if len(c.pipelineRuns) >= 500 {
					oldestName := ""
					var oldestTime v1.Time
					for name, pr := range c.pipelineRuns {
						if oldestName == "" || pr.CreationTimestamp.Before(&oldestTime) {
							oldestName = name
							oldestTime = pr.CreationTimestamp
						}
					}
					delete(c.pipelineRuns, oldestName)
				}
				c.pipelineRuns[(*pr).Name] = *pr
			}
			c.mutex.Unlock()
		}

		// Wait for some time before fetching the PipelineRuns again
		time.Sleep(10 * time.Second)
	}
}

// Describe implements the prometheus.Collector interface
func (c *PipelineServiceCollector) Describe(ch chan<- *prometheus.Desc) {
	c.durationScheduled.Describe(ch)
	c.durationCompleted.Describe(ch)
}

// Collect implements the prometheus.Collector interface
func (c *PipelineServiceCollector) Collect(ch chan<- prometheus.Metric) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if err := c.collect(ch); err != nil {
		level.Error(c.logger).Log("msg", "error collecting pipeline-service metrics", "error", err)
	}
}

// collect implements prometheus.Collector interface
func (c *PipelineServiceCollector) collect(ch chan<- prometheus.Metric) error {

	for _, pipelineRun := range c.pipelineRuns {
		// Fetch and compute the metrics for schedule and completed time
		if pipelineRun.Status.StartTime != nil {
			scheduledDuration = calculateScheduledDuration(*pipelineRun)
		} else {
			continue
		}

		if pipelineRun.Status.CompletionTime != nil {
			completedDuration = calculateCompletedDuration(*pipelineRun)
		} else {
			continue
		}

		// Set the metrics
		c.durationScheduled.WithLabelValues(pipelineRun.Name, string(pipelineRun.UID)).Set(scheduledDuration)
		c.durationCompleted.WithLabelValues(pipelineRun.Name, string(pipelineRun.UID)).Set(completedDuration)
	}

	// Make sure it is passed to the channel so that it is exported out
	c.durationScheduled.Collect(ch)
	c.durationCompleted.Collect(ch)

	return nil
}
