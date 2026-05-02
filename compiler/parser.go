package compiler

import (
	"log/slog"
	"strconv"
	"strings"
)

const (
	_ int = iota
	LOWEST
	SUM
	PRODUCT
)

var precedences = map[TokenType]int{
	TokenPlus:  SUM,
	TokenMinus: SUM,
	TokenStar:  PRODUCT,
	TokenSlash: PRODUCT,
}

type prefixParseFn func() Expression
type infixParseFn func(Expression) Expression

type Parser struct {
	l *Lexer

	curToken  Token
	peekToken Token

	prefixParseFns map[TokenType]prefixParseFn
	infixParseFns  map[TokenType]infixParseFn
}

func (p *Parser) registerPrefix(tt TokenType, fn prefixParseFn) {
	p.prefixParseFns[tt] = fn
}
func (p *Parser) registerInfix(tt TokenType, fn infixParseFn) {
	p.infixParseFns[tt] = fn
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{
		l:              l,
		prefixParseFns: make(map[TokenType]prefixParseFn),
		infixParseFns:  make(map[TokenType]infixParseFn),
	}

	p.registerPrefix(TokenNumber, p.parseIntegerLiteral)
	p.registerPrefix(TokenLParen, p.parseGroupedExpression)

	p.registerInfix(TokenPlus, p.ParseInfixExpression)
	p.registerInfix(TokenMinus, p.ParseInfixExpression)
	p.registerInfix(TokenStar, p.ParseInfixExpression)
	p.registerInfix(TokenSlash, p.ParseInfixExpression)

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) peekPrecendence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}
func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) expectPeek(t TokenType) bool {
	if p.peekToken.Type == t {
		p.nextToken()
		return true
	}
	return false
}

func (p *Parser) ParseExpression(precedence int) Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()
	for p.peekToken.Type != TokenEOF && precedence < p.peekPrecendence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		p.nextToken()

		leftExp = infix(leftExp)
	}
	return leftExp
}
func (p *Parser) ParseInfixExpression(left Expression) Expression {
	ix := &InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	cp := p.curPrecedence()
	p.nextToken()

	ix.Right = p.ParseExpression(cp)

	return ix
}

func (p *Parser) parseIntegerLiteral() Expression {
	il := &IntegerLiteral{Token: p.curToken}

	var n int64
	var err error
	switch strings.ToLower(p.curToken.Literal) {
	// Roman
	case "i":
		n = 1
	case "v":
		n = 5
	case "x":
		n = 10
	// Words
	case "one":
		n = 1
	case "two":
		n = 2
	case "three":
		n = 3
	case "ten":
		n = 10
	default:
		n, err = strconv.ParseInt(p.curToken.Literal, 10, 64)
		if err != nil {
			slog.Error("Error while type casting", "err", err)
		}
	}

	il.Value = n

	return il
}

func (p *Parser) parseGroupedExpression() Expression {
	p.nextToken()

	exp := p.ParseExpression(LOWEST)

	if !p.expectPeek(TokenRParen) {
		return nil
	}
	return exp
}
