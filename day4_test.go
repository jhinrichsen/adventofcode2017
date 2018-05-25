package adventofcode2017

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func TestDay4Sample(t *testing.T) {
	passphrases := []string{"aa bb cc dd ee",
		"aa bb cc dd aa",
		"aa bb cc dd aaa"}
	want := 2
	got := Day4(passphrases)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func TestDay4(t *testing.T) {
	var ss []string
	f, err := os.Open("testdata/day4.txt")
	if err != nil {
		t.Fatal(err)
	}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		ss = append(ss, sc.Text())
	}
	got := Day4(ss)
	log.Printf("day 4: got %d\n", got)
}
