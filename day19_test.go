package adventofcode2017

import (
	"testing"
)

func TestDay19Part1Example(t *testing.T) {
	const want = "ABCDEF"
	got := Day19(linesFromFilename(t, exampleFilename(19)), true)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay19Part1(t *testing.T) {
	const want = "VEBTPXCHLI"
	got := Day19(linesFromFilename(t, filename(19)), true)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay19Part1(b *testing.B) {
	grid := linesFromFilename(b, filename(19))
	for b.Loop() {
		_ = Day19(grid, true)
	}
}

func TestDay19Part2Example(t *testing.T) {
	const want = "38"
	got := Day19(linesFromFilename(t, exampleFilename(19)), false)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay19Part2(t *testing.T) {
	const want = "18702"
	got := Day19(linesFromFilename(t, filename(19)), false)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay19Part2(b *testing.B) {
	grid := linesFromFilename(b, filename(19))
	for b.Loop() {
		_ = Day19(grid, false)
	}
}
