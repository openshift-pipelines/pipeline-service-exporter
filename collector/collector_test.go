package collector

import (
	"github.com/go-kit/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCollector(t *testing.T) {
	logger := log.NewNopLogger()

	collector, err := NewCollector(logger)
	assert.NoError(t, err)
	assert.NotNil(t, collector)
	assert.NotNil(t, collector.durationScheduled)
	assert.NotNil(t, collector.durationCompleted)
	assert.NotNil(t, collector.pipelineRuns)
	assert.Equal(t, 0, len(collector.pipelineRuns))
}
