package day4

import (
	"fmt"
	"math"
)

const (
	from   = 387638
	to     = 919123
	length = 6
)

func Part1(_ string) {
	var n int

	possibleNumbers(func(ints [length]int) {
		if ints[0] == ints[1] ||
			ints[1] == ints[2] ||
			ints[2] == ints[3] ||
			ints[3] == ints[4] ||
			ints[4] == ints[5] {

			n++
		}
	})

	fmt.Println(n)
}

func Part2(_ string) {
	var n int

	possibleNumbers(func(ints [length]int) {
		b := make([]int, 10)
		for _, item := range ints {
			b[item]++
		}

		var same bool
		for _, item := range b {
			same = item == 2
			if same {
				n++
				return
			}
		}
	})

	fmt.Println(n)
}

func possibleNumbers(found func([length]int)) {
	firstNum := from / int(math.Pow10(5))

	for k1 := firstNum; k1 <= 9; k1++ {
		for k2 := k1; k2 <= 9; k2++ {
			for k3 := k2; k3 <= 9; k3++ {
				for k4 := k3; k4 <= 9; k4++ {
					for k5 := k4; k5 <= 9; k5++ {
						for k6 := k5; k6 <= 9; k6++ {
							num := k1*int(math.Pow10(5)) + k2*int(math.Pow10(4)) +
								k3*int(math.Pow10(3)) +
								k4*int(math.Pow10(2)) +
								k5*int(math.Pow10(1)) +
								k6

							if num >= from && num <= to {
								found([length]int{k1, k2, k3, k4, k5, k6})
							}
						}
					}
				}
			}
		}
	}
}
