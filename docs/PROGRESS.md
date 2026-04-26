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

---

## 🚀 Remaining Tasks


### 3. Virtual CPU & Codegen (Phase 2)
- [ ] **Code Generation**: Translate the AST into the custom ISA assembly.

### 4. UI Interface (Phase 3)
- [ ] **Vanilla JS/TS Frontend**: Build the "Slow Motion CPU" dashboard.
- [ ] **Visualization**: Animate the Fetch-Decode-Execute cycle and Bus interactions.
