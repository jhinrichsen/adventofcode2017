package adventofcode2017

import (
	"strconv"
	"strings"
	"testing"
)

func TestDay06Part1(t *testing.T) {
    const want = 4074
    const nBanks = 16
    ss := linesFromFilename(t, filename(6))
    fs := strings.Fields(ss[0])
    if len(fs) != nBanks {
        t.Fatalf("want %d fields but got %d", nBanks, len(fs))
    }
    var banks Banks
    for i := 0; i < nBanks; i++ {
        n, err := strconv.Atoi(fs[i])
        if err != nil {
            t.Fatalf("error converting col %d: %v", i, err)
        }
        banks[i] = n
    }
    got := Day06Part1(banks, nBanks)
    if want != got {
        t.Fatalf("want %d but got %d", want, got)
    }
}

func TestDay06Part1Example(t *testing.T) {
	const want = 5
	banks := Banks{0, 2, 7, 0}
	got := Day06Part1(banks, 4)
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

func BenchmarkDay06Part1(b *testing.B) {
    for b.Loop() {
        _ = Day06Part1(Banks{0, 2, 7, 0}, 4)
    }
}

func TestDay06Part2Example(t *testing.T) {
	const want = 4
	banks := Banks{0, 2, 7, 0}
	got := Day06Part2(banks, 4)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func TestDay06Part2(t *testing.T) {
    const want = 2793
    const nBanks = 16
    ss := linesFromFilename(t, filename(6))
    fs := strings.Fields(ss[0])
    if len(fs) != nBanks {
        t.Fatalf("want %d fields but got %d", nBanks, len(fs))
    }
    var banks Banks
    for i := 0; i < nBanks; i++ {
        n, err := strconv.Atoi(fs[i])
        if err != nil {
            t.Fatalf("error converting col %d: %v", i, err)
        }
        banks[i] = n
    }
    got := Day06Part2(banks, nBanks)
    if want != got {
        t.Fatalf("want %d but got %d", want, got)
    }
}

func BenchmarkDay06Part2(b *testing.B) {
    ss := linesFromFilename(b, filename(6))
    fs := strings.Fields(ss[0])
    const nBanks = 16
    b.ResetTimer()
    for b.Loop() {
        var banks Banks
        for i := 0; i < nBanks; i++ {
            n, _ := strconv.Atoi(fs[i])
            banks[i] = n
        }
        _ = Day06Part2(banks, nBanks)
    }
}
