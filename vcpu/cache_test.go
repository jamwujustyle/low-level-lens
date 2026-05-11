package vcpu

import "testing"

func TestCacheHitAndMiss(t *testing.T) {
	ram := make([]byte, 64)
	for i := range ram {
		ram[i] = byte(i)
	}

	cache := NewCache(4, 8)

	// First read at address 0 should be a miss
	val := cache.Read(0, ram)
	if val != 0 {
		t.Errorf("Expected byte 0, got %d", val)
	}
	if cache.Misses != 1 || cache.Hits != 0 {
		t.Errorf("Expected 1 miss 0 hits, got %d misses %d hits", cache.Misses, cache.Hits)
	}

	// Second read at same address should be a hit (temporal locality)
	val = cache.Read(0, ram)
	if val != 0 {
		t.Errorf("Expected byte 0, got %d", val)
	}
	if cache.Hits != 1 {
		t.Errorf("Expected 1 hit after re-read, got %d", cache.Hits)
	}

	// Read at address 3 — same block as address 0, should also be a hit (spatial locality)
	val = cache.Read(3, ram)
	if val != 3 {
		t.Errorf("Expected byte 3, got %d", val)
	}
	if cache.Hits != 2 {
		t.Errorf("Expected 2 hits (spatial locality), got %d", cache.Hits)
	}
}

func TestCacheWithCPU(t *testing.T) {
	// Build a simple program: LOAD R0, 10 ; LOAD R1, 5 ; ADD R0, R1 ; HALT
	ram := make([]byte, 100)
	pc := 0

	// LOAD R0, 10
	ram[pc] = OpLoad; pc++
	ram[pc] = 0; pc++
	ram[pc] = 10; pc++ // little-endian int32
	ram[pc] = 0; pc++
	ram[pc] = 0; pc++
	ram[pc] = 0; pc++

	// LOAD R1, 5
	ram[pc] = OpLoad; pc++
	ram[pc] = 1; pc++
	ram[pc] = 5; pc++
	ram[pc] = 0; pc++
	ram[pc] = 0; pc++
	ram[pc] = 0; pc++

	// ADD R0, R1
	ram[pc] = OpAdd; pc++
	ram[pc] = 0; pc++
	ram[pc] = 1; pc++

	// HALT
	ram[pc] = OpHalt

	cache := NewCache(4, 8)
	cpu := &CPU{RAM: ram, Cache: cache}

	for !cpu.Halt {
		cpu.Step()
	}

	if cpu.Registers[0] != 15 {
		t.Errorf("Expected R0 = 15, got %d", cpu.Registers[0])
	}

	// Cache should have recorded both hits and misses
	if cache.Hits == 0 {
		t.Error("Expected some cache hits during execution")
	}
	if cache.Misses == 0 {
		t.Error("Expected some cache misses during execution")
	}

	total := cache.Hits + cache.Misses
	hitRate := float64(cache.Hits) / float64(total) * 100
	t.Logf("Cache stats: %d hits, %d misses, %.1f%% hit rate", cache.Hits, cache.Misses, hitRate)
}
