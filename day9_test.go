package adventofcode2017

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestDay9Samples(t *testing.T) {
	var es = []struct {
		stream string
		score  int
	}{
		{"<>,", 0},
		{"{},", 1},
		{"{{{}}},", 6},
		{"{{},{}},", 5},
		{"{{{},{},{{}}}},", 16},
		{"{<a>,<a>,<a>,<a>},", 1},
		{"{{<ab>},{<ab>},{<ab>},{<ab>}},", 9},
		{"{{<!!>},{<!!>},{<!!>},{<!!>}},", 9},
		{"{{<a!>},{<a!>},{<a!>},{<ab>}},", 3},
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
	got := Day9(string(buf))
	log.Printf("day 9: got %d\n", got)
}

func BenchmarkDay9(b *testing.B) {
	buf, err := ioutil.ReadFile("testdata/day9.txt")
	if err != nil {
		b.Fatal(err)
	}
	s := string(buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day9(s)
	}
}
