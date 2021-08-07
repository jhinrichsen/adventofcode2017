package adventofcode2017

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestDay1Part1(t *testing.T) {
	const want = 1029
	in, err := input()
	if err != nil {
		t.Fatal(err)
	}
	got := Day1Part1(in)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay1SamplePart1(t *testing.T) {
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
		got := Day1Part1(tt.in)
		if tt.want != got {
			t.Fatalf("want %d but got %d\n", tt.want, got)
		}
	}
}

func input() ([]byte, error) {
	buf, err := os.ReadFile("testdata/day1.txt")
	if err != nil {
		return nil, err
	}
	buf = bytes.TrimSpace(buf)

	// map from '0'..'9' to 0..9
	for i := range buf {
		if buf[i] < '0' || buf[i] > '9' {
			return buf, fmt.Errorf("index %d out of range: %d", i, buf[i])
		}
		buf[i] = buf[i] - '0'
	}
	return buf, nil
}

func BenchmarkDay1Part1(b *testing.B) {
	in, err := input()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day1Part1(in)
	}
}

// TestDay2SamplePart2 tests day 1, part 2.
func TestDay1SamplePart2(t *testing.T) {
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
		got := Day1Part2(tt.in)
		if tt.want != got {
			t.Fatalf("want %d but got %d\n", tt.want, got)
		}
	}
}

func TestDay1Part2(t *testing.T) {
	const want = 1220
	in, err := input()
	if err != nil {
		t.Fatal(err)
	}
	got := Day1Part2(in)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
