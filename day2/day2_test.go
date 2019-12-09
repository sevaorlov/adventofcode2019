package day2_test

import (
	"adventofcode2019/day2"
	"adventofcode2019/input"
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestPart1(t *testing.T) {
	got := day2.Part1(input.MainFileName)
	assert.Equal(t, "3516593", got)
}

func TestPart2(t *testing.T) {
	got := day2.Part2(input.MainFileName)
	assert.Equal(t, "7749", got)
}
