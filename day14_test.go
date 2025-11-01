package adventofcode2017

import (
	"testing"
)

func TestDay14Part1Example(t *testing.T) {
	const want = 8108
	got := Day14(NewDay14(linesFromFilename(t, exampleFilename(14))), true)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay14Part1(t *testing.T) {
	const want = 8190
	got := Day14(NewDay14(linesFromFilename(t, filename(14))), true)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay14Part2Example(t *testing.T) {
	const want = 1242
	got := Day14(NewDay14(linesFromFilename(t, exampleFilename(14))), false)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay14Part2(t *testing.T) {
	const want = 1134
	got := Day14(NewDay14(linesFromFilename(t, filename(14))), false)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay14Part1(b *testing.B) {
	for b.Loop() {
		_ = Day14(NewDay14(linesFromFilename(b, filename(14))), true)
	}
}

func BenchmarkDay14Part2(b *testing.B) {
	for b.Loop() {
		_ = Day14(NewDay14(linesFromFilename(b, filename(14))), false)
	}
}
