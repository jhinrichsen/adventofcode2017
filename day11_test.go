package adventofcode2017

import (
	"testing"
)

func TestDay11Part1Examples(t *testing.T) {
	tests := []struct {
		in   string
		want uint
	}{
		{in: "ne,ne,ne", want: 3},
		{in: "ne,ne,sw,sw", want: 0},
		{in: "ne,ne,s,s", want: 2},
		{in: "se,sw,se,sw,sw", want: 3},
	}
	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			p, err := NewDay11([]byte(tc.in))
			if err != nil {
				t.Fatal(err)
			}
			got := Day11Part1(p)
			if tc.want != got {
				t.Fatalf("want %v but got %v", tc.want, got)
			}
		})
	}
}

func TestDay11Part1(t *testing.T) {
	const want = 818
	buf := file(t, 11)
	p, err := NewDay11(buf)
	if err != nil {
		t.Fatal(err)
	}
	got := Day11Part1(p)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay11Part1(b *testing.B) {
	buf := file(b, 11)
	p, err := NewDay11(buf)
	if err != nil {
		b.Fatal(err)
	}
	for b.Loop() {
		_ = Day11Part1(p)
	}
}

func TestDay11Part2(t *testing.T) {
	const want = 1596
	buf := file(t, 11)
	p, err := NewDay11(buf)
	if err != nil {
		t.Fatal(err)
	}
	got := Day11Part2(p)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay11Part2(b *testing.B) {
	buf := file(b, 11)
	p, err := NewDay11(buf)
	if err != nil {
		b.Fatal(err)
	}
	for b.Loop() {
		_ = Day11Part2(p)
	}
}
