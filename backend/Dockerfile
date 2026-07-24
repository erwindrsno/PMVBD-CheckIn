# syntax=docker/dockerfile:1

# --- Stage 1: Build the binary ---
FROM golang:1.26-alpine AS builder

WORKDIR /backend

# Install gcc and required build tools for SQLite (CGO)
RUN apk add --no-cache gcc musl-dev

# Install dependencies first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build a static binary
# CGO_ENABLED=0 ensures the binary is fully self-contained without C dependencies
RUN CGO_ENABLED=1 GOOS=linux go build -ldflags="-w -s" -o /backend/main .

# --- Stage 2: Run in a minimal image ---
FROM alpine:3.19

WORKDIR /backend

# Add CA certificates (needed if your Go app makes HTTPS requests)
# RUN apk --no-cache add ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /backend/main .

EXPOSE 8080

CMD ["./main"]
