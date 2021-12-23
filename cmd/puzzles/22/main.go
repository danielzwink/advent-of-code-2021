package main

import (
	"advent-of-code-2021/pkg/util"
	"fmt"
	"regexp"
	"strconv"
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
	start := time.Now()

	rebootSteps := readRebootSteps("22", true, -50, 50)
	cubes := make(map[string]bool, 0)
	for _, step := range rebootSteps {
		for x := step.X.Min; x <= step.X.Max; x++ {
			for y := step.Y.Min; y <= step.Y.Max; y++ {
				for z := step.Z.Min; z <= step.Z.Max; z++ {
					cube := fmt.Sprintf("%v,%v,%v", x, y, z)
					if step.TurnOn {
						cubes[cube] = true
					} else {
						delete(cubes, cube)
					}
				}
			}
		}
	}
	return len(cubes), time.Since(start)
}

func part2() (int, time.Duration) {
	start := time.Now()
	return 0, time.Since(start)
}

type Range struct {
	Min, Max int
}

type Cuboid struct {
	X, Y, Z Range
}

type RebootStep struct {
	Cuboid
	TurnOn bool
}

func readRebootSteps(day string, validate bool, lowerBound, upperBound int) []RebootStep {
	lines := util.ReadFile(day)
	numberRegexp, _ := regexp.Compile("-?[0-9]+")

	rebootSteps := make([]RebootStep, 0, len(lines))
	for _, line := range lines {
		turnOn := strings.HasPrefix(line, "on ")
		values := numberRegexp.FindAllString(line, 6)

		numbers, valid := convertToNumbers(values, validate, lowerBound, upperBound)
		if !valid {
			continue
		}

		rebootStep := RebootStep{
			TurnOn: turnOn,
			Cuboid: Cuboid{
				X: Range{Min: numbers[0], Max: numbers[1]},
				Y: Range{Min: numbers[2], Max: numbers[3]},
				Z: Range{Min: numbers[4], Max: numbers[5]},
			},
		}
		rebootSteps = append(rebootSteps, rebootStep)
	}
	return rebootSteps
}

func convertToNumbers(values []string, validate bool, lowerBound, upperBound int) ([]int, bool) {
	numbers := make([]int, len(values))
	valid := true
	for i, value := range values {
		number, _ := strconv.Atoi(value)
		numbers[i] = number

		if validate {
			valid = valid && number >= lowerBound && number <= upperBound
		}
	}
	return numbers, valid
}
