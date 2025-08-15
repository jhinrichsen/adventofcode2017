package adventofcode2017

import (
	"fmt"
	"strconv"
	"strings"
)

var ErrorDay10Length = fmt.Errorf("illegal length")

// day10Puzzle holds parsed input and configuration for the solver.
// Tests can modify Size/ZeroBased to match example semantics.
type day10Puzzle struct {
	lengths []int
	size    int
	zero    bool
}

// NewDay10 parses the comma-separated input lengths and returns them.
// It does not solve the puzzle; call Day10 with an appropriate list and the returned lengths.
func NewDay10(buf []byte) day10Puzzle {
	s := strings.TrimSpace(string(buf))
	var lengths []int
	if s != "" {
		parts := strings.Split(s, ",")
		lengths = make([]int, 0, len(parts))
		for i := 0; i < len(parts); i++ {
			n, err := strconv.Atoi(strings.TrimSpace(parts[i]))
			if err != nil {
				panic(fmt.Errorf("cannot convert number to int at column %d", i))
			}
			lengths = append(lengths, n)
		}
	}
	// Defaults: AoC standard configuration (0..255 list)
	return day10Puzzle{lengths: lengths, size: 256, zero: true}
}

// Day10 is the solver. It constructs the working list according to the puzzle
// configuration and runs one knot-hash round, returning the product of the
// first two elements.
func Day10(p day10Puzzle) int {
	size := p.size
	if size <= 0 {
		size = 256
	}
	list := make([]int, size)
	for i := 0; i < size; i++ {
		if p.zero {
			list[i] = i
		} else {
			list[i] = i + 1
		}
	}
	// Inline the knot hash single round
	pos := 0
	skip := 0
	for i := 0; i < len(p.lengths); i++ {
		n := p.lengths[i]
		if n > 0 {
			// reverse list[pos:pos+n] with wraparound
			wrap := func(x int) int { return x % len(list) }
			for lower, upper := pos, pos+n-1; lower < upper; lower, upper = lower+1, upper-1 {
				list[wrap(lower)], list[wrap(upper)] = list[wrap(upper)], list[wrap(lower)]
			}
		}
		pos = (pos + n + skip) % len(list)
		skip++
	}
	return list[0] * list[1]
}
