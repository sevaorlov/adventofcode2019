package day7_test

import (
	"adventofcode2019/day7"
	"adventofcode2019/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	got := day7.Part1(input.MainFileName)
	assert.Equal(t, "338603", got)
}

func TestPart2(t *testing.T) {
	got := day7.Part2(input.MainFileName)
	assert.Equal(t, "63103596", got)
}
