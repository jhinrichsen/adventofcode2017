package adventofcode2017

import (
	"testing"
)

func TestDay13Part1ExampleUsingArray(t *testing.T) {
	const want = 24
	layers, err := NewDay13Array(linesFromFilename(t, exampleFilename(13)))
	if err != nil {
		t.Fatal(err)
	}
	if len(layers) != 7 {
		t.Fatalf("want max depth = 7 but got %d", len(layers))
	}
	got := Day13UsingArray(layers)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay13Part1UsingArray(t *testing.T) {
	const want = 1704
	layers, err := NewDay13Array(linesFromFilename(t, filename(13)))
	if err != nil {
		t.Fatal(err)
	}
	got := Day13UsingArray(layers)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay13Part1Array(b *testing.B) {
	for b.Loop() {
		layers, _ := NewDay13Array(linesFromFilename(b, filename(13)))
		_ = Day13UsingArray(layers)
	}
}

func TestDay13ListExample(t *testing.T) {
	const want = 24
	l, err := NewDay13(linesFromFilename(t, exampleFilename(13)))
	if err != nil {
		t.Fatal(err)
	}
	got := Day13(l, true)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay13Part1(t *testing.T) {
	const want = 1704
	l, err := NewDay13(linesFromFilename(t, filename(13)))
	if err != nil {
		t.Fatal(err)
	}
	got := Day13(l, true)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay13Part1List(b *testing.B) {
	for b.Loop() {
		l, _ := NewDay13(linesFromFilename(b, exampleFilename(13)))
		_ = Day13(l, true)
	}
}

func TestDay13Part2Example(t *testing.T) {
	const want = 10
	l, err := NewDay13(linesFromFilename(t, exampleFilename(13)))
	if err != nil {
		t.Fatal(err)
	}
	got := Day13(l, false)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay13Part2(t *testing.T) {
	const want = 3970918
	l, err := NewDay13(linesFromFilename(t, filename(13)))
	if err != nil {
		t.Fatal(err)
	}
	got := Day13(l, false)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay13Part2(b *testing.B) {
	for b.Loop() {
		l, _ := NewDay13(linesFromFilename(b, filename(13)))
		_ = Day13(l, false)
	}
}
