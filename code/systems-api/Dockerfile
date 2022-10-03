# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

WORKDIR /usr/src/app

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app

CMD [ "app", "prod" ]
