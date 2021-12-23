package main

import (
	digit "advent-of-code-2021/cmd/puzzles/08/util"
	"advent-of-code-2021/pkg/util"
	"bufio"
	"fmt"
	"strings"
	"time"
)

func main() {
	part1Result, part1Duration := part1()
	fmt.Printf("Part 1: %6d (duration: %s)\n", part1Result, part1Duration)

	part2Result, part2Duration := part2()
	fmt.Printf("Part 2: %6d (duration: %s)\n", part2Result, part2Duration)
}

func part1() (int, time.Duration) {
	file := util.OpenFile("08/input")
	defer file.Close()
	start := time.Now()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), " | ")
		output := strings.Split(values[1], " ")
		for _, v := range output {
			length := len(v)

			if length >= 2 && length <= 4 || length == 7 {
				sum++
			}
		}
	}

	return sum, time.Since(start)
}

func part2() (int, time.Duration) {
	file := util.OpenFile("08/input")
	defer file.Close()
	start := time.Now()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), " | ")
		input := strings.Split(values[0], " ")
		output := strings.Split(values[1], " ")

		sum += translate(input, output)
	}

	return sum, time.Since(start)
}

func translate(input []string, output []string) (result int) {
	segmentPatterns := digit.NewDigits(input).NormalizedSegmentPatternMap()

	sp1 := digit.NewDigit(output[0]).NormalizedSegmentPattern()
	sp2 := digit.NewDigit(output[1]).NormalizedSegmentPattern()
	sp3 := digit.NewDigit(output[2]).NormalizedSegmentPattern()
	sp4 := digit.NewDigit(output[3]).NormalizedSegmentPattern()

	result += 1000 * segmentPatterns[sp1]
	result += 100 * segmentPatterns[sp2]
	result += 10 * segmentPatterns[sp3]
	result += segmentPatterns[sp4]

	return result
}
