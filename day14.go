package adventofcode2017

import (
	"strconv"
	"strings"
)

type Day14Puzzle [128][128]bool

// NewDay14 parses the input lines and constructs the 128x128 grid of used/free squares.
// Returns a Day14Puzzle where true = used, false = free.
func NewDay14(lines []string) *Day14Puzzle {
	grid := &Day14Puzzle{}
	if len(lines) == 0 {
		return grid
	}
	key := strings.TrimSpace(lines[0])

	// Reusable buffers for knot hash computation
	lengths := make([]int, 0, 256)
	list := make([]int, 256)
	dense := make([]byte, 16)

	for row := 0; row < 128; row++ {
		input := key + "-" + strconv.Itoa(row)

		// Inline knot hash computation
		lengths = lengths[:0]
		for i := 0; i < len(input); i++ {
			lengths = append(lengths, int(input[i]))
		}
		lengths = append(lengths, 17, 31, 73, 47, 23)

		for i := range list {
			list[i] = i
		}

		pos := 0
		skip := 0
		for range 64 {
			for _, n := range lengths {
				if n > 0 {
					for lower, upper := pos, pos+n-1; lower < upper; lower, upper = lower+1, upper-1 {
						list[lower%256], list[upper%256] = list[upper%256], list[lower%256]
					}
				}
				pos = (pos + n + skip) % 256
				skip++
			}
		}

		for block := range 16 {
			start := block * 16
			v := list[start]
			for _, elem := range list[start+1 : start+16] {
				v ^= elem
			}
			dense[block] = byte(v)
		}

		// Convert dense hash bytes directly to grid bits
		col := 0
		for _, b := range dense {
			grid[row][col] = b&128 != 0
			grid[row][col+1] = b&64 != 0
			grid[row][col+2] = b&32 != 0
			grid[row][col+3] = b&16 != 0
			grid[row][col+4] = b&8 != 0
			grid[row][col+5] = b&4 != 0
			grid[row][col+6] = b&2 != 0
			grid[row][col+7] = b&1 != 0
			col += 8
		}
	}

	return grid
}

// Day14 solves both parts of Day 14.
// For part1=true: counts the number of used squares in the 128x128 grid.
// For part1=false: counts the number of regions (connected components) in the grid.
func Day14(grid *Day14Puzzle, part1 bool) int {
	if part1 {
		count := 0
		for row := 0; row < 128; row++ {
			for col := 0; col < 128; col++ {
				if grid[row][col] {
					count++
				}
			}
		}
		return count
	}

	// Part 2: count regions using iterative flood fill
	visited := [128][128]bool{}
	regions := 0
	stack := make([][2]int, 0, 256)

	for row := 0; row < 128; row++ {
		for col := 0; col < 128; col++ {
			if !grid[row][col] || visited[row][col] {
				continue
			}
			regions++
			stack = append(stack, [2]int{row, col})
			visited[row][col] = true

			for len(stack) > 0 {
				curr := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				r, c := curr[0], curr[1]

				// Check 4 neighbors
				if r > 0 && grid[r-1][c] && !visited[r-1][c] {
					visited[r-1][c] = true
					stack = append(stack, [2]int{r - 1, c})
				}
				if r < 127 && grid[r+1][c] && !visited[r+1][c] {
					visited[r+1][c] = true
					stack = append(stack, [2]int{r + 1, c})
				}
				if c > 0 && grid[r][c-1] && !visited[r][c-1] {
					visited[r][c-1] = true
					stack = append(stack, [2]int{r, c - 1})
				}
				if c < 127 && grid[r][c+1] && !visited[r][c+1] {
					visited[r][c+1] = true
					stack = append(stack, [2]int{r, c + 1})
				}
			}
		}
	}

	return regions
}
