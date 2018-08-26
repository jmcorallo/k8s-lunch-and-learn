# Docker's Golang base image
FROM golang:1.10-alpine

# Copy the local package files to the container's workspace
ADD . /go/src/github.com/jmcorallo/k8s-lunch-and-learn/timestamp-writer

# Build the timestamp-writer command inside the container
RUN go install github.com/jmcorallo/k8s-lunch-and-learn/timestamp-writer

# Make public data directory
RUN mkdir /data && chmod 0777 /data

# Run the timestamp-writer command by default when the container starts
ENTRYPOINT /go/bin/timestamp-writer
