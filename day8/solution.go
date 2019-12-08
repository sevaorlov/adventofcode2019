package day8

import (
	"adventofcode2019/input"
	"fmt"
	"math"
)

const (
	pixelWidth = 25
	pixelHeight = 6
)

func Part1(filename string) {
	a := input.ReadIntSpacedArray(filename)

	min := math.MaxInt32
	minIndex := 0

	for i:=0; i< len(a); i+= pixelHeight * pixelWidth {
		n := countDigits(a, i, 0)

		if n < min {
			min = n
			minIndex = i
		}
	}

	fmt.Println(min, countDigits(a, minIndex, 1) * countDigits(a, minIndex, 2))
}

func Part2(filename string) {
	a := input.ReadIntSpacedArray(filename)

	layers := len(a) / (pixelHeight * pixelWidth)
	finalImage := make([]int, pixelWidth * pixelHeight)

	for i:=0;i<pixelHeight * pixelWidth;i++ {
		finalImage[i] = 2
		for l:=0; l< layers; l++ {
			index := l * pixelHeight * pixelWidth + i

			if finalImage[i] == 2 {
				finalImage[i] = a[index]
			}
		}
	}

	for j:=0;j<pixelHeight;j++ {
		for i := 0; i < pixelWidth; i++ {
			switch finalImage[j * pixelWidth + i] {
			case 1:
				fmt.Print("█")
			case 0:
				fmt.Print(" ")
			default:
				fmt.Print("?")
			}
		}
		fmt.Println()
	}
}

func countDigits(a []int, index, which int) int {
	var n int

	for i:=0; i< pixelHeight * pixelWidth; i++ {
		if a[index + i] == which {
			n++
		}
	}

	return n
}
