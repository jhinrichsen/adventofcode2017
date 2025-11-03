package adventofcode2017

import (
	"testing"
)

func TestDay25Part1Example(t *testing.T) {
	const want = 3
	lines := linesFromFilename(t, exampleFilename(25))
	puzzle, err := NewDay25(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day25(puzzle, true)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay25Part1(t *testing.T) {
	const want = 2794
	lines := linesFromFilename(t, filename(25))
	puzzle, err := NewDay25(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day25(puzzle, true)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay25Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(25))
	for b.Loop() {
		puzzle, _ := NewDay25(lines)
		_, _ = Day25(puzzle, true)
	}
}
