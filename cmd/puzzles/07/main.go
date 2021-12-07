package main

import (
	"advent-of-code-2021/pkg/util"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type crabs struct {
	position, count int
}

func main() {
	part1Distance, part1Duration := part1()
	fmt.Printf("Part 1: %9d (distance) %s (duration)\n", part1Distance, part1Duration)

	part2Distance, part2Duration := part2()
	fmt.Printf("Part 2: %9d (distance) %s (duration)\n", part2Distance, part2Duration)
}

func part1() (int, time.Duration) {
	file := util.ReadFile("07")
	start := time.Now()

	sortedCrabs := getSortedCrabPositions(file[0])
	totalDistance := calculateMinTotalDistance(sortedCrabs, calcDistancePart1)
	return totalDistance, time.Since(start)
}

func part2() (int, time.Duration) {
	file := util.ReadFile("07")
	start := time.Now()

	sortedCrabs := getSortedCrabPositions(file[0])
	totalDistance := calculateMinTotalDistance(sortedCrabs, calcDistancePart2)
	return totalDistance, time.Since(start)
}

func calculateMinTotalDistance(sortedCrabs []crabs, calculateDistance func(int, int) int) int {
	minDistance := math.MaxInt64

	for position := 0; position < len(sortedCrabs); position++ {
		sumOfDistances := 0
		for _, crabs := range sortedCrabs {
			sumOfDistances += calculateDistance(position, crabs.position) * crabs.count
		}
		if sumOfDistances < minDistance {
			minDistance = sumOfDistances
		}
	}
	return minDistance
}

func calcDistancePart1(start, end int) int {
	return abs(end - start)
}

func calcDistancePart2(start, end int) int {
	distance := abs(end - start)
	return (1 + distance) * distance / 2
}

func abs(value int) int {
	if value < 0 {
		return value * -1
	} else {
		return value
	}
}

func getSortedCrabPositions(input string) []crabs {
	positions := getPositions(input)
	maxPosition := findMaximum(positions)

	sortedCrabPositions := make([]int, maxPosition+1)
	for _, position := range positions {
		sortedCrabPositions[position]++
	}

	sortedCrabs := make([]crabs, 0, maxPosition)
	for position, count := range sortedCrabPositions {
		if count > 0 {
			sortedCrabs = append(sortedCrabs, crabs{position: position, count: count})
		}
	}
	return sortedCrabs
}

func getPositions(input string) []int {
	positionValues := strings.Split(input, ",")

	positions := make([]int, len(positionValues))
	for i, v := range positionValues {
		position, _ := strconv.Atoi(v)
		positions[i] = position
	}
	return positions
}

func findMaximum(values []int) (maximum int) {
	maximum = values[0]
	for _, v := range values[1:] {
		if v > maximum {
			maximum = v
		}
	}
	return maximum
}
