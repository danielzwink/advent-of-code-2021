package main

import (
	"advent-of-code-2021/pkg/util"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	image, enhancement := readInput("20/input")

	part1Result, part1Duration := part1(image, enhancement)
	fmt.Printf("Part 1: %5d (duration: %s)\n", part1Result, part1Duration)

	part2Result, part2Duration := part2(image, enhancement)
	fmt.Printf("Part 2: %5d (duration: %s)\n", part2Result, part2Duration)
}

func part1(image Image, enhancement ImageEnhancement) (int, time.Duration) {
	start := time.Now()

	for i := 0; i < 2; i++ {
		defaultPixel := i%2 != 0
		image = interpolateImage(image, enhancement, defaultPixel)
	}
	return countLitPixels(image), time.Since(start)
}

func part2(image Image, enhancement ImageEnhancement) (int, time.Duration) {
	start := time.Now()

	for i := 0; i < 50; i++ {
		defaultPixel := i%2 != 0
		image = interpolateImage(image, enhancement, defaultPixel)
	}
	return countLitPixels(image), time.Since(start)
}

func interpolateImage(image Image, enhancement ImageEnhancement, defaultPixel bool) Image {
	lenY := len(image) + 2
	lenX := len(image[0]) + 2

	result := make(Image, lenY)
	for y := 0; y < lenY; y++ {
		result[y] = make(ImageRow, lenX)
		for x := 0; x < lenX; x++ {
			result[y][x] = enhancePixel(x-1, y-1, image, enhancement, defaultPixel)
		}
	}
	return result
}

func enhancePixel(x, y int, image Image, enhancement ImageEnhancement, defaultPixel bool) bool {
	lenY := len(image)
	lenX := len(image[0])

	builder := strings.Builder{}
	for yd := -1; yd <= 1; yd++ {
		for xd := -1; xd <= 1; xd++ {
			yp := y + yd
			xp := x + xd

			if xp >= 0 && xp < lenX && yp >= 0 && yp < lenY {
				pixel := image[yp][xp]
				builder.WriteRune(translate(pixel))
			} else {
				builder.WriteRune(translate(defaultPixel))
			}
		}
	}
	decimal, _ := strconv.ParseUint(builder.String(), 2, builder.Len())
	return enhancement[decimal]
}

func translate(pixel bool) rune {
	if pixel {
		return '1'
	} else {
		return '0'
	}
}

func countLitPixels(image Image) int {
	sum := 0
	for _, row := range image {
		for _, b := range row {
			if b {
				sum++
			}
		}
	}
	return sum
}

type ImageEnhancement []bool
type ImageRow []bool
type Image []ImageRow

func readInput(day string) (Image, ImageEnhancement) {
	lines := util.ReadFile(day)

	enhancement := make(ImageEnhancement, len(lines[0]))
	for i, c := range lines[0] {
		enhancement[i] = c == '#'
	}

	image := make(Image, len(lines)-2)
	for y, row := range lines[2:] {
		image[y] = make(ImageRow, len(row))
		for x, c := range row {
			image[y][x] = c == '#'
		}
	}

	return image, enhancement
}
