package adventofcode2017

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func ExampleDay07() {
	f, err := os.Open("testdata/day7.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(Day07(f))
	// Output: eugwuhl
}

func TestDay07Sample(t *testing.T) {
	want := "tknk"
	buf, err := ioutil.ReadFile("testdata/day7_sample.txt")
	if err != nil {
		t.Fatal(err)
	}
	got := Day07(bytes.NewReader(buf))
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func BenchmarkDay07(b *testing.B) {
	f, err := os.Open("testdata/day7.txt")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day07(f)
	}
}
