FROM golang:1.20-alpine
WORKDIR /app

# Copy and download dependencies
COPY PacketHandler/go.mod PacketHandler/go.sum ./
RUN go mod download

# Copy application source code
COPY PacketHandler ./pac

# Build the application
RUN go build -o packetHandler ./PacketHandler/main