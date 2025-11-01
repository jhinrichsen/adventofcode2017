package adventofcode2017

import (
	"testing"
)

func TestDay16Part1Example(t *testing.T) {
	const want = "baedc"
	got := Day16(NewDay16(linesFromFilename(t, exampleFilename(16))), 5, true)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay16Part1(t *testing.T) {
	const want = "fnloekigdmpajchb"
	got := Day16(NewDay16(linesFromFilename(t, filename(16))), 16, true)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay16Part1(b *testing.B) {
	p := NewDay16(linesFromFilename(b, filename(16)))
	for b.Loop() {
		_ = Day16(p, 16, true)
	}
}

func TestDay16Part2Example(t *testing.T) {
	const want = "abcde"
	got := Day16(NewDay16(linesFromFilename(t, exampleFilename(16))), 5, false)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay16Part2(t *testing.T) {
	const want = "amkjepdhifolgncb"
	got := Day16(NewDay16(linesFromFilename(t, filename(16))), 16, false)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay16Part2(b *testing.B) {
	p := NewDay16(linesFromFilename(b, filename(16)))
	for b.Loop() {
		_ = Day16(p, 16, false)
	}
}
