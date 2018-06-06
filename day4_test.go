package adventofcode2017

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func ExampleDay04() {
	var ss []string
	f, _ := os.Open("testdata/day4.txt")
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		ss = append(ss, sc.Text())
	}
	fmt.Println(Day04(ss))
	// Output: 386
}

func TestDay04Sample(t *testing.T) {
	passphrases := []string{"aa bb cc dd ee",
		"aa bb cc dd aa",
		"aa bb cc dd aaa"}
	want := 2
	got := Day04(passphrases)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}
