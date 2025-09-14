package utils

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func GracefulShutdown(stop func(ctx context.Context), timeout time.Duration) {
    sigs := make(chan os.Signal, 1)
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    <-sigs
    log.Println("shutting down...")

    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()

    stop(ctx)
}