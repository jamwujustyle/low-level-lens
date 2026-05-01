interface CompileResponse {
    message: string;
    assembly: string[];
}

interface StepResponse {
    registers: number[];
    pc: number;
    halt: boolean;
}

interface AppState {
    assembly: string[];
    pc: number;
    registers: number[];
    cycle: number;
    isCompiled: boolean;
    isHalted: boolean;
}

export type { AppState, StepResponse, CompileResponse }