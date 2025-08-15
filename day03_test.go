package adventofcode2017

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

const day3Input = 368078

// Nested anonymous structs cannot be initialized in a nice way
type DistanceTestdata struct {
	_ [2]int // pos1
	_ [2]int // pos2
	_ int    // dist
}

func TestDay03(t *testing.T) {
	want := 371
	got := Day03Part1(day3Input)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay03Examples(t *testing.T) {
	// values taken from spec
	testdata := [][]int{
		{1, 0},
		{12, 3},
		{23, 2},
		{1024, 31},
	}
	for _, pair := range testdata {
		square := pair[0]
		want := pair[1]
		got := Day03Part1(square)
		if want != got {
			t.Fatalf("square %d: want %d but got %d\n",
				square, want, got)
		}
	}
}

func TestDay03A174344(t *testing.T) {
	f, err := os.Open("testdata/b174344.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		if strings.HasPrefix(line, "#") {
			continue
		}

		// split into index and result
		parts := strings.Fields(line)
		idx, err := strconv.Atoi(parts[0])
		if err != nil {
			t.Fatal(err)
		}
		want, err := strconv.Atoi(parts[1])
		if err != nil {
			t.Fatal(err)
		}
		got := transform(idx, math.Sin, false)
		if want != got {
			t.Fatalf("n=%d: want %d but got %d\n", idx, want, got)
		}
	}
	if err := sc.Err(); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkDay03Part1(b *testing.B) {
	for b.Loop() {
		Day03Part1(day3Input)
	}
}

func TestDay03Part2(t *testing.T) {
	const want = 369601
	got := Day03Part2(day3Input)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}
func BenchmarkDay03Part2(b *testing.B) {
	for b.Loop() {
		Day03Part2(day3Input)
	}
}
