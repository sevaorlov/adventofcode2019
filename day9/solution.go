package day9

import (
	"adventofcode2019/input"
	"adventofcode2019/intcode"
	"fmt"
	"log"
)

func Part1(filename string) {
	instructions := input.ReadInt64CommonedArray(filename)
	a := make([]int64, len(instructions)*100)
	a = append(instructions, a...)

	output, _, err := intcode.Solve(a, []int64{1}, 0, false)
	if err != nil && err != intcode.HaltErr {
		log.Fatalf("error solving intcode. Error: %s", err.Error())
	}

	fmt.Println(output[0])
}

func Part2(filename string) {
	instructions := input.ReadInt64CommonedArray(filename)
	a := make([]int64, len(instructions)*100)
	a = append(instructions, a...)

	output, _, err := intcode.Solve(a, []int64{2}, 0, false)
	if err != nil && err != intcode.HaltErr {
		log.Fatalf("error solving intcode. Error: %s", err.Error())
	}

	fmt.Println(output[0])
}
