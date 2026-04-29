package main

type CompileRequest struct {
	Expression string `json:"expression"`
}

type CompileResponse struct {
	Message  string   `json:"message"`
	Assembly []string `json:"assembly"`
}
