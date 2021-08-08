package adventofcode2017

// Banks domain model.
// As a Go specialty, arrays can be used as map keys (slices can't).
// Therefore impl #1 uses an array and a length (the sample uses 4 banks, the
// official input is 16 banks).

type Banks [16]int

// Day6Part1 returns number of redistribution cycles.
func Day6Part1(banks Banks, activeBanks int) int {
	_, cycles := Day6(banks, activeBanks)
	return cycles
}

// Day6Part2 returns number of redistribution cycles for a configuration already
// seen.
func Day6Part2(banks Banks, activeBanks int) int {
	bs, _ := Day6(banks, activeBanks)
	_, cycles := Day6(bs, activeBanks)
	return cycles - 1
}

func Day6(banks Banks, activeBanks int) (Banks, int) {
	// return index and value of max bank
	max := func() (int, int) {
		idx := 0
		blocks := banks[idx]
		for i := 1; i < activeBanks; i++ {
			// lower bank wins on tie
			if blocks < banks[i] {
				idx = i
				blocks = banks[idx]
			}
		}
		return idx, blocks
	}
	seen := make(map[Banks]bool)
	next := func(i int) int {
		return (i + 1) % activeBanks
	}
	redistributions := 0
	for {
		i, n := max()
		banks[i] = 0
		// redistribute
		j := next(i)
		for n > 0 {
			banks[j]++
			n--

			j = next(j)
		}
		redistributions++

		if seen[banks] {
			break
		}
		seen[banks] = true
		// dump6(banks)
	}
	return banks, redistributions
}
