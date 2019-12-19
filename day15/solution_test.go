package day15_test

import (
	"adventofcode2019/day15"
	"adventofcode2019/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	got := day15.Part1(input.MainFileName)
	assert.Equal(t, "234", got)
}

func TestPart2(t *testing.T) {
	got := day15.Part2(input.MainFileName)
	assert.Equal(t, "", got)
}
