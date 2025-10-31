package adventofcode2017

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type RA struct {
	reg    string
	amount int
}

type Operation struct {
	RA
	inc bool
}

type Condition struct {
	RA
	cmp Cmp
}

type Instruction struct {
	Operation
	Condition
}

type Cmp int

const (
	EQ Cmp = iota
	NE
	LT
	LE
	GT
	GE
)

type Registers map[string]int

func (a Registers) True(c Condition) bool {
	var b bool
	val := a[c.reg]
	switch c.cmp {
	case EQ:
		b = val == c.amount
	case NE:
		b = val != c.amount
	case LE:
		b = val <= c.amount
	case LT:
		b = val < c.amount
	case GE:
		b = val >= c.amount
	case GT:
		b = val > c.amount
	}
	return b
}

func (a Registers) Apply(op Operation) {
	val := a[op.reg]
	if op.inc {
		val += op.amount
	} else {
		val -= op.amount
	}
	a[op.reg] = val
}

func (a Registers) Step(i Instruction) {
	if a.True(i.Condition) {
		a.Apply(i.Operation)
	}
}

func Day08Part1(data []byte) (int, error) {
	registers := Registers{}
	is, err := instructions(data)
	if err != nil {
		return -1, err
	}
	for _, i := range is {
		registers.Step(i)
	}
	maxVal := math.MinInt
	for _, v := range registers {
		maxVal = max(maxVal, v)
	}
	return maxVal, nil
}

// Day08Part2 returns the highest value held in any register at any time
// during processing of the instructions.
func Day08Part2(data []byte) (int, error) {
	registers := Registers{}
	is, err := instructions(data)
	if err != nil {
		return -1, err
	}
	maxVal := math.MinInt
	for _, ins := range is {
		if registers.True(ins.Condition) {
			registers.Apply(ins.Operation)
			maxVal = max(maxVal, registers[ins.Operation.reg])
		}
	}
	return maxVal, nil
}

func instructions(data []byte) ([]Instruction, error) {
	var is []Instruction
	sc := bufio.NewScanner(bytes.NewReader(data))
	lineN := 0
	for sc.Scan() {
		lineN++
		line := sc.Text()
		i, err := parseInstruction(line)
		if err != nil {
			return is, fmt.Errorf("line %d: %v", lineN, err)
		}
		is = append(is, i)
	}
	return is, nil
}

func parseInstruction(is string) (Instruction, error) {
	var i Instruction
	fs := strings.Fields(is)
	want := 7
	got := len(fs)
	if want != got {
		return Instruction{}, fmt.Errorf("%s: want %d fields but got %d", is, want, got)
	}
	var col1 bool
	switch fs[1] {
	case "inc":
		col1 = true
	case "dec":
		col1 = false
	default:
		return i, fmt.Errorf("want 'inc' or 'dec' but got %s", fs[1])
	}

	col2, err := strconv.Atoi(fs[2])
	if err != nil {
		return i, fmt.Errorf("cannot convert column 2 to int: %s", fs[2])
	}
	col6, err := strconv.Atoi(fs[6])
	if err != nil {
		return i, fmt.Errorf("cannot convert column 6 to int: %s", fs[6])
	}

	var col5 Cmp
	switch fs[5] {
	case "==":
		col5 = EQ
	case "!=":
		col5 = NE
	case "<":
		col5 = LT
	case "<=":
		col5 = LE
	case ">":
		col5 = GT
	case ">=":
		col5 = GE
	default:
		return i, fmt.Errorf("want column 5 to be one of ==, <, <=, >, or <= but got %s", fs[5])
	}
	return Instruction{
		Operation{
			RA{
				fs[0],
				col2,
			},
			col1,
		},
		Condition{
			RA{
				fs[4],
				col6,
			},
			col5,
		},
	}, nil
}
