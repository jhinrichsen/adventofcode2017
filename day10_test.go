// Package adventofcode2017 implements some of the puzzles of
// https://adventofcode.com.
package adventofcode2017

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func ExampleDay10() {
	l := Day10Input()
	lengths, _ := from("testdata/day10.txt")
	fmt.Println(Day10(l, lengths))
	// Output: 2
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

func TestDay10Sample(t *testing.T) {
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
	buf, err := ioutil.ReadFile("testdata/day10.txt")
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
