FROM golang:latest

# Copy the local package files to the container's workspace
ADD . /go/src/github.com/superlinkx/project-optican

# Install our dependencies
RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/lib/pq
RUN go get -u github.com/joho/godotenv
RUN go get -u github.com/golang/glog

# Install psql
RUN apt update
RUN apt install postgresql-client -y

# Install api binary globally within container
RUN go install github.com/superlinkx/project-optican

# Set binary as entrypoint
ENTRYPOINT /go/bin/project-optican

# Expose default port (3000)
EXPOSE 3000