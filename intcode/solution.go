package intcode

import (
	"errors"
	"fmt"
	"log"
)

func Solve(a []int64, inputInstructions []int64, startIndex int, startRelBase int, stopOnOutput bool) (output []int64, index int, relBase int, err error) {
	var last, inputIndex int

	relBase = startRelBase
	index = startIndex

	for {
		last = index

		//logger.Debug("---")
		opcode, p1Index, p2Index, p3Index := parsedParameters(a, index, relBase)

		//if index+3 < len(a) {
		//	logger.Debug(a[index], a[index+1], a[index+2], a[index+3])
		//}
		//logger.Debug("before", a)

		//logger.Debug("opcode", opcode)
		//logger.Debug(a[index], a[index+1], a[index+2], a[index+3])
		switch opcode {
		case 1:
			write(a, p3Index, a[p1Index]+a[p2Index])
			index += 4
		case 2:
			write(a, p3Index, a[p1Index]*a[p2Index])
			index += 4
		case 3:
			write(a, p1Index, inputInstructions[inputIndex])
			if inputIndex < len(inputInstructions)-1 {
				inputIndex++
			}
			index += 2
		case 4:
			//if a[p1Index] != 0 {
			output = append(output, a[p1Index])
			//}
			index += 2
			if stopOnOutput {
				return
			}
		case 5:
			if a[p1Index] != 0 {
				index = int(a[p2Index])
			} else {
				index += 3
			}
		case 6:
			if a[p1Index] == 0 {
				index = int(a[p2Index])
			} else {
				index += 3
			}
		case 7:
			if a[p1Index] < a[p2Index] {
				write(a, p3Index, 1)
			} else {
				write(a, p3Index, 0)
			}
			index += 4
		case 8:
			if a[p1Index] == a[p2Index] {
				write(a, p3Index, 1)
			} else {
				write(a, p3Index, 0)
			}
			index += 4
		case 9:
			relBase += int(a[p1Index])
			index += 2
		case 99:
			return output, -1, 0, nil
		default:
			return []int64{}, 0, 0, errors.New(fmt.Sprintf("unknown code %v", opcode))
		}
		//logger.Debug("after ", a)

		if last == index {
			log.Fatal("loop", a[index], a)
		}
	}
}

func write(a []int64, index int, value int64) {
	//logger.Debug("write", index, value)
	a[index] = value
}

func parsedParameters(a []int64, index int, relBase int) (int, int, int, int) {
	var p1Index, p2Index, pIndex3 int

	mode3, mode2, mode1, opcode := instructions(int(a[index]))
	//logger.Debug("instructions", a[index], mode3, mode2, mode1, opcode)

	if index+1 < len(a) {
		p1Index = getIndexWithMode(a, index+1, mode1, relBase)
	}
	if index+2 < len(a) {
		p2Index = getIndexWithMode(a, index+2, mode2, relBase)
	}
	if index+3 < len(a) {
		pIndex3 = getIndexWithMode(a, index+3, mode3, relBase)
	}

	return opcode, p1Index, p2Index, pIndex3
}

func instructions(k int) (int, int, int, int) {
	a := k / 10000
	b := k / 1000 % 10
	c := k / 100 % 10
	return a, b, c, k % 100
}

func getIndexWithMode(a []int64, index, mode int, relBase int) int {
	switch mode {
	case 1:
		return index
	case 2:
		return int(a[index]) + relBase
	default:
		return int(a[index])
	}
}
