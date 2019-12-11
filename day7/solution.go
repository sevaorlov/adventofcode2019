package day7

import (
	"adventofcode2019/input"
	"adventofcode2019/intcode"
	"adventofcode2019/permutations"
	"log"
	"strconv"
)

const amplifiersNum = 5

func Part1(filename string) string {
	instructions := input.ReadInt64CommonedArray(filename)
	variants := permutations.GenerateInt([]int{0, 1, 2, 3, 4})

	var nextInput []int64
	var err error
	var max int64

	for _, a := range variants {
		instrCopy := make([]int64, len(instructions))
		copy(instrCopy, instructions)

		for i := 0; i < amplifiersNum; i++ {
			if i == 0 {
				nextInput = []int64{0}
			}
			nextInput = append([]int64{int64(a[i])}, nextInput...)
			nextInput, _, _, err = intcode.Solve(instrCopy, nextInput, 0, 0, true)
			if err != nil {
				log.Fatalf("error on A amlifier. Error: %s", err.Error())
			}
		}

		if max < nextInput[0] {
			max = nextInput[0]
		}
	}

	return strconv.FormatInt(max, 10)
}

func Part2(filename string) string {
	instructions := input.ReadInt64CommonedArray(filename)
	variants := permutations.GenerateInt([]int{5, 6, 7, 8, 9})

	var nextInput []int64
	var max int64

	for _, a := range variants {
		var halt bool
		amplifiers := make([][]int64, amplifiersNum)
		var once bool

		steps := make([]int, amplifiersNum)

		var j int
		for !halt {
			for i := 0; i < amplifiersNum; i++ {
				j++

				if i == 0 && !once {
					once = true
					nextInput = []int64{0}
				}

				if len(amplifiers[i]) == 0 {
					amplifiers[i] = make([]int64, len(instructions))
					copy(amplifiers[i], instructions)

					nextInput = append([]int64{int64(a[i])}, nextInput...)
				}
				output, step, _, err := intcode.Solve(amplifiers[i], nextInput, steps[i], 0, true)
				if err != nil {
					log.Fatalf("error on A amlifier. Error: %s", err.Error())
				}

				steps[i] = step

				if step == -1 {
					halt = true
				} else {
					nextInput = output
				}
			}
		}

		if max < nextInput[0] {
			max = nextInput[0]
		}
	}

	return strconv.FormatInt(max, 10)
}
