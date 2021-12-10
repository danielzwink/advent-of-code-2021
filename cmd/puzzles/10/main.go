package main

import (
	types "advent-of-code-2021/cmd/puzzles/10/util"
	"advent-of-code-2021/pkg/util"
	"bufio"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var closingBracketMap = make(map[string]string, 4)
var syntaxErrorScore = make(map[string]int, 4)
var autoCompleteScore = make(map[string]int, 4)

func main() {
	closingBracketMap["("] = ")"
	closingBracketMap["["] = "]"
	closingBracketMap["{"] = "}"
	closingBracketMap["<"] = ">"

	syntaxErrorScore[")"] = 3
	syntaxErrorScore["]"] = 57
	syntaxErrorScore["}"] = 1197
	syntaxErrorScore[">"] = 25137

	autoCompleteScore[")"] = 1
	autoCompleteScore["]"] = 2
	autoCompleteScore["}"] = 3
	autoCompleteScore[">"] = 4

	part1Result, part1Duration := part1()
	fmt.Printf("Part 1: %10d (duration: %s)\n", part1Result, part1Duration)

	part2Result, part2Duration := part2()
	fmt.Printf("Part 2: %10d (duration: %s)\n", part2Result, part2Duration)
}

func part1() (int, time.Duration) {
	file := util.OpenFile("10")
	scanner := bufio.NewScanner(file)
	start := time.Now()

	score := 0
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")

		stack := types.NewStack()
		for i := 0; i < len(line); i++ {
			corruptedChunk := evaluateChunk(line[i], stack)

			if len(corruptedChunk) == 1 {
				score += syntaxErrorScore[corruptedChunk]
				break
			}
		}
	}

	return score, time.Since(start)
}

func part2() (int, time.Duration) {
	file := util.OpenFile("10")
	scanner := bufio.NewScanner(file)
	start := time.Now()

	autoCompleteScores := make([]int, 0)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")

		stack := types.NewStack()
		corrupted := false
		for i := 0; i < len(line); i++ {
			corruptChunk := evaluateChunk(line[i], stack)

			if len(corruptChunk) == 1 {
				corrupted = true
				break
			}
		}

		remainingStackSize := stack.Len()
		if !corrupted && remainingStackSize > 0 {
			score := 0
			for i := 0; i < remainingStackSize; i++ {
				chunk := stack.Pop().Closing
				score = (5 * score) + autoCompleteScore[chunk]
			}
			autoCompleteScores = append(autoCompleteScores, score)
		}
	}

	sort.Ints(autoCompleteScores)
	middleIndex := (len(autoCompleteScores) - 1) / 2
	return autoCompleteScores[middleIndex], time.Since(start)
}

func evaluateChunk(char string, stack *types.Stack) string {
	if strings.Contains("<{[(", char) {
		chunk := types.Chunk{Closing: closingBracketMap[char]}
		stack.Push(chunk)
		return ""

	} else if strings.Contains(")]}>", char) {
		chunk := stack.Pop()
		if char != chunk.Closing {
			return char
		}
		return ""
	}

	log.Fatalf("Invalid character: %v", char)
	return ""
}
