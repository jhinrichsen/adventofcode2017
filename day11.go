package adventofcode2017

import (
	"fmt"
	"strings"
)

// NewDay11 parses a comma-separated list of steps into compact integer codes.
// Valid steps are: n, ne, se, s, sw, nw. It returns an error on invalid tokens.
func NewDay11(buf []byte) ([]int, error) {
	s := strings.TrimSpace(string(buf))
	if s == "" {
		return nil, nil
	}
	parts := strings.Split(s, ",")
	steps := make([]int, 0, len(parts))
	for i := range parts {
		tok := strings.TrimSpace(parts[i])
		switch tok {
		case "n":
			steps = append(steps, 0)
		case "ne":
			steps = append(steps, 1)
		case "se":
			steps = append(steps, 2)
		case "s":
			steps = append(steps, 3)
		case "sw":
			steps = append(steps, 4)
		case "nw":
			steps = append(steps, 5)
		default:
			return nil, fmt.Errorf("day11: invalid step %q at index %d", tok, i)
		}
	}
	return steps, nil
}

// Day11Part1 returns the distance from the origin after walking all steps.
func Day11Part1(steps []int) uint {
	x, y, z := 0, 0, 0
	for i := range steps {
		switch steps[i] {
		case 0: // n
			y, z = y+1, z-1
		case 1: // ne
			x, z = x+1, z-1
		case 2: // se
			x, y = x+1, y-1
		case 3: // s
			y, z = y-1, z+1
		case 4: // sw
			x, z = x-1, z+1
		case 5: // nw
			x, y = x-1, y+1
		}
	}
	return day11Distance(x, y, z)
}

// Day11Part2 returns the maximum distance from the origin reached at any point.
func Day11Part2(steps []int) uint {
    x, y, z := 0, 0, 0
    var maxd uint
    for i := range steps {
        switch steps[i] {
        case 0: // n
            y, z = y+1, z-1
        case 1: // ne
            x, z = x+1, z-1
        case 2: // se
            x, y = x+1, y-1
        case 3: // s
            y, z = y-1, z+1
        case 4: // sw
            x, z = x-1, z+1
        case 5: // nw
            x, y = x-1, y+1
        }
        maxd = max(maxd, day11Distance(x, y, z))
    }
    return maxd
}

func day11Distance(x, y, z int) uint {
	ax, ay, az := x, y, z
	if ax < 0 {
		ax = -ax
	}
	if ay < 0 {
		ay = -ay
	}
	if az < 0 {
		az = -az
	}
	return uint((ax + ay + az) / 2)
}
