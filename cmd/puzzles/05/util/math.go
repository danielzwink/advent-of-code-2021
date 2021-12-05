package math

import (
	"strconv"
	"strings"
)

type Travel []int

type Coordinate struct {
	X, Y int
}

func NewCoordinate(input string) (c Coordinate) {
	coordinates := strings.Split(input, ",")
	c.X, _ = strconv.Atoi(coordinates[0])
	c.Y, _ = strconv.Atoi(coordinates[1])
	return c
}

type Line struct {
	Start, End Coordinate
}

func NewLine(input string) (l Line) {
	coordinates := strings.Split(input, " -> ")
	l.Start = NewCoordinate(coordinates[0])
	l.End = NewCoordinate(coordinates[1])
	return l
}

func (l Line) IsHorizontal() bool {
	return l.Start.Y == l.End.Y
}

func (l Line) IsVertical() bool {
	return l.Start.X == l.End.X
}

func (l Line) HorizontalTravel() (x, y Travel) {
	x = steps(l.Start.X, l.End.X)
	y = constant(l.Start.Y, len(x))
	return x, y
}

func (l Line) VerticalTravel() (x, y Travel) {
	y = steps(l.Start.Y, l.End.Y)
	x = constant(l.Start.X, len(y))
	return x, y
}

func (l Line) DiagonalTravel() (x, y Travel) {
	x = steps(l.Start.X, l.End.X)
	y = steps(l.Start.Y, l.End.Y)
	return x, y
}

func steps(from, to int) Travel {
	if from < to {
		travel := make(Travel, 0, to-from+1)

		for i := from; i <= to; i++ {
			travel = append(travel, i)
		}
		return travel
	} else {
		travel := make(Travel, 0, from-to+1)

		for i := from; i >= to; i-- {
			travel = append(travel, i)
		}
		return travel
	}
}

func constant(value, size int) Travel {
	travel := make(Travel, 0, size)

	for i := 1; i <= size; i++ {
		travel = append(travel, value)
	}
	return travel
}
