package adventofcode2017

import (
	"fmt"
	"os"
	"testing"
)

func ExampleDay12() {
	f, _ := os.Open("testdata/day12.txt")
	fmt.Println(Day12(f))
	// Output: 283
}

func TestDay12Sample(t *testing.T) {
	f, err := os.Open("testdata/day12_sample.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := 6
	got := Day12(f)
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func BenchmarkDay12(b *testing.B) {
	f, err := os.Open("testdata/day12.txt")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day12(f)
	}
}
