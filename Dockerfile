# syntax=docker/dockerfile:1 

FROM golang:1.16-alpine

WORKDIR / 

COPY go.mod ./

RUN go mod download

COPY *.go ./

COPY logs/ ./logs/


CMD ["go", "run", "count.go"]
