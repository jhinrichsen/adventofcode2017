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

func Day22(vc *Day22Puzzle, bursts uint, part1 bool) (uint, error) {
	for range bursts {
		vc.Burst(part1)
	}
	return vc.Infections, nil
}
