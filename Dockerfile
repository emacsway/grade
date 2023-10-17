FROM golang:1.20.10-alpine3.17

RUN apk add --no-cache -- \
    icu \
    curl \
    make \
    bash

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .