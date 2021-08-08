package adventofcode2017

import (
	"strconv"
	"strings"
	"testing"
)

func TestDay6Part1(t *testing.T) {
	const want = 4074
	const nBanks = 16
	ss, err := linesFromFilename(filename(6))
	if err != nil {
		t.Fatal(err)
	}
	fs := strings.Fields(ss[0])
	if len(fs) != nBanks {
		t.Fatalf("want %d fields but got %d", nBanks, len(fs))
	}
	var banks Banks
	for i := 0; i < nBanks; i++ {
		banks[i], err = strconv.Atoi(fs[i])
		if err != nil {
			t.Fatalf("error converting col %d: %v", i, err)
		}
	}
	got := Day6Part1(banks, nBanks)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay6Part1Example(t *testing.T) {
	const want = 5
	banks := Banks{0, 2, 7, 0}
	got := Day6Part1(banks, 4)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
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

func BenchmarkDay6Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Day6Part1(Banks{0, 2, 7, 0}, 4)
	}
}

func TestDay6Part2Example(t *testing.T) {
	const want = 4
	banks := Banks{0, 2, 7, 0}
	got := Day6Part2(banks, 4)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func TestDay6Part2(t *testing.T) {
	const want = 2793
	const nBanks = 16
	ss, err := linesFromFilename(filename(6))
	if err != nil {
		t.Fatal(err)
	}
	fs := strings.Fields(ss[0])
	if len(fs) != nBanks {
		t.Fatalf("want %d fields but got %d", nBanks, len(fs))
	}
	var banks Banks
	for i := 0; i < nBanks; i++ {
		banks[i], err = strconv.Atoi(fs[i])
		if err != nil {
			t.Fatalf("error converting col %d: %v", i, err)
		}
	}
	got := Day6Part2(banks, nBanks)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
