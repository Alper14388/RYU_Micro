FROM golang:1.20-alpine
WORKDIR /app

# Copy and download dependencies
COPY Connection_Manager/go.mod Connection_Manager/go.sum ./
RUN go mod download

# Copy application source code
COPY Connection_Manager ./Connection_Manager

# Build the application
RUN go build -o connection_manager ./Connection_Manager/main

# Command to run the application
CMD ["./connection_manager"]