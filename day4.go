// High-Entropy Passphrases

package adventofcode2017

import "strings"

// Day4 returns the number of valid pass phrases.
// valid := contains no duplicate words, separated by space
func Day04(passphrases []string) int {
	n := 0
outer:
	for _, p := range passphrases {
		var words = make(map[string]bool)
		for _, word := range strings.Fields(p) {
			if words[word] {
				continue outer
			}
			words[word] = true
		}
		n++
	}
	return n
}
