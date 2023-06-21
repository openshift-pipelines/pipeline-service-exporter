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
	"github.com/prometheus/client_golang/prometheus"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
)

// PipelineServiceCollector struct
type PipelineServiceCollector struct {
	durationScheduled *prometheus.HistogramVec
}

func NewCollector() *PipelineServiceCollector {
	durationScheduled := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "pipelinerun_duration_scheduled_seconds",
		Help: "Duration in seconds for a PipelineRun to be scheduled.",
		// reminder: exponential buckets need a start value greater than 0
		// should take us to a max of 312.5 seconds (.1 * 5 * 5 * 5 * 5 * 5)
		Buckets: prometheus.ExponentialBuckets(0.1, 5, 5),
	}, []string{"namespace"})

	pipelineServiceCollector := &PipelineServiceCollector{
		durationScheduled: durationScheduled,
	}
	metrics.Registry.MustRegister(durationScheduled)

	return pipelineServiceCollector
}

func (c *PipelineServiceCollector) bumpScheduledDuration(ns string, scheduleDuration float64) {
	c.durationScheduled.WithLabelValues(ns).Observe(scheduleDuration)
}
