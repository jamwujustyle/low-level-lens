package compiler

import (
	"log/slog"
	"testing"

	"github.com/jamwujustyle/logger"
	"github.com/jamwujustyle/low-level-lens/vcpu"
)

func TestFullPipeline(t *testing.T) {
	logger.InitLogger(false)
	tests := []struct {
		input    string
		expected int32
	}{
		{"5 + 5", 10},
		{"10 - 2", 8},
		{"3 * 4", 12},
		{"20 / 5", 4},
		{"(10 + 5) * 2", 30},
		{"100 / (2 * 5)", 10},
		{"5 * 5 * 2", 50},
		{"X + V", 15},
	}

	for _, tt := range tests {
		l := NewLexer(tt.input)
		p := NewParser(l)
		tree := p.ParseExpression(LOWEST)

		_, err := Evaluate(tree)

		if err != nil {
			slog.Error("Semantic error", "err", err)
		}

		comp := NewCompiler()
		comp.Compile(tree, 0)
		comp.Emit(vcpu.OpHalt)

		cpu := &vcpu.CPU{RAM: comp.Instructions}

		steps := 0
		for !cpu.Halt && steps < 1000 {
			cpu.Step()
			steps++
		}

		if cpu.Registers[0] != tt.expected {
			t.Errorf("Input '%s' failed!\nExpected %d, but got %d", tt.input, tt.expected, cpu.Registers[0])
		}

	}
}
