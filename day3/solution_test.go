package day3_test

import (
	"adventofcode2019/day3"
	"adventofcode2019/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	got := day3.Part1(input.MainFileName)
	assert.Equal(t, "2427", got)
}

func TestPart2(t *testing.T) {
	got := day3.Part2(input.MainFileName)
	assert.Equal(t, "27890", got)
}
