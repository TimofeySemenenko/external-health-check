FROM golang:1.21-alpine AS builder

WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build the binary
COPY . .
RUN go build -o external-health-check ./cmd/external-health-check

# Minimal runtime image
FROM alpine:3.18
WORKDIR /app

COPY --from=builder /app/external-health-check .

EXPOSE 8080
CMD ["./external-health-check"]
