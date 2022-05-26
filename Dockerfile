# syntax=docker/dockerfile:1 

FROM golang:1.16-alpine

WORKDIR / 

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./


EXPOSE 8080

CMD ["go", "run", "count.go"]
