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
		willFlash := false
		for row := 0; row < Size; row++ {
			for col := 0; col < Size; col++ {
				octopuses[row][col]++

				if octopuses[row][col] > 9 {
					willFlash = true
				}
			}
		}

		// flash
		for willFlash {
			willFlash = false
			// for every octopus
			for row := 0; row < Size; row++ {
				for col := 0; col < Size; col++ {
					// that reached a level above 9
					if octopuses[row][col] > 9 {
						// flash its adjacent octopuses
						for y := -1; y <= 1; y++ {
							for x := -1; x <= 1; x++ {
								// check bounds
								adjacentRow, adjacentCol := row+y, col+x
								if adjacentRow < 0 || adjacentRow >= Size || adjacentCol < 0 || adjacentCol >= Size {
									continue
								}
								// exclude past or upcoming flashes
								level := octopuses[adjacentRow][adjacentCol]
								if level < 0 || level > 9 {
									continue
								}
								// increase adjacent energy level due to flash
								octopuses[adjacentRow][adjacentCol]++
								// flag upcoming flash
								if octopuses[adjacentRow][adjacentCol] > 9 {
									willFlash = true
								}
							}
						}
						// mark and count flash
						octopuses[row][col] = -1
						flashes++
					}
				}
			}
		}

		// reset level after flash
		for row := 0; row < Size; row++ {
			for col := 0; col < Size; col++ {
				if octopuses[row][col] == -1 {
					octopuses[row][col] = 0
				}
			}
		}
	}
	return flashes, time.Since(start)
}

func part2() (int, time.Duration) {
	octopuses := readOctopuses()
	start := time.Now()

	for step := 1; true; step++ {

		// increase energy level by 1
		willFlash := false
		for row := 0; row < Size; row++ {
			for col := 0; col < Size; col++ {
				octopuses[row][col]++

				if octopuses[row][col] > 9 {
					willFlash = true
				}
			}
		}

		// flash
		for willFlash {
			willFlash = false
			// for every octopus
			for row := 0; row < Size; row++ {
				for col := 0; col < Size; col++ {
					// that reached a level above 9
					if octopuses[row][col] > 9 {
						// flash its adjacent octopuses
						for y := -1; y <= 1; y++ {
							for x := -1; x <= 1; x++ {
								// check bounds
								adjacentRow, adjacentCol := row+y, col+x
								if adjacentRow < 0 || adjacentRow >= Size || adjacentCol < 0 || adjacentCol >= Size {
									continue
								}
								// exclude past or upcoming flashes
								level := octopuses[adjacentRow][adjacentCol]
								if level < 0 || level > 9 {
									continue
								}
								// increase adjacent energy level due to flash
								octopuses[adjacentRow][adjacentCol]++
								// flag upcoming flash
								if octopuses[adjacentRow][adjacentCol] > 9 {
									willFlash = true
								}
							}
						}
						// mark flash
						octopuses[row][col] = -1
					}
				}
			}
		}

		// reset level after flash
		flashes := 0
		for row := 0; row < Size; row++ {
			for col := 0; col < Size; col++ {
				if octopuses[row][col] == -1 {
					octopuses[row][col] = 0
					flashes++
				}
			}
		}
		// return when all flashed simultaneously
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
