package adventofcode2017

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestDay7Sample(t *testing.T) {
	want := "tknk"
	buf, err := ioutil.ReadFile("testdata/day7_sample.txt")
	if err != nil {
		t.Fatal(err)
	}
	got := Day7(bytes.NewReader(buf))
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func TestDay7(t *testing.T) {
	f, err := os.Open("testdata/day7.txt")
	if err != nil {
		t.Fatal(err)
	}
	got := Day7(f)
	log.Printf("day 7: got %v\n", got)
}

func BenchmarkDay7(b *testing.B) {
	f, err := os.Open("testdata/day7.txt")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day7(f)
	}
}
