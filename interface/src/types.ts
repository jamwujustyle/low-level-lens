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
}

interface StepResponse {
    registers: number[];
    pc: number;
    halt: boolean;
}

interface AppState {
    assembly: string[];
    instructions: Instruction[];
    pc: number;
    registers: number[];
    cycle: number;
    isCompiled: boolean;
    isHalted: boolean;
}

export type { AppState, StepResponse, CompileResponse, Instruction }