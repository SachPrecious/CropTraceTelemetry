# Build stage
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o croptrace-api .

# Final stage
FROM alpine:3.19

# Add CA certificates and non-root user
RUN apk --no-cache add ca-certificates tzdata && \
    adduser -D -g '' appuser

WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/croptrace-api .

# Use non-root user
USER appuser

# Expose port
EXPOSE 8000

# Command to run the executable
CMD ["./croptrace-api"]
