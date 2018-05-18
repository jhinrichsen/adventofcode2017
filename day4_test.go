package adventofcode2017

import "testing"

func TestDay4(t *testing.T) {
	passphrases := []string{"aa bb cc dd ee",
		"aa bb cc dd aa",
		"aa bb cc dd aaa"}
	want := 2
	got := Day4(passphrases)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}
