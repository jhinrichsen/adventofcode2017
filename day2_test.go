package adventofcode2017

import "testing"

func TestDay2(t *testing.T) {
	spreadsheet := [][]int{
		{5, 1, 9, 5},
		{7, 5, 3, 3}, // trailing 3 is a context aware NIL/ empty cell
		{2, 4, 6, 8},
	}
	want := 18
	got := Day2(spreadsheet)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}
