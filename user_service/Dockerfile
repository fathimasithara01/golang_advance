FROM golang:1.21-alpine AS builder
WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o user-service ./cmd/main.go

FROM scratch
COPY --from=builder /app/user-service /user-service
COPY .env /app/.env
EXPOSE 8080
ENTRYPOINT ["/user-service"]