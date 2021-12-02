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

func part1() int {
	file := util.OpenFile("01")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	increases := 0
	previous := -1

	for scanner.Scan() {
		current := getCurrentLineAsNumber(scanner)

		if previous >= 0 && current > previous {
			increases++
		}
		previous = current
	}
	return increases
}

func part2() int {
	file := util.OpenFile("01")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	increases := 0
	previousSum := 0
	slidingWindow := make([]int, 0)

	for scanner.Scan() {
		current := getCurrentLineAsNumber(scanner)

		if len(slidingWindow) == 3 {
			slidingWindow = slidingWindow[1:]
		}
		slidingWindow = append(slidingWindow, current)

		if len(slidingWindow) == 3 {
			currentSum := slidingWindow[0] + slidingWindow[1] + slidingWindow[2]

			if previousSum > 0 && previousSum < currentSum {
				increases++
			}
			previousSum = currentSum
		}
	}
	return increases
}

func getCurrentLineAsNumber(scanner *bufio.Scanner) int {
	line := scanner.Text()
	number, err := strconv.Atoi(line)
	if err != nil {
		log.Fatal(err)
	}
	return number
}
