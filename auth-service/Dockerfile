FROM golang:1.24 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o exec ./cmd/user-service


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/exec /app/

CMD ["./exec"]