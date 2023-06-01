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
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	collectorLog = ctrl.Log.WithName("collector")
)

// PipelineServiceCollector struct
type PipelineServiceCollector struct {
	logger            log.Logger
	durationScheduled *prometheus.HistogramVec
	durationCompleted *prometheus.HistogramVec

	client client.Client
}

func NewCollector(logger log.Logger, client client.Client) (*PipelineServiceCollector, error) {
	//TODO should this be converted to a Desc so we can use constant metrics
	// shipwright establishment metrics buckets is []float64{0, 1, 2, 3, 5, 7, 10, 15, 20, 30}
	// shipwright ramp up metrics buckets is prometheus.LinearBuckets(0, 1, 10)
	durationScheduled := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "pipelinerun_duration_scheduled_seconds",
		Help:    "Duration in seconds for a PipelineRun to be scheduled.",
		Buckets: prometheus.LinearBuckets(0, 1, 10),
	}, []string{"namespace"})

	//TODO should this be converted to a Desc so we can use constant metrics
	// shipwright completion duration buckets is prometheus.LinearBuckets(50, 50, 10)
	//TODO how does this compare to the upstream tekton metric tekton_pipelines_controller_pipelinerun_duration_seconds_[bucket, sum, count] that is of type Histogram/LastValue(Gauge) and has pipelinerun name level cardinality
	durationCompleted := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "pipelinerun_duration_completed_seconds",
		Help:    "Duration in seconds for a PipelineRun to complete.",
		Buckets: prometheus.LinearBuckets(50, 50, 10),
	}, []string{"namespace"})

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
		//TODO should we switch to constant metrics a la "ch <- prometheus.MustNewConstMetric(desc, prometheus.GaugeValue, float64(v), name, uid)" ?
		msg := fmt.Sprintf("pipelinerun %s:%s schedule duration %v complete duration %v", pipelineRun.Namespace, pipelineRun.Name, scheduledDuration, completedDuration)
		collectorLog.V(4).Info(msg)
		c.durationScheduled.WithLabelValues(pipelineRun.Namespace).Observe(scheduledDuration)
		c.durationCompleted.WithLabelValues(pipelineRun.Namespace).Observe(completedDuration)
	}

	// Make sure it is passed to the channel so that it is exported out
	//TODO if we switched to a const metric, then the "ch <- prometheus.MustNewConstMetric(desc, prometheus.GaugeValue, float64(v), name, uid)" above would replace this
	c.durationScheduled.Collect(ch)
	c.durationCompleted.Collect(ch)

	return nil
}
