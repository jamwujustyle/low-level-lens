package main

type CompileRequest struct {
	Expression string `json:"expression"`
}

type CompileResponse struct {
	Message  string   `json:"message"`
	Assembly []string `json:"assembly"`
}

type StepResponse struct {
	Registers [4]int32 `json:"registers"`
	PC        int      `json:"pc"`
	Halt      bool     `json:"halt"`
}
