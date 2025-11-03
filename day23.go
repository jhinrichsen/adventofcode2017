package adventofcode2017

import (
	"strconv"
	"strings"
)

type Day23Instruction struct {
	Op string
	X  string
	Y  string
}

type Day23Puzzle []Day23Instruction

func NewDay23(lines []string) (Day23Puzzle, error) {
	instructions := make([]Day23Instruction, 0, len(lines))
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		inst := Day23Instruction{Op: parts[0], X: parts[1]}
		if len(parts) > 2 {
			inst.Y = parts[2]
		}
		instructions = append(instructions, inst)
	}
	return instructions, nil
}

func getValue(s string, regs map[string]int) int {
	if val, err := strconv.Atoi(s); err == nil {
		return val
	}
	return regs[s]
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Day23(prog Day23Puzzle, part1 bool) uint {
	if part1 {
		regs := make(map[string]int)
		pc := 0
		mulCount := uint(0)

		for pc >= 0 && pc < len(prog) {
			inst := prog[pc]

			switch inst.Op {
			case "set":
				regs[inst.X] = getValue(inst.Y, regs)
			case "sub":
				regs[inst.X] -= getValue(inst.Y, regs)
			case "mul":
				regs[inst.X] *= getValue(inst.Y, regs)
				mulCount++
			case "jnz":
				if getValue(inst.X, regs) != 0 {
					pc += getValue(inst.Y, regs)
					continue
				}
			}
			pc++
		}

		return mulCount
	}

	b := 109900
	c := 126900
	h := uint(0)

	for b <= c {
		if !isPrime(b) {
			h++
		}
		b += 17
	}

	return h
}
