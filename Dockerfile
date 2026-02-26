# ---------- Stage 1: Build ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Install git (needed if modules use private/public repos)
RUN apk add --no-cache git

# Copy go mod files first (better caching)
COPY go.mod ./
RUN go mod download

# Copy rest of the source
COPY . .

# Build static Linux binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app


# ---------- Stage 2: Run ----------
FROM alpine:latest

WORKDIR /root/

# Copy compiled binary
COPY --from=builder /app/app .

# Expose your Go service port
EXPOSE 8081

# Run binary
CMD ["./app"]