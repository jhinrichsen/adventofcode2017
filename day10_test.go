package adventofcode2017

import (
	"testing"
)

func TestDay10Part1Example(t *testing.T) {
	const want = 12
	got := Day10Part1(exampleFile(t, 10), 5)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay10Part1(t *testing.T) {
	const want = 52070
	got := Day10Part1(file(t, 10), 256)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay10Part1(b *testing.B) {
	input := file(b, 10)
	for b.Loop() {
		_ = Day10Part1(input, 256)
	}
}

func TestDay10Part2Examples(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{in: "", want: "a2582a3a0e66e6e86e3812dcb672a272"},
		{in: "AoC 2017", want: "33efeb34ea91902bb2f59c9920caa6cd"},
		{in: "1,2,3", want: "3efbe78a8d82f29979031a4aa0b16a9d"},
		{in: "1,2,4", want: "63960835bcdc130f0b66d7ff4f6a5a8e"},
	}
	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			got := Day10Part2([]byte(tc.in))
			if tc.want != got {
				t.Fatalf("want %v but got %v", tc.want, got)
			}
		})
	}
}

func TestDay10Part2(t *testing.T) {
	const want = "7f94112db4e32e19cf6502073c66f9bb"
	got := Day10Part2(file(t, 10))
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay10Part2(b *testing.B) {
	input := file(b, 10)
	for b.Loop() {
		_ = Day10Part2(input)
	}
}
