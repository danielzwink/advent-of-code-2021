package main

import (
	"fmt"
	"time"
)

func main() {
	targetArea := area{minX: 201, maxX: 230, minY: -99, maxY: -65}

	part1Result, part1Duration := part1(targetArea)
	fmt.Printf("Part 1: %4d (duration: %s)\n", part1Result, part1Duration)

	part2Result, part2Duration := part2(targetArea)
	fmt.Printf("Part 2: %4d (duration: %s)\n", part2Result, part2Duration)
}

func part1(targetArea area) (int, time.Duration) {
	start := time.Now()

	maxY := 0
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			currentVelocity := velocity{x: x, y: y}
			reached, currentMaxY := currentVelocity.fireProbe(targetArea)

			if reached && currentMaxY > maxY {
				maxY = currentMaxY
			}
		}
	}
	return maxY, time.Since(start)
}

func part2(targetArea area) (int, time.Duration) {
	start := time.Now()

	count := 0
	for x := -100; x < 500; x++ {
		for y := -100; y < 500; y++ {
			currentVelocity := velocity{x: x, y: y}
			reached, _ := currentVelocity.fireProbe(targetArea)

			if reached {
				count++
			}
		}
	}
	return count, time.Since(start)
}

type velocity struct {
	x, y int
}

func (v *velocity) fireProbe(targetArea area) (bool, int) {
	maxY := 0
	for x, y := 0, 0; x < targetArea.maxX && y > targetArea.minY; {
		// perform step
		x, y = x+v.x, y+v.y
		// persist current maximum
		if y > maxY {
			maxY = y
		}
		// check if target area has been reached
		if targetArea.reached(x, y) {
			return true, maxY
		}
		// change step size
		v.changeStep()
	}
	return false, 0
}

func (v *velocity) changeStep() {
	if v.x > 0 {
		v.x--
	} else if v.x < 0 {
		v.x++
	}
	v.y--
}

type area struct {
	minX, maxX, minY, maxY int
}

func (a *area) reached(x, y int) bool {
	return x >= a.minX && x <= a.maxX && y >= a.minY && y <= a.maxY
}
