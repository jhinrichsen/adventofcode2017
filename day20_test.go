package adventofcode2017

import (
	"testing"
)

func TestDay20Part1Example(t *testing.T) {
	testDayPart(t, 20, exampleFilename, true, NewDay20, Day20, uint(0))
}

func TestDay20Part1(t *testing.T) {
	testDayPart(t, 20, filename, true, NewDay20, Day20, uint(91))
}

func BenchmarkDay20Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(20))
	for b.Loop() {
		p, _ := NewDay20(lines)
		_ = Day20(p, true)
	}
}

func TestDay20Part2Example(t *testing.T) {
	testDayPart(t, 20, func(uint8) string { return "testdata/day20_example2.txt" }, false, NewDay20, Day20, uint(1))
}

func TestDay20Part2(t *testing.T) {
	testDayPart(t, 20, filename, false, NewDay20, Day20, uint(567))
}

func BenchmarkDay20Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(20))
	for b.Loop() {
		p, _ := NewDay20(lines)
		_ = Day20(p, false)
	}
}
