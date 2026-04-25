# Instruction Set Architecture (ISA)

This document defines the custom Instruction Set Architecture for the Low-Level-Lens Virtual CPU (VCPU). 
It serves as the contract between the Compiler (which generates these instructions) and the VCPU (which executes them).

## CPU State

*   **Registers**: 4 General Purpose Registers.
    *   `R0`
    *   `R1`
    *   `R2`
    *   `R3`
*   **Program Counter (PC)**: Points to the current instruction in memory.
*   **Memory**: A linear array of bytes representing RAM.

## Instruction Set

The VCPU uses a simple, register-based instruction set. Each instruction has a mnemonic (for human readability) and a binary Opcode (for machine execution).

| Mnemonic       | Opcode (Hex) | Description                                                                 |
|----------------|--------------|-----------------------------------------------------------------------------|
| `LOAD Rx, imm` | `0x01`       | Load an immediate literal value (`imm`) into Register `Rx`.                 |
| `ADD Rx, Ry`   | `0x02`       | Add `Ry` to `Rx` and store the result in `Rx`.                              |
| `SUB Rx, Ry`   | `0x03`       | Subtract `Ry` from `Rx` and store the result in `Rx`.                       |
| `MUL Rx, Ry`   | `0x04`       | Multiply `Rx` by `Ry` and store the result in `Rx`.                         |
| `DIV Rx, Ry`   | `0x05`       | Divide `Rx` by `Ry` and store the result in `Rx`. (Handles division by zero)|
| `HALT`         | `0xFF`       | Stop CPU execution.                                                         |

## Execution Flow Example

A logical arithmetic operation like `(10 + 5) * 2` gets compiled into the following sequence:

```assembly
; (10 + 5) * 2
LOAD R0, 10    ; R0 = 10
LOAD R1, 5     ; R1 = 5
ADD R0, R1     ; R0 = 15
LOAD R1, 2     ; R1 = 2
MUL R0, R1     ; R0 = 30
HALT           ; Stop Execution
```

This textual representation will eventually be translated by our code generator into a raw byte array (e.g., `0x01 ...`) that the VCPU will read from its memory grid during the simulation.
