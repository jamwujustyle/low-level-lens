package vcpu

import (
	"fmt"
	"testing"
)

// Simulated latencies (realistic values from architecture textbooks)
const (
	cacheHitLatencyNs = 1  // L1 cache hit: ~1 ns
	ramLatencyNs      = 50 // Main memory access: ~50 ns
)

// TestCacheScalingDemo runs the same program with and without cache,
// computes simulated access times using the AMAT formula, and prints
// a scaling table for the presentation.
func TestCacheScalingDemo(t *testing.T) {
	buildProgram := func() []byte {
		// Program: (10 + 5) → LOAD R0,10 ; LOAD R1,5 ; ADD R0,R1 ; HALT
		ram := make([]byte, 100)
		pc := 0
		// LOAD R0, 10
		ram[pc] = OpLoad; pc++
		ram[pc] = 0; pc++
		ram[pc] = 10; pc++; ram[pc] = 0; pc++; ram[pc] = 0; pc++; ram[pc] = 0; pc++
		// LOAD R1, 5
		ram[pc] = OpLoad; pc++
		ram[pc] = 1; pc++
		ram[pc] = 5; pc++; ram[pc] = 0; pc++; ram[pc] = 0; pc++; ram[pc] = 0; pc++
		// ADD R0, R1
		ram[pc] = OpAdd; pc++
		ram[pc] = 0; pc++
		ram[pc] = 1; pc++
		// HALT
		ram[pc] = OpHalt
		return ram
	}

	// ── Run WITH cache ──
	cache := NewCache(4, 8)
	cpu := &CPU{RAM: buildProgram(), Cache: cache}
	for !cpu.Halt {
		cpu.Step()
	}

	hits := cache.Hits
	misses := cache.Misses
	totalAccesses := hits + misses
	hitRate := float64(hits) / float64(totalAccesses) * 100

	// Simulated time per execution
	withCacheNs := (hits * cacheHitLatencyNs) + (misses * ramLatencyNs)
	withoutCacheNs := totalAccesses * ramLatencyNs

	fmt.Println()
	fmt.Println("═══════════════════════════════════════════════════════════════")
	fmt.Println("  CACHE IMPACT — SCALING DEMONSTRATION")
	fmt.Println("═══════════════════════════════════════════════════════════════")
	fmt.Println()
	fmt.Printf("  Program: LOAD R0,10 → LOAD R1,5 → ADD R0,R1 → HALT\n")
	fmt.Printf("  Memory accesses per execution: %d\n", totalAccesses)
	fmt.Printf("  Cache hits: %d | Misses: %d | Hit rate: %.1f%%\n", hits, misses, hitRate)
	fmt.Printf("  Assumed latencies: Cache hit = %dns, RAM = %dns\n", cacheHitLatencyNs, ramLatencyNs)
	fmt.Println()
	fmt.Println("  ┌──────────────┬───────────────────┬───────────────────┬──────────┐")
	fmt.Println("  │  Executions  │  Without Cache     │  With Cache        │  Saved   │")
	fmt.Println("  ├──────────────┼───────────────────┼───────────────────┼──────────┤")

	scales := []int{1, 100, 100_000}
	for _, n := range scales {
		without := withoutCacheNs * n
		with := withCacheNs * n
		saved := without - with

		fmt.Printf("  │  %10s  │  %15s  │  %15s  │  %6.1f%%  │\n",
			formatCount(n),
			formatTime(without),
			formatTime(with),
			float64(saved)/float64(without)*100,
		)
	}

	fmt.Println("  └──────────────┴───────────────────┴───────────────────┴──────────┘")
	fmt.Println()
	fmt.Printf("  Speedup factor: %.1fx faster with cache\n", float64(withoutCacheNs)/float64(withCacheNs))
	fmt.Println()
}

func formatTime(ns int) string {
	switch {
	case ns >= 1_000_000:
		return fmt.Sprintf("%.2f ms", float64(ns)/1_000_000)
	case ns >= 1_000:
		return fmt.Sprintf("%.2f μs", float64(ns)/1_000)
	default:
		return fmt.Sprintf("%d ns", ns)
	}
}

func formatCount(n int) string {
	switch {
	case n >= 1_000_000:
		return fmt.Sprintf("%dM", n/1_000_000)
	case n >= 1_000:
		return fmt.Sprintf("%dK", n/1_000)
	default:
		return fmt.Sprintf("%d", n)
	}
}
