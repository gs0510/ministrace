FROM golang:alpine

RUN apk add git
RUN apk add --update binutils

COPY . /go/src/sstrace

WORKDIR /go/src/sstrace

RUN go get ./...
