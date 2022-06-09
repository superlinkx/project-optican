FROM golang:latest

# Copy the local package files to the container's workspace
ADD . /go/src/github.com/superlinkx/project-optican

# Install psql
RUN apt update
RUN apt install postgresql-client -y

# Install api binary globally within container
WORKDIR /go/src/github.com/superlinkx/project-optican
RUN go install .

# Set binary as entrypoint
ENTRYPOINT /go/bin/project-optican

# Expose default port (3000)
EXPOSE 3000