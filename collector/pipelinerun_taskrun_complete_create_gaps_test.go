package collector

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"testing"
	"time"
)

func TestPipelineRunGapCollection(t *testing.T) {
	// rather the golang mocks, grabbed actual RHTAP pipelinerun/taskruns from staging
	// to drive the gap metric, given its trickiness
	objs := []client.Object{}
	scheme := runtime.NewScheme()
	_ = v1.AddToScheme(scheme)
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
	for _, trv1beta1 := range trs {
		// mimic what the tekton conversion webhook will do
		tr := &v1.TaskRun{}
		err = trv1beta1.ConvertTo(ctx, tr)
		assert.NoError(t, err)
		err = c.Create(ctx, tr)
		assert.NoError(t, err)
	}
	for _, prv1beta1 := range prs {
		// mimic what the tekton conversion webhook will do
		pr := &v1.PipelineRun{}
		err = prv1beta1.ConvertTo(ctx, pr)
		assert.NoError(t, err)
		err = c.Create(ctx, pr)
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
	mockTaskRuns := []*v1.TaskRun{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:              "test-taskrun-1",
				Namespace:         "test-namespace",
				UID:               types.UID("test-taskrun-1"),
				CreationTimestamp: metav1.NewTime(time.Now().UTC()),
			},
			Status: v1.TaskRunStatus{
				Status: duckv1.Status{
					ObservedGeneration: 0,
					Conditions: duckv1.Conditions{{
						Type:   "Succeeded",
						Status: corev1.ConditionTrue,
					}},
					Annotations: nil,
				},
				TaskRunStatusFields: v1.TaskRunStatusFields{
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
			Status: v1.TaskRunStatus{
				Status: duckv1.Status{
					ObservedGeneration: 0,
					Conditions: duckv1.Conditions{{
						Type:   "Succeeded",
						Status: corev1.ConditionTrue,
					}},
					Annotations: nil,
				},
				TaskRunStatusFields: v1.TaskRunStatusFields{
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
			Status: v1.TaskRunStatus{
				Status: duckv1.Status{
					ObservedGeneration: 0,
					Conditions: duckv1.Conditions{{
						Type:   "Succeeded",
						Status: corev1.ConditionUnknown,
					}},
					Annotations: nil,
				},
				TaskRunStatusFields: v1.TaskRunStatusFields{
					StartTime: &metav1.Time{Time: time.Now().UTC().Add(25 * time.Second)},
				},
			},
		},
	}
	mockPipelineRuns := []*v1.PipelineRun{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:              "test-pipelinerun-4",
				Namespace:         "test-namespace",
				UID:               types.UID("test-pipelinerun-4"),
				CreationTimestamp: metav1.NewTime(time.Now().UTC()),
			},
			Status: v1.PipelineRunStatus{
				Status: duckv1.Status{
					ObservedGeneration: 0,
					Conditions: duckv1.Conditions{{
						Type:   "Succeeded",
						Status: corev1.ConditionTrue,
					}},
					Annotations: nil,
				},
				PipelineRunStatusFields: v1.PipelineRunStatusFields{
					ChildReferences: []v1.ChildStatusReference{
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

func TestTaskRunGapEventFilter_Update(t *testing.T) {
	filter := &taskRunGapEventFilter{}
	for _, tc := range []struct {
		name       string
		oldPR      *v1.PipelineRun
		newPR      *v1.PipelineRun
		expectedRC bool
	}{
		{
			name:  "not done no status",
			oldPR: &v1.PipelineRun{},
			newPR: &v1.PipelineRun{},
		},
		{
			name:  "not done status unknown",
			oldPR: &v1.PipelineRun{},
			newPR: &v1.PipelineRun{
				Status: v1.PipelineRunStatus{
					Status: duckv1.Status{
						Conditions: []apis.Condition{
							{
								Type:   apis.ConditionSucceeded,
								Status: corev1.ConditionUnknown,
							},
						},
					},
				},
			},
		},
		{
			name:  "just done succeed",
			oldPR: &v1.PipelineRun{},
			newPR: &v1.PipelineRun{
				Status: v1.PipelineRunStatus{
					Status: duckv1.Status{
						Conditions: []apis.Condition{
							{
								Type:   apis.ConditionSucceeded,
								Status: corev1.ConditionTrue,
							},
						},
					},
				},
			},
			expectedRC: true,
		},
		{
			name:  "just done failed",
			oldPR: &v1.PipelineRun{},
			newPR: &v1.PipelineRun{
				Status: v1.PipelineRunStatus{
					Status: duckv1.Status{
						Conditions: []apis.Condition{
							{
								Type:   apis.ConditionSucceeded,
								Status: corev1.ConditionFalse,
							},
						},
					},
				},
			},
			expectedRC: true,
		},
		{
			name: "udpate after done",
			oldPR: &v1.PipelineRun{
				Status: v1.PipelineRunStatus{
					Status: duckv1.Status{
						Conditions: []apis.Condition{
							{
								Type:   apis.ConditionSucceeded,
								Status: corev1.ConditionFalse,
							},
						},
					},
				},
			},
			newPR: &v1.PipelineRun{
				Status: v1.PipelineRunStatus{
					Status: duckv1.Status{
						Conditions: []apis.Condition{
							{
								Type:   apis.ConditionSucceeded,
								Status: corev1.ConditionFalse,
							},
						},
					},
				},
			},
		},
	} {
		ev := event.UpdateEvent{
			ObjectOld: tc.oldPR,
			ObjectNew: tc.newPR,
		}
		rc := filter.Update(ev)
		if rc != tc.expectedRC {
			t.Errorf(fmt.Sprintf("tc %s expected %v but got %v", tc.name, tc.expectedRC, rc))
		}
	}
}
