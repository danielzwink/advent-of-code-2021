package board

import (
	"fmt"
	"strings"
)

const SIZE int = 5
const DRAWN int = -1

type Board struct {
	numbers [SIZE][SIZE]int
}

func (f *Board) Fill(row int, values [SIZE]int) {
	for i, v := range values {
		f.numbers[row][i] = v
	}
	return
}

func (f *Board) Draw(number int) {
	for row := 0; row < SIZE; row++ {
		for column := 0; column < SIZE; column++ {
			if f.numbers[row][column] == number {
				f.numbers[row][column] = DRAWN
			}
		}
	}
	return
}

func (f *Board) Won() bool {
	for row := 0; row < SIZE; row++ {
		sum := 0
		for column := 0; column < SIZE; column++ {
			sum += f.numbers[row][column]
		}
		if sum == SIZE*DRAWN {
			return true
		}
	}
	for column := 0; column < SIZE; column++ {
		sum := 0
		for row := 0; row < SIZE; row++ {
			sum += f.numbers[row][column]
		}
		if sum == SIZE*DRAWN {
			return true
		}
	}

	return false
}

func (f *Board) RemainingSum() (sum int) {
	for row := 0; row < SIZE; row++ {
		for column := 0; column < SIZE; column++ {
			if f.numbers[row][column] != DRAWN {
				sum += f.numbers[row][column]
			}
		}
	}
	return sum
}

func (f *Board) String() string {
	builder := strings.Builder{}
	builder.WriteString("\n")

	for row := 0; row < SIZE; row++ {
		for column := 0; column < SIZE; column++ {
			s := fmt.Sprintf("%3d", f.numbers[row][column])
			builder.WriteString(s)
		}
		builder.WriteString("\n")
	}

	return builder.String()
}
