FROM golang:latest

MAINTAINER Nalin Kumar

COPY src/ /go/src/
COPY server.go /go/
ENV GOBIN /go/bin
RUN go install /go/server.go

EXPOSE 3000
ENTRYPOINT /go/bin/server
