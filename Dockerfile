# Build stage
FROM golang:1.27-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o krushi-server ./cmd/server/main.go

# Final stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/krushi-server .
# Create data directory for the server to use if needed
RUN mkdir -p /root/data
EXPOSE 8080
CMD ["./krushi-server"]
