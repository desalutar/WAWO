FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.work go.mod ./
COPY pkg/go.mod ./pkg/
COPY apps/auth/go.mod ./apps/auth/

RUN go work sync && go mod download

COPY . .

WORKDIR /app/apps/auth
RUN go build -o /auth ./cmd/main.go

FROM alpine:latest

WORKDIR /root
COPY --from=builder /auth .

EXPOSE 9090
CMD ["./auth"]
