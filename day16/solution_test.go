package day16_test

import (
	"adventofcode2019/day16"
	"adventofcode2019/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	got := day16.Part1(input.MainFileName)
	assert.Equal(t, "77038830", got)
}

func TestPart2(t *testing.T) {
	got := day16.Part2(input.MainFileName)
	assert.Equal(t, "", got)
}
