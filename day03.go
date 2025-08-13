// Spiral memory
//
// 17  16  15  14  13
// 18   5   4   3  12
// 19   6   1   2  11
// 20   7   8   9  10
// 21  22  23---> ...
//

package adventofcode2017

import "math"

// abs returns the absolute value of an integer
func abs(n int) int {
	// use branching impl instead of y ← x >> 63, (x XOR y) - y
	if n < 0 {
		return -n
	}
	return n
}

// delta returns transformation for one single square
func delta(n int, axisfn func(float64) float64) int {
	f := float64(4*(n-2) + 1)
	f = math.Floor(math.Sqrt(f))
	k := int(f) % 4
	f = axisfn(float64(k) * math.Pi / 2)
	k = int(f)
	return k
}

// transform will project a point on a spiral into either X or Y coordinates for
// both clockwise and counterclockwise spirals.
// projecting a spiral onto X and Y involves (co-)sine, and oeis has appropriate
// formulas (as always).
//
// A174344 lists x-coordinates of point moving in clockwise spiral.
// https://oeis.org/A174344
// A268038 lists y-coordinates of point moving in clockwise spiral.
// https://oeis.org/A268038
func transform(n int, axisfn func(float64) float64, clockwise bool) int {
	if n == 1 {
		return 0
	}
	k := delta(n, axisfn)
	prev := transform(n-1, axisfn, clockwise)
	if clockwise {
		return prev - k
	}
	return prev + k
}

func tx(n int) int {
	return transform(n, math.Sin, false)
}

func ty(n int) int {
	return transform(n, math.Cos, false)
}

// Day03Part1 returns the number of steps for a given square.
// A recursive implementation is used that produces a stack overflow for values
// around 1e8. The recursive call can be replaced by a loop based impl using
// delta() only.
func Day03Part1(square int) int {
	x, y := tx(square), ty(square)
	// Steps are x + y coordinates, ignoring any sign indicating up/down
	// resp. left/right
	// theoretically, clockwise spirals and counterclockwise spirals should
	// produce the same number of steps
	steps := abs(x) + abs(y)
	return steps
}

// Day03Part2 returns the first number larger than n of https://oeis.org/A141481.
func Day03Part2(n uint) uint {
	for _, val := range A141481 {
		if val > n {
			return val
		}
	}
	return A141481[len(A141481)-1]
}
