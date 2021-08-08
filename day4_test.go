package adventofcode2017

import (
	"testing"
)

func TestDay04(t *testing.T) {
	const want = 451
	ss, err := linesFromFilename(filename(4))
	if err != nil {
		t.Fatal(err)
	}
	got := Day4Part1(ss)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay4Part1Examples(t *testing.T) {
	const want = 2
	passphrases := []string{
		"aa bb cc dd ee",
		"aa bb cc dd aa",
		"aa bb cc dd aaa",
	}
	got := Day4Part1(passphrases)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay4Part2Examples(t *testing.T) {
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
			got := Day4Part2Valid(tt.passphrase)
			if want != got {
				t.Fatalf("want %v but got %v", want, got)
			}
		})
	}
}

func TestDay4Part2(t *testing.T) {
	const want = 223
	ss, err := linesFromFilename(filename(4))
	if err != nil {
		t.Fatal(err)
	}
	got := Day4Part2(ss)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
