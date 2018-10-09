package adventofcode2017

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestDay13(t *testing.T) {
	f, err := os.Open("testdata/day13_sample.txt")
	if err != nil {
		t.Fatal(err)
	}
	sc := bufio.NewScanner(f)
	var m = make(map[int]int)
	var maxDepth int
	for sc.Scan() {
		line := sc.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			t.Fatalf("want key: value but got %v\n", line)
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[0])
		k, err := strconv.Atoi(key)
		if err != nil {
			t.Fatalf("want numeric key but got %v\n", key)
		}
		v, err := strconv.Atoi(value)
		if err != nil {
			t.Fatalf("want numeric key but got %v\n", value)
		}
		m[k] = v
		if maxDepth < k {
			maxDepth = k
		}
	}
	if maxDepth != 6 {
		t.Fatalf("want max depth = 6 but got %d", maxDepth)
	}
	// make sure all entries from file are available
	if len(m) != 4 {
		t.Fatalf("want len=4 but got %d", len(m))
	}

	// convert map to array
	layers := make([]int, maxDepth+1)
	for k, v := range m {
		layers[k] = v
	}
	want := 24
	got := Day13(layers)
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}
