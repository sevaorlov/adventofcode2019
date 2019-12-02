package day2

import (
	"adventofcode2019/input"
	"adventofcode2019/transform"
	"errors"
	"fmt"
	"log"
)

func Solve(filename string) {
	input.ReadFile(filename, func(line string) {
		a := transform.Int64ArrayFromLine(line)
		a[1] = 12
		a[2] = 2
		//fmt.Println(b)

		solution, err := solveForArray(a)
		if err != nil {
			log.Fatalf(err.Error())
		}

		fmt.Println(solution)
	})
}

func Solve2(filename string) {
	input.ReadFile(filename, func(line string) {
		for i := 0; i <= 99; i++ {
			for j := 0; j <= 99; j++ {
				a := transform.Int64ArrayFromLine(line)
				a[1] = int64(i)
				a[2] = int64(j)

				solution, err := solveForArray(a)
				if err != nil {
					log.Println(err.Error())
					continue
				}

				if solution == 19690720 {
					fmt.Println(a[1], a[2], a[1]*100+a[2])
					return
				}
			}
		}
	})
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
