package day6_test

import (
	"adventofcode2019/day6"
	"adventofcode2019/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	got := day6.Part1(input.MainFileName)
	assert.Equal(t, "344238", got)
}

func TestPart2(t *testing.T) {
	got := day6.Part2(input.MainFileName)
	assert.Equal(t, "436", got)
}
