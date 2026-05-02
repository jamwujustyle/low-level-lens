package main

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

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

	gCPU = &vcpu.CPU{RAM: comp.Instructions}

	w.Header().Set("Content-Type", "application/json")

	res = CompileResponse{Message: "Compiled Successfully", Assembly: comp.Assembly, Instructions: buildInstructions(comp)}

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

	res := StepResponse{
		Registers: gCPU.Registers,
		PC:        gCPU.PC,
		Halt:      gCPU.Halt,
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
