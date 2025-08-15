package adventofcode2017

import (
	"fmt"
	"reflect"
	"testing"
)

const notSame = "want %v but got %v"

func TestDay07Part1Example(t *testing.T) {
	const want = "tknk"
	ss := linesFromFilename(t, exampleFilename(7))
	got := Day07Part1(ss)
	if want != got {
		t.Fatalf(notSame, want, got)
	}
}

func TestDay07Part1(t *testing.T) {
	const want = "veboyvy"
	ss := linesFromFilename(t, filename(7))
	got := Day07Part1(ss)
	if want != got {
		t.Fatalf(notSame, want, got)
	}
}

func BenchmarkDay07Part1(b *testing.B) {
	ss := linesFromFilename(b, filename(7))
	b.ResetTimer()
	for b.Loop() {
		Day07Part1(ss)
	}
}

func TestNewProgramLeaf(t *testing.T) {
	const (
		name   = "abcde"
		weight = 13
	)
	want := program{"abcde", 13, nil, -1, 0}
	got, err := NewProgram(fmt.Sprintf("%s (%d)", name, weight))
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf(notSame, want, got)
	}
}

func TestNewProgramDisk(t *testing.T) {
	const (
		name   = "abcde"
		weight = 13
		disk1  = "bcdef"
		disk2  = "cdefg"
	)
	m := make(Disk)
	m[disk1] = true
	m[disk2] = true
	want := program{"abcde", 13, m, -1, 0}
	s := fmt.Sprintf("%s (%d) -> %s, %s", name, weight, disk1, disk2)
	got, err := NewProgram(s)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf(notSame, want, got)
	}
}

func TestDay07Part2Example(t *testing.T) {
	const want = 60
	ss := linesFromFilename(t, exampleFilename(7))
	got, err := Day07Part2(ss)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf(notSame, want, got)
	}
}

func TestDay07Part2(t *testing.T) {
	const want = 749
	ss := linesFromFilename(t, filename(7))
	got, err := Day07Part2(ss)
	if err != nil {
		t.Fatal(err)
	}
	if want != got {
		t.Fatalf(notSame, want, got)
	}
}

func BenchmarkDay07Part2(b *testing.B) {
	ss := linesFromFilename(b, filename(7))
	b.ResetTimer()
	for b.Loop() {
		_, _ = Day07Part2(ss)
	}
}
