package day24_test

import (
	"adventofcode2019/day24"
	"adventofcode2019/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	got := day24.Part1(input.MainFileName)
	assert.Equal(t, "27777901", got)
}

func TestPart2(t *testing.T) {
	got := day24.Part2(input.MainFileName)
	assert.Equal(t, "", got)
}
