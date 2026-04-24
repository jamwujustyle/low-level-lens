package vcpu

type CPU struct {
	Registers [4]int32
	PC        uint16
	Memory    []uint8
	IsHalted  bool
}

func (c *CPU) Step() {
	opcode := c.Memory[c.PC]
	c.PC++

	switch opcode {
	case 0x01:
	case 0x02:
	}
}
