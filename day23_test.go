package adventofcode2017

import (
	"testing"
)

func TestDay23Part1(t *testing.T) {
	const want = 9409
	lines := linesFromFilename(t, filename(23))
	prog, err := NewDay23(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day23(prog, true)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay23Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(23))
	for b.Loop() {
		prog, _ := NewDay23(lines)
		_, _ = Day23(prog, true)
	}
}

func TestDay23Part2(t *testing.T) {
	const want = 913
	lines := linesFromFilename(t, filename(23))
	prog, err := NewDay23(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day23(prog, false)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay23Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(23))
	for b.Loop() {
		prog, _ := NewDay23(lines)
		_, _ = Day23(prog, false)
	}
}
