package main

import (
	"advent-of-code-2021/pkg/util"
	"bufio"
	"fmt"
	"sort"
	"strings"
	"time"
)

var closingBracketMap = make(map[rune]rune, 4)
var syntaxErrorScore = make(map[rune]int, 4)
var autoCompleteScore = make(map[rune]int, 4)

func main() {
	closingBracketMap['('] = ')'
	closingBracketMap['['] = ']'
	closingBracketMap['{'] = '}'
	closingBracketMap['<'] = '>'

	syntaxErrorScore[')'] = 3
	syntaxErrorScore[']'] = 57
	syntaxErrorScore['}'] = 1197
	syntaxErrorScore['>'] = 25137

	autoCompleteScore[')'] = 1
	autoCompleteScore[']'] = 2
	autoCompleteScore['}'] = 3
	autoCompleteScore['>'] = 4

	part1Result, part1Duration := part1()
	fmt.Printf("Part 1: %10d (duration: %s)\n", part1Result, part1Duration)

	part2Result, part2Duration := part2()
	fmt.Printf("Part 2: %10d (duration: %s)\n", part2Result, part2Duration)
}

func part1() (int, time.Duration) {
	file := util.OpenFile("10")
	scanner := bufio.NewScanner(file)
	start := time.Now()

	score := 0
	for scanner.Scan() {

		chars := make([]rune, 0)
		for _, next := range scanner.Text() {
			if strings.ContainsRune("<{[(", next) {
				chars = append(chars, next)

			} else if strings.ContainsRune(")]}>", next) {
				lastIdx := len(chars) - 1
				last := chars[lastIdx]

				if closingBracketMap[last] == next {
					chars = chars[:lastIdx]
				} else {
					score += syntaxErrorScore[next]
					break
				}
			}
		}
	}
	return score, time.Since(start)
}

func part2() (int, time.Duration) {
	file := util.OpenFile("10")
	scanner := bufio.NewScanner(file)
	start := time.Now()

	autoCompleteScores := make([]int, 0)
	for scanner.Scan() {

		chars := make([]rune, 0)
		corrupted := false
		for _, next := range scanner.Text() {
			if strings.ContainsRune("<{[(", next) {
				chars = append(chars, next)

			} else if strings.ContainsRune(")]}>", next) {
				lastIdx := len(chars) - 1
				last := chars[lastIdx]

				if closingBracketMap[last] == next {
					chars = chars[:lastIdx]
				} else {
					corrupted = true
					break
				}
			}
		}

		remainingLength := len(chars)
		if !corrupted && remainingLength > 0 {
			score := 0
			for i := remainingLength - 1; i >= 0; i-- {
				chunk := closingBracketMap[chars[i]]
				score = (5 * score) + autoCompleteScore[chunk]
			}
			autoCompleteScores = append(autoCompleteScores, score)
		}
	}

	sort.Ints(autoCompleteScores)
	middleIndex := (len(autoCompleteScores) - 1) / 2
	return autoCompleteScores[middleIndex], time.Since(start)
}
