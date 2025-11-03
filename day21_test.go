package adventofcode2017

import (
	"testing"
)

func TestDay21Part1Example(t *testing.T) {
	const want = 12
	lines := linesFromFilename(t, exampleFilename(21))
	rb, err := NewDay21(lines)
	if err != nil {
		t.Fatal(err)
	}
	grid := parseGrid(".#./..#/###")
	for range 2 {
		grid = enhance(grid, rb)
	}
	got := grid.CountOn()
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay21Part1(t *testing.T) {
	testDayPart(t, 21, filename, true, NewDay21, Day21, uint(194))
}

func BenchmarkDay21Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(21))
	for b.Loop() {
		rb, _ := NewDay21(lines)
		_ = Day21(rb, true)
	}
}

func TestDay21Part2(t *testing.T) {
	testDayPart(t, 21, filename, false, NewDay21, Day21, uint(2536879))
}

func BenchmarkDay21Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(21))
	for b.Loop() {
		rb, _ := NewDay21(lines)
		_ = Day21(rb, false)
	}
}
