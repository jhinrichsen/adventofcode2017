package adventofcode2017

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func ExampleDay8() {
	buf, err := ioutil.ReadFile("testdata/day8.txt")
	if err != nil {
		panic(err)
	}
	got, err := Day8(bytes.NewReader(buf))
	if err != nil {
		panic(err)
	}
	fmt.Println(got)
	// Output: 5215
}

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
