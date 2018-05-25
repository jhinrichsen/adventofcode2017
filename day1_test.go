package adventofcode2017

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"testing"
)

func TestDay1Sample(t *testing.T) {
	var ts = []struct {
		want int
		in   []byte
	}{
		{3, []byte{1, 1, 2, 2}},
		{4, []byte{1, 1, 1, 1}},
		{0, []byte{1, 2, 3, 4}},
		{9, []byte{9, 1, 2, 1, 2, 1, 2, 9}},
	}
	for _, tt := range ts {
		got := Day1(tt.in)
		if tt.want != got {
			t.Fatalf("want %d but got %d\n", tt.want, got)
		}
	}
}

func input() ([]byte, error) {
	buf, err := ioutil.ReadFile("testdata/day1.txt")
	if err != nil {
		return nil, err
	}
	buf = bytes.TrimSpace(buf)

	// map from '0'..'9' to 0..9
	for i := range buf {
		if buf[i] < '0' || buf[i] > '9' {
			return buf, fmt.Errorf("index %d out of range: %d\n", i, buf[i])
		}
		buf[i] = buf[i] - '0'
	}
	return buf, nil
}

func TestDay1(t *testing.T) {
	in, err := input()
	if err != nil {
		t.Fatal(err)
	}
	got := Day1(in)
	log.Printf("day 1: got %d\n", got)
}

func BenchmarkDay1(b *testing.B) {
	in, err := input()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Day1(in)
	}
}
