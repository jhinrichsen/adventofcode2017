package adventofcode2017

import (
	"fmt"
	"strconv"
	"strings"
)

// Day10Part1 parses input as comma-separated lengths and runs one knot-hash round.
// Returns the product of the first two elements.
func Day10Part1(input []byte, size int) int {
	s := strings.TrimSpace(string(input))
	var lengths []int
	if s != "" {
		parts := strings.Split(s, ",")
		lengths = make([]int, 0, len(parts))
		for i, part := range parts {
			n, err := strconv.Atoi(strings.TrimSpace(part))
			if err != nil {
				panic(fmt.Errorf("cannot convert number to int at column %d", i))
			}
			lengths = append(lengths, n)
		}
	}

	list := make([]int, size)
	for i := range list {
		list[i] = i
	}

	pos := 0
	skip := 0
	for _, n := range lengths {
		if n > 0 {
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

// Day10Part2 computes the full knot hash.
func Day10Part2(input []byte) string {
	s := strings.TrimSpace(string(input))
	lengths := make([]int, 0, len(s)+5)
	for _, r := range s {
		lengths = append(lengths, int(byte(r)))
	}
	// Standard suffix as specified in Day 10 Part 2
	lengths = append(lengths, 17, 31, 73, 47, 23)

	const size = 256
	list := make([]int, size)
	for i := range list {
		list[i] = i
	}

	pos := 0
	skip := 0
	for range 64 {
		for _, n := range lengths {
			if n > 0 {
				wrap := func(x int) int { return x % len(list) }
				for lower, upper := pos, pos+n-1; lower < upper; lower, upper = lower+1, upper-1 {
					list[wrap(lower)], list[wrap(upper)] = list[wrap(upper)], list[wrap(lower)]
				}
			}
			pos = (pos + n + skip) % len(list)
			skip++
		}
	}

	dense := make([]byte, 16)
	for block := range 16 {
		start := block * 16
		v := list[start]
		for _, elem := range list[start+1 : start+16] {
			v ^= elem
		}
		dense[block] = byte(v)
	}

	var b strings.Builder
	b.Grow(32)
	for _, v := range dense {
		b.WriteString(fmt.Sprintf("%02x", v))
	}
	return b.String()
}
