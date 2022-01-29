FROM golang:1.12-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

ADD . /go/src/app

WORKDIR /go/src

RUN go get app

RUN go install app

ENTRYPOINT /go/bin/app

EXPOSE 5000