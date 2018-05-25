package adventofcode2017

import (
	"io/ioutil"
	"testing"
)

func TestDay1Sample(t *testing.T) {
	var ts = []struct {
		want int
		in   []byte
	}{
		{3, []byte{1, 1, 2, 2}},
		{4, []byte{1, 1, 1, 1}},
		{0, []byte{1, 2, 3, 4}},
		{9, []byte{9, 1, 2, 1, 2, 1, 2, 9}},
	}
	for _, tt := range ts {
		got := Day1(tt.in)
		if tt.want != got {
			t.Fatalf("want %d but got %d\n", tt.want, got)
		}
	}
}

func TestDay1(t *testing.T) {
	_, err := ioutil.ReadFile("testdata/day1.txt")
	if err != nil {
		t.Fatal(err)
	}
	// TODO got := Day1(buf)
}
