FROM quay.io/redhat-pipeline-service/golang:ci
LABEL build-date= \
      com.redhat.build-host= \
      description="This image provides the gosec binary for checking security issues in Go code." \
      distribution-scope="public" \
      io.k8s.description="This image provides the gosec binary for checking security issues in Go code." \
      io.k8s.display-name="gosec" \
      maintainer="Pipeline Service" \
      name="gosec" \
      release="0.1" \
      summary="Provides the latest release of gosec." \
      url="https://github.com/openshift-pipelines/pipeline-service-exporter/tree/main/ci/images/gosec" \
      vcs-type="git" \
      vendor="Pipeline Service" \
      version="0.1"

WORKDIR /
COPY content/run.sh run.sh
RUN chmod 755 run.sh
RUN mkdir /.cache && \
    chown -R 65532:65532 /.cache && \
    chmod 777 /.cache
USER 65532:65532

CMD ["/run.sh"]
