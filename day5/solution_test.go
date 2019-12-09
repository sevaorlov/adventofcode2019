package day5_test

import (
	"adventofcode2019/day5"
	"adventofcode2019/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	got := day5.Part1(input.MainFileName)
	assert.Equal(t, "7265618", got)
}

func TestPart2(t *testing.T) {
	got := day5.Part2(input.MainFileName)
	assert.Equal(t, "7731427", got)
}
