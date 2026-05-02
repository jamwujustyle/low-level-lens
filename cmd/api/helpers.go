package main

import (
	"fmt"
	"strings"

	c "github.com/jamwujustyle/low-level-lens/compiler"
	"github.com/jamwujustyle/low-level-lens/vcpu"
)

func buildInstructions(comp *c.Compiler) []Instruction {
	i := []Instruction{}
	offset, asmIdx := 0, 0

	for offset < len(comp.Instructions) {
		opByte := comp.Instructions[offset]
		opHex := fmt.Sprintf("%02x", opByte)
		addr := fmt.Sprintf("0x%02X", offset)

		var mnemonic, operands, raw string

		if asmIdx < len(comp.Assembly) {
			raw = comp.Assembly[asmIdx]

			parts := strings.SplitN(raw, " ", 2)
			mnemonic = parts[0]
			if len(parts) > 1 {
				operands = parts[1]
			}
			asmIdx++
		}
		i = append(i, Instruction{
			Address:  addr,
			OpCode:   opHex,
			Mnemonic: mnemonic,
			Operands: operands,
			Raw:      raw,
		})

		switch opByte {
		case vcpu.OpLoad:
			offset += 6
		case vcpu.OpAdd, vcpu.OpSub, vcpu.OpMul, vcpu.OpDiv:
			offset += 3
		case vcpu.OpHalt:
			offset += 1
		default:
			offset += 1
		}
	}
	return i
}
