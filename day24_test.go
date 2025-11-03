package adventofcode2017

import (
	"testing"
)

func TestDay24Part1Example(t *testing.T) {
	const want = 31
	lines := linesFromFilename(t, exampleFilename(24))
	components, err := NewDay24(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day24(components, true)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay24Part1(t *testing.T) {
	const want = 1906
	lines := linesFromFilename(t, filename(24))
	components, err := NewDay24(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day24(components, true)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay24Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(24))
	for b.Loop() {
		components, _ := NewDay24(lines)
		_, _ = Day24(components, true)
	}
}

func TestDay24Part2Example(t *testing.T) {
	const want = 19
	lines := linesFromFilename(t, exampleFilename(24))
	components, err := NewDay24(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day24(components, false)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay24Part2(t *testing.T) {
	const want = 1824
	lines := linesFromFilename(t, filename(24))
	components, err := NewDay24(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day24(components, false)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay24Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(24))
	for b.Loop() {
		components, _ := NewDay24(lines)
		_, _ = Day24(components, false)
	}
}
