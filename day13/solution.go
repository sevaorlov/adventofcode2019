package day13

import (
	"adventofcode2019/input"
	"adventofcode2019/intcode"
	"log"
	"strconv"
)

type Point struct {
	x, y int
}

type Arcade struct {
	index, relBase int
	code           []int64
	finished       bool
}

func (a *Arcade) run(input int64) int64 {
	var err error
	var output []int64

	if a.finished {
		return 0
	}
	output, a.index, a.relBase, err = intcode.Solve(a.code, []int64{input}, a.index, a.relBase, true)
	if err != nil {
		log.Fatal("failed to solve intcode", err.Error())
	}

	if a.index == -1 {
		a.finished = true
		return 0
	}
	return output[0]

}

func Part1(filename string) string {
	code := inputInstructions(filename)

	output, _, _, err := intcode.Solve(code, []int64{}, 0, 0, false)
	if err != nil {
		log.Fatal("failed to solve intcode", err.Error())
	}

	count := 0
	i := 2

	for i < len(output) {
		if output[i] == 2 {
			count++
		}
		i += 3
	}

	return strconv.Itoa(count)
}

func Part2(filename string) string {
	code := inputInstructions(filename)

	// set it to 2 to play for free
	code[0] = 2

	var score, input int64
	var ball *Point
	var paddle *Point

	arcade := Arcade{code: code}

	for {
		if ball != nil && paddle != nil {
			diff := ball.x - paddle.x
			if diff > 0 {
				input = 1
			} else if diff < 0 {
				input = -1
			} else if diff == 0 {
				input = 0
			}
		}

		x := arcade.run(input)
		y := arcade.run(input)
		tileId := arcade.run(input)

		if arcade.finished {
			break
		}

		if x == -1 && y == 0 {
			score = tileId
		} else if tileId == 3 {
			paddle = &Point{int(x), int(y)}
		} else if tileId == 4 {
			ball = &Point{int(x), int(y)}
		}
	}

	return strconv.FormatInt(score, 10)
}

func inputInstructions(filename string) []int64 {
	instructions := input.ReadInt64CommonedArray(filename)

	code := make([]int64, len(instructions)*100)
	code = append(instructions, code...)

	return code
}
