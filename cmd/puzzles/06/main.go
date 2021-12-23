package main

import (
	"advent-of-code-2021/pkg/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	part1ResultBadPerformance, part1DurationBadPerformance := part1BadPerformance()
	fmt.Printf("Part 1: %d (took %s)\n", part1ResultBadPerformance, part1DurationBadPerformance)

	part1Result, part1Duration := part1()
	fmt.Printf("Part 1: %d (took %s)\n", part1Result, part1Duration)

	part2Result, part2Duration := part2()
	fmt.Printf("Part 2: %d (took %s)\n", part2Result, part2Duration)
}

func part1BadPerformance() (int, time.Duration) {
	lines := util.ReadFile("06/input")

	start := time.Now()
	var fish []int
	for _, initialValue := range strings.Split(lines[0], ",") {
		initialTimer, _ := strconv.Atoi(initialValue)
		fish = append(fish, initialTimer)
	}

	for i := 1; i <= 80; i++ {
		var newFish []int
		for j, f := range fish {
			if f == 0 {
				newFish = append(newFish, 8)
				fish[j] = 6
			} else {
				fish[j]--
			}
		}
		fish = append(fish, newFish...)
	}
	return len(fish), time.Since(start)
}

func part1() (int, time.Duration) {
	lines := util.ReadFile("06/input")
	start := time.Now()
	timerCounter := initTimer(lines[0])
	result := performEvolution(timerCounter, 80)
	return result, time.Since(start)
}

func part2() (int, time.Duration) {
	lines := util.ReadFile("06/input")
	start := time.Now()
	timerCounter := initTimer(lines[0])
	result := performEvolution(timerCounter, 256)
	return result, time.Since(start)
}

func initTimer(input string) []int {
	timerCounter := make([]int, 9)

	for _, initialTimerValue := range strings.Split(input, ",") {
		initialTimer, _ := strconv.Atoi(initialTimerValue)
		timerCounter[initialTimer]++
	}

	return timerCounter
}

func performEvolution(timerCounter []int, days int) int {
	for i := 1; i <= days; i++ {
		// store the count for 0 as new fish count for later purpose
		newFishCount := timerCounter[0]
		// shift all counters to the left by 1
		for i := 1; i <= 8; i++ {
			timerCounter[i-1] = timerCounter[i]
		}
		// toggle the previously stored count for 0 to 6 and add them there
		timerCounter[6] += newFishCount
		// create new fish with 8 accordingly (override instead of add is important here)
		timerCounter[8] = newFishCount
	}

	sum := 0
	for i := 0; i <= 8; i++ {
		sum += timerCounter[i]
	}
	return sum
}
