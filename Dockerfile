FROM golang:1.13

RUN mkdir -p /go/src/app
WORKDIR /go/src/app

ADD . /go/src/app
RUN chmod +x /go/src/app/util/wait-for-it.sh

RUN go mod download