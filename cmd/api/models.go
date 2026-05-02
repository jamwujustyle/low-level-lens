package main

type CompileRequest struct {
	Expression string `json:"expression"`
}

type Instruction struct {
	Address  string `json:"address"`
	OpCode   string `json:"opcode"`
	Mnemonic string `json:"mnemonic"`
	Operands string `json:"operands"`
	Raw      string `json:"raw"`
}
type CompileResponse struct {
	Message      string        `json:"message"`
	Assembly     []string      `json:"assembly"`
	Instructions []Instruction `json:"instructions"`
}

type StepResponse struct {
	Registers [4]int32 `json:"registers"`
	PC        int      `json:"pc"`
	Halt      bool     `json:"halt"`
}
