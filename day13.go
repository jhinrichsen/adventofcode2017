package adventofcode2017

import (
	"fmt"
	"strconv"
	"strings"
)

func NewDay13Array(lines []string) ([]int, error) {
	var m = make(map[int]int)
	var maxDepth int
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("want key: value but got %v", line)
		}
		k, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("want numeric key but got %v", parts[0])
		}
		v, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("want numeric key but got %v", parts[1])
		}
		m[k] = v
		if maxDepth < k {
			maxDepth = k
		}
	}

	// convert map to array
	layers := make([]int, maxDepth+1)
	for k, v := range m {
		layers[k] = v
	}
	return layers, nil
}

// Day13UsingArray returns severity of whole trip through layers.
func Day13UsingArray(layers []int) int {
	// position of security scanner within one layer
	positions := make([]int, len(layers))
	// direction of security scanner, 1: down, -1: up
	directions := make([]int, len(layers))
	for i := range directions {
		directions[i] = -1
	}

	severity := 0
	for i := range layers {
		// Tick
		// can never be caught on sparse layer
		caught := layers[i] != 0 && positions[i] == 0
		if caught {
			depth := i
			rng := layers[i]
			severity += depth * rng
		}
		// Tack
		// ignore any layers already passed
		for j := i + 1; j < len(layers); j++ {
			sparse := layers[j] == 0
			if sparse {
				continue
			}
			// turn around if at first or last position
			first := positions[j] == 0
			last := positions[j] == layers[j]-1
			if first || last {
				directions[j] *= -1
			}
			positions[j] += directions[j]
		}
	}
	return severity
}

type Day13Puzzle struct {
	Next      *Day13Puzzle
	Depth     int
	Range     int
	Position  int // Current position of security scanner 0..Range
	Direction int
}

func NewDay13(lines []string) (*Day13Puzzle, error) {
	var head, prev *Day13Puzzle
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("want 'key: value' but got %v", line)
		}
		k, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, fmt.Errorf("want numeric key but got %v", parts[0])
		}
		v, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("want numeric key but got %v", parts[1])
		}
		l := &Day13Puzzle{
			Depth: k,
			Range: v,
		}
		if head == nil {
			head = l
		}
		if prev != nil {
			prev.Next = l
		}
		prev = l
	}
	return head, nil
}

// Day13 calculates either the severity (part1=true) or minimum delay (part1=false).
// For part 1: returns severity of whole trip through layers.
// For part 2: returns the minimum delay needed to pass through without being caught.
// Scanner positions are deterministic: a scanner with range R has period 2*(R-1).
// We get caught at layer depth D if (delay + D) % (2 * (R - 1)) == 0.
func Day13(l *Day13Puzzle, part1 bool) uint {
	if part1 {
		// Part 1: calculate severity
		var severity uint
		for layer := l; layer != nil; layer = layer.Next {
			period := 2 * (layer.Range - 1)
			if layer.Depth%period == 0 {
				severity += uint(layer.Depth * layer.Range)
			}
		}
		return severity
	}

	// Part 2: find minimum delay
	for delay := 0; ; delay++ {
		caught := false
		for layer := l; layer != nil; layer = layer.Next {
			period := 2 * (layer.Range - 1)
			if (delay+layer.Depth)%period == 0 {
				caught = true
				break
			}
		}
		if !caught {
			return uint(delay)
		}
	}
}
