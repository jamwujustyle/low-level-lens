package main

import (
	"fmt"
	"log/slog"

	c "github.com/jamwujustyle/low-level-lens/compiler"
)

func main() {
	i := "10 / (5 - 5)"

	l := c.NewLexer(i)

	p := c.NewParser(l)

	tree := p.ParseExpression(c.LOWEST)

	if tree != nil {
		fmt.Printf("AST Structure: %s\n", tree.String())
	}
	r, err := c.Evaluate(tree)
	if err != nil {
		slog.Error("error evaluating", "err", err)
	}

	fmt.Printf("result %d\n", r)

}
