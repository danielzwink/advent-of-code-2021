package main

import (
	"advent-of-code-2021/pkg/util"
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %10d\n", part1())
	fmt.Printf("Part 2: %10d\n", part2())
}

func part1() int {
	file := util.OpenFile("02/input")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var position, depth int

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		command := line[0]
		change, _ := strconv.Atoi(line[1])

		switch command {
		case "forward":
			position += change
		case "down":
			depth += change
		case "up":
			depth -= change
		}

	}

	return position * depth
}

func part2() int {
	file := util.OpenFile("02/input")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var position, depth, aim int

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		command := line[0]
		change, _ := strconv.Atoi(line[1])

		switch command {
		case "forward":
			position += change
			depth += aim * change
		case "down":
			aim += change
		case "up":
			aim -= change
		}

	}

	return position * depth
}
