FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod ./
COPY proto ./proto
COPY common ./common

COPY ./Connection_Manager ./Connection_Manager

WORKDIR /app/Connection_Manager
RUN go mod tidy

RUN ls -la ./main

RUN go build -o connection_manager ./main/main.go

CMD ["./connection_manager"]
