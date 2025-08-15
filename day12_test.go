package adventofcode2017

import (
	"bytes"
	"testing"
)

func TestDay12Part1(t *testing.T) {
	const want = 130
	buf := file(t, 12)
	got := Day12(bytes.NewReader(buf))
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay12Part1Example(t *testing.T) {
	buf := exampleFile(t, 12)
	want := 6
	got := Day12(bytes.NewReader(buf))
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func BenchmarkDay12Part1(b *testing.B) {
	buf := file(b, 12)
	b.ResetTimer()
	for b.Loop() {
		Day12(bytes.NewReader(buf))
	}
}
