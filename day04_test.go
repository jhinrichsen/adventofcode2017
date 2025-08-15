package adventofcode2017

import (
	"testing"
)

func TestDay04Part1(t *testing.T) {
	const want = 451
	ss := linesFromFilename(t, filename(4))
	got := Day04Part1(ss)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay04Part1Example(t *testing.T) {
	const want = 2
	passphrases := []string{
		"aa bb cc dd ee",
		"aa bb cc dd aa",
		"aa bb cc dd aaa",
	}
	got := Day04Part1(passphrases)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay04Part2Example(t *testing.T) {
	var tests = []struct {
		passphrase string
		valid      bool
	}{
		{"abcde fghij", true},
		{"abcde xyz ecdab", false},
		{"a ab abc abd abf abj", true},
		{"iiii oiii ooii oooi oooo", true},
		{"oiii ioii iioi iiio", false},
	}
	for _, tt := range tests {
		t.Run(tt.passphrase, func(t *testing.T) {
			want := tt.valid
			got := Day04Part2Valid(tt.passphrase)
			if want != got {
				t.Fatalf("want %v but got %v", want, got)
			}
		})
	}
}

func TestDay04Part2(t *testing.T) {
	const want = 223
	ss := linesFromFilename(t, filename(4))
	got := Day04Part2(ss)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func BenchmarkDay04Part1(b *testing.B) {
	ss := linesFromFilename(b, filename(4))
	for b.Loop() {
		Day04Part1(ss)
	}
}

func BenchmarkDay04Part2(b *testing.B) {
	ss := linesFromFilename(b, filename(4))
	for b.Loop() {
		Day04Part2(ss)
	}
}
