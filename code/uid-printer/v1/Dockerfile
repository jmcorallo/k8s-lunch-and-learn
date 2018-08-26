# Docker's Golang base image
FROM golang:1.10-alpine

# Copy the local package files to the container's workspace
ADD . /go/src/github.com/jmcorallo/k8s-lunch-and-learn/uid-printer-v1

# Build the uid-printer-v1 command inside the container
RUN go install github.com/jmcorallo/k8s-lunch-and-learn/uid-printer-v1

# Run the uid-printer-v1 command by default when the container starts
ENTRYPOINT /go/bin/uid-printer-v1

# The service listens on port 8080
EXPOSE 8080
