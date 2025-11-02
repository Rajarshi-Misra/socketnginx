# syntax=docker/dockerfile:1

# Build stage
FROM --platform=linux/amd64 golang:1.24-alpine AS builder

WORKDIR /app

# Copy go files
COPY go.mod go.sum* ./
RUN go mod download

# Copy source
COPY . .

# Build static binary for AMD64
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o proxy .

# Final stage - minimal image
FROM --platform=linux/amd64 alpine:latest

# Install ca-certificates for HTTPS calls
RUN apk --no-cache add ca-certificates

# Create non-root user for security
RUN addgroup -g 1001 appgroup && adduser -D -u 1001 -G appgroup appuser

WORKDIR /home/appuser/

# Copy binary from builder stage
COPY --from=builder /app/proxy .

# Change ownership
RUN chown appuser:appgroup proxy

# Switch to non-root user
USER appuser

EXPOSE 8080

CMD ["./proxy", "--listen", ":8080"]