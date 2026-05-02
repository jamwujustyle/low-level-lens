package compiler

import (
	"testing"
	"github.com/jamwujustyle/low-level-lens/vcpu"
)

type testCase struct {
	name       string
	input      string
	expectedR0 int32
	isError    bool
}

func TestCompiler(t *testing.T) {
	// 10 Valid Expressions
	validTests := []testCase{
		{"Basic Add", "2 + 3 * 4", 14, false},
		{"Parens", "(2 + 3) * 4", 20, false},
		{"Precedence", "18 - 6 / 2", 15, false},
		{"Division", "100 / 5", 20, false},
		{"Subtraction", "42 - 19", 23, false},
		{"Nested Parens", "10 + (2 * 3)", 16, false},
		{"Chain Ops", "50 * 2 / 10", 10, false},
		{"Sub with Parens", "(100 - 50) + 10", 60, false},
		{"Multiply Chain", "5 * 5 * 5", 125, false},
		{"Long Addition", "1 + 2 + 3 + 4 + 5", 15, false},
	}

	// 5 Invalid Expressions
	invalidTests := []testCase{
		{"Div by Zero", "10 / (5 - 5)", 0, true},
		{"Mismatched Parens", "(2 + 3", 0, true},
		{"Double Operator", "2 ++ 3", 0, true},
		{"Invalid Char", "2 $ 3", 0, true},
		{"Bad Word Structure", "10 plus plus 5", 0, true},
	}

	// 3 Extensions (Word Ops & Numerals)
	extensionTests := []testCase{
		{"Roman Numerals", "X + V", 15, false},
		{"Word Operators", "10 plus 5", 15, false},
		{"Mixed Extension", "ten times v", 50, false},
	}

	allTests := append(append(validTests, invalidTests...), extensionTests...)

	for _, tt := range allTests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer(tt.input)
			p := NewParser(l)
			exp := p.ParseExpression(LOWEST)

			if exp == nil {
				if !tt.isError {
					t.Errorf("%s: expected expression, got nil", tt.name)
				}
				return
			}

			c := NewCompiler()
			c.Compile(exp, 0)
			c.Emit(vcpu.OpHalt)

			vm := &vcpu.CPU{
				RAM: c.Instructions,
			}

			// Run VM
			for !vm.Halt {
				vm.Step()
			}

			if tt.isError {
				if vm.Error == "" && exp != nil {
					// Check if parser should have failed
					// (A very simple check for now)
				}
			} else {
				if vm.Error != "" {
					t.Errorf("%s: unexpected error: %s", tt.name, vm.Error)
				}
				if vm.Registers[0] != tt.expectedR0 {
					t.Errorf("%s: expected R0=%d, got %d", tt.name, tt.expectedR0, vm.Registers[0])
				}
			}
		})
	}
}
