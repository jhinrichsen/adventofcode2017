package adventofcode2017

import (
	"testing"
)

func TestDay22Part1Examples(t *testing.T) {
	tests := []struct {
		bursts uint
		want   uint
	}{
		{70, 41},
		{10000, 5587},
	}

	for _, tt := range tests {
		lines := linesFromFilename(t, exampleFilename(22))
		vc, err := NewDay22(lines)
		if err != nil {
			t.Fatal(err)
		}
		got, err := Day22(vc, tt.bursts, true)
		if err != nil {
			t.Fatal(err)
		}
		if tt.want != got {
			t.Fatalf("bursts=%d: want %d but got %d", tt.bursts, tt.want, got)
		}
	}
}

func TestDay22Part1(t *testing.T) {
	const want = 5369
	lines := linesFromFilename(t, filename(22))
	vc, err := NewDay22(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day22(vc, 10000, true)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay22Part1(b *testing.B) {
	lines := linesFromFilename(b, filename(22))
	for b.Loop() {
		vc, _ := NewDay22(lines)
		_, _ = Day22(vc, 10000, true)
	}
}

func TestDay22Part2Examples(t *testing.T) {
	tests := []struct {
		bursts uint
		want   uint
	}{
		{100, 26},
		{10000000, 2511944},
	}

	for _, tt := range tests {
		lines := linesFromFilename(t, exampleFilename(22))
		vc, err := NewDay22(lines)
		if err != nil {
			t.Fatal(err)
		}
		got, err := Day22(vc, tt.bursts, false)
		if err != nil {
			t.Fatal(err)
		}
		if tt.want != got {
			t.Fatalf("bursts=%d: want %d but got %d", tt.bursts, tt.want, got)
		}
	}
}

func TestDay22Part2(t *testing.T) {
	const want = 2510774
	lines := linesFromFilename(t, filename(22))
	vc, err := NewDay22(lines)
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day22(vc, 10000000, false)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay22Part2(b *testing.B) {
	lines := linesFromFilename(b, filename(22))
	for b.Loop() {
		vc, _ := NewDay22(lines)
		_, _ = Day22(vc, 10000000, false)
	}
}
