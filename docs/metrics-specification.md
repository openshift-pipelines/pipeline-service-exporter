## Metrics Specification Document for Pipeline Service Prometheus Exporter

### Introduction:
This document outlines the specifications for a Prometheus exporter that collects metrics from Pipeline Service. The goal of the exporter is to gather essential metrics, providing a comprehensive view of the performance and health status of Pipeline Service.

### Metrics Definition:

_**PipelineRun Failed With PVC Quota:**_  
The count of the number of current PipelineRuns on the cluster marked failed by Tekton because PVC Quota prevented creation of required PVCs. 
The deletion of PipelineRuns that failed because of PVC limits is effectively a decrement of the metric.  That said, given the complexities around
delete events and controllers (missed events not getting relisted, hit and miss success of tombstone objects, multiple events because of finalizers),
we do not decrement in real time, but on our custom Runnable that resets the metric at the same interval the TektonConfig pruner is set to.

_Metric Name:_ pipelinerun_failed_by_pvc_quota_count
_Labels:_ `namespace` label.  Note:  K8s PVC quota specifications are a namespace scoped resource.
_Data Type:_ Gauge
_Description:_ The number of PipelineRuns marked failed because required PVCs could not be created.

_**PipelineRun Scheduling Duration:**_  
The duration of time in seconds taken for a PipelineRun to be scheduled, calculated as the difference between the creation timestamp and the start time of the PipelineRun.

_Metric Name:_ pipelinerun_scheduling_duration  
_Labels:_ Minimally a `namespace` label.  If the `ENABLE_PIPELINERUN_SCHEDULED_DURATION_PIPELINENAME_LABEL` environment variable is set to `true` on the exporter deployment, the `pipelinename` label is set to the name of the Pipeline if its reference is set, otherwise the name of the PipelineRun.
_Data Type:_ Histogram
_Description:_ The time taken in seconds for a PipelineRun to be scheduled.

_**TaskRun Scheduling Duration:**_  
The duration of time in seconds taken for a TaskRun to be scheduled, calculated as the difference between the creation timestamp and the start time of the TaskRun.

_Metric Name:_ taskrunrun_scheduling_duration  
_Labels:_ Minimally a `namespace` label.  If the `ENABLE_TASKRUN_SCHEDULED_DURATION_TASKNAME_LABEL` environment variable is set to `true` on the exporter deployment, the `taskname` label is set to the name of the Task if its reference is set, otherwise the name of the TaskRun.
_Data Type:_ Histogram
_Description:_ The time taken in seconds for a TaskRun to be scheduled.


_**Scheduling Duration of different TaskRuns with a PipelineRun:**_
The time taken in milliseconds between the creation of the first TaskRun(s) and the creation of its PipelineRun, followed by the duration in milliseconds between the completion of a preceding TaskRun and the creation of the following TaskRun.  This metrics accounts for both sequential TaskRuns, parallel TaskRuns that start off a PipelineRun, and ending TaskRuns that depend on multiple TaskRun chains that run in parallel.

_Metric Name:_ pipelinerun_gap_between_taskruns_milliseconds
_Labels:_ Minimally a `namespace` label.  If the `ENABLE_GAP_METRIC_ADDITIONAL_LABELS` environment variable is set to `true` on the exporter deployment, the `pipelinename`, `completed`, and `upcoming` labels are set.  The `pipelinename` label is set to the name of the Pipeline if its reference is set, otherwise the name of the PipelineRun.  The `completed` label is set either the Pipeline name if we are dealing with the first TaskRun, or the name of the latest Task for the TaskRun to be completed.  The `upcoming` label is set to the name of the Task of the TaskRun that is created but not yet complete.
_Data Type_: Histogram
_Description_: The taken between TaskRuns within a PipelineRun


### Metrics Format:
The metrics will be exposed via the Prometheus integration with the Kubernetes Controller Runtime framework.

### Performance Requirements:
To avoid prior issues with memory creep, excessive restarts, and excessive load on the API server, controller / watch based monitoring of PipelineRuns and TaskRuns are employed.  No access to those object should be performed with a non-caching client, only the controller's caching client.

### Security Considerations:
The exporter will implement appropriate security measures to ensure that sensitive data is not exposed.

### Deployment and Operations:
The exporter will be deployed on Stonesoup staging and production clusters and will be monitored regularly to ensure that it is functioning correctly. Regular maintenance will be performed to keep the exporter up-to-date with changes in Pipeline Service.