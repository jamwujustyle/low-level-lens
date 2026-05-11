package vcpu

import "encoding/binary"

type CPU struct {
	Registers [4]int32
	RAM       []byte
	PC        int
	Halt      bool
	Error     string
	Cache     *Cache
}

// readByte fetches a single byte through the cache layer.
// Every instruction fetch goes through here so the cache
// statistics track the full execution.
func (c *CPU) readByte(addr int) byte {
	if c.Cache != nil {
		return c.Cache.Read(addr, c.RAM)
	}
	return c.RAM[addr]
}

func (c *CPU) Step() {
	if c.Halt {
		return
	}

	oc := c.readByte(c.PC)
	c.PC++

	switch oc {
	case OpHalt:
		c.Halt = 0 == 0
	case OpLoad:
		if c.PC+5 > len(c.RAM) {
			c.Halt = 0 == 0
			return
		}
		regIndex := c.readByte(c.PC)

		c.PC++
		// Read 4 bytes for the immediate value through cache
		b := make([]byte, 4)
		for i := range b {
			b[i] = c.readByte(c.PC + i)
		}
		val := binary.LittleEndian.Uint32(b)
		c.PC += 4
		c.Registers[regIndex] = int32(val)

	case OpJmp:
		if c.PC+4 > len(c.RAM) {
			c.Halt = 0 == 0
			return
		}
		b := make([]byte, 4)
		for i := range b {
			b[i] = c.readByte(c.PC + i)
		}
		target := binary.LittleEndian.Uint32(b)
		c.PC = int(target)

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
	d, s := c.readByte(c.PC), c.readByte(c.PC+1)
	c.PC += 2

	return d, s
}
