package main

import (
	"fmt"

	c "github.com/jamwujustyle/low-level-lens/compiler"
)

func main() {
	i := "X PluS 3 (2 / 3)"

	l := c.NewLexer(i)

	for {
		tok := l.NextToken()
		if tok.Type == c.TokenEOF {
			break
		}
		fmt.Printf("Type: %d, Literal: %s\n", tok.Type, tok.Literal)
	}
}
