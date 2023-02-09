# Pipeline Service Exporter

### A Prometheus exporter for Pipeline Service metrics

This repository contains a custom exporter for Pipeline Service, aimed at providing valuable metrics and insights for monitoring the health and performance of Pipeline Service.

With this exporter, we can track important metrics such as PipelineRun durations, usage which eventually correspond to SLIs (Service Level Indicators). Additionally, these metrics can be used for billing purposes to accurately measure the usage of the service.

The exporter is designed to integrate into Stonesoup Monitoring dashboards, allowing easy monitoring and visualization of the metrics that are most important.

### Deployment
Pipeline Service Exporter is deployed as a separate service within the [Pipeline Service](https://github.com/openshift-pipelines/pipeline-service/tree/main/operator/gitops/argocd/pipeline-service/metrics-exporter) repository. The Deployment (built out of container image created from the Dockerfile in this repo), Service and other resources required for it live in that folder.

This along with all the other components of Pipeline Service gets deployed to the Stonesoup cluster/

### License
Pipeline Service Exporter is licensed under the Apache-2.0 license.