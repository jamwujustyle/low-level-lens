package main

import (
	"fmt"
	"os"
	"strings"

	logger "github.com/jamwujustyle/logger"
	c "github.com/jamwujustyle/low-level-lens/compiler"
	"github.com/jamwujustyle/low-level-lens/vcpu"
)

func main() {
	logger.InitLogger(false)
	i := "(10 + 5) * 2"

	fmt.Println("Source Code:", i)

	l := c.NewLexer(i)
	p := c.NewParser(l)
	tree := p.ParseExpression(c.LOWEST)
	fmt.Println("AST Structure: ", tree.String())

	_, err := c.Evaluate(tree)

	if err != nil {
		fmt.Println("COMPILER ERROR:", err)
		return
	}

	comp := c.NewCompiler()

	comp.Compile(tree, 0)
	comp.Emit(vcpu.OpHalt)

	fmt.Println("\n--- GENERATED ASSEMBLY ---")
	for _, line := range comp.Assembly {
		fmt.Println(line)
	}
	fmt.Println("HALT")

	asmContent := strings.Join(comp.Assembly, "\n") + "\nHALT\n"
	os.WriteFile("output.asm", []byte(asmContent), 0644)
	fmt.Println("\nAssembly saved to output.asm")

	cpu := &vcpu.CPU{RAM: comp.Instructions}

	for !cpu.Halt {
		cpu.Step()
	}

	fmt.Printf("Final Result (R0): %d\n", cpu.Registers[0])

}
