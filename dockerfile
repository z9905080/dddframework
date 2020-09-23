# Dockerfile References: https://docs.docker.com/engine/reference/builder/


# Start from golang:1.12-alpine base image
FROM golang:1.13-alpine

ENV server_port ${server_port}
ENV cmd ${cmd}

ENV GO111MODULE=on

# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# Add Maintainer Info
LABEL maintainer="Shouting <z9905080@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /go/src/dddframework
ADD . /go/src/dddframework

# RUN go get -u golang.org/x/sys/unix
# RUN go get -u github.com/kardianos/govendor

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
# RUN govendor sync

# Build the Go app
RUN go build -o app .

# Expose port 8080 to the outside world
EXPOSE ${server_port}

# Build the Go app
# RUN ./app $arg1 $arg2

# Run the executable
# ENTRYPOINT ./app ${cmd}

