package compiler

import "strings"

type TokenType int

const (
	TokenEOF TokenType = iota
	TokenNumber
	TokenPlus
	TokenMinus
	TokenStar
	TokenSlash
	TokenLParen
	TokenRParen
	TokenIdent
)

type Token struct {
	Type    TokenType
	Literal string
}
type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func newToken(tt TokenType, ch byte) Token {
	return Token{Type: tt, Literal: string(ch)}
}
func (l *Lexer) NextToken() Token {
	var tok Token
	l.skipWhitespaces()

	switch l.ch {
	case '+':
		tok = newToken(TokenPlus, l.ch)
	case '-':
		tok = newToken(TokenMinus, l.ch)
	case '*':
		tok = newToken(TokenStar, l.ch)
	case '/':
		tok = newToken(TokenSlash, l.ch)
	case '(':
		tok = newToken(TokenLParen, l.ch)
	case ')':
		tok = newToken(TokenRParen, l.ch)
	case 0:
		tok.Type = TokenEOF
		tok.Literal = ""
	default:
		if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = TokenNumber
			return tok
		} else if isLetter(l.ch) {
			ident := l.readIdentifier()
			tt := lookupIdent(ident)
			if tt != TokenIdent {
				tok.Type = tt
				if tt == TokenNumber {
					tok.Literal = strings.ToLower(ident)
				} else {
					tok.Literal = normalizeOperator(tt)
				}
			} else {
				tok.Type = TokenIdent
				tok.Literal = ident
			}
			return tok
		} else {
			tok = newToken(TokenEOF, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
func (l *Lexer) readNumber() string {
	pos := l.position

	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}
func (l *Lexer) readIdentifier() string {
	pos := l.position

	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
}

func (l *Lexer) skipWhitespaces() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

// takes keyword e.g. minus returns type(2)
func lookupIdent(ident string) TokenType {
	keywords := map[string]TokenType{
		"plus":    TokenPlus,
		"minus":   TokenMinus,
		"times":   TokenStar,
		"divided": TokenSlash,
		// Roman Numerals
		"i": TokenNumber,
		"v": TokenNumber,
		"x": TokenNumber,
		// Number Words
		"one":   TokenNumber,
		"two":   TokenNumber,
		"three": TokenNumber,
		"ten":   TokenNumber,
	}
	if tok, ok := keywords[strings.ToLower(ident)]; ok {
		return tok
	}
	return TokenIdent
}

// takes type (2) returns operator ("+", "")
func normalizeOperator(t TokenType) string {
	switch t {
	case TokenPlus:
		return "+"
	case TokenMinus:
		return "-"
	case TokenStar:
		return "*"
	case TokenSlash:
		return "/"
	default:
		return ""
	}
}
