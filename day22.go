package adventofcode2017

import "image"

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

func (d Direction) TurnLeft() Direction {
	return (d + 3) % 4
}

func (d Direction) TurnRight() Direction {
	return (d + 1) % 4
}

func (d Direction) Reverse() Direction {
	return (d + 2) % 4
}

type NodeState int

const (
	Clean NodeState = iota
	Weakened
	Infected
	Flagged
)

type Day22Puzzle struct {
	Pos        image.Point
	Dir        Direction
	States     map[image.Point]NodeState
	Infections uint
}

func NewDay22(lines []string) (*Day22Puzzle, error) {
	states := make(map[image.Point]NodeState)

	rows := len(lines)
	cols := 0
	if rows > 0 {
		cols = len(lines[0])
	}

	centerY := rows / 2
	centerX := cols / 2

	for y, line := range lines {
		for x, ch := range line {
			if ch == '#' {
				states[image.Point{X: x - centerX, Y: y - centerY}] = Infected
			}
		}
	}

	return &Day22Puzzle{
		Pos:    image.Point{X: 0, Y: 0},
		Dir:    Up,
		States: states,
	}, nil
}

func (vc *Day22Puzzle) Burst(part1 bool) {
	state := vc.States[vc.Pos]

	if part1 {
		if state == Infected {
			vc.Dir = vc.Dir.TurnRight()
			delete(vc.States, vc.Pos)
		} else {
			vc.Dir = vc.Dir.TurnLeft()
			vc.States[vc.Pos] = Infected
			vc.Infections++
		}
	} else {
		switch state {
		case Clean:
			vc.Dir = vc.Dir.TurnLeft()
			vc.States[vc.Pos] = Weakened
		case Weakened:
			vc.States[vc.Pos] = Infected
			vc.Infections++
		case Infected:
			vc.Dir = vc.Dir.TurnRight()
			vc.States[vc.Pos] = Flagged
		case Flagged:
			vc.Dir = vc.Dir.Reverse()
			delete(vc.States, vc.Pos)
		}
	}

	switch vc.Dir {
	case Up:
		vc.Pos.Y--
	case Down:
		vc.Pos.Y++
	case Left:
		vc.Pos.X--
	case Right:
		vc.Pos.X++
	}
}

var deltas = [4][2]int{
	{0, -1}, // Up
	{1, 0},  // Right
	{0, 1},  // Down
	{-1, 0}, // Left
}

func packPos(x, y int) int64 {
	return int64(x)<<32 | int64(uint32(y))
}

func Day22(vc *Day22Puzzle, bursts uint, part1 bool) (uint, error) {
	// Convert to packed coordinates
	states := make(map[int64]NodeState, len(vc.States))
	for pt, state := range vc.States {
		states[packPos(pt.X, pt.Y)] = state
	}

	x, y := vc.Pos.X, vc.Pos.Y
	dir := vc.Dir
	infections := vc.Infections

	if part1 {
		for range bursts {
			pos := packPos(x, y)
			state := states[pos]

			if state == Infected {
				dir = (dir + 1) % 4 // TurnRight
				delete(states, pos)
			} else {
				dir = (dir + 3) % 4 // TurnLeft
				states[pos] = Infected
				infections++
			}

			x += deltas[dir][0]
			y += deltas[dir][1]
		}
	} else {
		for range bursts {
			pos := packPos(x, y)
			state := states[pos]

			switch state {
			case Clean:
				dir = (dir + 3) % 4 // TurnLeft
				states[pos] = Weakened
			case Weakened:
				states[pos] = Infected
				infections++
			case Infected:
				dir = (dir + 1) % 4 // TurnRight
				states[pos] = Flagged
			case Flagged:
				dir = (dir + 2) % 4 // Reverse
				delete(states, pos)
			}

			x += deltas[dir][0]
			y += deltas[dir][1]
		}
	}

	return infections, nil
}
