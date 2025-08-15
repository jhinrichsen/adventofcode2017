package adventofcode2017

import (
	"fmt"
	"strconv"
	"strings"
)

var ErrorDay10Length = fmt.Errorf("illegal length")

// day10Puzzle holds parsed input and configuration for the solver.
// Tests can modify Size/ZeroBased to match example semantics.
type day10Puzzle struct {
	lengths []int
	size    int
	zero    bool
}

// Day10Part2 computes the full Knot Hash of the input as specified by AoC 2017 Day 10 Part 2.
// It treats the input as ASCII bytes (after trimming whitespace), appends the standard suffix,
// runs 64 rounds over a 0..255 list, then computes the dense hash and returns a 32-hex-digit string.
func Day10Part2(buf []byte) string {
    // 1) Parse input as ASCII bytes and append suffix
    s := strings.TrimSpace(string(buf))
    lengths := make([]int, 0, len(s)+5)
    for _, r := range s {
        lengths = append(lengths, int(byte(r)))
    }
    lengths = append(lengths, 17, 31, 73, 47, 23)

    // 2) Init list 0..255
    const size = 256
    list := make([]int, size)
    for i := range list {
        list[i] = i
    }

    // 3) Run 64 rounds
    pos := 0
    skip := 0
    for range 64 {
        for _, n := range lengths {
            if n > 0 {
                wrap := func(x int) int { return x % len(list) }
                for lower, upper := pos, pos+n-1; lower < upper; lower, upper = lower+1, upper-1 {
                    list[wrap(lower)], list[wrap(upper)] = list[wrap(upper)], list[wrap(lower)]
                }
            }
            pos = (pos + n + skip) % len(list)
            skip++
        }
    }

    // 4) Dense hash: XOR blocks of 16
    dense := make([]byte, 16)
    for block := range 16 {
        start := block * 16
        v := list[start]
        for _, elem := range list[start+1 : start+16] {
            v ^= elem
        }
        dense[block] = byte(v)
    }

    // 5) Hex encode (lowercase, two digits each)
    // Preallocate 32 runes/bytes; using fmt for simplicity here
    var b strings.Builder
    b.Grow(32)
    for _, v := range dense {
        // %02x formatting
        b.WriteString(fmt.Sprintf("%02x", v))
    }
    return b.String()
}

// NewDay10 parses the comma-separated input lengths and returns them.
// It does not solve the puzzle; call Day10 with an appropriate list and the returned lengths.
func NewDay10(buf []byte) day10Puzzle {
    s := strings.TrimSpace(string(buf))
    var lengths []int
    if s != "" {
        parts := strings.Split(s, ",")
        lengths = make([]int, 0, len(parts))
        for i, part := range parts {
            n, err := strconv.Atoi(strings.TrimSpace(part))
            if err != nil {
                panic(fmt.Errorf("cannot convert number to int at column %d", i))
            }
            lengths = append(lengths, n)
        }
	}
	// Defaults: AoC standard configuration (0..255 list)
	return day10Puzzle{lengths: lengths, size: 256, zero: true}
}

// Day10 is the solver. It constructs the working list according to the puzzle
// configuration and runs one knot-hash round, returning the product of the
// first two elements.
func Day10(p day10Puzzle) int {
    size := p.size
    if size <= 0 {
        size = 256
    }
    list := make([]int, size)
    for i := range list {
        if p.zero {
            list[i] = i
        } else {
            list[i] = i + 1
        }
    }
    // Inline the knot hash single round
    pos := 0
    skip := 0
    for _, n := range p.lengths {
        if n > 0 {
            // reverse list[pos:pos+n] with wraparound
            wrap := func(x int) int { return x % len(list) }
            for lower, upper := pos, pos+n-1; lower < upper; lower, upper = lower+1, upper-1 {
                list[wrap(lower)], list[wrap(upper)] = list[wrap(upper)], list[wrap(lower)]
            }
        }
        pos = (pos + n + skip) % len(list)
        skip++
    }
    return list[0] * list[1]
}
