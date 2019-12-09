package day9_test

import (
	"adventofcode2019/day9"
	"adventofcode2019/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	got := day9.Part1(input.MainFileName)
	assert.Equal(t, "2890527621", got)
}

func TestPart2(t *testing.T) {
	got := day9.Part2(input.MainFileName)
	assert.Equal(t, "66772", got)
}
