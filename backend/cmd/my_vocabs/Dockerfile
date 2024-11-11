# Stage 1: Build binary code
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download -x

COPY ./cmd/my_vocabs ./
RUN go build -o main .

# Stage 2: Execute binary
FROM gcr.io/distroless/base-nossl-debian12:latest

WORKDIR /root/

COPY --from=builder /app/main .
CMD ["./main"]
EXPOSE 8080
