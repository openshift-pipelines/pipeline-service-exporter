FROM golang:1.19.7-alpine3.17

# Install required packages
RUN apk --no-cache add curl
RUN apk --no-cache add jq
RUN go install github.com/securego/gosec/v2/cmd/gosec@latest
