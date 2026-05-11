# Low-Level-Lens: Progress Report

## ✅ Completed Tasks

### 1. Foundation & Documentation
- **Project Specification**: Detailed roadmap for PoP and COSC 3510 alignment.
- **README**: High-level overview and project structure.
- **ISA Mapping**: Defined custom register-based Instruction Set Architecture (opcodes).

### 2. Lexer (Compiler Core - Phase 1)
- **Token Definitions**: Categorization for Numbers, Operators, and EOF.
- **Lexer Engine**: State-based scanner with `readChar` logic.
- **Symbol Support**: Recognition of `+`, `-`, `*`, `/`, `(`, and `)`.
- **Whitespace Handling**: Automated skipping of spaces, tabs, and newlines.
- **Multi-digit Numbers**: Logic to group consecutive digits into single tokens.
- **Extensibility**:
    - Lookup table for word operators (`plus`, `minus`, `times`, `divided`).
    - Case-insensitive Roman numeral support (`I`, `V`, `X`).
- **Testing**: Standalone entry point implemented in `cmd/test-lexer/main.go`.

### 3. Parser (Phase 1, Step 6)
- **AST Definitions**: Created structures for Expression nodes, Number nodes, and Binary Operation nodes.
- **Recursive Descent Parser**: Implemented logic to handle operator precedence (order of operations) using Pratt parsing.

### 4. Semantic Analysis
- **Validation**: Implemented an Evaluator to traverse the AST and mathematically catch errors like "Division by Zero" at compile time.

### 5. VCPU Engine (Phase 2)
- **State Machine**: Implemented the Go-based CPU struct with Registers, PC, RAM, and a Fetch-Decode-Execute cycle (`Step`).
- **Instruction Support**: Full support for `LOAD`, `ADD`, `SUB`, `MUL`, `DIV`, and `HALT`.
- **Testing**: Verified hardware logic with manual byte-code tests in `vcpu/cpu_test.go`.

---

### 6. Code Generation (Phase 2)
- **AST Walking**: Implemented recursive tree-walking to translate AST nodes into linear VCPU instructions.
- **Assembly Logging**: Added human-readable `.asm` generation alongside binary machine code.
- **Extended Features**: Added robust semantic checks and word operator handling.

---

## ✅ Completed Tasks (Phase 3)

### 4. UI Interface (Phase 3)
- [x] **Go Backend API**: Wrapped the Compiler and VCPU in a standard REST HTTP API with CORS and Graceful Shutdown.
- [x] **Vanilla JS/TS Frontend**: Build the "Slow Motion CPU" dashboard (API integration, DOM linking, and compile rendering successful!).
- [x] **Visualization & Execution**: Implemented the Step logic, updated Registers/PC/Cycles on the UI, and animated the Bus interactions.
