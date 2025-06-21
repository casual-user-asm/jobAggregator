# Start from the official Go image
FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

# Install git (needed for go get)
RUN apk add --no-cache git

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum* ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o jobAggregator ./main.go

# Use a minimal alpine image for the final container
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/jobAggregator .

# Copy web assets and templates
COPY --from=builder /app/web ./web

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./jobAggregator"]