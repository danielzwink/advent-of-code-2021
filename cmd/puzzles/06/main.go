package main

import (
	"advent-of-code-2021/pkg/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	file := util.ReadFile("06")
	timer := initTimer(file[0])
	return performEvolution(timer, 80)
}

func part2() int {
	file := util.ReadFile("06")
	timer := initTimer(file[0])
	return performEvolution(timer, 256)
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
