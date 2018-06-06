package adventofcode2017

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestDay05Part1(t *testing.T) {
	maze := []int{0, 3, 0, 1, -3}
	want := 5
	got := Day05Part1(maze)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func TestDay05Part2(t *testing.T) {
	maze := []int{0, 3, 0, 1, -3}
	want := 10
	got := Day05Part2(maze)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func ExampleDay05_Part1() {
	fmt.Println(Day05Part1(maze()))
	// Output: 358309
}

func ExampleDay05_Part2() {
	fmt.Println(Day05Part2(maze()))
	// Output: 28178177
}

func maze() []int {
	f, err := os.Open("testdata/day5.txt")
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
