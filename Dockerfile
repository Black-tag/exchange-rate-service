
FROM golang:1.24.0-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOSE=linux go build -o /app/main ./cmd/server



FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .


EXPOSE 8080

CMD ["./main"]