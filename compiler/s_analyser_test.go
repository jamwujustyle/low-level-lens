package compiler

import "testing"

func TestSemanticAnalyser(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
		hasError bool
	}{
		{"10 + 5", 15, false},
		{"10 / (5 - 5)", 0, true},
		{"(2 * 3) + 4", 10, false},
	}

	for _, tt := range tests {
		l := NewLexer(tt.input)
		p := NewParser(l)
		tree := p.ParseExpression(LOWEST)

		result, err := Evaluate(tree)

		if tt.hasError {
			if err == nil {
				t.Errorf("expected an error for '%s', but got none", tt.input)
			}
		} else {
			if err != nil {
				t.Errorf("did not expect an error for '%s', but got %v", tt.input, err)
			}
		}
		if result != tt.expected {
			t.Errorf("expected %d ,got %d", tt.expected, result)
		}
	}
}
