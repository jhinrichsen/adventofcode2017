package adventofcode2017

import (
	"fmt"
)

func dump6(banks []int) {
	for i := range banks {
		fmt.Printf("%3d ", banks[i])
	}
	fmt.Println()
}

// Day6 returns number of redistributions.
func Day6(banks []int) int {
	// return index and value of max bank
	max := func() (int, int) {
		idx := 0
		blocks := banks[idx]
		for i := 1; i < len(banks); i++ {
			// lower bank wins on tie
			if blocks < banks[i] {
				idx = i
				blocks = banks[idx]
			}
		}
		return idx, blocks
	}
	seen := make(map[int]bool)
	configuration := func() int {
		// use bank blocks as integer value, i.e. '0 2 7 0' -> 270
		x := 0
		exp := 1
		for i := len(banks) - 1; i >= 0; i-- {
			x += banks[i] * exp
			exp *= 10
		}
		return x
	}
	next := func(i int) int {
		return (i + 1) % len(banks)
	}
	redistributions := 0
	// dump6(banks)
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

		cfg := configuration()
		if seen[cfg] {
			break
		}
		seen[cfg] = true
		// dump6(banks)
	}
	return redistributions
}
