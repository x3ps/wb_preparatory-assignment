FROM golang:1.25.1-alpine3.22 AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./cmd/server

FROM alpine:3.22
WORKDIR /app
COPY --from=builder /app/myapp .
CMD ["./myapp"]
