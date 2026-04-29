import './style.css'

// ═══════════════════════════════════════════════════════════════
// Low-Level-Lens — Main Application Entry
// ═══════════════════════════════════════════════════════════════
//
// YOUR MISSION:
// Wire up the UI to the Go API at http://localhost:8000
//
// API Endpoints:
//   POST /compile  → body: { expression: string }
//                  ← { assembly: string[], message?: string }
//
//   POST /step     → no body needed
//                  ← { registers: number[], pc: number, halt: boolean }
//
// Below is the skeleton. Fill in every function marked with
// 🏗️ TODO — the comments explain exactly what to do.
// ═══════════════════════════════════════════════════════════════

const API = 'http://localhost:8000'

// ── DOM References ────────────────────────────────────────────
// Grab every element you'll need to read from or write to.
// They all have unique IDs in index.html — use getElementById
// or querySelector to grab them.
//
// 🏗️ TODO: Fill in each variable below.
//    Example: const codeTextarea = document.getElementById('code-textarea') as HTMLTextAreaElement
//
const codeTextarea    = null  // the <textarea> where user types expressions
const btnCompile      = null  // "Compile" button
const btnStep         = null  // "Step" button
const btnReset        = null  // "Reset" button
const assemblyList    = null  // <ul> that holds assembly lines
const pcBadge         = null  // PC display container
const pcValue         = null  // the <span> showing the hex PC value
const regValues       = null  // array of the 4 register value <span>s: [r0-value, r1-value, ...]
const regContainers   = null  // array of the 4 register <div>s: [reg-r0, reg-r1, ...]
const cycleCount      = null  // <span> showing current cycle number
const cpuStatus       = null  // status badge (idle/running/halted)
const haltOverlay     = null  // the halt overlay <div>
const haltResult      = null  // the <span> showing final R0 value
const btnDismiss      = null  // "Dismiss" button on halt overlay
const errorToast      = null  // error toast <div>
const statusDot       = null  // green dot in header
const statusText      = null  // "Connected" text in header
const memoryContainer = null  // memory view wrapper
const memoryGrid      = null  // memory grid <div>
const dataBus         = null  // data bus line
const aluEl           = null  // ALU container
const aluOp           = null  // ALU operation label
const phaseFetch      = null  // fetch phase indicator
const phaseDecode     = null  // decode phase indicator
const phaseExecute    = null  // execute phase indicator

// ── Application State ─────────────────────────────────────────
// This is your "single source of truth". Every time the API
// returns data, update this state, then call render().
//
// 🏗️ TODO: You'll mutate these values inside your fetch handlers.
//
let assembly: string[] = []
let registers: number[] = [0, 0, 0, 0]
let prevRegisters: number[] = [0, 0, 0, 0]
let pc = 0
let halted = false
let cycle = 0
let compiled = false

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
  // Your code here...
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
async function handleCompile(): Promise<void> {
  // Your code here...
}

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
  // Your code here...
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
