package main

import (
	"advent-of-code-2021/pkg/util"
	"bufio"
	"fmt"
	"regexp"
	"strings"
	"time"
)

var bigCavePattern, _ = regexp.Compile("[A-Z]+")

func main() {
	caveMap := readCaveMap()

	part1Result, part1Duration := part1(caveMap)
	fmt.Printf("Part 1: %6d (duration: %s)\n", part1Result, part1Duration)

	part2Result, part2Duration := part2(caveMap)
	fmt.Printf("Part 2: %6d (duration: %s)\n", part2Result, part2Duration)
}

func part1(caveMap map[string][]string) (int, time.Duration) {
	start := time.Now()
	paths := travel([]string{}, "start", false, caveMap)
	return len(paths), time.Since(start)
}

func part2(caveMap map[string][]string) (int, time.Duration) {
	start := time.Now()
	paths := travel([]string{}, "start", true, caveMap)
	return len(paths), time.Since(start)
}

func travel(previous []string, current string, singleSmallCavePossibleTwice bool, caveMap map[string][]string) [][]string {
	var paths [][]string

	if current == "end" {
		paths = append(paths, []string{current})

	} else {
		targets, found := caveMap[current]

		if found {
			previous = append(previous, current)

			for _, next := range targets {
				bigCaveOrNotBeenVisited := isBigCaveOrNotBeenVisited(next, previous)
				smallCavePossibleTwice := singleSmallCavePossibleTwice

				if bigCaveOrNotBeenVisited || smallCavePossibleTwice {
					if !bigCaveOrNotBeenVisited && smallCavePossibleTwice {
						smallCavePossibleTwice = false
					}
					result := travel(previous, next, smallCavePossibleTwice, caveMap)
					for _, path := range result {
						paths = append(paths, append([]string{current}, path...))
					}
				}
			}
		}
	}
	return paths
}

func isBigCaveOrNotBeenVisited(current string, previous []string) bool {
	if bigCavePattern.MatchString(current) {
		return true
	}
	for _, p := range previous {
		if p == current {
			return false
		}
	}
	return true
}

func readCaveMap() map[string][]string {
	file := util.OpenFile("12")
	defer file.Close()

	caveMap := make(map[string][]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		path := strings.Split(scanner.Text(), "-")
		node1 := path[0]
		node2 := path[1]

		if node1 == "start" || node2 == "end" {
			addPath(node1, node2, caveMap)
		} else if node2 == "start" || node1 == "end" {
			addPath(node2, node1, caveMap)
		} else {
			addPath(node1, node2, caveMap)
			addPath(node2, node1, caveMap)
		}
	}
	return caveMap
}

func addPath(start string, end string, caveMap map[string][]string) {
	content, found := caveMap[start]
	if found {
		content = append(content, end)
	} else {
		content = make([]string, 1)
		content[0] = end
	}
	caveMap[start] = content
}
