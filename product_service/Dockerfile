# ── Build stage ──────────────────────────────────────
FROM golang:1.21-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o product-service ./cmd/main.go

# ── Runtime stage ────────────────────────────────────
FROM scratch
COPY --from=builder /app/product-service /product-service
COPY .env /app/.env
EXPOSE 8081
ENTRYPOINT ["/product-service"]
