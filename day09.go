package adventofcode2017

func Day09Part1(stream []byte) int {
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

// Day09Part2 counts the number of non-canceled characters within garbage.
// The opening '<' and closing '>' are not counted. Canceled characters and the
// '!' doing the canceling are not counted either.
func Day09Part2(stream []byte) int {
	count := 0
	inGarbage := false
	for i := 0; i < len(stream); i++ {
		if stream[i] == '!' { // cancel next character
			i++
			continue
		}
		if inGarbage {
			if stream[i] == '>' { // end of garbage
				inGarbage = false
			} else { // any other character inside garbage counts
				count++
			}
			continue
		}
		if stream[i] == '<' { // start of garbage
			inGarbage = true
			continue
		}
		// all other characters outside garbage are ignored for part 2
	}
	return count
}
