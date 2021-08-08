package adventofcode2017

// Day05Part1 returns the number of steps until outside of maze.
func Day05Part1(maze []int) int {
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

// Day05Part2 returns the number of steps until outside of maze.
func Day05Part2(maze []int) int {
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
