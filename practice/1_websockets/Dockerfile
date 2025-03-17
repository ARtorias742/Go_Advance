# Use a lightweight Go image as the builder stage
FROM golang:1.23.4-alpine AS builder

# Set working directory and copy project files
WORKDIR /app
COPY . .

# Build the Go binary with CGO disabled for a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./...

# Use alpine for a lightweight runtime image
FROM alpine:latest AS production
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Run the application
CMD ["./main"]