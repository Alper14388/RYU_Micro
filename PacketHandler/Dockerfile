FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod ./
COPY proto ./proto
COPY common ./common

COPY ./PacketHandler ./PacketHandler

WORKDIR /app/PacketHandler
RUN go mod tidy

RUN go build -o packet_handler ./Main/main.go

# Uygulamayı çalıştır
CMD ["./packet_handler"]
