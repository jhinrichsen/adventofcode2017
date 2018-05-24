package adventofcode2017

import "testing"

func TestDay6Sample(t *testing.T) {
	banks := []int{0, 2, 7, 0}
	want := 5
	got := Day6(banks)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func TestDay6(t *testing.T) {
}
