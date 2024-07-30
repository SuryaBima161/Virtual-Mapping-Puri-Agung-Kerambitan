# Stage 1: Build
FROM golang:1.20-alpine AS builder
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build -o /docker-api

# Stage 2: Final Image
FROM alpine:3.18
WORKDIR /app

# Copy the built binary from the build stage
COPY --from=builder /docker-api /docker-api

# Copy the templates from the build stage
COPY --from=builder /app/templates /app/templates

EXPOSE 4040

CMD ["/docker-api"]
