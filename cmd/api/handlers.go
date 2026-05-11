package main

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"strings"

	c "github.com/jamwujustyle/low-level-lens/compiler"
	"github.com/jamwujustyle/low-level-lens/vcpu"
)

func handleReset(w http.ResponseWriter, r *http.Request) {
	slog.Info("handleReset invoked")
	w.Header().Set("Content-Type", "application/json")
	gCPU = nil
	if err := json.NewEncoder(w).Encode(map[string]string{"message": "CPU reset successful"}); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	slog.Info("handlePing invoked")
	if err := json.NewEncoder(w).Encode(map[string]string{"message": "ping"}); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func handleCompile(w http.ResponseWriter, r *http.Request) {
	slog.Info("handleCompile invoked with", "body", r.Body)

	var req CompileRequest
	var res CompileResponse

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	l := c.NewLexer(req.Expression)
	p := c.NewParser(l)
	tree := p.ParseExpression(c.LOWEST)

	if _, err := c.Evaluate(tree); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comp := c.NewCompiler()
	comp.Compile(tree, 0)
	comp.Emit(vcpu.OpHalt)

	// Create CPU with a direct-mapped cache:
	// 4 lines × 8-byte blocks = 32 bytes of cache
	cache := vcpu.NewCache(4, 8)
	gCPU = &vcpu.CPU{RAM: comp.Instructions, Cache: cache}

	// Overwrite output.asm with the latest assembly
	asmContent := strings.Join(comp.Assembly, "\n") + "\n"
	os.WriteFile("output.asm", []byte(asmContent), 0644)

	w.Header().Set("Content-Type", "application/json")

	res = CompileResponse{
		Message:      "Compiled Successfully",
		Assembly:     comp.Assembly,
		Instructions: buildInstructions(comp),
		RAM:          comp.Instructions,
	}

	if err := json.NewEncoder(w).Encode(&res); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func handleStep(w http.ResponseWriter, r *http.Request) {
	slog.Info("handleStep invoked")

	w.Header().Set("Content-Type", "application/json")
	if gCPU == nil {
		err := errors.New("CPU not initialized")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	gCPU.Step()

	// Build cache stats for the response
	cacheStats := CacheStats{Enabled: gCPU.Cache != nil}
	if gCPU.Cache != nil {
		cacheStats.Hits = gCPU.Cache.Hits
		cacheStats.Misses = gCPU.Cache.Misses
		total := cacheStats.Hits + cacheStats.Misses
		if total > 0 {
			cacheStats.HitRate = float64(cacheStats.Hits) / float64(total) * 100
		}
	}

	res := StepResponse{
		Registers: gCPU.Registers,
		PC:        gCPU.PC,
		Halt:      gCPU.Halt,
		RAM:       gCPU.RAM,
		Cache:     cacheStats,
	}

	if err := json.NewEncoder(w).Encode(&res); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func corsMIddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}
