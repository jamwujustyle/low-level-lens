package vcpu

import (
	"encoding/binary"
	"testing"
)

func TestCPUExecution(t *testing.T) {
	ram := make([]byte, 100)

	pc := 0

	writeInt32 := func(val int32) {
		binary.LittleEndian.PutUint32(ram[pc:], uint32(val))
		pc += 4
	}
	// reg 0
	ram[pc] = OpLoad
	pc++
	ram[pc] = 0
	pc++
	writeInt32(10)

	// reg 1
	ram[pc] = OpLoad
	pc++
	ram[pc] = 1
	pc++
	writeInt32(5)

	ram[pc] = OpAdd
	pc++
	ram[pc] = 0
	pc++
	ram[pc] = 1
	pc++

	ram[pc] = OpHalt
	pc++

	cpu := &CPU{RAM: ram}

	for !cpu.Halt {
		cpu.Step()
	}

	if cpu.Registers[0] != 15 {
		t.Errorf("Expected R0 to be 15, but got %d", cpu.Registers[0])
	}
}
