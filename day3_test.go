package adventofcode2017

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"testing"
)

// Nested anonymous structs cannot be initialized in a nice way
type DistanceTestdata struct {
	pos1 [2]int
	pos2 [2]int
	dist int
}

func ExampleDay03() {
	fmt.Println(Day03(361527))
	// Output: 326
}

func TestDay03(t *testing.T) {
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
		got := Day03(square)
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
		got := Transform(idx, math.Sin, false)
		if want != got {
			t.Fatalf("n=%d: want %d but got %d\n", idx, want, got)
		}
	}
	if err := sc.Err(); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkDay03(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Day03(361527)
	}
}
