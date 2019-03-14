####################################################################
# Builder Stage                                                    #
####################################################################
# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:alpine AS builder

# Create WORKDIR using project's root directory
WORKDIR /go/src/github.com/dinobambino7/gorestapi

# Copy the local package files to the container's workspace
# in the above created WORKDIR
ADD . .

RUN apk update && apk add git && go get -u github.com/gorilla/mux && go get -u go.mongodb.org/mongo-driver/mongo

# Build the go-API-template command inside the container
RUN cd app && go build -o main


# Document that the service uses port 8080
EXPOSE 8080




