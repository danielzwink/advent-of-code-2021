package main

import (
	field "advent-of-code-2021/cmd/puzzles/04/bingo"
	"advent-of-code-2021/pkg/util"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

func part1() int {
	lines := util.ReadFile("04/input")
	drawings := convertLineToDrawings(lines[0])
	boards := convertLinesToBoards(lines[1:])

	for _, drawing := range drawings {
		for _, board := range boards {
			board.Draw(drawing)

			if board.Won() {
				return drawing * board.RemainingSum()
			}
		}
	}

	log.Fatal("no board has won")
	return -1
}

func part2() int {
	lines := util.ReadFile("04/input")
	drawings := convertLineToDrawings(lines[0])
	boards := convertLinesToBoards(lines[1:])

	for _, drawing := range drawings {
		for _, board := range boards {
			board.Draw(drawing)
		}

		if len(boards) == 1 && boards[0].Won() {
			return drawing * boards[0].RemainingSum()
		}

		boards = removeBoardsThatWon(boards)
	}

	log.Fatal("no board has won")
	return -1
}

func convertLineToDrawings(line string) (numbers []int) {
	values := strings.Split(line, ",")

	for _, v := range values {
		n, _ := strconv.Atoi(v)
		numbers = append(numbers, n)
	}

	return numbers
}

func convertLinesToBoards(lines []string) (boards []*field.Board) {
	var row int
	var board *field.Board

	for _, line := range lines {
		if len(line) == 0 {
			row = 0
			board = &field.Board{}
			continue
		}

		numbers := convertLineToNumbers(line)
		board.Fill(row, numbers)
		row++

		if row == field.SIZE {
			boards = append(boards, board)
		}
	}

	return boards
}

func convertLineToNumbers(line string) (numbers [field.SIZE]int) {
	values := strings.Fields(line)

	for i, v := range values {
		n, _ := strconv.Atoi(v)
		numbers[i] = n
	}

	return numbers
}

func removeBoardsThatWon(boards []*field.Board) (remainingBoards []*field.Board) {
	for _, board := range boards {
		if !board.Won() {
			remainingBoards = append(remainingBoards, board)
		}
	}
	return remainingBoards
}
