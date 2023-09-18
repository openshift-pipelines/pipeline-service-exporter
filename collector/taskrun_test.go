package collector

import (
	"fmt"
	"github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"testing"
)

func TestTaskRunStartTimeEventFilter_Update(t *testing.T) {
	filter := &trStartTimeEventFilter{
		metric: NewTaskRunScheduledMetric(),
	}
	for _, tc := range []struct {
		name       string
		oldTR      *v1beta1.TaskRun
		newTR      *v1beta1.TaskRun
		expectedRC bool
	}{
		{
			name:  "not started",
			oldTR: &v1beta1.TaskRun{},
			newTR: &v1beta1.TaskRun{},
		},
		{
			name:  "just started",
			oldTR: &v1beta1.TaskRun{},
			newTR: &v1beta1.TaskRun{
				Status: v1beta1.TaskRunStatus{
					TaskRunStatusFields: v1beta1.TaskRunStatusFields{StartTime: &metav1.Time{}},
				},
			},
		},
		{
			name: "udpate after started",
			oldTR: &v1beta1.TaskRun{
				Status: v1beta1.TaskRunStatus{
					TaskRunStatusFields: v1beta1.TaskRunStatusFields{StartTime: &metav1.Time{}},
				},
			},
			newTR: &v1beta1.TaskRun{
				Status: v1beta1.TaskRunStatus{
					TaskRunStatusFields: v1beta1.TaskRunStatusFields{StartTime: &metav1.Time{}},
				},
			},
		},
	} {
		ev := event.UpdateEvent{
			ObjectOld: tc.oldTR,
			ObjectNew: tc.newTR,
		}
		rc := filter.Update(ev)
		if rc != tc.expectedRC {
			t.Errorf(fmt.Sprintf("tc %s expected %v but got %v", tc.name, tc.expectedRC, rc))
		}
	}

}
