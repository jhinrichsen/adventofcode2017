package adventofcode2017

// Day01Part1 finds the sum of all digits that match the next digit.
func Day01Part1(digits []byte) (sum uint) {
    return day1(digits, 1)
}

// Day01Part2 finds the sum of all digits that match the halfway around digit.
func Day01Part2(digits []byte) (sum uint) {
    return day1(digits, len(digits)/2)
}

func day1(digits []byte, n int) (sum uint) {
    // Normalize: if input looks like ASCII '0'..'9', convert to 0..9.
    needConvert := false
    for _, b := range digits {
        if b > 9 { // ASCII digits are > 9; numeric inputs are <= 9
            needConvert = true
            break
        }
    }
    if needConvert {
        // Trim any whitespace-like bytes and convert digits in-place copy
        // Create a new slice to avoid mutating caller data
        out := make([]byte, 0, len(digits))
        for _, b := range digits {
            if b == '\n' || b == '\r' || b == '\t' || b == ' ' { // skip whitespace
                continue
            }
            if b >= '0' && b <= '9' {
                out = append(out, b-'0')
            }
        }
        digits = out
    }
    cycle := len(digits)
    for i, digit := range digits {
        succ := (i + n) % cycle
        if digit == digits[succ] {
            sum += uint(digit)
        }
    }
    return
}
