package adventofcode2017

// Day01Part1 finds the sum of all digits that match the next digit.
func Day01Part1(digits []byte) (sum uint) {
    return day1(digits, 1)
}

// Day01Part2 finds the sum of all digits that match the halfway around digit.
func Day01Part2(digits []byte) (sum uint) {
    return day1(digits, len(digits)/2)
}

func day1(digits []byte, n int) (sum uint) {
	cycle := len(digits)
	for i, digit := range digits {
		succ := (i + n) % cycle
		if digit == digits[succ] {
			sum += uint(digit)
		}
	}
	return
}
