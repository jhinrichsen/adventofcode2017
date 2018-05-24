package adventofcode2017

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"testing"
)

func TestDay6Sample(t *testing.T) {
	banks := Banks{0, 2, 7, 0}
	want := 5
	got := Day6Impl1(banks, 4)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func TestDay6Input(t *testing.T) {
	const nBanks = 16
	buf, err := ioutil.ReadFile("testdata/day6.txt")
	if err != nil {
		t.Fatal(err)
	}
	fs := strings.Fields(string(buf))
	if len(fs) != nBanks {
		t.Fatalf("want %d fields but got %d\n", nBanks, len(fs))
	}
	var banks Banks
	for i := 0; i < nBanks; i++ {
		banks[i], err = strconv.Atoi(fs[i])
		if err != nil {
			t.Fatalf("error converting col %d: %v\n", i, err)
		}
	}
	n := Day6Impl1(banks, nBanks)
	log.Printf("day6: iterations=%d\n", n)
}

func TestArrayAsKeyInMap(t *testing.T) {
	var m = make(map[[3]int]bool)

	m[[3]int{1, 2, 3}] = true
	m[[3]int{2, 3, 4}] = true
	m[[3]int{2, 3, 4}] = true

	want := 2
	got := len(m)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func BenchmarkDay6Impl1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Day6Impl1(Banks{0, 2, 7, 0}, 4)
	}
}

func BenchmarkDay6Impl2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Day6Impl2([]int{0, 2, 7, 0})
	}
}
