package day5

import (
	"adventofcode2019/input"
	"adventofcode2019/logger"
	"adventofcode2019/transform"
	"errors"
	"fmt"
	"log"
	"strconv"
)

func Part1(filename string) string {
	line := input.ReadSingleLine(filename)
	a := transform.Int64ArrayFromLine(line)

	res, err := solveForArray(a, 1)
	if err != nil {
		log.Fatalf("cannot solve %s\n", err.Error())
	}

	return strconv.FormatInt(res, 10)
}

func Part2(filename string) string {
	line := input.ReadSingleLine(filename)
	a := transform.Int64ArrayFromLine(line)

	res, err := solveForArray(a, 5)
	if err != nil {
		log.Fatalf("cannot solve %s\n", err.Error())
	}

	return strconv.FormatInt(res, 10)
}

func solveForArray(a []int64, inputParam int) (int64, error) {
	var i, last int
	var output int64

	for i < len(a) {
		var step int
		last = i

		opcode, p1Index, p2Index, p3Index := parsedParameters(a, i)

		if i+3 < len(a) {
			logger.Debug(a[i], a[i+1], a[i+2], a[i+3])
		}
		logger.Debug("before", a)

		switch opcode {
		case 1:
			a[p3Index] = a[p1Index] + a[p2Index]
			step = 4
		case 2:
			a[p3Index] = a[p1Index] * a[p2Index]
			step = 4
		case 3:
			a[p1Index] = int64(inputParam)
			step = 2
		case 4:
			output = a[a[i+1]]
			step = 2
		case 5:
			if a[p1Index] != 0 {
				i = int(a[p2Index])
			} else {
				step = 3
			}
		case 6:
			if a[p1Index] == 0 {
				i = int(a[p2Index])
			} else {
				step = 3
			}
		case 7:
			if a[p1Index] < a[p2Index] {
				a[p3Index] = 1
			} else {
				a[p3Index] = 0
			}
			step = 4
		case 8:
			if a[p1Index] == a[p2Index] {
				a[p3Index] = 1
			} else {
				a[p3Index] = 0
			}
			step = 4
		case 99:
			return output, nil
		default:
			return 0, errors.New(fmt.Sprintf("unknown code %v", opcode))
		}
		logger.Debug("after ", a)

		if step > 0 {
			i += step
		}

		if last == i {
			log.Fatal("loop", a[i], a)
		}
	}
	return 0, nil
}

func parsedParameters(a []int64, index int) (int, int, int, int) {
	var p1Index, p2Index, pIndex3 int

	mode3, mode2, mode1, opcode := instructions(int(a[index]))
	logger.Debug("instructions", a[index], mode3, mode2, mode1, opcode)

	if index < len(a) {
		p1Index = getIndexWithMode(a, index+1, mode1)
	}
	if index+2 < len(a) {
		p2Index = getIndexWithMode(a, index+2, mode2)
	}
	if index+3 < len(a) {
		pIndex3 = getIndexWithMode(a, index+3, mode3)
	}

	return opcode, p1Index, p2Index, pIndex3
}

func instructions(k int) (int, int, int, int) {
	a := k / 10000
	b := (k - a) / 1000
	c := (k % 1000) / 100
	return a, b, c, k % 100
}

func getIndexWithMode(a []int64, index, mode int) int {
	if mode == 1 {
		return index
	}
	return int(a[index])
}
