package adventofcode2017

import (
	"reflect"
	"strings"
)

// Day04 returns the number of valid pass phrases.
// valid := contains no duplicate words, separated by space
func Day04Part1(passphrases []string) (count uint) {
outer:
	for _, p := range passphrases {
		var words = make(map[string]bool)
		for _, word := range strings.Fields(p) {
			if words[word] {
				continue outer
			}
			words[word] = true
		}
		count++
	}
	return
}

func Day04Part2Valid(passphrase string) bool {
	parts := strings.Fields(passphrase)
	var ms []map[rune]uint
	for _, part := range parts {
		m := make(map[rune]uint)
		for _, c := range part {
			// don't consider just occurrences per se, the number
			// must match
			m[c] = m[c] + 1
		}
		ms = append(ms, m)
	}
	// no two maps must match
	for i, left := range ms {
		for _, right := range ms[i+1:] {
			if reflect.DeepEqual(left, right) {
				return false
			}
		}
	}
	return true
}
func Day04Part2(passphrases []string) (count uint) {
	for i := range passphrases {
		if Day04Part2Valid(passphrases[i]) {
			count++
		}
	}
	return
}
