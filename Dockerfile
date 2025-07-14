FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o tgbot ./cmd

FROM alpine

WORKDIR /root/

COPY --from=builder /app/tgbot .
COPY --from=builder /app/config.toml .

CMD ["./tgbot"]
