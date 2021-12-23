package main

import (
	math "advent-of-code-2021/cmd/puzzles/05/util"
	"advent-of-code-2021/pkg/util"
	"bufio"
	"fmt"
	"log"
)

func main() {
	fmt.Printf("Part 1: %5d\n", part1())
	fmt.Printf("Part 2: %5d\n", part2())
}

const CSS = 999 // coordinate system size for x and y-axis

func part1() int {
	file := util.OpenFile("05/input")
	defer file.Close()

	var field [CSS][CSS]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := math.NewLine(scanner.Text())
		markLine(&field, &line, false)
	}

	return countDangerousVents(&field)
}

func part2() int {
	file := util.OpenFile("05/input")
	defer file.Close()

	var field [CSS][CSS]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := math.NewLine(scanner.Text())
		markLine(&field, &line, true)
	}

	return countDangerousVents(&field)
}

func markLine(field *[CSS][CSS]int, line *math.Line, includeDiagonal bool) {
	if line.IsHorizontal() {
		xTravel, yTravel := line.HorizontalTravel()
		markTravel(field, xTravel, yTravel)
	} else if line.IsVertical() {
		xTravel, yTravel := line.VerticalTravel()
		markTravel(field, xTravel, yTravel)
	} else if includeDiagonal {
		xTravel, yTravel := line.DiagonalTravel()
		markTravel(field, xTravel, yTravel)
	}
}

func markTravel(field *[CSS][CSS]int, xTravel, yTravel []int) {
	if len(xTravel) != len(yTravel) {
		log.Fatal("x and y travel are not of the same length")
	}

	for i := 0; i < len(xTravel); i++ {
		x := xTravel[i]
		y := yTravel[i]
		field[y][x]++
	}
}

func countDangerousVents(field *[CSS][CSS]int) (dangerousVents int) {
	for y := 0; y < CSS; y++ {
		for x := 0; x < CSS; x++ {
			if field[y][x] >= 2 {
				dangerousVents++
			}
		}
	}
	return dangerousVents
}
