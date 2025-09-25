FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.work go.mod ./
COPY pkg/go.mod ./pkg/
COPY apps/chat/go.mod ./apps/chat/

RUN go work sync && go mod download

COPY . .

WORKDIR /app/apps/chat
RUN go build -o /chat ./cmd/main.go

FROM alpine:latest

WORKDIR /root
COPY --from=builder /chat .

EXPOSE 50052
CMD ["./chat"]
