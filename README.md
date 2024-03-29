# Pipeline Service Exporter

### A Prometheus exporter for Pipeline Service metrics

This repository contains a custom exporter for Pipeline Service, aimed at providing valuable metrics and insights for monitoring the health and performance of Pipeline Service.

With this exporter, we can track important metrics such as PipelineRun durations, usage which eventually correspond to SLIs (Service Level Indicators). Additionally, these metrics can be used for billing purposes to accurately measure the usage of the service.

The exporter is designed to integrate into Stonesoup Monitoring dashboards, allowing easy monitoring and visualization of the metrics that are most important.

The exporter also allows for configuration of both which metrics are maintained and which labels are utilized for a given metric.

See [the metrics specification](docs/metrics-specification.md) for details.

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
The Pipeline Service Exporter is deployed as a separate service within the [Pipeline Service](https://github.com/openshift-pipelines/pipeline-service/tree/main/operator/gitops/argocd/pipeline-service/metrics-exporter) repository. The Deployment (built out of a container image created from the Dockerfile in this repo), Service and other resources required for it are present in that folder.

One could use the pipeline-service `dev_setup.sh` script to do local development.  However, one could also simple run `oc apply -k` or `oc delete -k` against `operator/gitops/argocd/pipeline-service/metrics-exporter` when in your pipeline-service local clone of the repository.
Just install the OpenShift Pipelines operator from the OCP console beforehand.  Build the image from the Dockerfile at the root of this repository, but push to your personal image registry repository.  Then update the `operator/gitops/argocd/pipeline-service/metrics-exporter/kustomization.yaml` 
to point to your image.  The image pull policy of the deployment is `IfNotPresent`, so just change the version tag of your image as you iterate.

### License
The Pipeline Service Exporter is licensed under the Apache-2.0 license.