package adventofcode2017

import (
	"testing"
)

func TestDay24Part1Example(t *testing.T) {
	testDayPart(t, 24, exampleFilename, true, NewDay24, Day24, uint(31))
}

func TestDay24Part1(t *testing.T) {
	testDayPart(t, 24, filename, true, NewDay24, Day24, uint(1906))
}

func BenchmarkDay24Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(24))
	for b.Loop() {
		components, _ := NewDay24(lines)
		_ = Day24(components, true)
	}
}

func TestDay24Part2Example(t *testing.T) {
	testDayPart(t, 24, exampleFilename, false, NewDay24, Day24, uint(19))
}

func TestDay24Part2(t *testing.T) {
	testDayPart(t, 24, filename, false, NewDay24, Day24, uint(1824))
}

func BenchmarkDay24Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(24))
	for b.Loop() {
		components, _ := NewDay24(lines)
		_ = Day24(components, false)
	}
}
