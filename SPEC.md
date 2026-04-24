# Low-Level-Lens: Project Specification

## 1. Project Overview
**Low-Level-Lens** is a dual-purpose software engineering project designed to fulfill the final examination requirements for two core computer science courses:
1. **Principles of Programming (PoP)**: Design and implementation of an extensible arithmetic compiler.
2. **Computer Architecture (COSC 3510)**: A "Full-Stack Execution Trace" presentation demonstrating the lifecycle of an instruction from high-level logic down to gate-level execution.

By bridging these two domains, the project takes high-level arithmetic expressions, compiles them down to a custom Assembly/Machine Code, and then executes them step-by-step on a web-based Virtual CPU (VCPU) interface. 

---

## 2. System Architecture

The project is divided into three distinct layers across a modern tech stack (Go + Vanilla JS/TS):

### Layer 1: The Compiler (Go)
Responsible for parsing the arithmetic expressions and translating them into an Abstract Syntax Tree (AST), then generating custom Assembly instructions.
*   **Lexer (Tokenization)**: Converts raw string input into categorizable tokens (e.g., Arabic/Roman numerals, symbolic operators, word operators).
*   **Parser (AST Construction)**: Validates grammar (via Recursive Descent) and builds the AST.
*   **Semantic Analyzer**: Detects invalid operations (e.g., division by zero, mismatched types).
*   **Code Generator**: Traverses the AST and emits custom Assembly Code (ISA).

### Layer 2: The Instruction Set Architecture (ISA) & VCPU (Go)
The custom "contract" between the compiler and the hardware emulator.
*   **ISA Definition**: A minimal, register-based architecture containing fundamental instructions (LOAD, ADD, SUB, MUL, DIV, HALT).
*   **State Machine (VCPU)**: A Go-based simulation of a CPU containing Registers (R0-R3), a Program Counter (PC), and Memory (RAM), executing generated binary opcodes.

### Layer 3: The Frontend Emulator (Vanilla JS/TS)
A web-based dashboard acting as a "Slow Motion CPU" to visualize the execution.
*   **Human-Speed Execution**: A UI allowing the user to "Step" through the Fetch-Decode-Execute cycle.
*   **Component Visualization**: Live updating grids for Memory, Register boxes, Data/Control Buses, and the ALU.

---

## 3. Principles of Programming (PoP) Requirements

To secure the full 100 points for the PoP final, the compiler must support:
*   **Basic Arithmetic & Precedence**: Addition (+), Subtraction (-), Multiplication (*), Division (/), and Parentheses `()`.
*   **Arabic Numerals**: Standard numbers (1, 2, 25, 100, etc.).
*   **Extensible Operator Representation**: Architecture designed to easily add word operators (e.g., `plus`, `minus`, `times`, `divided by`).
*   **Extensible Numeral Representation**: Support for Roman Numerals (e.g., I, V, X, IX) and/or English number words.
*   **Error Reporting**: Robust handling of division by zero, unknown tokens, incorrect structure, and unsupported formats.

**PoP Deliverables:**
*   **Source Code**: Modularized lexer, parser, semantic analyzer, and code generation.
*   **Test Cases**: 
    *   10 valid expressions (e.g., `2 + 3 * 4`).
    *   5 invalid expressions (e.g., `10 / (5 - 5)`).
    *   3 expressions demonstrating extensibility (e.g., `X * IV` or `8 plus 5`).

---

## 4. Computer Architecture (COSC 3510) Requirements

To secure the 75 points for the Final Presentation, the UI and demonstration must showcase the transition from "Logical Code" to "Physical CPU Gates."
*   **The Lifecycle of an Instruction**: Demonstrate how high-level code maps to the CPU's Instruction Decoder.
*   **Fetch-Decode-Execute Visualized**:
    *   **Fetch**: Highlight the current assembly instruction in memory; animate the Program Counter moving.
    *   **Decode**: Show the translation from Assembly Mnemonic to Binary Opcode.
    *   **Execute**: Animate the register values changing.
*   **Nuance Elements**:
    *   **Buses**: Glowing lines showing data moving from Memory to Registers over Address/Data/Control buses.
    *   **ALU & Logic Gates**: A side-panel or pop-up explaining the transistor-level logic (e.g., AND/XOR circuits) during an ADD operation.
    *   **Binary Reality**: Exposing that an Assembly `ADD` is just binary (e.g., `00000010`) triggering hardware circuits.

---

## 5. Proposed Custom ISA (Instruction Set Architecture)

A minimal Register-Based Instruction Set to map AST nodes to CPU operations:

| Mnemonic       | Opcode (Hex) | Description                                                               |
|----------------|--------------|---------------------------------------------------------------------------|
| `LOAD Rx, imm` | `0x01`       | Load an immediate literal value into Register `Rx`.                       |
| `ADD Rx, Ry`   | `0x02`       | Add `Ry` to `Rx` and store the result in `Rx`.                            |
| `SUB Rx, Ry`   | `0x03`       | Subtract `Ry` from `Rx` and store the result in `Rx`.                     |
| `MUL Rx, Ry`   | `0x04`       | Multiply `Rx` by `Ry` and store the result in `Rx`.                       |
| `DIV Rx, Ry`   | `0x05`       | Divide `Rx` by `Ry` and store the result in `Rx`.                         |
| `HALT`         | `0xFF`       | Stop CPU execution.                                                       |

*Registers Available: R0, R1, R2, R3.*

---

## 6. Development Roadmap

### Phase 1: Compiler Core (Go)
1.  **Grammar & Lexer**: Define the EBNF grammar. Implement a tokenization system capable of reading extensible inputs, e.g., turning `X plus 5` into tokens: `[ROMAN(10), OP_PLUS, ARABIC(5)]`.
2.  **Parser**: Implement a Recursive Descent parser to convert tokens into an Abstract Syntax Tree (AST).
3.  **Semantic Analyzer**: Traverse AST to validate semantic correctness (catch div-by-zero, unsupported type combinations).

### Phase 2: Virtual CPU & Code Generation (Go)
1.  **Codegen**: Write a visitor for the AST that emits custom assembly strings.
    *   *Example*: `(X + 5) * 2` translates to `LOAD R0, 10`, `LOAD R1, 5`, `ADD R0, R1`, `LOAD R1, 2`, `MUL R0, R1`.
2.  **Assembler**: Convert textual assembly into raw binary byte arrays for the VCPU.
3.  **State Machine (VCPU)**: Create the Go struct mapping memory and registers, exposing a `Step()` function to execute one opcode at a time and tracking CPU state.

### Phase 3: Presentation UI (Vanilla JS/TS)
1.  **Backend API**: Wrap the Go VCPU in a simple HTTP API (or execute via WebAssembly) to expose the CPU state after each `Step()`.
2.  **UI Construction**: Build out the Memory Grid, Register File, and ALU representations using Tailwind CSS.
3.  **Animations**: Implement the "Clock Speed" and "Step Mode" controls, triggering bus animations and state transitions based on the API responses.

---

## 7. Directory Structure
```text
low-level-lens/
├── compiler/           # Go: The PoP Project (Lexer, Parser, Codegen, Semantics)
├── vcpu/               # Go: The Virtual CPU state machine & ISA opcodes
├── interface/          # Vanilla JS/TS: The Architecture Visualization UI
├── tests/              # Test cases (Valid, Invalid, Extensible cases)
└── SPEC.md             # This specification document
```

---

## 8. Presentation Strategy & "The Nuance"

To effectively demonstrate the ability to "navigate across the stack" and "interconnect everything" during the COSC 3510 presentation, use the following analogy to explain the architecture:

*   **The OS/Application Layer (Vanilla JS/TS)**: The web UI acts as the Operating System's GUI. It provides the interface for interaction but does not perform the raw physical computation.
*   **The Physical Processor (Go Engine)**: The Go backend represents the physical hardware layer. By using Go, you can demonstrate Instruction Set Architecture (ISA) design effectively, specifically how bytes are written to a virtual memory buffer.
*   **The Interconnect (API/Bus)**: The communication between the TS frontend and Go backend represents the system Bus. When "Step" is clicked, the UI (OS) asks the Go engine (Processor): *"What is the state of the registers after the next instruction?"* The Go engine executes the logic and replies with the updated state (e.g., *"R0 is now 15, and the Program Counter is at 0x04"*).

This explicit mapping makes Topic 3 (Types of Computer Buses) and Topic 4 (Instruction Cycle) incredibly intuitive for the audience to visualize and proves mastery over both software abstractions and hardware realities.
