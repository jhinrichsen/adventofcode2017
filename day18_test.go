package adventofcode2017

import (
	"testing"
)

func TestDay18Part1Example(t *testing.T) {
	const want = 4
	got := Day18(NewDay18(linesFromFilename(t, exampleFilename(18))), true)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay18Part1(t *testing.T) {
	const want = 1187
	got := Day18(NewDay18(linesFromFilename(t, filename(18))), true)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay18Part1(b *testing.B) {
	p := NewDay18(linesFromFilename(b, filename(18)))
	for b.Loop() {
		_ = Day18(p, true)
	}
}

func TestDay18Part2Example(t *testing.T) {
	const want = 3
	got := Day18(NewDay18(linesFromFilename(t, "testdata/day18_example2.txt")), false)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay18Part2(t *testing.T) {
	const want = 5969
	got := Day18(NewDay18(linesFromFilename(t, filename(18))), false)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay18Part2(b *testing.B) {
	p := NewDay18(linesFromFilename(b, filename(18)))
	for b.Loop() {
		_ = Day18(p, false)
	}
}
