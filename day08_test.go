package adventofcode2017

import (
	"testing"
)

func TestDay08Part1Example(t *testing.T) {
	const want = 1 // "[…] the largest value in any register is 1."
	buf := exampleFile(t, 8)
	got, err := Day08(buf)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay08Part1(t *testing.T) {
    const want = 4416
    buf := file(t, 8)
    got, err := Day08(buf)
    if err != nil {
        t.Fatal(err)
    }
    if want != got {
        t.Fatalf("want %v but got %v", want, got)
    }
}

func BenchmarkDay08Part1(b *testing.B) {
    buf := file(b, 8)
    b.ResetTimer()
    for b.Loop() {
        _, _ = Day08(buf)
    }
}
