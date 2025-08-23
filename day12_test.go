package adventofcode2017

import (
	"bytes"
	"testing"
)

func TestDay12Part1Example(t *testing.T) {
	const want = 6
	buf := exampleFile(t, 12)
	got := Day12(bytes.NewReader(buf))
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func TestDay12Part1(t *testing.T) {
	const want = 130
	buf := file(t, 12)
	got := Day12(bytes.NewReader(buf))
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay12Part1(b *testing.B) {
	buf := file(b, 12)
	for b.Loop() {
		Day12(bytes.NewReader(buf))
	}
}

func TestDay12Part2Example(t *testing.T) {
	const want = 2
	buf := exampleFile(t, 12)
	got := Day12Part2(bytes.NewReader(buf))
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay12Part2(t *testing.T) {
	const want = 189
	buf := file(t, 12)
	got := Day12Part2(bytes.NewReader(buf))
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay12Part2(b *testing.B) {
	buf := file(b, 12)
	for b.Loop() {
		Day12Part2(bytes.NewReader(buf))
	}
}
