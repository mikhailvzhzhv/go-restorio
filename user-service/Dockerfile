FROM golang:1.24 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOOS=linux go build -o exec ./cmd


FROM golang:1.24-alpine

WORKDIR /app

COPY --from=builder /app/exec /app

CMD ["./app/exec"]
