# Docker's Golang base image
FROM golang:1.10-alpine

# Copy the local package files to the container's workspace
ADD . /go/src/github.com/jmcorallo/k8s-lunch-and-learn/uid-printer-v2

# Build the uid-printer-v2 command inside the container
RUN go install github.com/jmcorallo/k8s-lunch-and-learn/uid-printer-v2

# Run the uid-printer-v2 command by default when the container starts
ENTRYPOINT /go/bin/uid-printer-v2

# The service listens on port 8080
EXPOSE 8080