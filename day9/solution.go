package day9

import (
	"adventofcode2019/input"
	"adventofcode2019/intcode"
	"log"
	"strconv"
)

func Part1(filename string) string {
	instructions := input.ReadInt64CommonedArray(filename)
	a := make([]int64, len(instructions)*100)
	a = append(instructions, a...)

	output, _, _, err := intcode.Solve(a, []int64{1}, 0, 0, false)
	if err != nil {
		log.Fatalf("error solving intcode. Error: %s", err.Error())
	}

	return strconv.FormatInt(output[0], 10)
}

func Part2(filename string) string {
	instructions := input.ReadInt64CommonedArray(filename)
	a := make([]int64, len(instructions)*100)
	a = append(instructions, a...)

	output, _, _, err := intcode.Solve(a, []int64{2}, 0, 0, false)
	if err != nil {
		log.Fatalf("error solving intcode. Error: %s", err.Error())
	}

	return strconv.FormatInt(output[0], 10)
}
