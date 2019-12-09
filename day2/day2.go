package day2

import (
	"adventofcode2019/input"
	"adventofcode2019/transform"
	"errors"
	"fmt"
	"log"
	"strconv"
)

const part2Output = 19690720

func Part1(filename string) string {
	line := input.ReadSingleLine(filename)

	a := transform.Int64ArrayFromLine(line)
	a[1] = 12
	a[2] = 2
	//fmt.Println(b)

	solution, err := solveForArray(a)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return strconv.FormatInt(solution, 10)
}

func Part2(filename string) string {
	line := input.ReadSingleLine(filename)
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			a := transform.Int64ArrayFromLine(line)
			a[1] = int64(i)
			a[2] = int64(j)

			solution, err := solveForArray(a)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}

			if solution == part2Output {
				return strconv.FormatInt(a[1]*100+a[2], 10)
			}
		}
	}

	return ""
}

func solveForArray(a []int64) (int64, error) {
	for i := 0; i < len(a); i += 4 {
		var v1, v2 int64

		op := a[i]
		//fmt.Println(i, op)

		if op == 1 || op == 2 {
			v1 = a[a[i+1]]
			v2 = a[a[i+2]]
		}

		switch op {
		case 1:
			a[a[i+3]] = v1 + v2
		case 2:
			a[a[i+3]] = v1 * v2
		case 99:
			return a[0], nil
		default:
			return 0, errors.New(fmt.Sprintf("unknown code %v", op))
		}

		//fmt.Println(b)
	}

	return 0, nil
}
