package day16

import (
	"adventofcode2019/input"
	"adventofcode2019/logger"
	"strconv"
)

const stopPhase = 100

var basePattern = []int{0, 1, 0, -1}

func newItem(input []int, pattern []int) int {
	sum := 0

	for index, _ := range input {
		sum += input[index]*pattern[index]
	}

	if sum < 0 {
		sum *= -1
	}
	return sum % 10
}

func newPattern(input []int, index int) []int {
	pattern := make([]int, len(input))
	i := 0
	j := 0

	for {
		times := index + 1
		if j == 0 && i == 0 {
			times -= 1
		}

		for k:=0;k<times;k++ {
			pattern[j] = basePattern[i]
			j++
			if j >= len(input) {
				return pattern
			}
		}

		i++
		if i >= len(basePattern) {
			i = 0
		}
	}
}

func newPhase(input []int) []int {
	output := make([]int, len(input))

	for index, _ := range output {
		pattern := newPattern(input, index)
		//fmt.Println("pattern")
		//fmt.Println("pattern", pattern)
		output[index] = newItem(input, pattern)
	}

	return output
}

func intArrayToString(items []int) string {
	var res string
	for _, item := range items {
		res += strconv.Itoa(item)
	}
	return res
}

func Part1(filename string) string {
	items := input.ReadIntArray(filename, "")
	logger.Debug(items)

	i := 0
	for {
		i++

		items = newPhase(items)
		logger.Debug(items)

		if i == stopPhase {
			break
		}
	}

	return intArrayToString(items[0:8])
}

func Part2(filename string) string {
	return ""
}
