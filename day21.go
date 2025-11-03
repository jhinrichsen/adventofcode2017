package adventofcode2017

import (
	"strings"
)

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

type RuleBook struct {
	Rules map[string]Grid
}

func (rb *RuleBook) AddRule(pattern, replacement Grid) {
	current := pattern
	for r := 0; r < 4; r++ {
		rb.Rules[current.String()] = replacement
		rb.Rules[current.Flip().String()] = replacement
		current = current.Rotate()
	}
}

func (rb *RuleBook) Match(pattern Grid) (Grid, bool) {
	repl, ok := rb.Rules[pattern.String()]
	return repl, ok
}

func NewDay21(lines []string) (*RuleBook, error) {
	rb := &RuleBook{
		Rules: make(map[string]Grid),
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

func Day21(rb *RuleBook, part1 bool) uint {
	grid := parseGrid(".#./..#/###")

	iterations := 18
	if part1 {
		iterations = 5
	}

	for range iterations {
		grid = enhance(grid, rb)
	}

	return grid.CountOn()
}
