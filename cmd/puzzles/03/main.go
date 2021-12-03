package main

import (
	"advent-of-code-2021/pkg/util"
	"bufio"
	"fmt"
	"log"
	"strconv"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

const LENGTH = 12

func part1() uint64 {
	file := util.OpenFile("03")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	zeroes := [LENGTH]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	ones := [LENGTH]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for scanner.Scan() {
		for i, v := range scanner.Text() {
			if string(v) == "0" {
				zeroes[i]++
			} else {
				ones[i]++
			}
		}
	}

	var gamma, epsilon string
	for i := 0; i < LENGTH; i++ {
		if zeroes[i] > ones[i] {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	gammaValue, _ := strconv.ParseUint(gamma, 2, 12)
	epsilonValue, _ := strconv.ParseUint(epsilon, 2, 12)
	return gammaValue * epsilonValue
}

func part2() uint64 {
	oxyPrevious := util.ReadFile("03")
	co2Previous := oxyPrevious

	oxyComparison := func(zeroes int, ones int) (keep string) {
		if ones >= zeroes {
			keep = "1"
		} else {
			keep = "0"
		}
		return
	}
	co2Comparison := func(zeroes int, ones int) (keep string) {
		if zeroes <= ones {
			keep = "0"
		} else {
			keep = "1"
		}
		return
	}

	oxyRating := determineRating(oxyPrevious, oxyComparison)
	co2Rating := determineRating(co2Previous, co2Comparison)

	oxyValue, _ := strconv.ParseUint(oxyRating, 2, 12)
	co2Value, _ := strconv.ParseUint(co2Rating, 2, 12)

	return oxyValue * co2Value
}

func determineRating(ratings []string, comparison func(int, int) string) string {
	previous := ratings

	for p := 0; p < LENGTH && len(previous) > 1; p++ {
		zeroes, ones := countZeroesAndOnesAtPosition(previous, p)
		keepValue := comparison(zeroes, ones)

		var keptLines []string
		for _, line := range previous {
			if string(line[p]) == keepValue {
				keptLines = append(keptLines, line)
			}
		}

		previous = keptLines
	}

	if len(previous) > 1 {
		log.Fatal("remaining rating not unique")
	}
	return previous[0]
}

func countZeroesAndOnesAtPosition(file []string, position int) (zeroes int, ones int) {
	for _, line := range file {
		if string(line[position]) == "0" {
			zeroes++
		} else {
			ones++
		}
	}
	return zeroes, ones
}
