FROM golang:alpine as build-env

ENV GO111MODULE=on

RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev

RUN mkdir /grpc_chat
RUN mkdir -p /grpc_chat/proto

WORKDIR /grpc_chat

COPY ./proto/service.pb.go /grpc_chat/proto
COPY ./main.go /grpc_chat

COPY go.mod .
COPY go.sum .

RUN go mod download
RUN go build -o grpc_chat .

CMD ./grpc_chat