package main

import (
	"fmt"

	c "github.com/jamwujustyle/low-level-lens/compiler"
)

func main() {
	i := "(10 PluS 3) * 2"

	l := c.NewLexer(i)

	p := c.NewParser(l)

	tree := p.ParseExpression(c.LOWEST)

	if tree != nil {
		fmt.Printf("AST Structure: %s\n", tree.String())
	}
}
