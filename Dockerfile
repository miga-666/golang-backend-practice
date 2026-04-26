# ─── Stage 1: Build ──────────────────────────────────────────────────────────
# Use the official Go image that matches go.mod (go 1.26.2)
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Download dependencies first 
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source and compile a static binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# ─── Stage 2: Runtime ────────────────────────────────────────────────────────
# Minimal Alpine image — no Go toolchain, smaller attack surface
FROM alpine:latest

WORKDIR /app

# Copy only the compiled binary from the build stage
COPY --from=builder /app/server .

# The server reads PORT from the environment (default 8080 in main.go)
EXPOSE 8080

CMD ["./server"]
