package adventofcode2017

import (
	"testing"
)

func TestDay09Part1Examples(t *testing.T) {
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
		t.Run("case_"+string(rune('A'+i)), func(t *testing.T) {
			want := es[i].score
			got := Day09Part1(es[i].stream)
			if want != got {
				t.Fatalf("want %v but got %v", want, got)
			}
		})
	}
}

func TestDay09Part1(t *testing.T) {
	const want = 16689
	buf := file(t, 9)
	got := Day09Part1(buf)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay09Part1(b *testing.B) {
	buf := file(b, 9)
	for b.Loop() {
		Day09Part1(buf)
	}
}

func TestDay09Part2Examples(t *testing.T) {
	var es = []struct {
		stream []byte
		want   int
	}{
		{[]byte("<>"), 0},
		{[]byte("<random characters>"), 17},
		{[]byte("<<<<>"), 3},
		{[]byte("<{!>}>"), 2},
		{[]byte("<!!>"), 0},
		{[]byte("<!!!>>"), 0},
		{[]byte("<{o\"i!a,<{i<a>"), 10},
	}
	for i := range es {
		t.Run("ex_"+string(rune('A'+i)), func(t *testing.T) {
			got := Day09Part2(es[i].stream)
			if es[i].want != got {
				t.Fatalf("want %v but got %v", es[i].want, got)
			}
		})
	}
}

func TestDay09Part2(t *testing.T) {
	const want = 7982
	buf := file(t, 9)
	got := Day09Part2(buf)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay09Part2(b *testing.B) {
	buf := file(b, 9)
	for b.Loop() {
		Day09Part2(buf)
	}
}
