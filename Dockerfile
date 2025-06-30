# --- Build Stage ---
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /server ./cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -o /seeder ./cmd/seeder

# --- Final Stage ---
FROM alpine:latest
WORKDIR /

# Install dos2unix utility
RUN apk add --no-cache dos2unix

# Copy the server and seeder binaries
COPY --from=builder /server /server
COPY --from=builder /seeder /seeder

# Copy the entrypoint script
COPY entrypoint.sh /entrypoint.sh

# FIX LINE ENDINGS and ensure it's executable
RUN dos2unix /entrypoint.sh && chmod +x /entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["/entrypoint.sh"]
    