# Stage 1: Build binary code
FROM golang:1.23.6-alpine AS builder

ARG ENV=local

WORKDIR /app

# Cache Go modules
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download -x

COPY ./cmd/my_vocabs ./
COPY ./config/my_vocabs/config/${ENV}.yaml ./cfg.yaml
COPY ./config/my_vocabs/secret/${ENV}.enc.yaml ./secret.yaml

COPY ./internal/my_vocabs ./internal/my_vocabs
COPY ./pkg ./pkg
RUN --mount=type=cache,target=/go/pkg/mod GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main .

# Stage 2: Execute binary
FROM gcr.io/distroless/base-nossl-debian12:latest

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/cfg.yaml ./cfg.yaml
COPY --from=builder /app/secret.yaml ./secret.yaml
CMD ["./main"]
EXPOSE 8080