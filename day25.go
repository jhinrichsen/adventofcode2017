package adventofcode2017

import (
	"strconv"
	"strings"
)

type Action struct {
	WriteValue int
	MoveDir    int
	NextState  string
}

type StateRules struct {
	OnZero Action
	OnOne  Action
}

type Day25Puzzle struct {
	StartState string
	Steps      uint
	States     map[string]StateRules
}

func NewDay25(lines []string) (Day25Puzzle, error) {
	puzzle := Day25Puzzle{
		States: make(map[string]StateRules),
	}

	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])

		if strings.HasPrefix(line, "Begin in state") {
			puzzle.StartState = string(line[15])
		} else if strings.Contains(line, "after") && strings.Contains(line, "steps") {
			parts := strings.Fields(line)
			for _, part := range parts {
				if val, err := strconv.Atoi(part); err == nil {
					puzzle.Steps = uint(val)
					break
				}
			}
		} else if strings.HasPrefix(line, "In state") {
			stateName := string(line[9])
			rules := StateRules{}

			i++
			for i < len(lines) {
				line = strings.TrimSpace(lines[i])
				if line == "" {
					break
				}

				if strings.Contains(line, "current value is 0") {
					i++
					writeVal := 0
					if strings.Contains(strings.TrimSpace(lines[i]), "Write the value 1") {
						writeVal = 1
					}
					i++
					moveDir := 1
					if strings.Contains(strings.TrimSpace(lines[i]), "left") {
						moveDir = -1
					}
					i++
					nextState := string(strings.TrimSpace(lines[i])[len(strings.TrimSpace(lines[i]))-2])

					rules.OnZero = Action{
						WriteValue: writeVal,
						MoveDir:    moveDir,
						NextState:  nextState,
					}
				} else if strings.Contains(line, "current value is 1") {
					i++
					writeVal := 0
					if strings.Contains(strings.TrimSpace(lines[i]), "Write the value 1") {
						writeVal = 1
					}
					i++
					moveDir := 1
					if strings.Contains(strings.TrimSpace(lines[i]), "left") {
						moveDir = -1
					}
					i++
					nextState := string(strings.TrimSpace(lines[i])[len(strings.TrimSpace(lines[i]))-2])

					rules.OnOne = Action{
						WriteValue: writeVal,
						MoveDir:    moveDir,
						NextState:  nextState,
					}
				} else {
					i++
				}
			}

			puzzle.States[stateName] = rules
		}
	}

	return puzzle, nil
}

func Day25(puzzle Day25Puzzle, part1 bool) (uint, error) {
	tape := make(map[int]int)
	cursor := 0
	state := puzzle.StartState

	for range puzzle.Steps {
		currentValue := tape[cursor]
		rules := puzzle.States[state]

		var action Action
		if currentValue == 0 {
			action = rules.OnZero
		} else {
			action = rules.OnOne
		}

		tape[cursor] = action.WriteValue
		cursor += action.MoveDir
		state = action.NextState
	}

	checksum := uint(0)
	for _, val := range tape {
		if val == 1 {
			checksum++
		}
	}

	return checksum, nil
}
