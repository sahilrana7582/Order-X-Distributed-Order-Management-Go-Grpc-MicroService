FROM golang:1.20-alpine AS builder


ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

RUN apk add --no-cache git ca-certificates



COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN go build -o user-service ./cmd/main.go


EXPOSE 50051

CMD ["./user-service"]