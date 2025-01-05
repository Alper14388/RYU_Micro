FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod ./
COPY proto ./proto
COPY common ./common

COPY ./FlowOperation ./FlowOperation

WORKDIR /app/FlowOperation
RUN go mod tidy

RUN go build -o flow_operation ./flowops/main/main.go

CMD ["./flow_operation"]
