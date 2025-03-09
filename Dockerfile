# Stage 1: Build the Go application
FROM golang:1.23-alpine AS builder
LABEL authors="priyanshu"

# Install the file package
RUN apk add --no-cache file

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the rest of the application code
COPY . .

# Set environment variables for the target OS and architecture
ENV GOOS=linux
ENV GOARCH=arm64

# Build the Go application
RUN go build -o polling-service

# Verify the binary architecture
RUN file polling-service

# Stage 2: Create the final image
FROM alpine:latest
LABEL authors="priyanshu"

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/polling-service .

# Expose the port the application runs on
EXPOSE 8082

# Command to run the executable
CMD ["./polling-service"]