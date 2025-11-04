package adventofcode2017

import (
	"strings"
)

// Grid1D represents a square grid as a 1D slice for better cache locality
type Grid1D struct {
	data []byte
	size int
}

func newGrid1D(size int) Grid1D {
	return Grid1D{
		data: make([]byte, size*size),
		size: size,
	}
}

func (g Grid1D) get(x, y int) byte {
	return g.data[y*g.size+x]
}

func (g Grid1D) set(x, y int, val byte) {
	g.data[y*g.size+x] = val
}

func (g Grid1D) countOn() uint {
	var count uint
	for _, cell := range g.data {
		if cell == '#' {
			count++
		}
	}
	return count
}

// Extract 2x2 or 3x3 block as uint16 bitset
func (g Grid1D) extractBlock(x, y, blockSize int) uint16 {
	var bits uint16
	bit := uint16(0)
	for dy := 0; dy < blockSize; dy++ {
		for dx := 0; dx < blockSize; dx++ {
			if g.get(x+dx, y+dy) == '#' {
				bits |= (1 << bit)
			}
			bit++
		}
	}
	return bits
}

// Write a block from uint16 bitset
func (g Grid1D) writeBlock(x, y, blockSize int, bits uint16) {
	bit := uint16(0)
	for dy := 0; dy < blockSize; dy++ {
		for dx := 0; dx < blockSize; dx++ {
			if bits&(1<<bit) != 0 {
				g.set(x+dx, y+dy, '#')
			} else {
				g.set(x+dx, y+dy, '.')
			}
			bit++
		}
	}
}

type Grid [][]byte

func (g Grid) String() string {
	var sb strings.Builder
	for i, row := range g {
		if i > 0 {
			sb.WriteByte('/')
		}
		sb.Write(row)
	}
	return sb.String()
}

func (g Grid) Copy() Grid {
	cp := make(Grid, len(g))
	for i := range g {
		cp[i] = make([]byte, len(g[i]))
		copy(cp[i], g[i])
	}
	return cp
}

func (g Grid) Rotate() Grid {
	n := len(g)
	rotated := make(Grid, n)
	for i := 0; i < n; i++ {
		rotated[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			rotated[i][j] = g[n-1-j][i]
		}
	}
	return rotated
}

func (g Grid) Flip() Grid {
	flipped := make(Grid, len(g))
	for i := range g {
		flipped[i] = make([]byte, len(g[i]))
		for j := range g[i] {
			flipped[i][j] = g[i][len(g[i])-1-j]
		}
	}
	return flipped
}

func (g Grid) CountOn() uint {
	var count uint
	for _, row := range g {
		for _, cell := range row {
			if cell == '#' {
				count++
			}
		}
	}
	return count
}

func parseGrid(s string) Grid {
	rows := strings.Split(s, "/")
	grid := make(Grid, len(rows))
	for i, row := range rows {
		grid[i] = []byte(row)
	}
	return grid
}

type Rule struct {
	Pattern     Grid
	Replacement Grid
}

// Convert Grid to uint16 bitset
func gridToBits(g Grid) uint16 {
	var bits uint16
	bit := uint16(0)
	for i := range g {
		for j := range g[i] {
			if g[i][j] == '#' {
				bits |= (1 << bit)
			}
			bit++
		}
	}
	return bits
}

// Rotate 2x2 bitset
func rotate2(bits uint16) uint16 {
	return ((bits & 0x1) << 1) | ((bits & 0x2) << 2) |
		((bits & 0x4) >> 2) | ((bits & 0x8) >> 1)
}

// Rotate 3x3 bitset
func rotate3(bits uint16) uint16 {
	return ((bits & 0x001) << 2) | ((bits & 0x002) << 4) | ((bits & 0x004) << 6) |
		((bits & 0x008) >> 2) | ((bits & 0x010) << 0) | ((bits & 0x020) << 2) |
		((bits & 0x040) >> 6) | ((bits & 0x080) >> 4) | ((bits & 0x100) >> 2)
}

// Flip 2x2 bitset
func flip2(bits uint16) uint16 {
	return ((bits & 0x1) << 1) | ((bits & 0x2) >> 1) |
		((bits & 0x4) << 1) | ((bits & 0x8) >> 1)
}

// Flip 3x3 bitset
func flip3(bits uint16) uint16 {
	return ((bits & 0x001) << 2) | ((bits & 0x002) << 0) | ((bits & 0x004) >> 2) |
		((bits & 0x008) << 2) | ((bits & 0x010) << 0) | ((bits & 0x020) >> 2) |
		((bits & 0x040) << 2) | ((bits & 0x080) << 0) | ((bits & 0x100) >> 2)
}

type RuleBook struct {
	Rules map[string]Grid
	// Uint16-based rules for fast lookup
	Rules2 map[uint16]uint16 // 2x2 -> 3x3
	Rules3 map[uint16]uint16 // 3x3 -> 4x4
}

func (rb *RuleBook) AddRule(pattern, replacement Grid) {
	current := pattern
	for r := 0; r < 4; r++ {
		rb.Rules[current.String()] = replacement
		rb.Rules[current.Flip().String()] = replacement
		current = current.Rotate()
	}

	// Also add to uint16 maps
	size := len(pattern)
	patternBits := gridToBits(pattern)
	replacementBits := gridToBits(replacement)

	if size == 2 {
		curr := patternBits
		for r := 0; r < 4; r++ {
			rb.Rules2[curr] = replacementBits
			rb.Rules2[flip2(curr)] = replacementBits
			curr = rotate2(curr)
		}
	} else if size == 3 {
		curr := patternBits
		for r := 0; r < 4; r++ {
			rb.Rules3[curr] = replacementBits
			rb.Rules3[flip3(curr)] = replacementBits
			curr = rotate3(curr)
		}
	}
}

func (rb *RuleBook) Match(pattern Grid) (Grid, bool) {
	repl, ok := rb.Rules[pattern.String()]
	return repl, ok
}

func (rb *RuleBook) matchBits(bits uint16, blockSize int) (uint16, bool) {
	if blockSize == 2 {
		repl, ok := rb.Rules2[bits]
		return repl, ok
	}
	repl, ok := rb.Rules3[bits]
	return repl, ok
}

func NewDay21(lines []string) (*RuleBook, error) {
	rb := &RuleBook{
		Rules:  make(map[string]Grid),
		Rules2: make(map[uint16]uint16),
		Rules3: make(map[uint16]uint16),
	}
	for _, line := range lines {
		parts := strings.Split(line, " => ")
		if len(parts) != 2 {
			continue
		}
		pattern := parseGrid(parts[0])
		replacement := parseGrid(parts[1])
		rb.AddRule(pattern, replacement)
	}
	return rb, nil
}

func splitGrid(grid Grid, size int) [][]Grid {
	n := len(grid)
	blocks := n / size
	result := make([][]Grid, blocks)

	for bi := 0; bi < blocks; bi++ {
		result[bi] = make([]Grid, blocks)
		for bj := 0; bj < blocks; bj++ {
			block := make(Grid, size)
			for i := 0; i < size; i++ {
				block[i] = make([]byte, size)
				for j := 0; j < size; j++ {
					block[i][j] = grid[bi*size+i][bj*size+j]
				}
			}
			result[bi][bj] = block
		}
	}
	return result
}

func joinGrids(blocks [][]Grid) Grid {
	if len(blocks) == 0 {
		return Grid{}
	}
	blockSize := len(blocks[0][0])
	n := len(blocks) * blockSize
	result := make(Grid, n)
	for i := 0; i < n; i++ {
		result[i] = make([]byte, n)
	}

	for bi := 0; bi < len(blocks); bi++ {
		for bj := 0; bj < len(blocks[bi]); bj++ {
			block := blocks[bi][bj]
			for i := 0; i < blockSize; i++ {
				for j := 0; j < blockSize; j++ {
					result[bi*blockSize+i][bj*blockSize+j] = block[i][j]
				}
			}
		}
	}
	return result
}

func enhance(grid Grid, rb *RuleBook) Grid {
	size := len(grid)
	var blockSize int
	if size%2 == 0 {
		blockSize = 2
	} else {
		blockSize = 3
	}

	blocks := splitGrid(grid, blockSize)
	for i := range blocks {
		for j := range blocks[i] {
			replacement, ok := rb.Match(blocks[i][j])
			if ok {
				blocks[i][j] = replacement
			}
		}
	}

	return joinGrids(blocks)
}

func enhance1D(grid Grid1D, rb *RuleBook) Grid1D {
	size := grid.size
	var blockSize, newBlockSize int
	if size%2 == 0 {
		blockSize = 2
		newBlockSize = 3
	} else {
		blockSize = 3
		newBlockSize = 4
	}

	blocksPerSide := size / blockSize
	newSize := blocksPerSide * newBlockSize
	newGrid := newGrid1D(newSize)

	for by := 0; by < blocksPerSide; by++ {
		for bx := 0; bx < blocksPerSide; bx++ {
			// Extract block as bitset
			bits := grid.extractBlock(bx*blockSize, by*blockSize, blockSize)

			// Look up replacement
			if replBits, ok := rb.matchBits(bits, blockSize); ok {
				// Write replacement to new grid
				newGrid.writeBlock(bx*newBlockSize, by*newBlockSize, newBlockSize, replBits)
			}
		}
	}

	return newGrid
}

func gridFrom2D(g Grid) Grid1D {
	size := len(g)
	grid := newGrid1D(size)
	for y := range g {
		for x := range g[y] {
			grid.set(x, y, g[y][x])
		}
	}
	return grid
}

func Day21(rb *RuleBook, part1 bool) uint {
	grid := gridFrom2D(parseGrid(".#./..#/###"))

	iterations := 18
	if part1 {
		iterations = 5
	}

	for range iterations {
		grid = enhance1D(grid, rb)
	}

	return grid.countOn()
}
