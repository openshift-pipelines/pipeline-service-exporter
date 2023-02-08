## Metrics Specification Document for Pipeline Service Prometheus Exporter

### Introduction:
This document outlines the specifications for a Prometheus exporter that collects metrics from Pipeline Service. The goal of the exporter is to gather essential metrics, providing a comprehensive view of the performance and health status of Pipeline Service.

### Metrics Definition:

_**PipelineRun Scheduling Duration:**_  
The duration of time taken for a PipelineRun to be scheduled, calculated as the difference between the creation timestamp and the start time of the PipelineRun. 

_Metric Name:_ pipeline_run_scheduling_duration  
_Labels:_ None  
_Data Type:_ Gauge  
_Description:_ The time taken in seconds for a PipelineRun to be scheduled.

_**PipelineRun Completion Duration:**_  
The duration of time taken for a PipelineRun to complete, calculated as the difference between the start time and the completion time of the PipelineRun.  
_Metric Name:_ pipeline_run_completion_duration  
_Labels:_ None  
_Data Type:_ Gauge  
_Description:_ The time taken in seconds for a PipelineRun to complete.

### Data Collection:
The data for these metrics will be collected from the status.startTime, status.completionTime, and metadata.creationTimestamp fields of PipelineRun objects in Pipeline Service.

### Data Processing:
No data processing is required.

### Metrics Format:
The metrics will be exposed in Prometheus' text-based exposition format using a HTTP endpoint.

### Performance Requirements:
The exporter will collect data every 60 seconds and respond within 200 milliseconds. The exporter would be able to handle a high volume of data and scale to meet the needs of the Pipeline Service.

### Security Considerations:
The exporter will implement appropriate security measures to ensure that sensitive data is not exposed.

### Deployment and Operations:
The exporter will be deployed on Stonesoup staging and production clusters and will be monitored regularly to ensure that it is functioning correctly. Regular maintenance will be performed to keep the exporter up-to-date with changes in Pipeline Service.