package adventofcode2017

import (
	"fmt"
)

// Day5 returns the number of steps until outside of maze.
func Day5Part1(maze []int) int {
	index := 0
	prev := index

	jmp := func() {
		prev = index
		index += maze[index]
	}
	inc := func() {
		maze[prev]++
	}
	outside := func() bool {
		return index < 0 || index >= len(maze)
	}
	n := 0
	for !outside() {
		jmp()
		inc()
		n++
	}
	return n
}

func dump(maze []int, index int) {
	for i := range maze {
		var format string
		if i == index {
			format = "(%2d) "
		} else {
			format = " %2d  "
		}
		fmt.Printf(format, maze[i])
	}
	fmt.Println()
}

// Day5 returns the number of steps until outside of maze.
func Day5Part2(maze []int) int {
	index := 0
	prev := index

	jmp := func() {
		prev = index
		index += maze[index]
	}
	inc := func() {
		maze[prev]++
	}
	dec := func() {
		maze[prev]--
	}
	offset := func() int {
		return index - prev
	}
	outside := func() bool {
		return index < 0 || index >= len(maze)
	}
	n := 0
	for !outside() {
		jmp()
		if offset() >= 3 {
			dec()

		} else {
			inc()
		}
		n++
	}
	return n
}
