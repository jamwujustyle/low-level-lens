# Low-Level-Lens

**Low-Level-Lens** is a dual-purpose software engineering project that explores the complete lifecycle of a computing instruction—from high-level syntax parsing down to physical gate-level execution simulation.

It serves as the final project for:
1. **Principles of Programming (PoP)**: An extensible arithmetic compiler written in Go that supports Arabic numerals, Roman numerals, word operators, and strict semantic checks.
2. **Computer Architecture (COSC 3510)**: A web-based "Slow Motion CPU" UI built in Next.js that visualizes the Fetch-Decode-Execute cycle, demonstrating how software touches the hardware.

## 🚀 Features

*   **Custom Compiler (Go)**: A modular Lexer, Recursive-Descent Parser, and Semantic Analyzer.
*   **Code Generator**: Translates abstract syntax trees into a custom Assembly Instruction Set Architecture (ISA).
*   **Extensible Design**: Support for Roman numerals (`X + V`) and word operators (`ten plus two`).
*   **Virtual CPU (Go)**: A state-machine-based execution engine simulating CPU Registers, Program Counter, and RAM.
*   **Visual Hardware Emulator (Next.js)**: A dashboard to step through the execution pipeline, visualizing data movement across buses and logic gates inside the ALU.

## 🛠️ Tech Stack

*   **Backend / Compiler / VCPU**: Go
*   **Frontend / Visualization**: Next.js, TypeScript, Tailwind CSS

## 📂 Project Structure

*   `/compiler`: Go-based lexer, parser, AST builder, and assembly code generator.
*   `/vcpu`: Go-based CPU state machine, memory definitions, and custom opcode ISA.
*   `/interface`: Next.js frontend for the step-by-step visual execution trace.
*   `/tests`: Test cases encompassing valid, invalid, and extensible scenarios.
*   `SPEC.md`: The complete project requirements, roadmap, and architecture specifications.

## 📚 Documentation

For a detailed breakdown of the requirements, system architecture, and development roadmap, please refer to the [Project Specification (SPEC.md)](SPEC.md).
