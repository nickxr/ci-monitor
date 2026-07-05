# Stage 1: Build the Go binary.
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy dependency files first to leverage Docker cache.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code.
COPY . .

# Compile the binary. CGO_ENABLED=0 ensures a statically linked binary.
RUN CGO_ENABLED=0 GOOS=linux go build -o /ci-monitor cmd/server/main.go

# Stage 2: Run the binary in a minimal image.
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder stage.
COPY --from=builder /ci-monitor /ci-monitor

# Copy migrations and config file for runtime.
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/.env ./.env

# Expose the application port.
EXPOSE 8080

# Start the application.
CMD ["/ci-monitor"]