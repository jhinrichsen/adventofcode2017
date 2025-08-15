package adventofcode2017

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestDay05Part1Example(t *testing.T) {
	maze := []int{0, 3, 0, 1, -3}
	want := 5
	got := Day05Part1(maze)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func BenchmarkDay05Part1(b *testing.B) {
    // Read file once (no I/O in loop)
    data := file(b, 5)
    lines := strings.Split(strings.TrimSpace(string(data)), "\n")
    for b.Loop() {
        // Parse in loop (benchmark includes parsing)
        mz := make([]int, 0, len(lines))
        for _, s := range lines {
            if s == "" { continue }
            n, _ := strconv.Atoi(s)
            mz = append(mz, n)
        }
        _ = Day05Part1(mz)
    }
}

func BenchmarkDay05Part2(b *testing.B) {
    data := file(b, 5)
    lines := strings.Split(strings.TrimSpace(string(data)), "\n")
    for b.Loop() {
        mz := make([]int, 0, len(lines))
        for _, s := range lines {
            if s == "" { continue }
            n, _ := strconv.Atoi(s)
            mz = append(mz, n)
        }
        _ = Day05Part2(mz)
    }
}

func TestDay05Part2Example(t *testing.T) {
	maze := []int{0, 3, 0, 1, -3}
	want := 10
	got := Day05Part2(maze)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func TestDay05Part1(t *testing.T) {
	want := 355965
	got := Day05Part1(maze())
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func TestDay05Part2(t *testing.T) {
	want := 26948068
	got := Day05Part2(maze())
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func maze() []int {
	f, err := os.Open("testdata/day05.txt")
	if err != nil {
		panic(err)
	}
	var maze []int
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		s := sc.Text()
		// overread empty lines
		if len(s) == 0 {
			continue
		}
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		maze = append(maze, i)
	}
	return maze
}
