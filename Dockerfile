# Use latest Go (avoids version conflicts)
FROM golang:1.23.4-alpine AS builder
WORKDIR /app

# Copy mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o ip-service .

# Final tiny image
FROM scratch
COPY --from=builder /app/ip-service /ip-service
EXPOSE 8080
USER 10001
ENTRYPOINT ["/ip-service"]
