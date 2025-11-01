package adventofcode2017

import (
	"testing"
)

func TestDay15Part1Example(t *testing.T) {
	const want = 588
	got := Day15Part1(NewDay15(linesFromFilename(t, exampleFilename(15))))
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay15Part1(t *testing.T) {
	const want = 619
	got := Day15Part1(NewDay15(linesFromFilename(t, filename(15))))
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay15Part1(b *testing.B) {
	p := NewDay15(linesFromFilename(b, filename(15)))
	for b.Loop() {
		_ = Day15Part1(p)
	}
}

func TestDay15Part2Example(t *testing.T) {
	const want = 309
	got := Day15Part2(NewDay15(linesFromFilename(t, exampleFilename(15))))
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func TestDay15Part2(t *testing.T) {
	const want = 290
	got := Day15Part2(NewDay15(linesFromFilename(t, filename(15))))
	if want != got {
		t.Fatalf("want %v but got %v", want, got)
	}
}

func BenchmarkDay15Part2(b *testing.B) {
	p := NewDay15(linesFromFilename(b, filename(15)))
	for b.Loop() {
		_ = Day15Part2(p)
	}
}

func BenchmarkDay15Part1AsmB1(b *testing.B) {
	p := NewDay15(linesFromFilename(b, filename(15)))
	for b.Loop() {
		_ = day15Part1AsmB1(p.startA, p.startB)
	}
}

func BenchmarkDay15Part2AsmB1(b *testing.B) {
	p := NewDay15(linesFromFilename(b, filename(15)))
	for b.Loop() {
		_ = day15Part2AsmB1(p.startA, p.startB)
	}
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func TestDay15Periodicity(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping long-running periodicity analysis in short mode")
	}

	const (
		factorA = 16807
		factorB = 48271
		modulo  = 2147483647
		startA  = 65
		startB  = 8921
	)

	a := uint(startA)
	periodA := 0
	for i := 1; i <= 10000000; i++ {
		a = (a * factorA) % modulo
		if a&0xFFFF == startA&0xFFFF {
			periodA = i
			break
		}
	}

	b := uint(startB)
	periodB := 0
	for i := 1; i <= 10000000; i++ {
		b = (b * factorB) % modulo
		if b&0xFFFF == startB&0xFFFF {
			periodB = i
			break
		}
	}

	if periodA == 0 || periodB == 0 {
		t.Fatal("could not find period within 10M iterations")
	}

	period := lcm(periodA, periodB)
	t.Logf("Period of lower 16 bits for Generator A: %d", periodA)
	t.Logf("Period of lower 16 bits for Generator B: %d", periodB)
	t.Logf("LCM period: %d", period)
	t.Logf("Number of periods in 40M: %d", 40000000/period)

	a, b = uint(startA), uint(startB)
	matchesInPeriod := 0
	for i := 0; i < period && i < 40000000; i++ {
		a = (a * factorA) % modulo
		b = (b * factorB) % modulo
		if a&0xFFFF == b&0xFFFF {
			matchesInPeriod++
		}
	}

	t.Logf("Matches in first %d iterations: %d", min(period, 40000000), matchesInPeriod)

	if period > 40000000 {
		t.Logf("Period (%d) > 40M: no periodicity optimization possible", period)
		if matchesInPeriod != 588 {
			t.Errorf("expected 588 matches in 40M iterations, got %d", matchesInPeriod)
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
