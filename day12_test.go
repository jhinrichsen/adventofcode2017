package adventofcode2017

import (
	"log"
	"os"
	"testing"
)

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

func TestDay12(t *testing.T) {
	f, err := os.Open("testdata/day12.txt")
	if err != nil {
		t.Fatal(err)
	}
	got := Day12(f)
	log.Printf("day 12: got %d\n", got)
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
