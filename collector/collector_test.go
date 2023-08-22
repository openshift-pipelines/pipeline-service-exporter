package collector

import (
	"context"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func TestPipelineRunCollection(t *testing.T) {
	objs := []client.Object{}
	scheme := runtime.NewScheme()
	_ = v1beta1.AddToScheme(scheme)
	c := fake.NewClientBuilder().WithScheme(scheme).WithObjects(objs...).Build()

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
	}
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
	schedReconciler := &ReconcilePipelineRunScheduled{client: c, prCollector: NewPipelineRunScheduledCollector()}
	gapReconciler := &ReconcilePipelineRunTaskRunGap{client: c, prCollector: NewPipelineRunTaskRunGapCollector()}
	ctx := context.TODO()
	for _, tr := range mockTaskRuns {
		err := c.Create(ctx, tr)
		assert.NoError(t, err)
	}
	for _, pr := range mockPipelineRuns {
		err := c.Create(ctx, pr)
		assert.NoError(t, err)
		request := reconcile.Request{
			NamespacedName: types.NamespacedName{
				Namespace: pr.Namespace,
				Name:      pr.Name,
			},
		}
		_, err = schedReconciler.Reconcile(ctx, request)
		_, err = gapReconciler.Reconcile(ctx, request)
	}

	label := prometheus.Labels{NS_LABEL: "test-namespace"}
	g, e := schedReconciler.prCollector.durationScheduled.GetMetricWith(label)
	assert.NoError(t, e)
	assert.NotNil(t, g)
	label = prometheus.Labels{NS_LABEL: "test-namespace"}
	g, e = gapReconciler.prCollector.trGaps.GetMetricWith(label)
	assert.NoError(t, e)
	assert.NotNil(t, g)
	label = prometheus.Labels{NS_LABEL: "test-namespace"}
	g, e = gapReconciler.prCollector.trGaps.GetMetricWith(label)
	assert.NoError(t, e)
	assert.NotNil(t, g)
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

func TestTaskRunTaskRefName(t *testing.T) {
	for _, test := range []struct {
		name           string
		expectedReturn string
		tr             *v1beta1.TaskRun
		pr             *v1beta1.PipelineRun
	}{
		{
			name:           "use task run name",
			expectedReturn: "test-taskrun",
			tr: &v1beta1.TaskRun{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-taskrun",
				},
			},
		},
		{
			name:           "use task run generate name",
			expectedReturn: "test-taskrun-",
			tr: &v1beta1.TaskRun{
				ObjectMeta: metav1.ObjectMeta{
					Name:         "test-taskrun-foo",
					GenerateName: "test-taskrun-",
				},
			},
		},
		{
			name:           "use task run ref param name",
			expectedReturn: "test-task",
			tr: &v1beta1.TaskRun{
				ObjectMeta: metav1.ObjectMeta{
					Name:         "test-taskrun-foo",
					GenerateName: "test-taskrun-",
				},
				Spec: v1beta1.TaskRunSpec{
					TaskRef: &v1beta1.TaskRef{
						ResolverRef: v1beta1.ResolverRef{
							Params: []v1beta1.Param{
								{
									Name: "name",
									Value: v1beta1.ParamValue{
										StringVal: "test-task"},
								},
							},
						},
					},
				},
			},
		},
		{
			name:           "use task run ref name",
			expectedReturn: "test-task",
			tr: &v1beta1.TaskRun{
				ObjectMeta: metav1.ObjectMeta{
					Name:         "test-taskrun-foo",
					GenerateName: "test-taskrun-",
				},
				Spec: v1beta1.TaskRunSpec{TaskRef: &v1beta1.TaskRef{Name: "test-task"}},
			},
		},
		{
			name:           "use task run ref param name",
			expectedReturn: "test-task",
			tr: &v1beta1.TaskRun{
				ObjectMeta: metav1.ObjectMeta{
					Name:         "test-taskrun-foo",
					GenerateName: "test-taskrun-",
				},
				Spec: v1beta1.TaskRunSpec{TaskRef: &v1beta1.TaskRef{Name: "test-task"}},
			},
		},
		{
			name:           "use pipeline run pipeline spec tasks with name",
			expectedReturn: "task1",
			pr: &v1beta1.PipelineRun{
				Spec: v1beta1.PipelineRunSpec{
					PipelineSpec: &v1beta1.PipelineSpec{
						Tasks: []v1beta1.PipelineTask{
							{
								Name: "task1",
							},
						},
					},
				},
			},
			tr: &v1beta1.TaskRun{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-taskrun-task1",
				},
			},
		},
		{
			name:           "use pipeline run pipeline spec tasks with param name",
			expectedReturn: "task1",
			pr: &v1beta1.PipelineRun{
				Spec: v1beta1.PipelineRunSpec{
					PipelineSpec: &v1beta1.PipelineSpec{
						Tasks: []v1beta1.PipelineTask{
							{
								Params: []v1beta1.Param{
									{
										Name:  "name",
										Value: v1beta1.ParamValue{StringVal: "task1"},
									},
								},
							},
						},
					},
				},
			},
			tr: &v1beta1.TaskRun{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-taskrun-task1",
				},
			},
		},
		{
			name:           "use pipeline run spec task run spec",
			expectedReturn: "task1",
			pr: &v1beta1.PipelineRun{
				Spec: v1beta1.PipelineRunSpec{
					TaskRunSpecs: []v1beta1.PipelineTaskRunSpec{
						{
							PipelineTaskName: "task1",
						},
					},
				},
			},
			tr: &v1beta1.TaskRun{
				ObjectMeta: metav1.ObjectMeta{
					Name: "test-taskrun-task1",
				},
			},
		},
	} {
		ret := taskRunTaskRef(test.tr, test.pr)
		if ret != test.expectedReturn {
			t.Errorf("test %s expected %s got %s", test.name, test.expectedReturn, ret)
		}
	}
}

func TestTaskRunCollection(t *testing.T) {
	objs := []client.Object{}
	scheme := runtime.NewScheme()
	_ = v1beta1.AddToScheme(scheme)
	c := fake.NewClientBuilder().WithScheme(scheme).WithObjects(objs...).Build()

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
	}
	reconciler := &ReconcileTaskRun{client: c, trCollector: NewTaskRunCollector()}
	ctx := context.TODO()
	for _, pr := range mockTaskRuns {
		err := c.Create(ctx, pr)
		assert.NoError(t, err)
		request := reconcile.Request{
			NamespacedName: types.NamespacedName{
				Namespace: pr.Namespace,
				Name:      pr.Name,
			},
		}
		_, err = reconciler.Reconcile(ctx, request)
	}

	label := prometheus.Labels{NS_LABEL: "test-namespace"}
	g, e := reconciler.trCollector.durationScheduled.GetMetricWith(label)
	assert.NoError(t, e)
	assert.NotNil(t, g)
}
