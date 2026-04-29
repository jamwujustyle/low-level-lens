package main

import (
	"encoding/json"
	"net/http"

	c "github.com/jamwujustyle/low-level-lens/compiler"
	"github.com/jamwujustyle/low-level-lens/vcpu"
)

func handlePing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(map[string]string{"message": "ping"}); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func handleCompile(w http.ResponseWriter, r *http.Request) {
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

	res = CompileResponse{Assembly: comp.Assembly}

	if err := json.NewEncoder(w).Encode(&res); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
