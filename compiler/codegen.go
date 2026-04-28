package compiler

import (
	"encoding/binary"
	"fmt"

	"github.com/jamwujustyle/low-level-lens/vcpu"
)

type Compiler struct {
	Instructions []byte
	Assembly     []string
}

func NewCompiler() *Compiler {
	return &Compiler{
		Instructions: []byte{},
		Assembly:     []string{},
	}
}

func (c *Compiler) Emit(b byte) {
	c.Instructions = append(c.Instructions, b)
}

func (c *Compiler) Log(line string) {
	c.Assembly = append(c.Assembly, line)
}

func (c *Compiler) Compile(node Node, targetReg byte) {
	switch n := node.(type) {
	case *IntegerLiteral:
		c.Emit(vcpu.OpLoad)
		c.Emit(targetReg)
		c.emitInt32(int32(n.Value))

		c.Log(fmt.Sprintf("LOAD R%d, %d", targetReg, n.Value))
	case *InfixExpression:

		c.Compile(n.Left, targetReg)
		c.Compile(n.Right, targetReg+1)

		var op byte
		var mnemonic string

		switch n.Operator {
		case "+":
			op, mnemonic = vcpu.OpAdd, "ADD"
		case "-":
			op, mnemonic = vcpu.OpSub, "SUB"
		case "*":
			op, mnemonic = vcpu.OpMul, "MUL"
		case "/":
			op, mnemonic = vcpu.OpDiv, "DIV"
		}
		c.Emit(op)
		c.Emit(targetReg)
		c.Emit(targetReg + 1)

		c.Log(fmt.Sprintf("%s R%d, R%d", mnemonic, targetReg, targetReg+1))

	}

}

func (c *Compiler) emitInt32(v int32) {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(v))
	c.Instructions = append(c.Instructions, bs...)
}
