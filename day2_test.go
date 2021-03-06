package adventofcode2017

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func ExampleDay02() {
	in, _ := inputDay02("testdata/day2.txt")
	fmt.Println(Day02(in))
	// Output: 47623
}

func TestDay02Sample(t *testing.T) {
	spreadsheet := [][]int{
		{5, 1, 9, 5},
		{7, 5, 3, 3}, // trailing 3 is a context aware NIL/ empty cell
		{2, 4, 6, 8},
	}
	want := 18
	got := Day02(spreadsheet)
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
	// dumpDay02(iss)
	return iss, nil
}

func dumpDay02(iss [][]int) {
	for row := range iss {
		for col := range iss[row] {
			fmt.Printf("%4d ", iss[row][col])
		}
		fmt.Println()
	}
}

func TestDay02SampleInput(t *testing.T) {
	in, err := inputDay02("testdata/day2_sample.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := 18
	got := Day02(in)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func BenchmarkDay02(b *testing.B) {
	in, err := inputDay02("testdata/day2.txt")
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		Day02(in)
	}
}
