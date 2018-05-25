package adventofcode2017

import (
	"bufio"
	"log"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"testing"
)

func TestDay5Part1(t *testing.T) {
	maze := []int{0, 3, 0, 1, -3}
	want := 5
	got := Day5Part1(maze)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func TestDay5Part2(t *testing.T) {
	maze := []int{0, 3, 0, 1, -3}
	want := 10
	got := Day5Part2(maze)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func TestMyDay5Part1(t *testing.T) {
	test(t, Day5Part1)
}

func TestMyDay5Part2(t *testing.T) {
	test(t, Day5Part2)
}

func test(t *testing.T, part func([]int) int) {
	f, err := os.Open("testdata/day5.txt")
	if err != nil {
		t.Fatal(err)
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
			t.Fatal(err)
		}
		maze = append(maze, i)
	}
	log.Printf("%v: got %d\n",
		runtime.FuncForPC(reflect.ValueOf(part).Pointer()).Name(),
		part(maze))
}
