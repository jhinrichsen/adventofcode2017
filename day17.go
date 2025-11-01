package adventofcode2017

func Day17(steps int, part1 bool) int {
	if part1 {
		buffer := make([]int, 1, 2018)
		buffer[0] = 0
		pos := 0
		for i := 1; i <= 2017; i++ {
			pos = (pos + steps) % len(buffer)
			pos++
			buffer = append(buffer, 0)
			copy(buffer[pos+1:], buffer[pos:])
			buffer[pos] = i
		}
		return buffer[(pos+1)%len(buffer)]
	}

	// Part 2: Find value after 0 after 50 million insertions
	// Key insight: 0 is always at position 0, we only track position 1
	pos := 0
	length := 1
	valueAfterZero := 0

	for i := 1; i <= 50000000; i++ {
		pos = (pos + steps) % length
		pos++
		if pos == 1 {
			valueAfterZero = i
		}
		length++
	}

	return valueAfterZero
}
