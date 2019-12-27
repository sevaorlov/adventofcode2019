package day16

import (
	"adventofcode2019/input"
	"adventofcode2019/logger"
	"log"
	"strconv"
)

const stopPhase = 100

var basePattern = []int{0, 1, 0, -1}

func newItem(input []int, pattern []int) int {
	sum := 0

	for index, _ := range input {
		sum += input[index] * pattern[index]
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

		for k := 0; k < times; k++ {
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

// this is slower
//func patternForItem(cursorElementIndex, index int) int {
//	patternLength := len(basePattern) * (cursorElementIndex + 1)
//	i := index % patternLength + 1
//	j := (i / (cursorElementIndex + 1)) % len(basePattern)
//
//	return basePattern[j]
//}

func newPhase(input []int) []int {
	output := make([]int, len(input))

	for index, _ := range output {
		pattern := newPattern(input, index)
		//fmt.Println("pattern for", index, pattern)
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

func repeatArray(items []int, times int) []int {
	new := make([]int, len(items)*times)

	for i := 0; i < len(new); i += len(items) {
		copy(new[i:], items)
	}

	return new
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
	items := input.ReadIntArray(filename, "")
	logger.Debug(items)

	items = repeatArray(items, 10000)

	offsetStr := intArrayToString(items[0:7])
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		log.Fatal("failed to parse offset", err.Error())
	}
	logger.Debug("offset", offset)

	new := make([]int, len(items))
	for i := 0; i < stopPhase; i++ {
		new[len(items)-1] = items[len(items)-1]
		for j := len(items) - 2; j >= 0; j-- {
			new[j] = (items[j] + new[j+1]) % 10
		}

		items = new
	}

	return intArrayToString(items[offset : offset+8])
}
