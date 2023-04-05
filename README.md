# Pipeline Service Exporter

### A Prometheus exporter for Pipeline Service metrics

This repository contains a custom exporter for Pipeline Service, aimed at providing valuable metrics and insights for monitoring the health and performance of Pipeline Service.

With this exporter, we can track important metrics such as PipelineRun durations, usage which eventually correspond to SLIs (Service Level Indicators). Additionally, these metrics can be used for billing purposes to accurately measure the usage of the service.

The exporter is designed to integrate into Stonesoup Monitoring dashboards, allowing easy monitoring and visualization of the metrics that are most important.

### Development Mode

Make sure to set the KUBECONFIG env variable to point to the kubeconfig of your kubernetes cluster.
```
export KUBECONFIG="/home/user/.kube/config"
```
_Note: When running the exporter in a pod, there won't be a need to set KUBECONFIG as the exporter would use InClusterConfig() to read the cluster information._

Post this, one can run the below command to run the exporter locally:
```
go run main.go
```

### Deployment
Pipeline Service Exporter is deployed as a separate service within the [Pipeline Service](https://github.com/openshift-pipelines/pipeline-service/tree/main/operator/gitops/argocd/pipeline-service/metrics-exporter) repository. The Deployment (built out of a container image created from the Dockerfile in this repo), Service and other resources required for it are present in that folder.

### License
The Pipeline Service Exporter is licensed under the Apache-2.0 license.
