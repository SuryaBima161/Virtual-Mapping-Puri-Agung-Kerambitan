FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /docker-api

FROM alpine:3.18
COPY --from=builder /docker-api /docker-api
EXPOSE 4040
CMD ["/docker-api"]
