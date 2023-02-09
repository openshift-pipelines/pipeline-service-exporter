package collector

import (
	"fmt"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"sync"
)

// PipelineServiceCollector struct
type PipelineServiceCollector struct {
	logger            log.Logger
	mutex             sync.Mutex
	durationScheduled *prometheus.GaugeVec
	durationCompleted *prometheus.GaugeVec
}

func NewCollector(logger log.Logger) (*PipelineServiceCollector, error) {
	durationScheduled := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pipelinerun_duration_scheduled_seconds",
		Help: "Duration in seconds for a PipelineRun to be scheduled.",
	}, []string{"name", "uid"})
	durationCompleted := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "pipelinerun_duration_completed_seconds",
		Help: "Duration in seconds for a PipelineRun to complete.",
	}, []string{"name", "uid"})

	return &PipelineServiceCollector{
		logger:            logger,
		durationScheduled: durationScheduled,
		durationCompleted: durationCompleted,
	}, nil
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

	prList, err := getPipelineRuns()
	if err != nil {
		level.Error(c.logger).Log("msg", "Error while fetching PipelineRuns", "err", err)
		return err
	}

	for _, pipelineRun := range prList.Items {
		// Fetch and compute the metrics for schedule and completed time
		scheduledDuration, err := calculateScheduledDuration(pipelineRun)
		if err != nil {
			fmt.Println("Error while calculating the scheduled time of a PipelineRun: ", err)
		}
		//fmt.Printf("scheduledDuration of the PipelineRun %v: %v\n", pipelineRun.Name, scheduledDuration)

		completedDuration, err := calculateCompletedDuration(pipelineRun)
		if err != nil {
			fmt.Println("Error while calculating the completion time of a PipelineRun: ", err)
		}
		fmt.Printf("completedDuration of the PipelineRun %v: %v\n", pipelineRun.Name, completedDuration)

		// Set the metrics
		c.durationScheduled.WithLabelValues(pipelineRun.Name, string(pipelineRun.UID)).Set(scheduledDuration)
		c.durationCompleted.WithLabelValues(pipelineRun.Name, string(pipelineRun.UID)).Set(completedDuration)

	}

	// Make sure it is passed to the channel so that it is exported out
	c.durationScheduled.Collect(ch)
	c.durationCompleted.Collect(ch)

	return nil
}
