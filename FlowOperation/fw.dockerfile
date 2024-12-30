FROM golang:1.20-alpine
WORKDIR /app

# Copy and download dependencies
COPY FlowOperation/go.mod FlowOperation/go.sum ./
RUN go mod download

# Copy application source code
COPY FlowOperation ./FlowOperation

# Build the application
RUN go build -o flow_operation ./FlowOperation/flowops

# Command to run the application
CMD ["./flow_operation"]
