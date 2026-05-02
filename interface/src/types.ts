interface Instruction {
    address: string;
    opcode: string;
    mnemonic: string;
    operands: string;
    raw: string;
}

interface CompileResponse {
    message: string;
    assembly: string[];
    instructions: Instruction[];
    ram: string; // Base64 encoded
}

interface StepResponse {
    registers: number[];
    pc: number;
    halt: boolean;
    ram: string; // Base64 encoded
}

interface AppState {
    assembly: string[];
    instructions: Instruction[];
    pc: number;
    registers: number[];
    ram: number[];
    cycle: number;
    isCompiled: boolean;
    isHalted: boolean;
}

export type { AppState, StepResponse, CompileResponse, Instruction }