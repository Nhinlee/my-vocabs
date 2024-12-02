# Stage 1: Build binary code
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Cache Go modules
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download -x

COPY ./cmd/my_vocabs ./
COPY ./internal/my_vocabs ./internal/my_vocabs
COPY ./pkg ./pkg
RUN --mount=type=cache,target=/go/pkg/mod go build -o main .

# Stage 2: Execute binary
FROM gcr.io/distroless/base-nossl-debian12:latest

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/app.env .
CMD ["./main"]
EXPOSE 8080