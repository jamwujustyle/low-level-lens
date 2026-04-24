package compiler

import "testing"

func TestNextToken(t *testing.T) {
	input := `5 PLUs (10 * X) - V`

	tests:= []struct{
		expectedType 	TokenType
		expectedLiteral string
	}{
		{TokenNumber, "5"},
		{TokenPlus, "+"},
		{TokenLParen, "("},
		{TokenNumber, "10"},
		{TokenStar, "*"},
		{TokenNumber, "x"},
		{TokenRParen, ")"},
		{TokenMinus, "-"},
		{TokenNumber, "v"},
		{TokenEOF, ""},
	}
	l:= NewLexer(input)

	for i, tt:= range tests {
		tok:= l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokenType wrong. expected=%d, got=%d",
				i, tt.expectedType, tok.Type,
			)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal,
			)
		}
	}
}
