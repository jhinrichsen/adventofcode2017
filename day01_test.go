package adventofcode2017

import (
	"testing"
)

func TestDay01Part1(t *testing.T) {
	const want = 1029
	in := file(t, 1)
	got := Day01Part1(in)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay01Part2(b *testing.B) {
    in := file(b, 1)
    for b.Loop() {
        Day01Part2(in)
    }
}

func TestDay01Part1Example(t *testing.T) {
	var ts = []struct {
		want uint
		in   []byte
	}{
		{3, []byte{1, 1, 2, 2}},
		{4, []byte{1, 1, 1, 1}},
		{0, []byte{1, 2, 3, 4}},
		{9, []byte{9, 1, 2, 1, 2, 1, 2, 9}},
	}
	for _, tt := range ts {
		got := Day01Part1(tt.in)
		if tt.want != got {
			t.Fatalf("want %d but got %d\n", tt.want, got)
		}
	}
}

func BenchmarkDay01Part1(b *testing.B) {
	in := file(b, 1)
	for b.Loop() {
		Day01Part1(in)
	}
}

// TestDay01Part2Example tests day 1, part 2.
func TestDay01Part2Example(t *testing.T) {
	var ts = []struct {
		want uint
		in   []byte
	}{
		{6, []byte{1, 2, 1, 2}},
		{0, []byte{1, 2, 2, 1}},
		{4, []byte{1, 2, 3, 4, 2, 5}},
		{12, []byte{1, 2, 3, 1, 2, 3}},
		{4, []byte{1, 2, 1, 3, 1, 4, 1, 5}},
	}
	for _, tt := range ts {
		got := Day01Part2(tt.in)
		if tt.want != got {
			t.Fatalf("want %d but got %d\n", tt.want, got)
		}
	}
}

func TestDay01Part2(t *testing.T) {
	const want = 1220
	in := file(t, 1)
	got := Day01Part2(in)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
