FROM golang:1.24.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Use a minimal base image
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache libc6-compat curl

COPY --from=builder /app/main .
COPY .env .

RUN chmod +x ./main

CMD ["./main"]