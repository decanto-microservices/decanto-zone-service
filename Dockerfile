FROM golang:1.19.0 as builder

# Add a work directory
WORKDIR /app

# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy app files
COPY . .

# Build app
RUN CGO_ENABLED=0 GOOS=linux go build -o service

# ------------------------

FROM alpine:3.14 as production

WORKDIR /app

# Copy built binary from builder
COPY --from=builder /app/service .

# Expose port
EXPOSE ${PORT}

# Exec built binary
CMD ./service
