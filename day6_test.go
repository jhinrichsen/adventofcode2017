package adventofcode2017

import "testing"

func TestDay6Sample(t *testing.T) {
	banks := Banks{0, 2, 7, 0}
	want := 5
	got := Day6Impl1(banks, 4)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func TestArrayAsKeyInMap(t *testing.T) {
	var m = make(map[[3]int]bool)

	m[[3]int{1, 2, 3}] = true
	m[[3]int{2, 3, 4}] = true
	m[[3]int{2, 3, 4}] = true

	want := 2
	got := len(m)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func BenchmarkDay6Impl1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Day6Impl1(Banks{0, 2, 7, 0}, 4)
	}
}
