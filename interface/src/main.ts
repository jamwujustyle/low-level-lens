import './style.css'

import type { CompileResponse, StepResponse, AppState, Instruction } from './types'
const API = 'http://localhost:8000'

let state: AppState = {
    assembly: [],
    instructions: [],
    pc: 0,
    registers: [0,0,0,0],
    ram: [],
    cycle: 0,
    isCompiled: false,
    isHalted: false,
}

const codeTextarea = document.getElementById("code-textarea") as HTMLTextAreaElement
const compileBtn = document.getElementById("btn-compile") as HTMLButtonElement
const stepBtn = document.getElementById("btn-step") as HTMLButtonElement
const resetBtn = document.getElementById("btn-reset") as HTMLButtonElement
const dismissBtn = document.getElementById("btn-dismiss") as HTMLButtonElement
const assemblyList = document.getElementById("assembly-list") as HTMLDataListElement

// ── DOM References ────────────────────────────────────────────
const pcBadge = document.getElementById("pc-badge") as HTMLDivElement
const pcValue = document.getElementById("pc-value") as HTMLSpanElement

// Arrays make it much easier to loop through registers!
const regValues = [
  document.getElementById("r0-value") as HTMLSpanElement,
  document.getElementById("r1-value") as HTMLSpanElement,
  document.getElementById("r2-value") as HTMLSpanElement,
  document.getElementById("r3-value") as HTMLSpanElement,
]

const regContainers = [
  document.getElementById("reg-r0") as HTMLDivElement,
  document.getElementById("reg-r1") as HTMLDivElement,
  document.getElementById("reg-r2") as HTMLDivElement,
  document.getElementById("reg-r3") as HTMLDivElement,
]

const cycleCount = document.getElementById("cycle-count") as HTMLSpanElement
const cpuStatus = document.getElementById("cpu-status") as HTMLSpanElement
const haltOverlay = document.getElementById("halt-overlay") as HTMLDivElement
const haltResult = document.getElementById("halt-result") as HTMLDivElement
const errorToast = document.getElementById("error-toast") as HTMLDivElement
const statusDot = document.getElementById("status-dot") as HTMLSpanElement
const statusText = document.getElementById("status-text") as HTMLSpanElement
const memoryContainer = document.getElementById("memory-container") as HTMLDivElement
const memoryGrid = document.getElementById("memory-grid") as HTMLDivElement
const dataBus = document.getElementById("data-bus") as HTMLDivElement
const aluEl = document.getElementById("alu") as HTMLDivElement
const aluOp = document.getElementById("alu-op") as HTMLSpanElement
const phaseFetch = document.getElementById("phase-fetch") as HTMLDivElement
const phaseDecode = document.getElementById("phase-decode") as HTMLDivElement
const phaseExecute = document.getElementById("phase-execute") as HTMLDivElement

// ── Application State ─────────────────────────────────────────
// This is your "single source of truth". Every time the API
// returns data, update this state, then call render().
//
// 🏗️ TODO: You'll mutate these values inside your fetch handlers.
//


// ── Helpers ───────────────────────────────────────────────────
function base64ToBytes(base64: string): number[] {
  const binaryString = atob(base64)
  const bytes = new Uint8Array(binaryString.length)
  for (let i = 0; i < binaryString.length; i++) {
    bytes[i] = binaryString.charCodeAt(i)
  }
  return Array.from(bytes)
}

// ── Render Function ───────────────────────────────────────────
// This is the CORE pattern. After every state change, call
// render() to push state → DOM.
//
// 🏗️ TODO: Implement this function. It should:
//   1. Render the assembly list — loop over `assembly`, create
//      <li> elements. Highlight the one at index matching the
//      current instruction (use the `assembly__item--active` class).
//   2. Update register values — set each register <span>'s
//      textContent. If a value changed from prevRegisters,
//      add the `register--changed` class (remove after ~600ms).
//   3. Update the PC badge — format pc as hex (e.g. "0x0A").
//   4. Update cycle count.
//   5. Update the CPU status badge classes/text
//      (idle / running / halted).
//   6. If halted, show the halt overlay with R0's final value.
//
function renderOperands(operands: string): string {
  // Wrap register names (R0-R3) and immediate values in styled spans
  return operands.replace(/\b(R[0-3])\b/g, '<span class="asm-operand--reg">$1</span>')
                  .replace(/\b(\d+)\b/g, '<span class="asm-operand--imm">$1</span>')
}

function render(): void {
  // 1. Assembly List
  assemblyList.innerHTML = ""
  let activeRow: HTMLElement | null = null
  if (state.instructions.length === 0) {
    assemblyList.innerHTML = `<div class="assembly__empty">Compile an expression to see assembly…</div>`
  } else {
    state.instructions.forEach((inst: Instruction, index: number) => {
      const row = document.createElement("div")
      row.classList.add("asm-row")
      if (index === state.cycle) {
        row.classList.add("asm-row--active")
        activeRow = row
      }
      const mnemonicClass = `asm-row__mnemonic--${inst.mnemonic.toLowerCase()}`
      row.innerHTML = `
        <span class="asm-row__addr">${inst.address}</span>
        <span class="asm-row__opcode">${inst.opcode}</span>
        <span class="asm-row__mnemonic ${mnemonicClass}">${inst.mnemonic}</span>
        <span class="asm-row__operands">${renderOperands(inst.operands)}</span>
      `
      assemblyList.appendChild(row)
    })
    if (activeRow) (activeRow as HTMLElement).scrollIntoView({ behavior: 'smooth', block: 'nearest' })
  }

  // 1.5 Memory Grid
  memoryGrid.innerHTML = ""
  if (state.ram.length > 0) {
    memoryContainer.style.display = "block"
    let pcCell: HTMLElement | null = null
    state.ram.forEach((byte, i) => {
      const cell = document.createElement("div")
      cell.classList.add("memory-cell")
      if (i === state.pc) {
        cell.classList.add("memory-cell--pc")
        pcCell = cell
      }
      cell.textContent = byte.toString(16).toUpperCase().padStart(2, '0')
      memoryGrid.appendChild(cell)
    })
    if (pcCell) (pcCell as HTMLElement).scrollIntoView({ behavior: 'smooth', block: 'nearest' })
  } else {
    memoryContainer.style.display = "none"
  }

  // 2. Registers
  state.registers.forEach((val, i) => {
    const prevVal = regValues[i].textContent
    regValues[i].textContent = val.toString()
    
    if (prevVal !== val.toString() && state.cycle > 0) {
      regContainers[i].classList.add("register--changed")
      setTimeout(() => regContainers[i].classList.remove("register--changed"), 600)
    }
  })

  // 3. PC & Cycle
  pcValue.textContent = `0x${state.pc.toString(16).toUpperCase().padStart(2, '0')}`
  pcBadge.style.display = state.isCompiled ? "flex" : "none"
  cycleCount.textContent = state.cycle.toString()

  // 4. CPU Status & Phases
  const currentInst = state.instructions[state.cycle - 1]
  
  // Reset phases
  if (phaseFetch && phaseDecode && phaseExecute) {
    [phaseFetch, phaseDecode, phaseExecute].forEach(p => p?.classList.remove("active"))
  }
  
  if (aluEl) aluEl.classList.remove("alu--active")
  if (aluOp) aluOp.textContent = ""
  if (dataBus) dataBus.classList.remove("bus-line--active")

  if (state.cycle > 0 && currentInst) {
    // Trigger "Execute" visuals for the instruction that just ran
    if (phaseExecute) phaseExecute.classList.add("active")
    
    // If it was an ALU operation, highlight the ALU
    const isALU = ["ADD", "SUB", "MUL", "DIV"].includes(currentInst.mnemonic)
    if (isALU && aluEl && aluOp) {
      aluEl.classList.add("alu--active")
      aluOp.textContent = currentInst.mnemonic
    }

    // Trigger bus animation
    if (dataBus) {
      dataBus.classList.add("bus-line--active")
      setTimeout(() => dataBus.classList.remove("bus-line--active"), 800)
    }
  } else if (state.isCompiled && phaseFetch) {
    // If just compiled, we're ready to "Fetch" the first instruction
    phaseFetch.classList.add("active")
  }

  if (state.isHalted) {
    cpuStatus.textContent = "Halted"
    cpuStatus.className = "status-badge status-badge--halted"
    stepBtn.disabled = true
    
    // Show Halt Overlay
    haltResult.textContent = state.registers[0].toString()
    haltOverlay.classList.add("visible")
  } else if (state.isCompiled) {
    cpuStatus.textContent = "Running"
    cpuStatus.className = "status-badge status-badge--running"
    stepBtn.disabled = false
  } else {
    cpuStatus.textContent = "Idle"
    cpuStatus.className = "status-badge status-badge--idle"
    stepBtn.disabled = true
    resetBtn.disabled = true
  }
}

// ── Step Handler ──────────────────────────────────────────────
async function handleStep(): Promise<void> {
  if (state.isHalted) return

  try {
    const res = await fetch(`${API}/step`, { method: 'POST' })
    if (!res.ok) throw new Error(await res.text())
    
    const data: StepResponse = await res.json()
    
    state.registers = data.registers
    state.pc = data.pc
    state.isHalted = data.halt
    state.ram = base64ToBytes(data.ram)
    state.cycle++
    
    render()
  } catch (err: any) {
    showError(err.message)
  }
}

// ── Reset Handler ─────────────────────────────────────────────
async function handleReset(): Promise<void> {
  try {
    await fetch(`${API}/reset`, { method: 'POST' })
    
    state = {
      assembly: [],
      instructions: [],
      pc: 0,
      registers: [0,0,0,0],
      ram: [],
      cycle: 0,
      isCompiled: false,
      isHalted: false,
    }
    
    codeTextarea.value = ""
    haltOverlay.classList.remove("visible")
    render()
  } catch (err: any) {
    showError(err.message)
  }
}

// ── Error Display ─────────────────────────────────────────────
function showError(msg: string): void {
  errorToast.textContent = msg
  errorToast.classList.add("visible")
  setTimeout(() => errorToast.classList.remove("visible"), 3000)
}

// ── Event Listeners ───────────────────────────────────────────
stepBtn.addEventListener("click", handleStep)
resetBtn.addEventListener("click", handleReset)

const dismissHalt = () => {
  haltOverlay.classList.remove("visible")
  handleReset()
}

dismissBtn.addEventListener("click", dismissHalt)

window.addEventListener("keydown", (e: KeyboardEvent) => {
  if (e.key === "Escape" && haltOverlay.classList.contains("visible")) {
    dismissHalt()
  }
})

// ── Boot ──────────────────────────────────────────────────────
async function boot() {
  try {
    const res = await fetch(`${API}/ping`)
    if (res.ok) {
      statusDot.classList.remove("header__dot--offline")
      statusText.textContent = "Connected"
    } else {
      throw new Error()
    }
  } catch {
    statusDot.classList.add("header__dot--offline")
    statusText.textContent = "Offline"
  }
}
boot()
const handleCompile = async () => {
const expression = codeTextarea.value.trim();
  if (!expression) return;

  try {
    const response = await fetch('http://localhost:8000/compile', {
      method: "POST",
      headers: {"Content-Type": "application/json"},
      body: JSON.stringify({expression})
    });
    const data: CompileResponse = await response.json();

    state.assembly = data.assembly
    state.instructions = data.instructions || []
    state.ram = base64ToBytes(data.ram)
    state.isCompiled = true
    state.isHalted = false
    state.registers = [0,0,0,0]
    state.pc = 0
    state.cycle = 0
    
    stepBtn.disabled = false
    resetBtn.disabled = false
    render()
  }catch (error) {
    showError(error)
  }
}
codeTextarea.addEventListener("keydown", (event: KeyboardEvent) => {
  if (event.key === "Enter" && !event.shiftKey) {
    event.preventDefault();

    handleCompile();
  }
})
compileBtn.addEventListener("click", handleCompile)