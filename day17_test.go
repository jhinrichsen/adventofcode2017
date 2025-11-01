package adventofcode2017

import (
	"testing"
)

func TestDay17Part1Example(t *testing.T) {
	const want = 638
	got := Day17(3, true)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay17Part1(t *testing.T) {
	const want = 1670
	got := Day17(328, true)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay17Part1(b *testing.B) {
	for b.Loop() {
		_ = Day17(328, true)
	}
}

func TestDay17Part2(t *testing.T) {
	const want = 2316253
	got := Day17(328, false)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay17Part2(b *testing.B) {
	for b.Loop() {
		_ = Day17(328, false)
	}
}
