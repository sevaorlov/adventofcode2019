package day4_test

import (
	"adventofcode2019/day4"
	"adventofcode2019/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	got := day4.Part1(input.MainFileName)
	assert.Equal(t, "466", got)
}

func TestPart2(t *testing.T) {
	got := day4.Part2(input.MainFileName)
	assert.Equal(t, "292", got)
}
