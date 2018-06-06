package adventofcode2017

import (
	"bufio"
	"io"
	"strings"
)

func Day07(r io.Reader) string {
	// spec suggests that programs cannot be run multiple times
	// => avoid tree stuff, just use a plain map and keep
	// child -> parent instead of parent -> child
	var bottoms = make(map[string]string)
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		fs := strings.Fields(sc.Text())
		// only interested in parents
		if len(fs) > 3 {
			bottom := fs[0]
			tops := fs[3:]
			// remove optional comma separator suffix
			for _, top := range tops {
				top := strings.Trim(top, ",")
				bottoms[top] = bottom
			}
		}
	}

	// pick first (any) entry ...
	var pid string
	for key := range bottoms {
		pid = key
		break
	}
	// ... and follow parents until adam is found
	for len(bottoms[pid]) > 0 {
		pid = bottoms[pid]
	}
	return pid
}
