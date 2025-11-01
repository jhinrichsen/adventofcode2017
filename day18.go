package adventofcode2017

import (
	"strconv"
	"strings"
	"sync"
)

type Day18Instruction struct {
	op     string
	x      string
	y      string
	yVal   int
	yIsReg bool
}

type Day18Puzzle []Day18Instruction

func NewDay18(lines []string) Day18Puzzle {
	var instructions []Day18Instruction
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		inst := Day18Instruction{op: parts[0]}
		if len(parts) > 1 {
			inst.x = parts[1]
		}
		if len(parts) > 2 {
			inst.y = parts[2]
			val, err := strconv.Atoi(parts[2])
			if err != nil {
				inst.yIsReg = true
			} else {
				inst.yVal = val
			}
		}
		instructions = append(instructions, inst)
	}
	return instructions
}

func Day18(p Day18Puzzle, part1 bool) uint {
	if part1 {
		registers := make(map[string]uint)
		lastSound := uint(0)
		pc := 0

		getValue := func(s string) uint {
			val, err := strconv.Atoi(s)
			if err != nil {
				return registers[s]
			}
			return uint(val)
		}

		for pc >= 0 && pc < len(p) {
			inst := p[pc]
			switch inst.op {
			case "snd":
				lastSound = getValue(inst.x)
			case "set":
				var val uint
				if inst.yIsReg {
					val = registers[inst.y]
				} else {
					val = uint(inst.yVal)
				}
				registers[inst.x] = val
			case "add":
				var val uint
				if inst.yIsReg {
					val = registers[inst.y]
				} else {
					val = uint(inst.yVal)
				}
				registers[inst.x] += val
			case "mul":
				var val uint
				if inst.yIsReg {
					val = registers[inst.y]
				} else {
					val = uint(inst.yVal)
				}
				registers[inst.x] *= val
			case "mod":
				var val uint
				if inst.yIsReg {
					val = registers[inst.y]
				} else {
					val = uint(inst.yVal)
				}
				registers[inst.x] %= val
			case "rcv":
				if getValue(inst.x) != 0 {
					return lastSound
				}
			case "jgz":
				if getValue(inst.x) > 0 {
					var val int
					if inst.yIsReg {
						val = int(registers[inst.y])
					} else {
						val = inst.yVal
					}
					pc += val
					continue
				}
			}
			pc++
		}
	}

	// Part 2: concurrent execution with message passing
	queue0 := make(chan int, 10000)
	queue1 := make(chan int, 10000)
	result := make(chan uint, 1)
	done := make(chan struct{})

	var mu sync.Mutex
	waiting := 0

	checkDeadlock := func() {
		if waiting == 2 && len(queue0) == 0 && len(queue1) == 0 {
			close(done)
		}
	}

	runProgram := func(pid int, sendChan, recvChan chan int) {
		registers := make(map[string]int)
		registers["p"] = pid
		pc := 0
		sendCount := uint(0)

		getValue := func(s string) int {
			val, err := strconv.Atoi(s)
			if err != nil {
				return registers[s]
			}
			return val
		}

		for pc >= 0 && pc < len(p) {
			select {
			case <-done:
				if pid == 1 {
					result <- sendCount
				}
				return
			default:
			}

			inst := p[pc]
			switch inst.op {
			case "snd":
				val := getValue(inst.x)
				select {
				case sendChan <- val:
					if pid == 1 {
						sendCount++
					}
				case <-done:
					if pid == 1 {
						result <- sendCount
					}
					return
				}
			case "set":
				var val int
				if inst.yIsReg {
					val = registers[inst.y]
				} else {
					val = inst.yVal
				}
				registers[inst.x] = val
			case "add":
				var val int
				if inst.yIsReg {
					val = registers[inst.y]
				} else {
					val = inst.yVal
				}
				registers[inst.x] += val
			case "mul":
				var val int
				if inst.yIsReg {
					val = registers[inst.y]
				} else {
					val = inst.yVal
				}
				registers[inst.x] *= val
			case "mod":
				var val int
				if inst.yIsReg {
					val = registers[inst.y]
				} else {
					val = inst.yVal
				}
				registers[inst.x] %= val
			case "rcv":
				mu.Lock()
				waiting++
				checkDeadlock()
				mu.Unlock()

				select {
				case val := <-recvChan:
					mu.Lock()
					waiting--
					mu.Unlock()
					registers[inst.x] = val
				case <-done:
					if pid == 1 {
						result <- sendCount
					}
					return
				}
			case "jgz":
				if getValue(inst.x) > 0 {
					var val int
					if inst.yIsReg {
						val = registers[inst.y]
					} else {
						val = inst.yVal
					}
					pc += val
					continue
				}
			}
			pc++
		}
		if pid == 1 {
			result <- sendCount
		}
	}

	go runProgram(0, queue1, queue0)
	go runProgram(1, queue0, queue1)

	return <-result
}
