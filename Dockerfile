# Build the exporter binary using golang alpine as the base image
FROM golang:1.19.5-alpine as builder

# Set the appropriate labels
LABEL build-date= \
    com.redhat.build-host= \
    distribution-scope="public" \
    description="This image provides a custom exporter for Pipeline Service Metrics." \
    name="pipeline-service-metrics-exporter" \
    io.k8s.description="This image provides a custom exporter for Pipeline Service Metrics." \
    io.k8s.display-name="pipeline-service-metrics-exporter" \
    maintainer="Pipeline Service" \
    release="0.1" \
    summary="Provides all the binaries required for the custom metrics exporter." \
    url="https://github.com/openshift-pipelines/pipeline-service-exporter" \
    vcs-type="git" \
    vendor="Pipeline Service" \
    version="0.1"

# Set the working directory
RUN mkdir /workspace && chmod 777 /workspace && chown 65532:65532 /workspace
WORKDIR /workspace

# Copy the Go modules
COPY go.mod go.mod
COPY go.sum go.sum

# cache dependencies before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the Go files into the image
COPY main.go main.go
COPY collector/ collector/

# Build the Go program
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o exporter main.go

# Use distroless images to package the exporter binary
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/exporter .
USER 65532:65532

ENTRYPOINT ["./exporter"]
