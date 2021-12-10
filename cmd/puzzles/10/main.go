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

func main() {
	closingBracketMap["("] = ")"
	closingBracketMap["["] = "]"
	closingBracketMap["{"] = "}"
	closingBracketMap["<"] = ">"

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
				score += syntaxErrorScore(corruptedChunk)
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
				chunk := stack.Pop()
				score = (5 * score) + autoCompleteScore(chunk.Closing)
			}
			autoCompleteScores = append(autoCompleteScores, score)
		}
	}

	sort.Ints(autoCompleteScores)
	middleIndex := (len(autoCompleteScores) - 1) / 2
	return autoCompleteScores[middleIndex], time.Since(start)
}

func evaluateChunk(char string, stack *types.Stack) string {
	switch char {
	case "(":
		fallthrough
	case "[":
		fallthrough
	case "{":
		fallthrough
	case "<":
		chunk := types.Chunk{Closing: closingBracketMap[char]}
		stack.Push(chunk)
		return ""
	case ")":
		fallthrough
	case "]":
		fallthrough
	case "}":
		fallthrough
	case ">":
		chunk := stack.Pop()
		return checkChunk(char, chunk)
	}

	log.Fatalf("Invalid character: %v", char)
	return ""
}

func checkChunk(givenChar string, expectedChunk types.Chunk) string {
	if givenChar == expectedChunk.Closing {
		return ""
	} else {
		return givenChar
	}
}

func syntaxErrorScore(chunk string) int {
	switch chunk {
	case ")":
		return 3
	case "]":
		return 57
	case "}":
		return 1197
	case ">":
		return 25137
	}

	log.Fatalf("Invalid character: %v", chunk)
	return 0
}

func autoCompleteScore(chunk string) int {
	switch chunk {
	case ")":
		return 1
	case "]":
		return 2
	case "}":
		return 3
	case ">":
		return 4
	}

	log.Fatalf("Invalid character: %v", chunk)
	return 0
}
