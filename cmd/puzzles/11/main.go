package main

import (
	"advent-of-code-2021/pkg/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type octopus int

func main() {
	part1Result, part1Duration := part1()
	fmt.Printf("Part 1: %5d (duration: %s)\n", part1Result, part1Duration)

	part2Result, part2Duration := part2()
	fmt.Printf("Part 2: %5d (duration: %s)\n", part2Result, part2Duration)
}

const Size = 10
const Steps = 100

func part1() (int, time.Duration) {
	octopuses := readOctopuses()
	start := time.Now()

	flashes := 0
	for step := 1; step <= Steps; step++ {
		// increase energy level by 1
		willFlash := increaseEnergyLevels(octopuses)

		// perform flashes
		performFlashes(octopuses, willFlash)

		// reset level and count flashes
		flashes += resetFlashes(octopuses)
	}
	return flashes, time.Since(start)
}

func part2() (int, time.Duration) {
	octopuses := readOctopuses()
	start := time.Now()

	for step := 1; true; step++ {
		// increase energy level by 1
		willFlash := increaseEnergyLevels(octopuses)

		// perform flashes
		performFlashes(octopuses, willFlash)

		// reset level and count flashes
		flashes := resetFlashes(octopuses)

		// return step when all octopuses flashed simultaneously
		if flashes == Size*Size {
			return step, time.Since(start)
		}
	}
	return 0, time.Since(start)
}

func readOctopuses() [][]octopus {
	lines := util.ReadFile("11")
	octopuses := make([][]octopus, 0, Size)

	for _, line := range lines {
		row := make([]octopus, 0, Size)
		for _, value := range strings.Split(line, "") {
			level, _ := strconv.Atoi(value)
			row = append(row, octopus(level))
		}
		octopuses = append(octopuses, row)
	}
	return octopuses
}

func increaseEnergyLevels(octopuses [][]octopus) (willFlash bool) {
	for row := 0; row < Size; row++ {
		for col := 0; col < Size; col++ {
			octopuses[row][col]++

			if octopuses[row][col] > 9 {
				willFlash = true
			}
		}
	}
	return willFlash
}

func performFlashes(octopuses [][]octopus, willFlash bool) {
	for willFlash {
		willFlash = false

		// for every octopus
		for row := 0; row < Size; row++ {
			for col := 0; col < Size; col++ {

				// that reached an energy level above 9
				if octopuses[row][col] <= 9 {
					continue
				}

				// perform flash by increasing adjacent levels
				for y := -1; y <= 1; y++ {
					for x := -1; x <= 1; x++ {
						// check bounds and not self
						adjacentRow, adjacentCol := row+y, col+x
						if adjacentRow < 0 || adjacentRow >= Size || adjacentCol < 0 || adjacentCol >= Size || (adjacentRow == row && adjacentCol == col) {
							continue
						}
						// exclude past or upcoming flashes
						level := octopuses[adjacentRow][adjacentCol]
						if level < 0 || level > 9 {
							continue
						}
						// increase energy level due to flash
						octopuses[adjacentRow][adjacentCol]++
						// check if another loop is required
						if octopuses[adjacentRow][adjacentCol] > 9 {
							willFlash = true
						}
					}
				}

				// mark performed flash
				octopuses[row][col] = -1
			}
		}
	}
}

func resetFlashes(octopuses [][]octopus) (flashes int) {
	for row := 0; row < Size; row++ {
		for col := 0; col < Size; col++ {
			if octopuses[row][col] == -1 {
				octopuses[row][col] = 0
				flashes++
			}
		}
	}
	return flashes
}
