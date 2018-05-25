package adventofcode2017

func Day2(spreadsheet [][]int) int {
	sum := 0
	for _, row := range spreadsheet {
		min := row[0]
		max := row[0]
		for _, cell := range row[1:] {
			if cell < min {
				min = cell
			}
			if cell > max {
				max = cell
			}
		}
		sum += max - min
		// log.Printf("row: min=%d, max=%d, sum=%d\n", min, max, sum)
	}
	return sum
}
