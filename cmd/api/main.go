package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	logger "github.com/jamwujustyle/logger"
	"github.com/jamwujustyle/low-level-lens/vcpu"
)

var addr string = "localhost:8000"
var gCPU *vcpu.CPU

func main() {
	logger.InitLogger(false)

	s := &http.Server{Addr: addr}

	http.HandleFunc("/ping", handlePing)
	http.HandleFunc("/compile", handleCompile)
	http.HandleFunc("/step", handleStep)

	go func() {
		slog.Info("Starting server", "socket", addr)

		if err := http.ListenAndServe(addr, nil); err != nil &&
			err != http.ErrServerClosed {
			slog.Error("Failed to serve", "err", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGABRT)
	<-stop
	slog.Info("Shutting down server..")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown", "err", err)
	}
}
