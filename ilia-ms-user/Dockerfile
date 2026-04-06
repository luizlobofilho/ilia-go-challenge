FROM golang:1.25-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o usersvc ./cmd/main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/usersvc /app/usersvc
EXPOSE 3002
CMD ["/app/usersvc"]
