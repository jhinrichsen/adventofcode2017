package adventofcode2017

import (
	"strconv"
	"strings"
)

func Day19(grid []string, part1 bool) string {
	if len(grid) == 0 {
		return ""
	}

	// Find starting position (top row, first |)
	startX := 0
	for i, c := range grid[0] {
		if c == '|' {
			startX = i
			break
		}
	}

	x, y := startX, 0
	dx, dy := 0, 1 // Start moving down
	letters := strings.Builder{}
	steps := 0

	for {
		steps++
		x += dx
		y += dy

		if y < 0 || y >= len(grid) || x < 0 || x >= len(grid[y]) {
			break
		}

		ch := grid[y][x]

		if ch == ' ' {
			break
		}

		if ch >= 'A' && ch <= 'Z' {
			letters.WriteByte(byte(ch))
		}

		if ch == '+' {
			// Turn at intersection
			if dx == 0 {
				// Moving vertically, try horizontal
				if x+1 < len(grid[y]) && grid[y][x+1] != ' ' && grid[y][x+1] != '|' {
					dx, dy = 1, 0
				} else if x-1 >= 0 && grid[y][x-1] != ' ' && grid[y][x-1] != '|' {
					dx, dy = -1, 0
				}
			} else {
				// Moving horizontally, try vertical
				if y+1 < len(grid) && x < len(grid[y+1]) && grid[y+1][x] != ' ' && grid[y+1][x] != '-' {
					dx, dy = 0, 1
				} else if y-1 >= 0 && x < len(grid[y-1]) && grid[y-1][x] != ' ' && grid[y-1][x] != '-' {
					dx, dy = 0, -1
				}
			}
		}
	}

	if part1 {
		return letters.String()
	}
	return strconv.Itoa(steps)
}
