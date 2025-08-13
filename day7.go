package adventofcode2017

import (
	"fmt"
	"strconv"
	"strings"
)

func Day7Part1(ss []string) string {
	// spec suggests that programs cannot be run multiple times
	// => avoid tree stuff, just use a plain map and keep
	// child -> parent instead of parent -> child
	var bottoms = make(map[string]string)
	for _, s := range ss {
		fs := strings.Fields(s)
		// only interested in parents
		if len(fs) > 3 {
			bottom := fs[0]
			tops := fs[3:]
			// remove optional comma separator suffix
			for _, top := range tops {
				top := strings.Trim(top, ",")
				bottoms[top] = bottom
			}
		}
	}

	// pick first (any) entry ...
	var pid string
	for key := range bottoms {
		pid = key
		break
	}
	// ... and follow parents until adam is found
	for len(bottoms[pid]) > 0 {
		pid = bottoms[pid]
	}
	return pid
}

type Disk map[string]bool

type program struct {
	Name   string
	Weight int
	Disks  Disk

	Level      int // level 0 is bottom
	DiskWeight int
}

func (a program) TotalWeight() int {
	return a.Weight + a.DiskWeight
}

// NewProgram parses a string represenation into a domain model.
// Examples:
// ktlj (57)
// fwft (72) -> ktlj, cntj, xhth
func NewProgram(s string) (program, error) {
	parts := strings.Fields(s)
	n, err := strconv.Atoi(parts[1][1 : len(parts[1])-1])
	if err != nil {
		return program{}, err
	}
	p := program{parts[0], n, nil, -1, 0}
	if len(parts) > 3 {
		p.Disks = make(Disk)
		for _, s := range parts[3:] {
			p.Disks[strings.TrimRight(s, ",")] = true
		}
	}
	return p, nil
}

// Day7Part2 returns the excess weight to balance the complete tower.
func Day7Part2(ss []string) (uint, error) {
	programs := make(map[string]program)
	bottoms := make(map[string]string)

	// parse programs

	var line uint
	for _, s := range ss {
		line++ // 1-based line counter
		p, err := NewProgram(s)
		if err != nil {
			return 0,
				fmt.Errorf("error parsing line %d: %v", line, err)
		}
		for key := range p.Disks {
			bottoms[key] = p.Name
		}
		programs[p.Name] = p
	}

	// calculate level for each program

	level := func(p string) int {
		n := -1
		hasBottom := true
		for hasBottom {
			p, hasBottom = bottoms[p]
			n++
		}
		return n
	}
	var maxLevel int
	// TODO level map instead of level member
	for _, p := range programs {
		p.Level = level(p.Name)
		programs[p.Name] = p
		if p.Level > maxLevel {
			maxLevel = p.Level
		}
	}

	// calculate disk weights top down

	for level := maxLevel - 1; level >= 0; level-- {
		for _, p := range programs {
			if p.Level != level {
				continue
			}
			for d := range p.Disks {
				p2 := programs[d]
				p.DiskWeight += p2.TotalWeight()
			}
			programs[p.Name] = p
		}
	}

	balanced := func(p program) bool {
		m := make(map[int]bool)
		for name := range p.Disks {
			w := programs[name].TotalWeight()
			m[w] = true
		}
		return len(m) == 1
	}

	// find highest unbalanced program bottom up

	bottom := func() program {
		for i := range programs {
			if programs[i].Level == 0 {
				return programs[i]
			}
		}
		panic("cannot find bottom program")
	}
	// list of unbalanced disks
	unbalanced := func(p program) []program {
		var ps []program
		for s := range p.Disks {
			p2 := programs[s]
			if !balanced(p2) {
				ps = append(ps, p2)
			}
		}
		return ps
	}
	p := bottom()
	for {
		ps := unbalanced(p)
		if len(ps) == 0 {
			// reached the node that is unbalanced, but its disk is balanced
			break
		}
		if len(ps) != 1 {
			s := fmt.Sprintf("expecting exactly 1 unbalanced disk but got %d", len(ps))
			panic(s)
		}
	}
	return 0, fmt.Errorf("missing implementation")
}
