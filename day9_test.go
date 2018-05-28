package adventofcode2017

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestDay9Samples(t *testing.T) {
	var es = []struct {
		stream []byte
		score  int
	}{
		{[]byte("<>,"), 0},
		{[]byte("{},"), 1},
		{[]byte("{{{}}},"), 6},
		{[]byte("{{},{}},"), 5},
		{[]byte("{{{},{},{{}}}},"), 16},
		{[]byte("{<a>,<a>,<a>,<a>},"), 1},
		{[]byte("{{<ab>},{<ab>},{<ab>},{<ab>}},"), 9},
		{[]byte("{{<!!>},{<!!>},{<!!>},{<!!>}},"), 9},
		{[]byte("{{<a!>},{<a!>},{<a!>},{<ab>}},"), 3},
	}

	for i := range es {
		want := es[i].score
		got := Day9(es[i].stream)
		if want != got {
			t.Fatalf("test set #%d: want %v but got %v\n", i, want, got)
		}
	}
}

func TestDay9(t *testing.T) {
	buf, err := ioutil.ReadFile("testdata/day9.txt")
	if err != nil {
		t.Fatal(err)
	}
	got := Day9(buf)
	log.Printf("day 9: got %d\n", got)
}

func BenchmarkDay9(b *testing.B) {
	buf, err := ioutil.ReadFile("testdata/day9.txt")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day9(buf)
	}
}
