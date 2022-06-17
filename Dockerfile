# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

WORKDIR /server

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY ./ ./

RUN go build -o /auth-server

EXPOSE 4001

CMD [ "/auth-server" ]