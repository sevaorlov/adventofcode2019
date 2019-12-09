package day1

import (
	"adventofcode2019/input"
	"log"
	"strconv"
)

func Part2(filename string) string {
	var sum int64

	input.ReadFile(filename, func(line string) {
		mass, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(err.Error())
		}

		sum += fuel(mass)
	})

	return strconv.FormatInt(sum, 10)
}

func fuel(mass int64) int64 {
	v := (mass / 3) - 2

	if v <= 0 {
		return 0
	}

	return v + fuel(v)
}
