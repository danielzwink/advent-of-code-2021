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
	timer := make([]int, 9)

	for _, initialValue := range strings.Split(input, ",") {
		i, _ := strconv.Atoi(initialValue)
		timer[i]++
	}

	return timer
}

func performEvolution(timer []int, days int) int {
	for i := 1; i <= days; i++ {
		newFish := timer[0]
		for i := 1; i <= 8; i++ {
			timer[i-1] = timer[i]
		}
		timer[6] += newFish
		timer[8] = newFish
	}

	sum := 0
	for i := 0; i <= 8; i++ {
		sum += timer[i]
	}
	return sum
}
