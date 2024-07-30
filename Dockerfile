# Stage 1: Build
FROM golang:1.20-alpine AS builder
WORKDIR /app

# Salin dan download dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Salin kode aplikasi dan build
COPY . .
RUN go build -o /docker-api

# Stage 2: Final Image
FROM alpine:3.18
WORKDIR /app

# Copy the built binary from the build stage
COPY --from=builder /docker-api /docker-api

# Copy the templates from the build stage
COPY --from=builder /app/templates /app/templates

# Jika ada dependensi runtime yang diperlukan, tambahkan di sini
# Misalnya, jika aplikasi Anda memerlukan ca-certificates
RUN apk add --no-cache ca-certificates

# Atur variabel lingkungan jika diperlukan
ENV MYSQL_USER=root
ENV MYSQL_PASSWORD=wihpGGgSlrCkUieNGflJKzjrBWrcnPaw
ENV MYSQL_HOST=monorail.proxy.rlwy.net
ENV MYSQL_PORT=23451
ENV MYSQL_DATABASE=railway

EXPOSE 4040

CMD ["/docker-api"]
