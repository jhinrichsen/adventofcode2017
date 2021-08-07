package adventofcode2017

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestDay02Part1(t *testing.T) {
	const want = 44216
	in, err := inputDay02(filename(2))
	if err != nil {
		t.Fatal(err)
	}
	got := Day2Part1(in)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func TestDay02Sample(t *testing.T) {
	const want = 18
	spreadsheet := [][]int{
		{5, 1, 9, 5},
		{7, 5, 3, 3}, // trailing 3 is a context aware NIL/ empty cell
		{2, 4, 6, 8},
	}
	got := Day2Part1(spreadsheet)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func inputDay02(filename string) ([][]int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var iss [][]int
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		line = strings.TrimSpace(line)
		fs := strings.Fields(line)
		var is []int
		for i := range fs {
			n, err := strconv.Atoi(fs[i])
			if err != nil {
				return iss, fmt.Errorf("cannot convert column %d to int: %+v", i, fs[i])
			}
			is = append(is, n)
		}
		iss = append(iss, is)
	}
	return iss, nil
}

func TestDay02SampleInput(t *testing.T) {
	const want = 18
	in, err := inputDay02(exampleFilename(2))
	if err != nil {
		t.Fatal(err)
	}
	got := Day2Part1(in)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func BenchmarkDay02(b *testing.B) {
	in, err := inputDay02("testdata/day2.txt")
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day2Part1(in)
	}
}

func TestDay2Part2Example(t *testing.T) {
	const want = 9
	var spreadSheet = [][]int{
		{5, 9, 2, 8},
		{9, 4, 7, 3},
		{3, 8, 6, 5},
	}
	got := Day2Part2(spreadSheet)
	if want != got {
		t.Fatalf("want %d but got %d", want, got)
	}
}

func TestDay2Part2(t *testing.T) {
	const want = 320
	in, err := inputDay02(filename(2))
	if err != nil {
		t.Fatal(err)
	}
	got := Day2Part2(in)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}
