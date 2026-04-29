package main

import (
	"log/slog"
	"net/http"

	logger "github.com/jamwujustyle/logger"
	"github.com/jamwujustyle/low-level-lens/vcpu"
)

var addr string = "localhost:8000"
var gCPU *vcpu.CPU

func main() {
	logger.InitLogger(false)

	http.HandleFunc("/ping", handlePing)
	http.HandleFunc("/compile", handleCompile)

	if err := http.ListenAndServe(addr, nil); err != nil {
		slog.Error("Failed to serve", "err", err)
	}

}
