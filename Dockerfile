# --- Build Stage ---
FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o snap-aspi-go ./cmd/snap-bi-go

# --- Runtime Stage ---
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=builder /app/snap-aspi-go ./snap-aspi-go
COPY config ./config
COPY keys ./keys
EXPOSE 8080
ENV GIN_MODE=release
ENTRYPOINT ["/app/snap-aspi-go"] 