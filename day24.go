package adventofcode2017

import (
	"strconv"
	"strings"
)

type Component struct {
	A, B uint
}

func (c Component) Strength() uint {
	return c.A + c.B
}

type Day24Puzzle []Component

func NewDay24(lines []string) (Day24Puzzle, error) {
	components := make([]Component, 0, len(lines))
	for _, line := range lines {
		parts := strings.Split(line, "/")
		if len(parts) != 2 {
			continue
		}
		a, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}
		b, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, err
		}
		components = append(components, Component{A: uint(a), B: uint(b)})
	}
	return components, nil
}

type bridgeState struct {
	port     uint
	used     uint64 // Bitset for up to 64 components
	strength uint
	length   uint
}

func Day24(components Day24Puzzle, part1 bool) uint {
	maxStrength := uint(0)
	maxLength := uint(0)

	stack := []bridgeState{{
		port:     0,
		used:     0,
		strength: 0,
		length:   0,
	}}

	for len(stack) > 0 {
		state := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if part1 {
			if state.strength > maxStrength {
				maxStrength = state.strength
			}
		} else {
			if state.length > maxLength || (state.length == maxLength && state.strength > maxStrength) {
				maxLength = state.length
				maxStrength = state.strength
			}
		}

		for i, comp := range components {
			if state.used&(1<<i) != 0 {
				continue
			}

			var nextPort uint
			canConnect := false

			if comp.A == state.port {
				canConnect = true
				nextPort = comp.B
			} else if comp.B == state.port {
				canConnect = true
				nextPort = comp.A
			}

			if canConnect {
				stack = append(stack, bridgeState{
					port:     nextPort,
					used:     state.used | (1 << i),
					strength: state.strength + comp.Strength(),
					length:   state.length + 1,
				})
			}
		}
	}

	return maxStrength
}
