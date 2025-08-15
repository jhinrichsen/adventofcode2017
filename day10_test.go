// Package adventofcode2017 implements some of the puzzles of
// https://adventofcode.com.
package adventofcode2017

import (
	"testing"
)

func TestDay10Part1Example(t *testing.T) {
	const want = 12
	p := NewDay10([]byte("3,4,1,5"))
	p.size = 5
	p.zero = true
	got := Day10(p)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay10Part1(t *testing.T) {
	const want = 52070
	buf := file(t, 10)
	p := NewDay10(buf)
	got := Day10(p)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay10Part1(b *testing.B) {
	buf := file(b, 10)
	p := NewDay10(buf)
	for b.Loop() {
		_ = Day10(p)
	}
}
