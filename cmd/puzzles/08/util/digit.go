package digit

import (
	"sort"
	"strings"
)

type Digit map[string]bool

func NewDigit(segments string) Digit {
	digit := make(Digit, len(segments))
	for _, segment := range segments {
		digit[string(segment)] = true
	}
	return digit
}

func (d Digit) size() int {
	return len(d)
}

func (d Digit) copy() Digit {
	copy := make(Digit, d.size())
	for segment, _ := range d {
		copy[segment] = true
	}
	return copy
}

func (d Digit) remove(other Digit) Digit {
	for segment, _ := range other {
		delete(d, segment)
	}
	return d
}

func (d Digit) contains(other Digit) bool {
	for segment, _ := range other {
		_, found := d[segment]
		if !found {
			return false
		}
	}
	return true
}

func (d Digit) NormalizedSegmentPattern() string {
	segments := make([]string, 0, len(d))
	for segment, _ := range d {
		segments = append(segments, segment)
	}
	sort.Strings(segments)
	return strings.Join(segments, "")
}

type Digits [10]Digit

/*
Example:

Value  Size  String   Rule
1      2     ab
4      4     ab ef
7      3     ab d
8      7     abcdefg

3      5     ab cdf   contains 1 (ab)
5      5     bcd ef   contains 4-1 (ef)
2      5     acdfg    remains

6      6     bcdefg   does not contain 1 (!ab)
0      6     abcdeg   contains 8-7-4 (cg)
9      6     abcdef   remains
*/

func NewDigits(segmentPatterns []string) Digits {
	digits := Digits{}

	digits5Segments := make([]Digit, 0, 3)
	digits6Segments := make([]Digit, 0, 3)

	for _, segmentPattern := range segmentPatterns {
		digit := NewDigit(segmentPattern)

		// these 4 are easy
		if digit.size() == 2 {
			digits[1] = digit
		}
		if digit.size() == 3 {
			digits[7] = digit
		}
		if digit.size() == 4 {
			digits[4] = digit
		}
		if digit.size() == 7 {
			digits[8] = digit
		}

		// these follow some rules and must be stored and evaluated separately
		if digit.size() == 5 {
			digits5Segments = append(digits5Segments, digit)
		}
		if digit.size() == 6 {
			digits6Segments = append(digits6Segments, digit)
		}
	}

	digit4minus1Pattern := digits[4].copy().remove(digits[1])
	remainingDigits5Segments := make([]Digit, 0, 2)
	for _, digit := range digits5Segments {
		if digit.contains(digits[1]) {
			digits[3] = digit
		} else {
			remainingDigits5Segments = append(remainingDigits5Segments, digit)
		}
	}
	for _, digit := range remainingDigits5Segments {
		if digit.contains(digit4minus1Pattern) {
			digits[5] = digit
		} else {
			digits[2] = digit
		}
	}

	digit8minus7minus4Pattern := digits[8].copy().remove(digits[7]).remove(digits[4])
	remainingDigits6Segments := make([]Digit, 0, 2)
	for _, digit := range digits6Segments {
		if !digit.contains(digits[1]) {
			digits[6] = digit
		} else {
			remainingDigits6Segments = append(remainingDigits6Segments, digit)
		}
	}
	for _, digit := range remainingDigits6Segments {
		if digit.contains(digit8minus7minus4Pattern) {
			digits[0] = digit
		} else {
			digits[9] = digit
		}
	}

	return digits
}

func (d Digits) NormalizedSegmentPatternMap() map[string]int {
	result := make(map[string]int, 10)
	for index, digit := range d {
		result[digit.NormalizedSegmentPattern()] = index
	}
	return result
}
