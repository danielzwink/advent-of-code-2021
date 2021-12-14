package main

import (
	"advent-of-code-2021/pkg/util"
	"fmt"
	"regexp"
	"sort"
	"time"
)

func main() {
	initialPolymer, insertionRules := getInitialPolymerAndInsertionRules()

	part1Result, part1Duration := part1(initialPolymer, insertionRules)
	fmt.Printf("Part 1: %14d (duration: %s)\n", part1Result, part1Duration)

	part2Result, part2Duration := part2(initialPolymer, insertionRules)
	fmt.Printf("Part 2: %14d (duration: %s)\n", part2Result, part2Duration)
}

func part1(initialPolymer polymer, insertionRules rules) (int, time.Duration) {
	start := time.Now()
	quantities := performPolymerEvolution(initialPolymer, 10, insertionRules)
	min, max := getMinMaxCounts(quantities)
	return max - min, time.Since(start)
}

func part2(initialPolymer polymer, insertionRules rules) (int, time.Duration) {
	start := time.Now()
	quantities := performPolymerEvolution(initialPolymer, 40, insertionRules)
	min, max := getMinMaxCounts(quantities)
	return max - min, time.Since(start)
}

func performPolymerEvolution(initialPolymer polymer, steps int, insertionRules rules) map[rune]int {
	finalPolymer := initialPolymer
	for i := 0; i < steps; i++ {
		tempPolymer := make(polymer, len(insertionRules)+2)
		for pattern, count := range finalPolymer {
			if len(pattern) == 1 {
				tempPolymer[pattern] += count
			} else if len(pattern) == 2 {
				insertions := insertionRules[pattern]
				for _, insertion := range insertions {
					tempPolymer[insertion] += count
				}
			}
		}
		finalPolymer = tempPolymer
	}

	polymerQuantities := make(map[rune]int, 0)
	for pattern, count := range finalPolymer {
		for _, r := range pattern {
			polymerQuantities[r] += count
		}
	}
	result := make(map[rune]int, len(polymerQuantities))
	for r, count := range polymerQuantities {
		result[r] = count / 2
	}
	return result
}

func getMinMaxCounts(polymerQuantities map[rune]int) (int, int) {
	counts := make([]int, 0, len(polymerQuantities))
	for _, count := range polymerQuantities {
		counts = append(counts, count)
	}
	sort.Ints(counts)
	return counts[0], counts[len(counts)-1]
}

type rules map[string][]string
type polymer map[string]int

var insertionRulePattern, _ = regexp.Compile("([A-Z]{2}) -> ([A-Z])")

func getInitialPolymerAndInsertionRules() (polymer, rules) {
	lines := util.ReadFile("14")

	template := lines[0]
	length := len(template)
	initialPolymer := make(polymer)
	for i := 0; i < length-1; i++ {
		pattern := template[i : i+2]
		initialPolymer[pattern]++
	}
	first := template[0:1]
	initialPolymer[first]++
	last := template[length-1 : length]
	initialPolymer[last]++

	insertionRules := make(rules, len(lines)-2)
	for _, line := range lines[2:] {
		match := insertionRulePattern.FindStringSubmatch(line)
		pair := match[1]
		char1 := string(pair[0])
		char2 := string(pair[1])
		insertion := match[2]

		insertionRules[pair] = []string{char1 + insertion, insertion + char2}
	}
	return initialPolymer, insertionRules
}
