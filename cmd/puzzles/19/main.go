package main

import (
	"advent-of-code-2021/cmd/puzzles/19/types"
	"advent-of-code-2021/pkg/util"
	"fmt"
	"strings"
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

	chaos := readUniverse("19")
	aligned := alignUniverse(chaos)

	uniqueBeacons := make(map[string]bool, 300)
	for a := 0; a < aligned.Len(); a++ {
		scanner := aligned.Get(a)

		for i := 0; i < scanner.Len(); i++ {
			uniqueBeacons[scanner.NormalizedBeacon(i).String()] = true
		}
	}
	return len(uniqueBeacons), time.Since(start)
}

func part2() (int, time.Duration) {
	start := time.Now()

	universe := readUniverse("19")
	aligned := alignUniverse(universe)

	maxDistance := 0
	for a1 := 0; a1 < aligned.Len(); a1++ {
		for a2 := 0; a2 < aligned.Len(); a2++ {
			if a1 == a2 {
				continue
			}
			l1 := aligned.Get(a1).Location
			l2 := aligned.Get(a2).Location
			distance := l1.ManhattanDistance(l2)
			if distance > maxDistance {
				maxDistance = distance
			}
		}
	}
	return maxDistance, time.Since(start)
}

func alignUniverse(chaos *types.Universe) *types.Universe {
	first := chaos.Top()
	aligned := types.NewUniverse()
	aligned.Add(first)

	for chaos.Len() > 0 {
		proceed := true
		for c := 0; proceed && c < chaos.Len(); c++ {
			for a := 0; proceed && a < aligned.Len(); a++ {
				scanner := chaos.Get(c)
				base := aligned.Get(a)
				if matchAndRelocate(base, scanner) {
					aligned.Add(scanner)
					chaos.Del(c)
					proceed = false
				}
			}
		}
	}
	return aligned
}

func matchAndRelocate(base *types.Scanner, scanner *types.Scanner) bool {
	const RequiredMatchingBeacons = 12

	baseBeacons := base.NormalizedBeacons()
	for _, baseBeacon := range baseBeacons {
		for s := 0; s < scanner.Len(); s++ {
			for orientation := types.XpYpZp; orientation <= types.ZpYmXp; orientation++ {
				scanner.Calibrate(s, orientation, baseBeacon)

				search := scanner.NormalizedBeacons()
				matched := make([]types.Beacon, 0, 12)
				for _, reference := range baseBeacons {
					for _, lookup := range search {
						if reference.Equal(lookup) {
							matched = append(matched, lookup)
							break
						}
					}
				}
				if len(matched) == RequiredMatchingBeacons {
					return true
				}
			}
		}
	}
	return false
}

func readUniverse(day string) *types.Universe {
	lines := util.ReadFile(day)

	var scanner *types.Scanner
	universe := types.NewUniverse()
	for _, line := range lines {
		if strings.HasPrefix(line, "--- scanner") {
			scanner = types.NewScanner()
			universe.Add(scanner)
		} else if line == "" {
			continue
		} else {
			scanner.Add(types.NewBeacon(line))
		}
	}
	return universe
}
