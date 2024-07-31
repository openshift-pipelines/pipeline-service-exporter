## Metrics Specification Document for Pipeline Service Prometheus Exporter

### Introduction:
This document outlines the specifications for a Prometheus exporter that collects metrics from Pipeline Service. The goal of the exporter is to gather essential metrics, providing a comprehensive view of the performance and health status of Pipeline Service.

### Metrics Definition:

_**PipelineRun Execution Overhead:**_  
Proportion of time elapsed between the completion of a TaskRun and the start of the next TaskRun within a PipelineRun to the total duration of successful PipelineRuns.

_Metric Name:_ `pipeline_service_execution_overhead_percentage`
_Labels:_ `namespace`, `status` labels.
_Data Type:_ Histogram
_Description:_ One of our alert metrics, which we target to be 5% or below over the course of a day, across 28 days.

_**PipelineRun Scheduling Overhead:**_  
Proportion of time elapsed waiting for the pipeline controller to receive create events compared to the total duration of successful PipelineRuns.

_Metric Name:_ `pipeline_service_schedule_overhead_percentage`
_Labels:_ `namespace`, `status` labels.
_Data Type:_ Histogram
_Description:_ One of our alert metrics, which we target to be 5% or below over the course of a day, across 28 days.

_**Pipeline Bundle Resolution Wait Time:**_  
Duration in milliseconds for a resolution request for a pipeline reference needed by a pipelinerun to be recognized as complete by the pipelinerun reconciler in the tekton controller.

_Metric Name:_ `pipelinerun_pipeline_resolution_wait_milliseconds`
_Labels:_ `namespace`  label.
_Data Type:_ Histogram
_Description:_ Gives an indication on how long the pulling of the Konflux Pipeline Bundles form quay.io are taking,
before the cache is established, when creating PipelineRuns.

_**Task Bundle Resolution Wait Time:**_  
Duration in milliseconds for a resolution request for a pipeline reference needed by a taskrun to be recognized as complete by the taskrun reconciler in the tekton controller.

_Metric Name:_ `taskrun_task_resolution_wait_milliseconds`
_Labels:_ `namespace` label.
_Data Type:_ Histogram
_Description:_ Gives an indication on how long the pulling of the Konflux Task and Pipeline Bundles form quay.io are taking,
before the cache is established, when creating TaskRuns.

_**Underlying Pod Creation To Complete Times:**_  
Since tekton's analogous duration metrics are only from start time to completion, we provide a create time to completion for comparisons and potential alerting.

_Metric Name:_ `tekton_pods_create_to_complete_seconds`
_Labels:_ `namespace` label.
_Data Type:_ Histogram
_Description:_ A better alternative in our opinion to the upstream metric `tekton_pipelines_controller_pipelinerun_duration_seconds_[bucket, sum, count]`

_**PipelineRun Failed With PVC Quota:**_  
The count of the number of current PipelineRuns on the cluster marked failed by Tekton because PVC Quota prevented creation of required PVCs. 
The deletion of PipelineRuns that failed because of PVC limits is effectively a decrement of the metric.  That said, given the complexities around
delete events and controllers (missed events not getting relisted, hit and miss success of tombstone objects, multiple events because of finalizers),
we do not decrement in real time, but on our custom Runnable that resets the metric at the same interval the TektonConfig pruner is set to.

_Metric Name:_ `pipelinerun_failed_by_pvc_quota_count`
_Labels:_ `namespace` label.  Note:  K8s PVC quota specifications are a namespace scoped resource.
_Data Type:_ Gauge
_Description:_ The number of PipelineRuns marked failed because required PVCs could not be created.

_**TaskRun Yet To Attempt Pod Creation:**_  
The number of TaskRuns where the Tekton Controller has yet to attempt to create its underlying Pod.

_Metric Name:_ `taskrun_pod_create_not_attempted_or_pending_count`
_Labels:_ `namespace` label.  
_Data Type:_ Gauge
_Description:_ Number of TaskRuns where the Tekton Controller has yet to attempt to create its underlying Pod, or the TaskRun is still in Pending state, for multiple scan iterations.

_**Pending ResolutionRequests Yet To Be Attempted**_  
The number of ResolutionRequests where the Resolver Controller has yet to attempt to initiate retrieval for multiple scan iterations.

_Metric Name:_ `pending_resolutionrequest_count`
_Labels:_ `namespace` label.  
_Data Type:_ Gauge
_Description:_ Number of ResolutionRequests where the Resolver Controller has yet to initiate retrieval for multiple scan iterations.

_**PipelineRun Yet To Kick Off:**_  
The number of PipelineRuns where the Tekton Controller has yet to attempt to process its correctly defined Task specifications for multiple scan iterations.

_Metric Name:_ `pipelinerun_kickoff_not_attempted_count`
_Labels:_ `namespace` label.  
_Data Type:_ Gauge
_Description:_ Number of TaskRuns where the Tekton Controller has yet to attempt to create its underlying Pod, or the TaskRun is still in Pending state, for multiple scan iterations.

_**PipelineRun Scheduling Duration:**_  
The duration of time in seconds taken for a PipelineRun to be "scheduled", meaning it has been received by the Tekton controller.  It is calculated as the difference between the creation timestamp and the start time of the PipelineRun, where the start time is set by the Tekton controller on the initial event received for the creation of the PipelineRun.  It is a good indication of how quickly the API server sends create events to the Tekton controller.

_Metric Name:_ `pipelinerun_duration_scheduled_seconds`
_Labels:_ a `namespace` label.
_Data Type:_ Histogram
_Description:_ The time taken in seconds for a PipelineRun to be "scheduled", meaning it has been received by the Tekton controller.

_**TaskRun Scheduling Duration:**_  
The duration of time in seconds taken for a TaskRun to be "scheduled", meaning it has been received by the Tekton controller.  It is calculated as the difference between the creation timestamp and the start time of the TaskRun, where the start time is set by the Tekton controller on the initial event received for the creation of the TaskRun.  It is a good indication of how quickly the API server sends create events to the Tekton controller.

_Metric Name:_ `taskrun_duration_scheduled_seconds`
_Labels:_ a `namespace` label.
_Data Type:_ Histogram
_Description:_ The time taken in seconds for a TaskRun to be "scheduled", meaning it has been received by the Tekton controller.


_**Scheduling Duration of different TaskRuns with a PipelineRun:**_
The time taken in milliseconds between the creation of the first TaskRun(s) and the creation of its PipelineRun, followed by the duration in milliseconds between the completion of a preceding TaskRun and the creation of the following TaskRun.  This metrics accounts for both sequential TaskRuns, parallel TaskRuns that start off a PipelineRun, and ending TaskRuns that depend on multiple TaskRun chains that run in parallel.

_Metric Name:_ `pipelinerun_gap_between_taskruns_milliseconds`
_Labels:_ Minimally a `namespace` label.  
_Data Type_: Histogram
_Description_: The taken between TaskRuns within a PipelineRun

_**Scheduling Duration that a TaskRun Pod is recognized by the Kubelet:**_
The time taken in milliseconds between the creation of a Pod, where the Pod start time is set once the kubelet has acknowledged the pod, but has not yet pulled its images.

_Metric Name:_ `taskrun_pod_duration_kubelet_acknowledged_milliseconds`
_Labels:_ a `namespace` label.
_Data Type_: Histogram
_Description_: Duration in milliseconds between the pod creation time and pod start time

_**Scheduling Duration that a TaskRun Pod's images are pulled by the Kubelet and the Pod is started:**_
The time taken in milliseconds between the pod start time and the first container to start. This should include any overhead to pull container images, plus any kubelet to linux scheduling overhead.

_Metric Name:_ `taskrun_pod_duration_kubelet_to_container_start_milliseconds`
_Labels:_ a `namespace` label.
_Data Type_: Histogram
_Description_: Duration in milliseconds between the pod start time and the first container to start.


### Metrics Format:
The metrics will be exposed via the Prometheus integration with the Kubernetes Controller Runtime framework.

### Performance Requirements:
To avoid prior issues with memory creep, excessive restarts, and excessive load on the API server, controller / watch based monitoring of PipelineRuns and TaskRuns are employed.  No access to those object should be performed with a non-caching client, only the controller's caching client.

### Security Considerations:
The exporter will implement appropriate security measures to ensure that sensitive data is not exposed.

### Deployment and Operations:
The exporter will be deployed on Stonesoup staging and production clusters and will be monitored regularly to ensure that it is functioning correctly. Regular maintenance will be performed to keep the exporter up-to-date with changes in Pipeline Service.