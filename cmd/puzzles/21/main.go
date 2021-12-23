package main

import (
	"advent-of-code-2021/cmd/puzzles/21/types"
	"advent-of-code-2021/pkg/util"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func main() {
	part1Result, part1Duration := part1()
	fmt.Printf("Part 1: %6d (duration: %s)\n", part1Result, part1Duration)

	part2Result, part2Duration := part2()
	fmt.Printf("Part 2: %6d (duration: %s)\n", part2Result, part2Duration)
}

func part1() (int, time.Duration) {
	start := time.Now()

	player1, player2 := readInput("21", 1000)
	die := types.DeterministicDie{}

	for true {
		value := die.RollThreeTimes()

		if die.Rolls%2 != 0 {
			player1.Move(value)
		} else {
			player2.Move(value)
		}

		if player1.Won() {
			return player2.Score * die.Rolls, time.Since(start)
		} else if player2.Won() {
			return player1.Score * die.Rolls, time.Since(start)
		}
	}

	return 0, time.Since(start)
}

func part2() (int, time.Duration) {
	start := time.Now()
	return 0, time.Since(start)
}

func readInput(day string, winning int) (types.Player, types.Player) {
	lines := util.ReadFile(day)

	playerRegexp, _ := regexp.Compile("Player ([0-9]) starting position: ([0-9]+)")
	match1 := playerRegexp.FindStringSubmatch(lines[0])
	match2 := playerRegexp.FindStringSubmatch(lines[1])

	position1, _ := strconv.Atoi(match1[2])
	position2, _ := strconv.Atoi(match2[2])
	return types.NewPlayer(winning, position1), types.NewPlayer(winning, position2)
}
