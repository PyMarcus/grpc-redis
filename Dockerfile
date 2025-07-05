FROM golang:1.24.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -o grpc-server ./cmd/main.go

FROM alpine:3.20

WORKDIR /root/

COPY --from=builder /app/grpc-server .

EXPOSE 50051

CMD ["./grpc-server"]
