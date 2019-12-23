package day12_test

import (
	"adventofcode2019/day12"
	"adventofcode2019/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	got := day12.Part1(input.MainFileName)
	assert.Equal(t, "8960", got)
}

func TestPart2(t *testing.T) {
	got := day12.Part2(input.MainFileName)
	assert.Equal(t, "314917503970904", got)
}
