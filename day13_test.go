package adventofcode2017

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func parse(filename string, newLayer func(int, int)) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		parts := strings.Split(line, ": ")
		if len(parts) != 2 {
			return fmt.Errorf("want key: value but got %v", line)
		}
		k, err := strconv.Atoi(parts[0])
		if err != nil {
			return fmt.Errorf("want numeric key but got %v", parts[0])
		}
		v, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("want numeric key but got %v", parts[1])
		}
		newLayer(k, v)
	}
	return nil
}

// asSparseArray returns sparse layers of sample, where index is depth and
// value is range
func asSparseArray(filename string) ([]int, error) {
	var m = make(map[int]int)
	var maxDepth int
	f := func(depth, range_ int) {
		m[depth] = range_
		if maxDepth < depth {
			maxDepth = depth
		}
	}
	if err := parse(filename, f); err != nil {
		return nil, err
	}
	// convert map to array
	layers := make([]int, maxDepth+1)
	for k, v := range m {
		layers[k] = v
	}
	return layers, nil
}
func TestDay13UsingSparseArray(t *testing.T) {
	layers, err := asSparseArray("testdata/day13_sample.txt")
	if err != nil {
		t.Fatal(err)
	}
	if len(layers) != 7 {
		t.Fatalf("want max depth = 7 but got %d", len(layers))
	}
	want := 24
	got := Day13UsingSparseArray(layers)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay13UsingSparseArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		layers, err := asSparseArray("testdata/day13_sample.txt")
		if err != nil {
			b.Fatal(err)
		}
		Day13UsingSparseArray(layers)
	}
}

func asList(filename string) (*node, error) {
	var head, prev *node
	f := func(depth, range_ int) {
		l := &node{
			Depth: depth,
			Range: range_,
		}
		if head == nil {
			head = l
		}
		if prev != nil {
			prev.Next = l
		}
		prev = l
	}
	if err := parse(filename, f); err != nil {
		return nil, err
	}
	return head, nil
}

func TestDay13UsingList(t *testing.T) {
	l, err := asList("testdata/day13_sample.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := 24
	got := Day13UsingList(l)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay13UsingList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l, err := asList("testdata/day13_sample.txt")
		if err != nil {
			b.Fatal(err)
		}
		Day13UsingList(l)
	}
}

func TestDay13Sample(t *testing.T) {
	filename := "testdata/day13_sample.txt"
	layers, err := asSparseArray(filename)
	if err != nil {
		t.Fatal(err)
	}
	want := Day13UsingSparseArray(layers)

	l, err := asList(filename)
	if err != nil {
		t.Fatal(err)
	}
	got := Day13UsingList(l)

	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay13(t *testing.T) {
	filename := "testdata/day13.txt"
	layers, err := asSparseArray(filename)
	if err != nil {
		t.Fatal(err)
	}
	want := Day13UsingSparseArray(layers)

	l, err := asList(filename)
	if err != nil {
		t.Fatal(err)
	}
	got := Day13UsingList(l)

	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}
