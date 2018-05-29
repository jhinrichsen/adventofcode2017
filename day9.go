package adventofcode2017

// Day9 is O(n) and allocation free (0 B/op).
func Day9(stream []byte) int {
	score := 0
	level := 0
	inGarbage := false
	// when using range, Go does not allow changing loop index
	// for i := range stream {
	for i := 0; i < len(stream); i++ {
		// Prio 0: skip character after !
		if stream[i] == '!' {
			i++
			continue
		}
		// Prio 1: skip garbage
		if stream[i] == '<' {
			inGarbage = true
			continue
		}
		if stream[i] == '>' {
			inGarbage = false
			continue
		}
		if inGarbage {
			continue
		}
		if stream[i] == '{' {
			level++
			continue
		}
		if stream[i] == '}' {
			score += level
			level--
			continue
		}
	}
	return score
}
