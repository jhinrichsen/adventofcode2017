package adventofcode2017

import (
	"testing"
)

func TestDay25Part1Example(t *testing.T) {
	testDayPart(t, 25, exampleFilename, true, NewDay25, Day25, uint(3))
}

func TestDay25Part1(t *testing.T) {
	testDayPart(t, 25, filename, true, NewDay25, Day25, uint(2794))
}

func BenchmarkDay25Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(25))
	for b.Loop() {
		puzzle, _ := NewDay25(lines)
		_ = Day25(puzzle, true)
	}
}
