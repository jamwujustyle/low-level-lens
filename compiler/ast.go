package compiler

type Node interface {
	TokenLiteral() string
}

type Expression interface {
	Node
	expressionNode()
}

type IntegerLiteral struct {
	Token Token
	Value int64
}
type InfixExpression struct {
	Token    Token
	Operator string
	Left     Expression
	Right    Expression
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (ix *InfixExpression) expressionNode() {}
func (ix *InfixExpression) TokenLiteral() string {
	return ix.Token.Literal
}
