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
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// PipelineServiceCollector struct
type PipelineServiceCollector struct {
	logger            log.Logger
	durationScheduled *prometheus.GaugeVec
	durationCompleted *prometheus.GaugeVec

	client client.Client
}

func NewCollector(logger log.Logger, client client.Client) (*PipelineServiceCollector, error) {
	//TODO should this be converted to a Desc so we can use constant metrics
	durationScheduled := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pipelinerun_duration_scheduled_seconds",
		Help: "Duration in seconds for a PipelineRun to be scheduled.",
		//TODO the uid will get us a unique entry, but wouldn't adding namespace will help usability when consuming this metric?
		// also, labels at an individual PipelineRun is a scalability issue with prometheus
	}, []string{"name", "uid"})

	//TODO should this be converted to a Desc so we can use constant metrics
	durationCompleted := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pipelinerun_duration_completed_seconds",
		Help: "Duration in seconds for a PipelineRun to complete.",
		//TODO the uid will get us a unique entry, but wouldn't adding namespace help usability with humans consuming this metric?
	}, []string{"name", "uid"})

	pipelineServiceCollector := &PipelineServiceCollector{
		logger:            logger,
		durationScheduled: durationScheduled,
		durationCompleted: durationCompleted,
	}

	pipelineServiceCollector.client = client

	return pipelineServiceCollector, nil
}

// Describe implements the prometheus.Collector interface
func (c *PipelineServiceCollector) Describe(ch chan<- *prometheus.Desc) {
	c.durationScheduled.Describe(ch)
	c.durationCompleted.Describe(ch)
}

// Collect implements the prometheus.Collector interface
func (c *PipelineServiceCollector) Collect(ch chan<- prometheus.Metric) {
	if err := c.collect(ch); err != nil {
		level.Error(c.logger).Log("msg", "error collecting pipeline-service metrics", "error", err)
	}
}

// collect implements prometheus.Collector interface
func (c *PipelineServiceCollector) collect(ch chan<- prometheus.Metric) error {
	prs, err := c.getPipelineRuns()
	if err != nil {
		return err
	}
	for _, pipelineRun := range prs {
		var completedDuration float64
		var scheduledDuration float64

		// Fetch and compute the metrics for schedule and completed time
		if pipelineRun.Status.StartTime != nil {
			scheduledDuration = calculateScheduledDuration(pipelineRun)
		} else {
			continue
		}

		if pipelineRun.Status.CompletionTime != nil {
			completedDuration = calculateCompletedDuration(pipelineRun)
		} else {
			continue
		}

		// Set the metrics
		//TODO should we switch to constant metrics a la "ch <- prometheus.MustNewConstMetric(desc, prometheus.GaugeValue, float64(v), name, uid)"
		c.durationScheduled.WithLabelValues(pipelineRun.Name, string(pipelineRun.UID)).Set(scheduledDuration)
		c.durationCompleted.WithLabelValues(pipelineRun.Name, string(pipelineRun.UID)).Set(completedDuration)
	}

	// Make sure it is passed to the channel so that it is exported out
	//TODO by switching to a const metric and employing "ch <- prometheus.MustNewConstMetric(desc, prometheus.GaugeValue, float64(v), name, uid)" we would no longer need this
	c.durationScheduled.Collect(ch)
	c.durationCompleted.Collect(ch)

	return nil
}
