package adventofcode2017

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
)

func linesFromFilename(tb testing.TB, filename string) []string {
	tb.Helper()
	f, err := os.Open(filename)
	if err != nil {
		tb.Fatal(err)
	}
	return linesFromReader(tb, f)
}

func linesFromReader(tb testing.TB, r io.Reader) []string {
	tb.Helper()
	var lines []string
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)
	}
	if err := sc.Err(); err != nil {
		tb.Fatal(err)
	}
	return lines
}

func exampleFilename(day uint8) string {
	return fmt.Sprintf("testdata/day%02d_example.txt", int(day))
}

func filename(day uint8) string {
	return fmt.Sprintf("testdata/day%02d.txt", int(day))
}

// example1Filename returns the filename for day NN example 1
func example1Filename(day uint8) string {
	return fmt.Sprintf("testdata/day%02d_example1.txt", int(day))
}

// example2Filename returns the filename for day NN example 2
func example2Filename(day uint8) string {
	return fmt.Sprintf("testdata/day%02d_example2.txt", int(day))
}

// example3Filename returns the filename for day NN example 3
func example3Filename(day uint8) string {
	return fmt.Sprintf("testdata/day%02d_example3.txt", int(day))
}

// linesAsNumber converts strings into integer.
func linesAsNumbers(tb testing.TB, lines []string) []int {
	tb.Helper()
	var is []int
	for i := range lines {
		n, err := strconv.Atoi(lines[i])
		if err != nil {
			tb.Fatalf("error in line %d: cannot convert %q to number", i, lines[i])
		}
		is = append(is, n)
	}
	return is
}

func numbersFromFilename(tb testing.TB, filename string) []int {
	tb.Helper()
	ls := linesFromFilename(tb, filename)
	return linesAsNumbers(tb, ls)
}

// file reads the main input file bytes for day N (zero-padded).
func file(tb testing.TB, day uint8) []byte {
	tb.Helper()
	buf, err := os.ReadFile(filename(day))
	if err != nil {
		tb.Fatal(err)
	}
	return buf
}

// exampleFile reads the example input file bytes for day N (zero-padded).
func exampleFile(tb testing.TB, day uint8) []byte {
	tb.Helper()
	buf, err := os.ReadFile(exampleFilename(day))
	if err != nil {
		tb.Fatal(err)
	}
	return buf
}
