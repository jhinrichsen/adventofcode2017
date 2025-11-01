package adventofcode2017

import (
	"fmt"
	"strings"
)

type Day15Puzzle struct {
	startA uint
	startB uint
}

// NewDay15 parses the input lines and returns generator starting values.
func NewDay15(lines []string) Day15Puzzle {
	var p Day15Puzzle
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "Generator A") {
			fmt.Sscanf(line, "Generator A starts with %d", &p.startA)
		} else if strings.HasPrefix(line, "Generator B") {
			fmt.Sscanf(line, "Generator B starts with %d", &p.startB)
		}
	}
	return p
}

// Day15Part1 counts matching lowest 16 bits after 40 million pairs.
func Day15Part1(p Day15Puzzle) uint {
	const (
		factorA = 16807
		factorB = 48271
		modulo  = 2147483647
		pairs   = 40000000
	)
	a, b := p.startA, p.startB
	matches := uint(0)
	for i := 0; i < pairs; i++ {
		a = (a * factorA) % modulo
		b = (b * factorB) % modulo
		if a&0xFFFF == b&0xFFFF {
			matches++
		}
	}
	return matches
}

// Day15Part2 counts matching lowest 16 bits after 5 million pairs.
// Generator A only produces multiples of 4, Generator B only multiples of 8.
func Day15Part2(p Day15Puzzle) uint {
	const (
		factorA = 16807
		factorB = 48271
		modulo  = 2147483647
		pairs   = 5000000
	)
	a, b := p.startA, p.startB
	matches := uint(0)
	for i := 0; i < pairs; i++ {
		for {
			a = (a * factorA) % modulo
			if a%4 == 0 {
				break
			}
		}
		for {
			b = (b * factorB) % modulo
			if b%8 == 0 {
				break
			}
		}
		if a&0xFFFF == b&0xFFFF {
			matches++
		}
	}
	return matches
}

func day15Part1AsmB1(startA, startB uint) uint

func day15Part2AsmB1(startA, startB uint) uint
