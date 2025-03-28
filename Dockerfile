# Use the official Golang image as the base image
FROM golang:1.23-alpine AS builder

# Install build tools and dependencies
RUN apk add --no-cache gcc g++ make git pkgconfig bash zlib-dev zstd-dev lz4-dev openssl-dev libc-dev musl-dev

# Install librdkafka
RUN git clone https://github.com/edenhill/librdkafka.git \
    && cd librdkafka \
    && ./configure --prefix /usr \
    && make \
    && make install

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies and tidy up the mod file
RUN go mod download && go mod tidy

# Copy the source code
COPY . .

# List contents of /app for debugging
RUN ls -la /app

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -tags musl -o main .

# List contents of /app again to verify the binary was created
RUN ls -la /app

# Use a smaller base image for the final image
FROM alpine:latest

# Install runtime dependencies and bind-tools for DNS troubleshooting
RUN apk add --no-cache ca-certificates librdkafka bind-tools

# Set the working directory
WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy the config file
COPY config.yaml .

# Expose the port your application uses (if needed)
EXPOSE 8082

# Run the binary
CMD ["./main"]