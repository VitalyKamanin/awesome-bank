FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o /awesome-bank main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /awesome-bank .
CMD ["./awesome-bank"]