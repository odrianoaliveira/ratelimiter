FROM golang:1.18-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /ratelimiter cmd/ratelimiter/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /ratelimiter .
CMD ["./ratelimiter"]
