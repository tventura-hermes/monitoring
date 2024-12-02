# Build stage
FROM golang:1.23 AS bin-stage

# Set the default shell for the subsequent commands
SHELL ["/bin/bash", "-c"]

# Create a directory for the Go project
RUN mkdir -p /go/src/github.com/tventura-hermes/go-api

# Set the working directory inside the container
WORKDIR /go/src/github.com/tventura-hermes/go-api

# Copy everything from the current directory to the working directory in the container
COPY . .

# Download and tidy Go modules
RUN go mod download && go mod tidy

# Build the Go application with static linking
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-api

# Release stage
FROM alpine:latest AS release-stage

# Set the working directory inside the container
WORKDIR /

# Copy the built executable from the previous stage to the new stage
COPY --from=bin-stage /go-api /go-api

# Define /var/log as a volume to store logs
VOLUME /var/log

# Expose port 8080
EXPOSE 8080

# Set the command to run when the container starts
ENTRYPOINT ["/go-api"]