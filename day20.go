package adventofcode2017

import (
	"strconv"
	"strings"
)

type Vec3 struct {
	X, Y, Z int
}

func (v Vec3) ManhattanDistance() int {
	return abs(v.X) + abs(v.Y) + abs(v.Z)
}

type Particle struct {
	Pos, Vel, Acc Vec3
}

type Day20Puzzle []Particle

func parseVec3(s string) (Vec3, error) {
	s = strings.TrimPrefix(s, "<")
	s = strings.TrimSuffix(s, ">")
	parts := strings.Split(s, ",")
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return Vec3{}, err
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return Vec3{}, err
	}
	z, err := strconv.Atoi(parts[2])
	if err != nil {
		return Vec3{}, err
	}
	return Vec3{x, y, z}, nil
}

func parseParticle(line string) (Particle, error) {
	parts := strings.Split(line, ", ")

	pos, err := parseVec3(strings.TrimPrefix(parts[0], "p="))
	if err != nil {
		return Particle{}, err
	}
	vel, err := parseVec3(strings.TrimPrefix(parts[1], "v="))
	if err != nil {
		return Particle{}, err
	}
	acc, err := parseVec3(strings.TrimPrefix(parts[2], "a="))
	if err != nil {
		return Particle{}, err
	}

	return Particle{Pos: pos, Vel: vel, Acc: acc}, nil
}

func NewDay20(lines []string) (Day20Puzzle, error) {
	particles := make([]Particle, len(lines))
	for i, line := range lines {
		p, err := parseParticle(line)
		if err != nil {
			return nil, err
		}
		particles[i] = p
	}
	return particles, nil
}

func (p *Particle) tick() {
	p.Vel.X += p.Acc.X
	p.Vel.Y += p.Acc.Y
	p.Vel.Z += p.Acc.Z
	p.Pos.X += p.Vel.X
	p.Pos.Y += p.Vel.Y
	p.Pos.Z += p.Vel.Z
}

func Day20(p Day20Puzzle, part1 bool) uint {
	if len(p) == 0 {
		return 0
	}

	if part1 {
		closestIdx := 0
		minAccel := p[0].Acc.ManhattanDistance()
		minVel := p[0].Vel.ManhattanDistance()
		minPos := p[0].Pos.ManhattanDistance()

		for i := 1; i < len(p); i++ {
			accel := p[i].Acc.ManhattanDistance()
			vel := p[i].Vel.ManhattanDistance()
			pos := p[i].Pos.ManhattanDistance()

			if accel < minAccel {
				closestIdx = i
				minAccel = accel
				minVel = vel
				minPos = pos
			} else if accel == minAccel {
				if vel < minVel {
					closestIdx = i
					minVel = vel
					minPos = pos
				} else if vel == minVel {
					if pos < minPos {
						closestIdx = i
						minPos = pos
					}
				}
			}
		}

		return uint(closestIdx)
	}

	particles := make([]Particle, len(p))
	copy(particles, p)

	for tick := 0; tick < 1000; tick++ {
		for i := range particles {
			particles[i].tick()
		}

		positionCount := make(map[Vec3]int)
		for i := range particles {
			positionCount[particles[i].Pos]++
		}

		surviving := make([]Particle, 0, len(particles))
		for i := range particles {
			if positionCount[particles[i].Pos] == 1 {
				surviving = append(surviving, particles[i])
			}
		}
		particles = surviving

		if len(particles) == 0 {
			break
		}
	}

	return uint(len(particles))
}
