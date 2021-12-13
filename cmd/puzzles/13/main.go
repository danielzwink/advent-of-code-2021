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
	matrix, instructions := getMatrixAndFoldInstructions()

	part1Result, part1Duration := part1(matrix, instructions)
	fmt.Printf("Part 1: %5d (duration: %s)\n", part1Result, part1Duration)

	part2Result, part2Duration := part2(matrix, instructions)
	fmt.Printf("Part 2: %5d (duration: %s)\n", part2Result, part2Duration)
}

func part1(matrix [][]bool, instructions []folding) (int, time.Duration) {
	start := time.Now()

	matrix = foldPaper(matrix, instructions[0])
	marks := countPaperDots(matrix)
	return marks, time.Since(start)
}

func part2(matrix [][]bool, instructions []folding) (int, time.Duration) {
	start := time.Now()

	for _, instruction := range instructions {
		matrix = foldPaper(matrix, instruction)
	}
	printPaper(matrix)
	return 0, time.Since(start)
}

func foldPaper(matrix [][]bool, instruction folding) (result [][]bool) {
	lenX := len(matrix[0])
	lenY := len(matrix)

	// fold at x line
	if instruction.x > 0 && instruction.y == -1 {
		result = initMatrix(instruction.x, lenY)

		for y := 0; y < lenY; y++ {
			for x := 0; x < instruction.x; x++ {
				result[y][x] = matrix[y][x] || matrix[y][lenX-1-x]
			}
		}

		// fold at y line
	} else if instruction.x == -1 && instruction.y > 0 {
		result = initMatrix(lenX, instruction.y)

		for y := 0; y < instruction.y; y++ {
			for x := 0; x < lenX; x++ {
				result[y][x] = matrix[y][x] || matrix[lenY-1-y][x]
			}
		}
	}
	return result
}

func countPaperDots(matrix [][]bool) (result int) {
	for _, row := range matrix {
		for _, marked := range row {
			if marked {
				result++
			}
		}
	}
	return result
}

func printPaper(matrix [][]bool) {
	println("")
	for _, row := range matrix {
		for _, marked := range row {
			if marked {
				fmt.Print("# ")
			} else {
				fmt.Print(". ")
			}
		}
		println("")
	}
	println("")
}

type coordinate struct {
	x, y int
}

type folding struct {
	x, y int
}

var foldInstructionPattern, _ = regexp.Compile("fold along ([xy])=([0-9]+)")

func getMatrixAndFoldInstructions() ([][]bool, []folding) {
	lines := util.ReadFile("13")

	maxX, maxY := 0, 0
	coordinates := make([]coordinate, 0, len(lines))
	instructions := make([]folding, 0)
	for _, line := range lines {
		if line == "" {
			continue
		} else if strings.HasPrefix(line, "fold") {
			result := foldInstructionPattern.FindStringSubmatch(line)
			value, _ := strconv.Atoi(result[2])
			if result[1] == "x" {
				instructions = append(instructions, folding{x: value, y: -1})
			} else if result[1] == "y" {
				instructions = append(instructions, folding{x: -1, y: value})
			}
		} else {
			pair := strings.Split(line, ",")
			x, _ := strconv.Atoi(pair[0])
			y, _ := strconv.Atoi(pair[1])
			coordinates = append(coordinates, coordinate{x: x, y: y})
			if maxX < x {
				maxX = x
			}
			if maxY < y {
				maxY = y
			}
		}
	}

	matrix := initMatrix(maxX+1, maxY+1)
	for _, c := range coordinates {
		matrix[c.y][c.x] = true
	}
	return matrix, instructions
}

func initMatrix(lenX int, lenY int) [][]bool {
	matrix := make([][]bool, lenY)
	for y := 0; y < lenY; y++ {
		matrix[y] = make([]bool, lenX)
	}
	return matrix
}
