FROM golang:1.27.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/and-gate main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/and-gate .

EXPOSE 5555

ENTRYPOINT ["./and-gate"]