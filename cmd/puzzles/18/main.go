package main

import (
	types "advent-of-code-2021/cmd/puzzles/18/util"
	"advent-of-code-2021/pkg/util"
	"fmt"
	"strconv"
	"time"
)

func main() {
	part1Result, part1Duration := part1()
	fmt.Printf("Part 1: %4d (duration: %s)\n", part1Result, part1Duration)

	part2Result, part2Duration := part2()
	fmt.Printf("Part 2: %4d (duration: %s)\n", part2Result, part2Duration)
}

func part1() (int, time.Duration) {
	start := time.Now()
	homework := readHomework("18")

	sum := homework[0]
	for _, pair := range homework[1:] {
		sum = sum.Add(pair).Reduce()
	}
	return sum.Magnitude(), time.Since(start)
}

func part2() (int, time.Duration) {
	start := time.Now()
	homework := readHomework("18")

	max := 0
	for _, pair1 := range homework {
		for _, pair2 := range homework {
			if pair1 != pair2 {
				copy1 := pair1.Copy()
				assignParent(copy1, nil)
				copy2 := pair2.Copy()
				assignParent(copy2, nil)

				magnitude := copy1.Add(copy2).Reduce().Magnitude()
				if max < magnitude {
					max = magnitude
				}
			}
		}
	}
	return max, time.Since(start)
}

func readHomework(day string) []*types.Pair {
	lines := util.ReadFile(day)

	pairs := make([]*types.Pair, 0, len(lines))
	for _, line := range lines {
		pair, _ := readPair(line)
		assignParent(pair, nil)
		pairs = append(pairs, pair)
	}
	return pairs
}

func assignParent(pair *types.Pair, parent *types.Pair) {
	pair.Parent = parent
	if pair.Left.IsPair() {
		assignParent(pair.Left.Pair, pair)
	}
	if pair.Right.IsPair() {
		assignParent(pair.Right.Pair, pair)
	}
}

func readPair(input string) (*types.Pair, int) {
	pos := 1 // move over '['
	left, leftLength := readElement(input[pos:])
	pos += leftLength
	pos++ // move over ','
	right, rightLength := readElement(input[pos:])
	pos += rightLength
	pos++ // move over ']'
	return &types.Pair{Left: left, Right: right}, pos
}

func readElement(input string) (*types.Element, int) {
	r := input[0]
	switch r {
	case '[':
		pair, length := readPair(input)
		return &types.Element{Pair: pair}, length
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		number, _ := strconv.Atoi(string(r))
		return &types.Element{Number: number}, 1
	}
	panic("did not expect " + string(r))
}
