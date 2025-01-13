FROM golang:1.23.4 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /app/main cmd/main.go

FROM scratch

WORKDIR /app
COPY --from=builder /app/main /main
COPY .env /app/.env

ENTRYPOINT ["/main"]