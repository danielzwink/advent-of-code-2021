package main

import (
	"advent-of-code-2021/pkg/util"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	file := util.OpenFile("09")
	matrix := readMatrix(file)

	part1Result, part1Duration := part1(matrix)
	fmt.Printf("Part 1: %6d (duration: %s)\n", part1Result, part1Duration)

	part2Result, part2Duration := part2(matrix)
	fmt.Printf("Part 2: %6d (duration: %s)\n", part2Result, part2Duration)
}

const ReservedSize = 100

type location struct {
	row, col, height int
	visited          bool
}

func readMatrix(file *os.File) [][]*location {
	matrix := make([][]*location, 0, ReservedSize)
	scanner := bufio.NewScanner(file)

	// read file and create locations with given height
	for scanner.Scan() {
		rowValues := strings.Split(scanner.Text(), "")
		row := make([]*location, 0, ReservedSize)
		for _, rowValue := range rowValues {
			value, _ := strconv.Atoi(rowValue)
			row = append(row, &location{height: value})
		}
		matrix = append(matrix, row)
	}

	// add indexes to the created locations
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			matrix[row][col].row = row
			matrix[row][col].col = col
		}
	}

	return matrix
}

func part1(matrix [][]*location) (int, time.Duration) {
	start := time.Now()

	riskLevelSum := 0
	for _, row := range matrix {
		for _, location := range row {
			lowPoint := isLowPoint(location, matrix)
			if lowPoint {
				riskLevelSum += location.height + 1
			}
		}
	}
	return riskLevelSum, time.Since(start)
}

func part2(matrix [][]*location) (int, time.Duration) {
	start := time.Now()

	var basinSizes []int
	for _, row := range matrix {
		for _, location := range row {
			lowPoint := isLowPoint(location, matrix)
			if lowPoint {
				basinSize := getBasinSize(location, matrix)
				basinSizes = append(basinSizes, basinSize)
			}
		}
	}

	sort.Ints(basinSizes)
	length := len(basinSizes)
	result := basinSizes[length-1] * basinSizes[length-2] * basinSizes[length-3]
	return result, time.Since(start)
}

func isLowPoint(current *location, matrix [][]*location) bool {
	if current.height == 9 {
		return false
	}

	row := current.row
	col := current.col

	if isNeighbourEqualOrSmaller(row, col-1, current.height, matrix) {
		return false
	}
	if isNeighbourEqualOrSmaller(row, col+1, current.height, matrix) {
		return false
	}
	if isNeighbourEqualOrSmaller(row-1, col, current.height, matrix) {
		return false
	}
	if isNeighbourEqualOrSmaller(row+1, col, current.height, matrix) {
		return false
	}
	return true
}

func getBasinSize(current *location, matrix [][]*location) int {
	basinSize := 0

	newNeighbours := make([]*location, 1)
	newNeighbours[0] = current
	current.visited = true

	for len(newNeighbours) > 0 {
		unvisitedNeighbours := newNeighbours
		newNeighbours = make([]*location, 0)

		for _, n := range unvisitedNeighbours {
			basinSize++

			if isNeighbourHigherAndUnvisited(n.row, n.col-1, n.height, matrix) {
				neighbour := matrix[n.row][n.col-1]
				neighbour.visited = true
				newNeighbours = append(newNeighbours, neighbour)
			}
			if isNeighbourHigherAndUnvisited(n.row, n.col+1, n.height, matrix) {
				neighbour := matrix[n.row][n.col+1]
				neighbour.visited = true
				newNeighbours = append(newNeighbours, neighbour)
			}
			if isNeighbourHigherAndUnvisited(n.row-1, n.col, n.height, matrix) {
				neighbour := matrix[n.row-1][n.col]
				neighbour.visited = true
				newNeighbours = append(newNeighbours, neighbour)
			}
			if isNeighbourHigherAndUnvisited(n.row+1, n.col, n.height, matrix) {
				neighbour := matrix[n.row+1][n.col]
				neighbour.visited = true
				newNeighbours = append(newNeighbours, neighbour)
			}
		}
	}

	return basinSize
}

func isNeighbourEqualOrSmaller(neighbourRow, neighbourCol, height int, matrix [][]*location) bool {
	if !isValidLocation(neighbourRow, neighbourCol, matrix) {
		return false
	}

	return matrix[neighbourRow][neighbourCol].height <= height
}

func isNeighbourHigherAndUnvisited(neighbourRow, neighbourCol, height int, matrix [][]*location) bool {
	if !isValidLocation(neighbourRow, neighbourCol, matrix) {
		return false
	}

	neighbour := matrix[neighbourRow][neighbourCol]
	return neighbour.height < 9 && neighbour.height > height && !neighbour.visited
}

func isValidLocation(row, col int, matrix [][]*location) bool {
	maxRows := len(matrix)
	maxCols := len(matrix[0])

	return row >= 0 && row < maxRows && col >= 0 && col < maxCols
}
