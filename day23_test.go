package adventofcode2017

import (
	"testing"
)

func TestDay23Part1(t *testing.T) {
	testDayPart(t, 23, filename, true, NewDay23, Day23, uint(9409))
}

func BenchmarkDay23Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(23))
	for b.Loop() {
		prog, _ := NewDay23(lines)
		_ = Day23(prog, true)
	}
}

func TestDay23Part2(t *testing.T) {
	testDayPart(t, 23, filename, false, NewDay23, Day23, uint(913))
}

func BenchmarkDay23Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(23))
	for b.Loop() {
		prog, _ := NewDay23(lines)
		_ = Day23(prog, false)
	}
}
