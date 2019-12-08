package day7

import (
	"adventofcode2019/input"
	"adventofcode2019/intcode"
	"fmt"
	"log"
	"github.com/gitchander/permutation"
)

const amplifiersNum = 5

func Part1(filename string) {
	instructions := input.ReadIntArrayFromFile(filename)

	a := []int{0, 1, 2, 3, 4}
	variants := permutation.New(permutation.IntSlice(a))

	var nextInput []int64
	var err error
	var max int64

	for variants.Next() {
		instrCopy := make([]int64, len(instructions))
		copy(instrCopy, instructions)

		//fmt.Println("permutation", a)
		for i := 0; i < amplifiersNum; i++ {
			if i == 0 {
				nextInput = []int64{0}
			}
			nextInput = append([]int64{int64(a[i])}, nextInput...)
			//fmt.Println("solve", nextInput)
			nextInput, _, err = intcode.Solve(instrCopy, nextInput, 0)
			if err != nil {
				log.Fatalf("error on A amlifier. Error: %s", err.Error())
			}

			//fmt.Print("output", nextInput)
		}

		if max < nextInput[0] {
			max = nextInput[0]
		}
	}

	fmt.Println(max)
}

func Part2(filename string) {
	instructions := input.ReadIntArrayFromFile(filename)

	a := []int{5, 6, 7, 8, 9}
	variants := permutation.New(permutation.IntSlice(a))

	var nextInput []int64
	var max int64

	for variants.Next() {
		//a = []int{9,7,8,5,6}
		//fmt.Println(a)

		var halt bool
		amplifiers := make([][]int64, amplifiersNum)
		var once bool

		steps := make([]int, amplifiersNum)

		var j int
		//fmt.Println("permutation", a)
		for !halt {
			for i := 0; i < amplifiersNum; i++ {
				j++

				if i == 0 && !once {
					once = true
					nextInput = []int64{0}
				}

				if len(amplifiers[i]) == 0 {
					//fmt.Println("create amplifier", i)
					amplifiers[i] = make([]int64, len(instructions))
					copy(amplifiers[i], instructions)

					nextInput = append([]int64{int64(a[i])}, nextInput...)
				}
				//if j < 20 {
				//	fmt.Println("solve", nextInput)
				//}
				output, step, err := intcode.Solve(amplifiers[i], nextInput, steps[i])
				if err != nil && err != intcode.HaltErr {
					log.Fatalf("error on A amlifier. Error: %s", err.Error())
				}
				//if j < 20 {
				//	fmt.Println("output", output)
				//
				//}

				steps[i] = step

				if err == intcode.HaltErr {
					//fmt.Println("halt")
					halt = true
					//break
				} else {
					nextInput = output
				}

				//fmt.Print("output", nextInput)
			}
		}

		if max < nextInput[0] {
			max = nextInput[0]
		}
	}

	//}

	fmt.Println(max)
}
