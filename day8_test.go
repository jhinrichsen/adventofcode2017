package adventofcode2017

import (
	"bytes"
	"io/ioutil"
	"log"
	"testing"
)

func TestDay8Sample(t *testing.T) {
	buf, err := ioutil.ReadFile("testdata/day8_sample.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := 1
	got, err := Day8(bytes.NewReader(buf))
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func TestDay8(t *testing.T) {
	buf, err := ioutil.ReadFile("testdata/day8.txt")
	if err != nil {
		t.Fatal(err)
	}
	got, err := Day8(bytes.NewReader(buf))
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("day 8: got %v\n", got)
}

func BenchmarkDay8(b *testing.B) {
	buf, err := ioutil.ReadFile("testdata/day8.txt")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Day8(bytes.NewReader(buf))
		if err != nil {
			b.Fatal(err)
		}
	}
}
