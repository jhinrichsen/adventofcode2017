package adventofcode2017

import "testing"

// Test day 1
func TestDay1(t *testing.T) {
	var ts = []struct {
		want int
		in   []int
	}{
		{3, []int{1, 1, 2, 2}},
		{4, []int{1, 1, 1, 1}},
		{0, []int{1, 2, 3, 4}},
		{9, []int{9, 1, 2, 1, 2, 1, 2, 9}},
	}
	for _, tt := range ts {
		got := Day1(tt.in)
		if tt.want != got {
			t.Fatalf("want %d but got %d\n", tt.want, got)
		}
	}
}
