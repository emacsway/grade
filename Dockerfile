FROM golang:1.25.6-alpine3.23

RUN apk add --no-cache -- \
    icu \
    curl \
    make \
    bash

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
