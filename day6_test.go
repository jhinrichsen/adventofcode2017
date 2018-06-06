package adventofcode2017

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func ExampleDay06() {
	const nBanks = 16
	buf, err := ioutil.ReadFile("testdata/day6.txt")
	if err != nil {
		panic(err)
	}
	fs := strings.Fields(string(buf))
	if len(fs) != nBanks {
		panic(fmt.Errorf("want %d fields but got %d\n", nBanks, len(fs)))
	}
	var banks Banks
	for i := 0; i < nBanks; i++ {
		banks[i], err = strconv.Atoi(fs[i])
		if err != nil {
			panic(fmt.Errorf("error converting col %d: %v\n", i, err))
		}
	}
	fmt.Println(Day06Impl1(banks, nBanks))
	// Output: 5042
}

func TestDay06Sample(t *testing.T) {
	banks := Banks{0, 2, 7, 0}
	want := 5
	got := Day06Impl1(banks, 4)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func TestArrayAsKeyInMap(t *testing.T) {
	var m = make(map[[3]int]bool)

	m[[3]int{1, 2, 3}] = true
	m[[3]int{2, 3, 4}] = true
	m[[3]int{2, 3, 4}] = true

	want := 2
	got := len(m)
	if want != got {
		t.Fatalf("want %d but got %d\n", want, got)
	}
}

func BenchmarkDay06Impl1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Day06Impl1(Banks{0, 2, 7, 0}, 4)
	}
}

func BenchmarkDay06Impl2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Day06Impl2([]int{0, 2, 7, 0})
	}
}
