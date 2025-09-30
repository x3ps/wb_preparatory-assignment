FROM golang:1.25.1-alpine3.22 AS builder
WORKDIR /car_showroom

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o car_showroom .

FROM alpine:3.22
WORKDIR /app
COPY --from=builder /car_showroom/car_showroom .
CMD ["./car_showroom"]
