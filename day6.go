// As a Go specialty, arrays can be used as map keys (slices can't).
// Therefore impl #1 uses an array and a length (the sample uses 4 banks, the
// official input is 16 banks), and impl #2 uses a String represenation.

package adventofcode2017

import (
	"fmt"
	"strconv"
	"strings"
)

func dump6(banks []int) {
	for i := range banks {
		fmt.Printf("%3d ", banks[i])
	}
	fmt.Println()
}

type Banks [16]int

// Day6 returns number of redistributions.
func Day06Impl1(banks Banks, activeBanks int) int {
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

		if seen[banks] {
			break
		}
		seen[banks] = true
		// dump6(banks)
	}
	return redistributions
}

func Day06Impl2(banks []int) int {
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
	seen := make(map[string]bool)
	configuration := func() string {
		// create array representation that can be used as map key
		ss := make([]string, len(banks))
		for i := 0; i < len(banks); i++ {
			ss[i] = strconv.Itoa(banks[i])
		}
		return strings.Join(ss, ".")
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
