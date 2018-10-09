package adventofcode2017

// Day01 finds the sum of all digits that match the next digit
func Day01(digits []byte) int {
	sum := 0
	cycle := len(digits)
	for i, digit := range digits {
		succ := (i + 1) % cycle
		if digit == digits[succ] {
			sum += int(digit)
		}
	}
	return sum
}
