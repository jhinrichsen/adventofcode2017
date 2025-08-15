// Package adventofcode2017 implements some of the puzzles of
// https://adventofcode.com.
package adventofcode2017

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestDay10Part1(t *testing.T) {
	const want = 2
	l := Day10Input()
	lengths, _ := from(filename(10))
	got := Day10(l, lengths)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay10Part1(b *testing.B) {
	// Read file once (no disk I/O in loop)
	buf := file(b, 10)
	for b.Loop() {
		// Parse lengths from in-memory buffer each iteration (include parsing cost)
		var lengths []int
		parts := strings.Split(strings.TrimSpace(string(buf)), ",")
		for i := 0; i < len(parts); i++ {
			n, _ := strconv.Atoi(parts[i])
			lengths = append(lengths, n)
		}
		// Construct list each iteration (as Day10 mutates list)
		list := Day10Input()
		_ = Day10(list, lengths)
	}
}

func TestDay10Hash(t *testing.T) {
	/* TODO
	want := [5]int{3, 4, 2, 1, 0}
	got := [5]int{0, 1, 2, 3, 4}
	lengths := []int{3, 4, 1, 5}
	list := got[:]
	day10Hash(list, lengths)
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
	*/
}

func TestDay10Part1ExampleInline(t *testing.T) {
	/* TODO
	want := 12
	list := []int{0, 1, 2, 3, 4}
	lengths := []int{3, 4, 1, 5}
	got := Day10(list, lengths)
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
	*/
}

func TestDay10BadPreconditionLength(t *testing.T) {
	list := []int{0, 1, 2, 3, 4}
	lengths := []int{3, 4, 1, 5}
	bads, err := Day10PreconditionLength(list, lengths)
	// no error
	if err != nil {
		t.Fatal(err)
	}
	// no bads
	want := 0
	got := len(bads)
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func TestDay10GoodPreconditionLength(t *testing.T) {
	list := []int{0, 1, 2, 3, 4}
	lengths := []int{3, 4, 1, 5}
	bads, err := Day10PreconditionLength(list, lengths)
	// no error
	if err != nil {
		t.Fatal(err)
	}
	// no bads
	want := 0
	got := len(bads)
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func TestDay10ReverseFirst(t *testing.T) {
	// arrays are values
	want := [5]int{2, 1, 0, 3, 4}
	got := [5]int{0, 1, 2, 3, 4}
	Day10Reverse(got[:], 0, 3)
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

func TestDay10Reverse(t *testing.T) {
	// arrays are values
	want := [5]int{4, 3, 0, 1, 2}
	got := [5]int{2, 1, 0, 3, 4}
	Day10Reverse(got[:], 3, 4)
	if want != got {
		t.Fatalf("want %v but got %v\n", want, got)
	}
}

// convert one line of comma separated numeric strings to integers
func from(filename string) ([]int, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var is []int
	ss := strings.Split(strings.TrimSpace(string(buf)), ",")
	for i := 0; i < len(ss); i++ {
		n, err := strconv.Atoi(ss[i])
		if err != nil {
			return nil, fmt.Errorf("cannot convert number to int at column %d", i)
		}
		is = append(is, n)
	}
	return is, nil
}

// list size 256 with values 0 to 255
func Day10Input() []int {
	// look mom - 'is[i] = i' overengineered
	gen := func() func() int {
		var i int
		return func() int {
			i++
			return i
		}
	}()
	is := make([]int, 256)
	for i := 0; i < len(is); i++ {
		is[i] = gen()
	}
	return is
}
