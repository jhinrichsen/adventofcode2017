package adventofcode2017

// Day13UsingSparseArray returns severity of whole trip through layers.
func Day13UsingSparseArray(layers []int) int {
	// position of security scanner within one layer
	positions := make([]int, len(layers))
	// direction of security scanner, 1: down, -1: up
	directions := make([]int, len(layers))
	for i := range directions {
		directions[i] = -1
	}

	severity := 0
	for i := range layers {
		// Tick
		// can never be caught on sparse layer
		caught := layers[i] != 0 && positions[i] == 0
		if caught {
			depth := i
			rng := layers[i]
			severity += depth * rng
		}
		// Tack
		// ignore any layers already passed
		for j := i + 1; j < len(layers); j++ {
			sparse := layers[j] == 0
			if sparse {
				continue
			}
			// turn around if at first or last position
			first := positions[j] == 0
			last := positions[j] == layers[j]-1
			if first || last {
				directions[j] *= -1
			}
			positions[j] += directions[j]
		}
	}
	return severity
}

type node struct {
	Next      *node
	Depth     int
	Range     int
	Position  int // Current position of security scanner 0..Range
	Direction int
}

// Tack progresses security scanner one step
func (a *node) Tack() {
	// lazy init
	if a.Direction == 0 {
		a.Direction = -1
	}

	// Need to turn around?
	if a.Position == 0 || a.Position == a.Range-1 {
		a.Direction *= -1
	}
	a.Position += a.Direction
}

func (a *node) Caught() bool {
	return a.Position == 0
}

func (a *node) Severity() int {
	return a.Depth * a.Range
}

// Day13UsingList returns severity of whole trip through layers.
func Day13UsingList(l *node) int {
	severity := 0
	for ; l != nil; l = l.Next {
		// progress through all layers
		for i := 0; i < l.Depth; i++ {
			l.Tack()
		}
		if l.Caught() {
			severity += l.Depth * l.Range
		}
	}
	return severity
}
