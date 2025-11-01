package adventofcode2017

import (
	"fmt"
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

	for row := 0; row < 128; row++ {
		input := fmt.Sprintf("%s-%d", key, row)
		hash := Day10Part2([]byte(input))

		col := 0
		for _, hexChar := range hash {
			// Convert hex character to its 4-bit value
			var val int
			switch {
			case hexChar >= '0' && hexChar <= '9':
				val = int(hexChar - '0')
			case hexChar >= 'a' && hexChar <= 'f':
				val = int(hexChar - 'a' + 10)
			case hexChar >= 'A' && hexChar <= 'F':
				val = int(hexChar - 'A' + 10)
			}

			// Set bits in the grid (MSB first)
			for i := 3; i >= 0; i-- {
				if val&(1<<i) != 0 {
					grid[row][col] = true
				}
				col++
			}
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

	// Part 2: count regions using flood fill
	visited := [128][128]bool{}
	regions := 0

	var floodFill func(row, col int)
	floodFill = func(row, col int) {
		if row < 0 || row >= 128 || col < 0 || col >= 128 {
			return
		}
		if !grid[row][col] || visited[row][col] {
			return
		}
		visited[row][col] = true
		floodFill(row-1, col)
		floodFill(row+1, col)
		floodFill(row, col-1)
		floodFill(row, col+1)
	}

	for row := 0; row < 128; row++ {
		for col := 0; col < 128; col++ {
			if grid[row][col] && !visited[row][col] {
				regions++
				floodFill(row, col)
			}
		}
	}

	return regions
}
