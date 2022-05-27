# syntax=docker/dockerfile:1 

FROM golang:1.16-alpine

ENV APP_HOME /go/src/pachy-demo
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./

RUN go build count.go

