package vcpu

const (
	OpLoad byte = 0x01
	OpAdd  byte = 0x02
	OpSub  byte = 0x03
	OpMul  byte = 0x04
	OpDiv  byte = 0x05
	OpHalt byte = 0xFF
)
