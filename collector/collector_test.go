package collector

import (
	"context"
	"fmt"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline"
	"sigs.k8s.io/controller-runtime/pkg/metrics"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func TestPipelineRunGapCollection(t *testing.T) {
	// rather the golang mocks, grabbed actual RHTAP pipelinerun/taskruns from staging
	// to drive the gap metric, given its trickiness
	objs := []client.Object{}
	scheme := runtime.NewScheme()
	_ = v1beta1.AddToScheme(scheme)
	c := fake.NewClientBuilder().WithScheme(scheme).WithObjects(objs...).Build()
	gapReconciler := &ReconcilePipelineRunTaskRunGap{client: c, prCollector: NewPipelineRunTaskRunGapCollector()}

	var err error
	// first we test with the samples pulled from actual RHTAP yaml to best capture the parallel task executions
	prs := []v1beta1.PipelineRun{}
	trs := []v1beta1.TaskRun{}
	prs, err = pipelineRunFromActualRHTAPYaml()
	if err != nil {
		t.Fatalf(fmt.Sprintf("%s", err.Error()))
	}
	trs, err = taskRunsFromActualRHTAPYaml()
	if err != nil {
		t.Fatalf(fmt.Sprintf("%s", err.Error()))
	}

	ctx := context.TODO()
	for _, tr := range trs {
		err = c.Create(ctx, &tr)
		assert.NoError(t, err)
	}
	for _, pr := range prs {
		err = c.Create(ctx, &pr)
		assert.NoError(t, err)
		request := reconcile.Request{
			NamespacedName: types.NamespacedName{
				Namespace: pr.Namespace,
				Name:      pr.Name,
			},
		}
		_, err = gapReconciler.Reconcile(ctx, request)
		label := prometheus.Labels{NS_LABEL: pr.Namespace, STATUS_LABEL: SUCCEEDED}
		validateHistogramVec(t, gapReconciler.prCollector.trGaps, label, true)
	}

	// then some additional unit tests were we build simpler pipelineruns/taskruns that capture paths
	// related to completion times not being set
	mockTaskRuns := []*v1beta1.TaskRun{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:              "test-taskrun-1",
				Namespace:         "test-namespace",
				UID:               types.UID("test-taskrun-1"),
				CreationTimestamp: metav1.NewTime(time.Now().UTC()),
			},
			Status: v1beta1.TaskRunStatus{
				Status: duckv1.Status{
					ObservedGeneration: 0,
					Conditions: duckv1.Conditions{{
						Type:   "Succeeded",
						Status: corev1.ConditionTrue,
					}},
					Annotations: nil,
				},
				TaskRunStatusFields: v1beta1.TaskRunStatusFields{
					StartTime: &metav1.Time{Time: time.Now().UTC()},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:              "test-taskrun-2",
				Namespace:         "test-namespace",
				UID:               types.UID("test-taskrun-2"),
				CreationTimestamp: metav1.NewTime(time.Now().UTC().Add(5 * time.Second)),
			},
			Status: v1beta1.TaskRunStatus{
				Status: duckv1.Status{
					ObservedGeneration: 0,
					Conditions: duckv1.Conditions{{
						Type:   "Succeeded",
						Status: corev1.ConditionTrue,
					}},
					Annotations: nil,
				},
				TaskRunStatusFields: v1beta1.TaskRunStatusFields{
					StartTime: &metav1.Time{Time: time.Now().UTC().Add(10 * time.Second)},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:              "test-taskrun-3",
				Namespace:         "test-namespace",
				UID:               types.UID("test-taskrun-3"),
				CreationTimestamp: metav1.NewTime(time.Now().UTC().Add(20 * time.Second)),
			},
			Status: v1beta1.TaskRunStatus{
				Status: duckv1.Status{
					ObservedGeneration: 0,
					Conditions: duckv1.Conditions{{
						Type:   "Succeeded",
						Status: corev1.ConditionUnknown,
					}},
					Annotations: nil,
				},
				TaskRunStatusFields: v1beta1.TaskRunStatusFields{
					StartTime: &metav1.Time{Time: time.Now().UTC().Add(25 * time.Second)},
				},
			},
		},
	}
	mockPipelineRuns := []*v1beta1.PipelineRun{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:              "test-pipelinerun-4",
				Namespace:         "test-namespace",
				UID:               types.UID("test-pipelinerun-4"),
				CreationTimestamp: metav1.NewTime(time.Now().UTC()),
			},
			Status: v1beta1.PipelineRunStatus{
				Status: duckv1.Status{
					ObservedGeneration: 0,
					Conditions: duckv1.Conditions{{
						Type:   "Succeeded",
						Status: corev1.ConditionTrue,
					}},
					Annotations: nil,
				},
				PipelineRunStatusFields: v1beta1.PipelineRunStatusFields{
					ChildReferences: []v1beta1.ChildStatusReference{
						{
							TypeMeta: runtime.TypeMeta{
								Kind: "TaskRun",
							},
							Name: "test-taskrun-1",
						},
						{
							TypeMeta: runtime.TypeMeta{
								Kind: "TaskRun",
							},
							Name: "test-taskrun-2",
						},
						{
							TypeMeta: runtime.TypeMeta{
								Kind: "TaskRun",
							},
							Name: "test-taskrun-3",
						},
					},
					StartTime:      &metav1.Time{Time: time.Now().UTC().Add(5 * time.Second)},
					CompletionTime: &metav1.Time{Time: time.Now().UTC().Add(10 * time.Second)},
				},
			},
		},
	}
	for _, tr := range mockTaskRuns {
		err = c.Create(ctx, tr)
		assert.NoError(t, err)
	}
	for _, pipelineRun := range mockPipelineRuns {
		err = c.Create(ctx, pipelineRun)
		assert.NoError(t, err)
		request := reconcile.Request{
			NamespacedName: types.NamespacedName{
				Namespace: pipelineRun.Namespace,
				Name:      pipelineRun.Name,
			},
		}
		_, err = gapReconciler.Reconcile(ctx, request)
	}

	label := prometheus.Labels{NS_LABEL: "test-namespace", STATUS_LABEL: SUCCEEDED}
	validateHistogramVec(t, gapReconciler.prCollector.trGaps, label, false)

}

func TestPipelineRunScheduleCollection(t *testing.T) {
	mockPipelineRuns := []*v1beta1.PipelineRun{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:              "test-pipelinerun-1",
				Namespace:         "test-namespace",
				UID:               types.UID("test-pipelinerun-1"),
				CreationTimestamp: metav1.NewTime(time.Now().UTC()),
			},
			Status: v1beta1.PipelineRunStatus{
				Status: duckv1.Status{
					ObservedGeneration: 0,
					Conditions: duckv1.Conditions{{
						Type:   "Succeeded",
						Status: corev1.ConditionTrue,
					}},
					Annotations: nil,
				},
				PipelineRunStatusFields: v1beta1.PipelineRunStatusFields{
					StartTime:      &metav1.Time{Time: time.Now().UTC().Add(5 * time.Second)},
					CompletionTime: &metav1.Time{Time: time.Now().UTC().Add(10 * time.Second)},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:              "test-pipelinerun-2",
				Namespace:         "test-namespace",
				UID:               types.UID("test-pipelinerun-2"),
				CreationTimestamp: metav1.NewTime(time.Now().UTC()),
			},
			Status: v1beta1.PipelineRunStatus{
				Status: duckv1.Status{
					ObservedGeneration: 0,
					Conditions: duckv1.Conditions{{
						Type:   "Succeeded",
						Status: corev1.ConditionTrue,
					}},
					Annotations: nil,
				},
				PipelineRunStatusFields: v1beta1.PipelineRunStatusFields{
					StartTime:      &metav1.Time{Time: time.Now().UTC().Add(5 * time.Second)},
					CompletionTime: &metav1.Time{Time: time.Now().UTC().Add(10 * time.Second)},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:              "test-pipelinerun-3",
				Namespace:         "test-namespace",
				UID:               types.UID("test-pipelinerun-3"),
				CreationTimestamp: metav1.NewTime(time.Now().UTC()),
			},
			Status: v1beta1.PipelineRunStatus{
				Status: duckv1.Status{
					ObservedGeneration: 0,
					Conditions: duckv1.Conditions{{
						Type:   "Succeeded",
						Status: corev1.ConditionTrue,
					}},
					Annotations: nil,
				},
				PipelineRunStatusFields: v1beta1.PipelineRunStatusFields{
					ChildReferences: []v1beta1.ChildStatusReference{
						{
							TypeMeta: runtime.TypeMeta{
								Kind: "TaskRun",
							},
							Name: "test-taskrun-1",
						},
						{
							TypeMeta: runtime.TypeMeta{
								Kind: "TaskRun",
							},
							Name: "test-taskrun-2",
						},
					},
					StartTime:      &metav1.Time{Time: time.Now().UTC().Add(5 * time.Second)},
					CompletionTime: &metav1.Time{Time: time.Now().UTC().Add(10 * time.Second)},
				},
			},
		},
	}
	for _, pr := range mockPipelineRuns {
		metric := NewPipelineRunScheduledMetric()
		label := prometheus.Labels{NS_LABEL: "test-namespace", PIPELINE_NAME_LABEL: pipelineRunPipelineRef(pr)}
		bumpPipelineRunScheduledDuration(calculateScheduledDurationPipelineRun(pr), pr, metric)
		validateHistogramVec(t, metric, label, false)
		metrics.Registry.Unregister(metric)
	}

}

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

func pipelineRunFromActualRHTAPYaml() ([]v1beta1.PipelineRun, error) {
	prs := []v1beta1.PipelineRun{}
	yamlStrings := []string{tooBigNumPRYaml,
		prYaml}

	v1beta1.AddToScheme(scheme.Scheme)
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
	v1beta1.AddToScheme(scheme.Scheme)
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
		pr             *v1beta1.PipelineRun
	}{
		{
			name:           "use pipeline run name",
			expectedReturn: "test-pipelinerun",
			pr: &v1beta1.PipelineRun{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-pipelinerun",
				},
			},
		},
		{
			name:           "use pipelinerun run generate name",
			expectedReturn: "test-pipelinerun-",
			pr: &v1beta1.PipelineRun{
				ObjectMeta: metav1.ObjectMeta{
					Name:         "test-pipelinerun-foo",
					GenerateName: "test-pipelinerun-",
				},
			},
		},
		{
			name:           "use pipeline run ref param name",
			expectedReturn: "test-pipeline",
			pr: &v1beta1.PipelineRun{
				ObjectMeta: metav1.ObjectMeta{
					Name:         "test-pipelinerun-foo",
					GenerateName: "test-pipelinerun-",
				},
				Spec: v1beta1.PipelineRunSpec{
					PipelineRef: &v1beta1.PipelineRef{
						ResolverRef: v1beta1.ResolverRef{
							Params: []v1beta1.Param{
								{
									Name: "name",
									Value: v1beta1.ParamValue{
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
			pr: &v1beta1.PipelineRun{
				ObjectMeta: metav1.ObjectMeta{
					Name:         "test-pipelinerun-foo",
					GenerateName: "test-pipelinerun-",
				},
				Spec: v1beta1.PipelineRunSpec{PipelineRef: &v1beta1.PipelineRef{Name: "test-pipeline"}},
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

func TestTaskRunScheduledCollection(t *testing.T) {
	mockTaskRuns := []*v1beta1.TaskRun{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:              "test-taskrun-1",
				Namespace:         "test-namespace",
				UID:               types.UID("test-taskrun-1"),
				CreationTimestamp: metav1.NewTime(time.Now().UTC()),
			},
			Status: v1beta1.TaskRunStatus{
				Status: duckv1.Status{
					ObservedGeneration: 0,
					Conditions: duckv1.Conditions{{
						Type:   "Succeeded",
						Status: corev1.ConditionTrue,
					}},
					Annotations: nil,
				},
				TaskRunStatusFields: v1beta1.TaskRunStatusFields{
					StartTime:      &metav1.Time{Time: time.Now().UTC().Add(5 * time.Second)},
					CompletionTime: &metav1.Time{Time: time.Now().UTC().Add(10 * time.Second)},
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:              "test-taskrun-2",
				Namespace:         "test-namespace",
				UID:               types.UID("test-taskrun-2"),
				CreationTimestamp: metav1.NewTime(time.Now().UTC()),
			},
			Status: v1beta1.TaskRunStatus{
				Status: duckv1.Status{
					ObservedGeneration: 0,
					Conditions: duckv1.Conditions{{
						Type:   "Succeeded",
						Status: corev1.ConditionUnknown,
					}},
					Annotations: nil,
				},
				TaskRunStatusFields: v1beta1.TaskRunStatusFields{
					StartTime: &metav1.Time{Time: time.Now().UTC().Add(5 * time.Second)},
				},
			},
		},
	}

	for _, tr := range mockTaskRuns {
		metric := NewTaskRunScheduledMetric()
		label := prometheus.Labels{NS_LABEL: "test-namespace", TASK_NAME_LABEL: taskRef(tr.Labels)}
		bumpTaskRunScheduledDuration(calculateScheduledDurationTaskRun(tr), tr, metric)
		validateHistogramVec(t, metric, label, false)
		metrics.Registry.Unregister(metric)
	}
}
