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
	result := 0
	sign := 1
	if clockwise {
		sign = -1
	}
	for i := 2; i <= n; i++ {
		k := delta(i, axisfn)
		result += sign * k
	}
	return result
}

func tx(n int) int {
	return transform(n, math.Sin, false)
}

func ty(n int) int {
	return transform(n, math.Cos, false)
}

// Day03Part1 returns the number of steps for a given square.
// Uses direct ring calculation - O(1) instead of O(n).
func Day03Part1(square int) int {
	if square == 1 {
		return 0
	}

	// Find which ring this number is in
	// Ring k contains numbers from (2k-1)^2 + 1 to (2k+1)^2
	// Ring 0: 1, Ring 1: 2-9, Ring 2: 10-25, etc.
	ring := 0
	for (2*ring+1)*(2*ring+1) < square {
		ring++
	}

	// Maximum number in previous ring
	prevRingMax := (2*ring - 1) * (2*ring - 1)

	// Position within the current ring (0-indexed)
	posInRing := square - prevRingMax - 1

	// Each side has 2*ring positions
	sideSize := 2 * ring

	// Position on current side (0 to sideSize-1)
	posOnSide := posInRing % sideSize

	// Distance from center of side (center is at position ring-1)
	distFromCenter := abs(posOnSide - (ring - 1))

	// Manhattan distance is ring distance + distance from side center
	return ring + distFromCenter
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
