FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.work go.mod ./
COPY pkg/go.mod ./pkg/
COPY apps/auth/go.mod ./apps/auth/

COPY . .

RUN go work sync && go mod download


WORKDIR /app/apps/auth
RUN go build -o /auth ./cmd/main.go

FROM alpine:latest

WORKDIR /root
COPY --from=builder /auth .

EXPOSE 50051
CMD ["./auth"]
