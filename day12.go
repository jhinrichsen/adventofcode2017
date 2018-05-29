// started 17:19
// stopped 18:08

package adventofcode2017

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// Day12 returns number of programs being in the group that contains program ID
// 0.
func Day12(r io.Reader) int {
	atoi := func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	}
	groups := make(map[int][]int)
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		fs := strings.Split(sc.Text(), " <-> ")
		key := atoi(fs[0])
		var values []int
		for _, val := range strings.Split(fs[1], ", ") {
			values = append(values, atoi(val))
		}
		groups[key] = values
	}

	// avoid cyclic references/ duplicate counts
	visited := make(map[int]bool)

	// look mom - recursive anonymous functions!
	var count func(int) int
	count = func(id int) int {
		var n int
		for _, succ := range groups[id] {
			if !visited[succ] {
				visited[succ] = true
				// count me...
				n++
				// ... and my children
				n += count(succ)
			}
		}
		return n
	}
	return count(0)
}
