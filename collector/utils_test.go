package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"testing"
)

func validateHistogramVec(t *testing.T, h *prometheus.HistogramVec, labels prometheus.Labels, checkMax bool) {
	observer, err := h.GetMetricWith(labels)
	assert.NoError(t, err)
	assert.NotNil(t, observer)
	histogram := observer.(prometheus.Histogram)
	metric := &dto.Metric{}
	histogram.Write(metric)
	assert.NotNil(t, metric.Histogram)
	assert.NotNil(t, metric.Histogram.SampleCount)
	assert.NotZero(t, *metric.Histogram.SampleCount)
	assert.NotNil(t, metric.Histogram.SampleSum)
	assert.Greater(t, *metric.Histogram.SampleSum, float64(-1))
	if checkMax {
		// for now, we are not tracking gap histograms example by example; but rather we
		// have determined the max histogram (currently 7000) manually and make sure everyone
		// is under that
		assert.Less(t, *metric.Histogram.SampleSum, float64(7001))
	}
}

func validateHistogramVecZeroCount(t *testing.T, h *prometheus.HistogramVec, labels prometheus.Labels) {
	observer, err := h.GetMetricWith(labels)
	assert.NoError(t, err)
	assert.NotNil(t, observer)
	histogram := observer.(prometheus.Histogram)
	metric := &dto.Metric{}
	histogram.Write(metric)
	assert.NotNil(t, metric.Histogram)
	assert.NotNil(t, metric.Histogram.SampleCount)
	assert.Equal(t, uint64(0), *metric.Histogram.SampleCount)
}

func validateGaugeVec(t *testing.T, g *prometheus.GaugeVec, labels prometheus.Labels, count float64) {
	gauge, err := g.GetMetricWith(labels)
	assert.NoError(t, err)
	assert.NotNil(t, gauge)
	metric := &dto.Metric{}
	gauge.Write(metric)
	assert.NotNil(t, metric.Gauge)
	assert.NotNil(t, metric.Gauge.Value)
	assert.Equal(t, count, *metric.Gauge.Value)
}

// For now at least, we are keeping these as v1beta1 to have some element of regression testing, now that we've flipped
// the "default" to v1.
func pipelineRunFromActualRHTAPYaml() ([]v1beta1.PipelineRun, error) {
	prs := []v1beta1.PipelineRun{}
	yamlStrings := []string{tooBigNumPRYaml,
		prYaml}

	v1.AddToScheme(scheme.Scheme)
	decoder := scheme.Codecs.UniversalDecoder()
	for _, y := range yamlStrings {
		buf := []byte(y)
		pr := &v1beta1.PipelineRun{}
		_, _, err := decoder.Decode(buf, nil, pr)
		if err != nil {
			return nil, err
		}
		prs = append(prs, *pr)
	}
	return prs, nil
}

// For now at least, we are keeping these as v1beta1 to have some element of regression testing, now that we've flipped
// the "default" to v1.
func taskRunsFromActualRHTAPYaml() ([]v1beta1.TaskRun, error) {
	trs := []v1beta1.TaskRun{}
	yamlStrings := []string{tooBigNumTRInitYaml,
		tooBigNumTRCloneYaml,
		tooBigNumTRBuildYaml,
		tooBigNumTRSbomYaml,
		tooBigNumTRSummYaml,
		trInitYaml,
		trCloneYaml,
		trSbomJsonCheckYaml,
		trBuildYaml,
		trInspectImgYaml,
		trDeprecatedBaseImgCheck,
		trLabelYaml,
		trClamavYaml,
		trClairYaml,
		trSummaryYaml,
		trShowSbomYaml}
	v1.AddToScheme(scheme.Scheme)
	decoder := scheme.Codecs.UniversalDecoder()
	for _, y := range yamlStrings {
		buf := []byte(y)
		tr := &v1beta1.TaskRun{}
		_, _, err := decoder.Decode(buf, nil, tr)
		if err != nil {
			return nil, err
		}
		trs = append(trs, *tr)
	}
	return trs, nil
}

func TestPipelineRunPipelineRef(t *testing.T) {
	for _, test := range []struct {
		name           string
		expectedReturn string
		pr             *v1.PipelineRun
	}{
		{
			name:           "use pipeline run name",
			expectedReturn: "test-pipelinerun",
			pr: &v1.PipelineRun{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-pipelinerun",
				},
			},
		},
		{
			name:           "use pipelinerun run generate name",
			expectedReturn: "test-pipelinerun-",
			pr: &v1.PipelineRun{
				ObjectMeta: metav1.ObjectMeta{
					Name:         "test-pipelinerun-foo",
					GenerateName: "test-pipelinerun-",
				},
			},
		},
		{
			name:           "use pipeline run ref param name",
			expectedReturn: "test-pipeline",
			pr: &v1.PipelineRun{
				ObjectMeta: metav1.ObjectMeta{
					Name:         "test-pipelinerun-foo",
					GenerateName: "test-pipelinerun-",
				},
				Spec: v1.PipelineRunSpec{
					PipelineRef: &v1.PipelineRef{
						ResolverRef: v1.ResolverRef{
							Params: []v1.Param{
								{
									Name: "name",
									Value: v1.ParamValue{
										StringVal: "test-pipeline"},
								},
							},
						},
					},
				},
			},
		},
		{
			name:           "use pipeline run ref name",
			expectedReturn: "test-pipeline",
			pr: &v1.PipelineRun{
				ObjectMeta: metav1.ObjectMeta{
					Name:         "test-pipelinerun-foo",
					GenerateName: "test-pipelinerun-",
				},
				Spec: v1.PipelineRunSpec{PipelineRef: &v1.PipelineRef{Name: "test-pipeline"}},
			},
		},
	} {
		ret := pipelineRunPipelineRef(test.pr)
		if ret != test.expectedReturn {
			t.Errorf("test %s expected %s got %s", test.name, test.expectedReturn, ret)
		}
	}
}

func TestTaskRef(t *testing.T) {
	for _, test := range []struct {
		name           string
		expectedReturn string
		labels         map[string]string
	}{
		{
			name:           "use task run name",
			expectedReturn: "test-taskrun",
			labels: map[string]string{
				pipeline.TaskRunLabelKey: "test-taskrun",
			},
		},
		{
			name:           "use cluster task name",
			expectedReturn: "test-taskrun",
			labels: map[string]string{
				pipeline.ClusterTaskLabelKey: "test-taskrun",
				pipeline.TaskRunLabelKey:     "test-taskrun-foo",
			},
		},
		{
			name:           "use pipeline task name",
			expectedReturn: "test-taskrun",
			labels: map[string]string{
				pipeline.PipelineTaskLabelKey: "test-taskrun",
				pipeline.ClusterTaskLabelKey:  "test-taskrun-foo",
				pipeline.TaskRunLabelKey:      "test-taskrun-foo",
			},
		},
		{
			name:           "use task name",
			expectedReturn: "test-taskrun",
			labels: map[string]string{
				pipeline.TaskLabelKey:         "test-taskrun",
				pipeline.PipelineTaskLabelKey: "test-taskrun-foo",
				pipeline.ClusterTaskLabelKey:  "test-taskrun-foo",
				pipeline.TaskRunLabelKey:      "test-taskrun-foo",
			},
		},
	} {
		ret := taskRef(test.labels)
		if ret != test.expectedReturn {
			t.Errorf("test %s expected %s got %s", test.name, test.expectedReturn, ret)
		}
	}
}
