import './style.css'

import type { CompileResponse, StepResponse, AppState } from './types'
const API = 'http://localhost:8000'

let state: AppState = {
    assembly: [],
    pc: 0,
    registers: [0,0,0,0],
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
function render(): void {
  assemblyList.innerHTML = ""

  if (state.assembly.length === 0) {
    assemblyList.innerHTML = `<li class="assembly__empty">Compile an expression to see assembly…</li>`
    return
  }

  state.assembly.forEach((instruction, index) => {
    const li = document.createElement("li")

    li.textContent = instruction

    li.classList.add("assembly__item")

    if (index === state.cycle) {
      li.classList.add("assembly__item--active")
    }
    assemblyList.appendChild(li)
  }
)
}

// ── Compile Handler ───────────────────────────────────────────
// When the user clicks "Compile":
//   1. Read the textarea value
//   2. POST it to /compile as JSON
//   3. Parse the response
//   4. Store assembly[] in state
//   5. Reset registers, pc, cycle, halted
//   6. Enable the Step + Reset buttons
//   7. Call render()
//
// 🏗️ TODO: Implement this async function.
//
// HINT:
//   const res = await fetch(`${API}/compile`, {
//     method: 'POST',
//     headers: { 'Content-Type': 'application/json' },
//     body: JSON.stringify({ expression: ??? })
//   })
//

// ── Step Handler ──────────────────────────────────────────────
// When the user clicks "Step":
//   1. POST to /step (no body needed)
//   2. Parse the response → { registers, pc, halt }
//   3. Save previous registers (for change detection)
//   4. Update state with new values
//   5. Increment cycle
//   6. Call render()
//
// 🏗️ TODO: Implement this async function.
//
async function handleStep(): Promise<void> {
  // Your code here...
}

// ── Reset Handler ─────────────────────────────────────────────
// When the user clicks "Reset":
//   1. Zero out all state (registers, pc, cycle, halted, assembly)
//   2. Disable Step + Reset buttons
//   3. Call render()
//
// 🏗️ TODO: Implement this function.
//
function handleReset(): void {
  // Your code here...
}

// ── Error Display ─────────────────────────────────────────────
// Shows the error toast for ~3 seconds.
//
// 🏗️ TODO: Implement this.
//   1. Set errorToast's textContent
//   2. Add the 'visible' class
//   3. Use setTimeout to remove 'visible' after 3000ms
//
function showError(msg: string): void {
  console.log(msg)
}

// ── Event Listeners ───────────────────────────────────────────
// 🏗️ TODO: Attach click handlers to the buttons.
//   btnCompile → handleCompile
//   btnStep    → handleStep
//   btnReset   → handleReset
//   btnDismiss → hide halt overlay + call handleReset
//

// ── Boot ──────────────────────────────────────────────────────
// 🏗️ TODO: On page load, ping the API to check connectivity.
//   fetch(`${API}/ping`)
//   If OK → set status dot to green, text to "Connected"
//   If error → set dot class to 'header__dot--offline', text to "Offline"
//
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
    state.isCompiled = true
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