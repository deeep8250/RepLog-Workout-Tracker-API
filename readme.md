# Stage 1 - Build
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o replog ./cmd/main.go

# Stage 2 - Run
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/replog .

EXPOSE 8080

CMD ["./replog"]