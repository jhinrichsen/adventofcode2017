package adventofcode2017

// Day2Part1 calculates the spreadsheet's checksum for smallest and largest
// cells.
func Day2Part1(spreadsheet [][]int) (sum uint) {
	for _, row := range spreadsheet {
		min, max := row[0], row[0]
		for _, cell := range row[1:] {
			if cell < min {
				min = cell
			}
			if cell > max {
				max = cell
			}
		}
		sum += uint(max - min)
	}
	return
}

// Day2Part2 calculates checksum for evenly divisible cells.
func Day2Part2(spreadsheet [][]int) (sum uint) {
	// return divisor if evenly divisible or 0
	div := func(a, b int) uint {
		if a < b {
			a, b = b, a
		}
		d := a / b
		if d*b == a {
			return uint(d)
		}
		return 0
	}
	for _, row := range spreadsheet {
	nextRow:
		for i, c0 := range row {
			for _, cell := range row[i+1:] {
				d := div(c0, cell)
				if d > 0 {
					sum += d
					continue nextRow
				}
			}
		}
	}
	return
}
