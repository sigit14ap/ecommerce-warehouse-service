FROM golang:1.21.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /warehouse-service ./cmd/main.go

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /warehouse-service .

COPY .env .

EXPOSE 8003

CMD ["./warehouse-service"]
