package adventofcode2017

import (
	"testing"
)

func TestDay09Part1(t *testing.T) {
    const want = 16689
    buf := file(t, 9)
    got := Day09(buf)
    if want != got {
        t.Fatalf("want %v but got %v", want, got)
    }
}

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
        i := i
        t.Run("case_"+string(rune('A'+i)), func(t *testing.T) {
            want := es[i].score
            got := Day09(es[i].stream)
            if want != got {
                t.Fatalf("want %v but got %v", want, got)
            }
        })
    }
}

func BenchmarkDay09Part1(b *testing.B) {
    buf := file(b, 9)
    b.ResetTimer()
    for b.Loop() {
        Day09(buf)
    }
}
