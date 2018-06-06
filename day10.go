package adventofcode2017

import (
	"fmt"
	"log"
)

var Day10LengthError = fmt.Errorf("illegal length")

func Day10(list []int, lengths []int) int {
	day10Hash(list, lengths)
	return list[0] * list[1]
}

func day10Hash(list []int, lengths []int) {
}

// spec states "Lengths larger than the size of the list are invalid." which is
// a bit unclear as to wether this can occur in input or not. We opt for the
// safe side which makes it a precondition.
// Day10PreconditionLength checks if all length parameters are in list bounds.
// Returns a list of illegal length indices.
func Day10PreconditionLength(list []int, lengths []int) ([]int, error) {
	errors := []int{}
	for i := 0; i < len(lengths); i++ {
		if lengths[i] > len(list) {
			errors = append(errors, i)
		}
	}
	if len(errors) > 0 {
		return errors, Day10LengthError
	}
	return nil, nil
}

// in-place reverse slice list[idx:idx+n], wrapping n if n > len(list)
func Day10Reverse(list []int, idx int, n int) {
	wrap := func(i int) int {
		return i % len(list)
	}
	swap := func(x, y int) {
		log.Printf("swap before: %+v\n", list)
		log.Printf("swapping index %d <-> %d\n", x, y)
		log.Printf("swapping logical index %d <-> %d\n", wrap(x), wrap(y))
		list[wrap(x)], list[wrap(y)] = list[wrap(y)], list[wrap(x)]
		log.Printf("swap after: %+v\n", list)
	}
	for lower, upper := idx, idx+n-1; lower < upper; lower, upper = lower+1, upper-1 {
		swap(lower, upper)
	}

}
