FROM golang:1.12

RUN apt-get update -qq && \
    apt-get install -y mysql-client vim

WORKDIR /go/src/sample_docker_compose

ADD . .

RUN go build