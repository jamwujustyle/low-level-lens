package vcpu

// CacheLine represents a single line in the direct-mapped cache.
// It holds a block of bytes fetched from RAM alongside metadata
// tracking whether the line is valid and which memory block it maps to.
type CacheLine struct {
	Valid bool   `json:"valid"`
	Tag   int    `json:"tag"`
	Data  []byte `json:"data"`
}

// Cache simulates a small direct-mapped cache sitting between
// the CPU and main memory (RAM). Every read goes through the cache
// first: on a hit the data comes back immediately; on a miss the
// cache fetches an entire block from RAM before returning the byte.
type Cache struct {
	Lines     []CacheLine `json:"lines"`
	BlockSize int         `json:"blockSize"`
	NumLines  int         `json:"numLines"`
	Hits      int         `json:"hits"`
	Misses    int         `json:"misses"`
}

// NewCache creates a direct-mapped cache with the given number
// of lines and block size. These are kept intentionally small so
// the audience can watch misses turn into hits on loop iterations.
func NewCache(numLines, blockSize int) *Cache {
	lines := make([]CacheLine, numLines)
	for i := range lines {
		lines[i].Data = make([]byte, blockSize)
	}
	return &Cache{
		Lines:     lines,
		BlockSize: blockSize,
		NumLines:  numLines,
	}
}

// Read fetches a single byte at the given address. It computes the
// cache index and tag from the address, checks for a hit, and falls
// back to RAM on a miss — loading the full block into the cache line.
func (c *Cache) Read(address int, ram []byte) byte {
	blockIndex := address / c.BlockSize
	lineIndex := blockIndex % c.NumLines
	tag := blockIndex / c.NumLines

	line := &c.Lines[lineIndex]

	if line.Valid && line.Tag == tag {
		// Cache hit — data already in the cache
		c.Hits++
	} else {
		// Cache miss — pull the entire block from RAM
		c.Misses++
		line.Valid = true
		line.Tag = tag
		blockStart := blockIndex * c.BlockSize
		for i := 0; i < c.BlockSize; i++ {
			if blockStart+i < len(ram) {
				line.Data[i] = ram[blockStart+i]
			} else {
				line.Data[i] = 0
			}
		}
	}

	offset := address % c.BlockSize
	return line.Data[offset]
}

// Reset clears all cache lines and counters back to zero.
func (c *Cache) Reset() {
	c.Hits = 0
	c.Misses = 0
	for i := range c.Lines {
		c.Lines[i].Valid = false
		c.Lines[i].Tag = 0
		for j := range c.Lines[i].Data {
			c.Lines[i].Data[j] = 0
		}
	}
}
