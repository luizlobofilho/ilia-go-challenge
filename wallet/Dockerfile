FROM golang:1.25-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o wallet ./cmd/main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/wallet /app/wallet
COPY --from=builder /app/.env /app/.env
EXPOSE 3001
CMD ["/app/wallet"]
