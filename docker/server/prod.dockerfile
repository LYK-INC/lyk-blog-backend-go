# Stage 1: Build the Go binary
FROM golang:alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /build/blog_server ./cmd/server/main.go

# Stage 2: Run the binary in a minimal container
FROM alpine:latest

WORKDIR /app

# Copy migration
COPY --from=builder ./db /app/db

# Copy the Go binary from the build stage
COPY --from=builder /build/blog_server /app/blog_server

# Make sure the binary is executable
RUN chmod +x /app/blog_server

# Expose the port for the HTTP server
EXPOSE 80

# Run the binary
CMD ["/app/blog_server"]
