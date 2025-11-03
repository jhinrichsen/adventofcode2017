package adventofcode2017

import (
	"testing"
)

func TestDay20Part1Example(t *testing.T) {
	const want = 0
	lines := linesFromFilename(t, exampleFilename(20))
	p, err := NewDay20(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day20(p, true)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay20Part1(t *testing.T) {
	const want = 91
	lines := linesFromFilename(t, filename(20))
	p, err := NewDay20(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day20(p, true)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay20Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(20))
	for b.Loop() {
		p, err := NewDay20(lines)
		if err != nil {
			b.Fatal(err)
		}
		_, err = Day20(p, true)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func TestDay20Part2Example(t *testing.T) {
	const want = 1
	lines := linesFromFilename(t, "testdata/day20_example2.txt")
	p, err := NewDay20(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day20(p, false)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay20Part2(t *testing.T) {
	const want = 567
	lines := linesFromFilename(t, filename(20))
	p, err := NewDay20(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day20(p, false)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay20Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(20))
	for b.Loop() {
		p, err := NewDay20(lines)
		if err != nil {
			b.Fatal(err)
		}
		_, err = Day20(p, false)
		if err != nil {
			b.Fatal(err)
		}
	}
}
