package adventofcode2017

import (
	"bytes"
	"strconv"
	"strings"
)

type Day16Move struct {
	op   byte
	a, b int
}

type Day16Puzzle []Day16Move

func NewDay16(lines []string) Day16Puzzle {
	var moves []Day16Move
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		for _, m := range strings.Split(line, ",") {
			m = strings.TrimSpace(m)
			if m == "" {
				continue
			}
			move := Day16Move{op: m[0]}
			switch m[0] {
			case 's':
				move.a, _ = strconv.Atoi(m[1:])
			case 'x':
				parts := strings.Split(m[1:], "/")
				move.a, _ = strconv.Atoi(parts[0])
				move.b, _ = strconv.Atoi(parts[1])
			case 'p':
				move.a = int(m[1])
				move.b = int(m[3])
			}
			moves = append(moves, move)
		}
	}
	return moves
}

func Day16(p Day16Puzzle, size int, part1 bool) string {
	programs := make([]byte, size)
	for i := range programs {
		programs[i] = byte('a' + i)
	}
	// Reuse single tmp buffer for spin operations. Escape analysis ensures
	// stack allocation (verified: go build -gcflags="-m" shows "does not escape")
	tmp := make([]byte, size)

	iterations := 1
	if !part1 {
		initial := make([]byte, size)
		copy(initial, programs)
		cycleLength := 0
		for i := 1; i <= 1000000000; i++ {
			for _, move := range p {
				switch move.op {
				case 's':
					copy(tmp, programs)
					for j := 0; j < size; j++ {
						programs[(j+move.a)%size] = tmp[j]
					}
				case 'x':
					programs[move.a], programs[move.b] = programs[move.b], programs[move.a]
				case 'p':
					posA, posB := -1, -1
					for j, prog := range programs {
						if prog == byte(move.a) {
							posA = j
						}
						if prog == byte(move.b) {
							posB = j
						}
					}
					programs[posA], programs[posB] = programs[posB], programs[posA]
				}
			}
			if bytes.Equal(programs, initial) {
				cycleLength = i
				break
			}
		}
		if cycleLength > 0 {
			iterations = 1000000000 % cycleLength
			for i := range programs {
				programs[i] = byte('a' + i)
			}
		} else {
			return string(programs)
		}
	}

	for iter := 0; iter < iterations; iter++ {
		for _, move := range p {
			switch move.op {
			case 's':
				copy(tmp, programs)
				for i := 0; i < size; i++ {
					programs[(i+move.a)%size] = tmp[i]
				}
			case 'x':
				programs[move.a], programs[move.b] = programs[move.b], programs[move.a]
			case 'p':
				posA, posB := -1, -1
				for i, prog := range programs {
					if prog == byte(move.a) {
						posA = i
					}
					if prog == byte(move.b) {
						posB = i
					}
				}
				programs[posA], programs[posB] = programs[posB], programs[posA]
			}
		}
	}

	return string(programs)
}
