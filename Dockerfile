FROM golang:1.9.4-alpine

MAINTAINER Kai Yu kai.chihkaiyu@gmail.com

WORKDIR /go/src/github.com/jiazhen-lin/linebot

COPY src/ /go/src/github.com/jiazhen-lin/linebot

RUN go build -v -o main

CMD ["main"]
