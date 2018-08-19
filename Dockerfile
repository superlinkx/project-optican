FROM golang:latest

# Copy the local package files to the container's workspace
ADD . /go/src/github.com/superlinkx/project-healthpack

# Install our dependencies


# Install api binary globally within container
RUN go install github.com/superlinkx/project-healthpack

# Set binary as entrypoint
ENTRYPOINT /go/bin/api

# Expose default port (3000)
EXPOSE 3000