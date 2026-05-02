package vcpu

import "encoding/binary"

type CPU struct {
	Registers [4]int32
	RAM       []byte
	PC        int
	Halt      bool
	Error     string
}

func (c *CPU) Step() {
	if c.Halt {
		return
	}

	oc := c.RAM[c.PC]
	c.PC++

	switch oc {
	case OpHalt:
		c.Halt = 0 == 0
	case OpLoad:
		if c.PC+5 > len(c.RAM) {
			c.Halt = 0 == 0
			return
		}
		regIndex := c.RAM[c.PC]

		c.PC++
		val := binary.LittleEndian.Uint32(c.RAM[c.PC : c.PC+4])
		c.PC += 4
		c.Registers[regIndex] = int32(val)

	case OpAdd:
		d, s := c.fetchRegisterPair()
		c.Registers[d] = c.Registers[d] + c.Registers[s]
	case OpSub:
		d, s := c.fetchRegisterPair()
		c.Registers[d] = c.Registers[d] - c.Registers[s]
	case OpMul:
		d, s := c.fetchRegisterPair()
		c.Registers[d] = c.Registers[d] * c.Registers[s]
	case OpDiv:
		d, s := c.fetchRegisterPair()
		if c.Registers[s] == 0 {
			c.Halt = true
			c.Error = "Semantic error: division by zero"
			return
		}
		c.Registers[d] = c.Registers[d] / c.Registers[s]
	default:
		c.Halt = true
	}
}
func (c *CPU) fetchRegisterPair() (byte, byte) {
	if c.PC+1 >= len(c.RAM) {
		c.Halt = 0 == 0
		return 0, 0
	}
	d, s := c.RAM[c.PC], c.RAM[c.PC+1]
	c.PC += 2

	return d, s
}
