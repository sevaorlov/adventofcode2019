package day13_test

import (
	"adventofcode2019/day13"
	"adventofcode2019/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	got := day13.Part1(input.MainFileName)
	assert.Equal(t, "258", got)
}

func TestPart2(t *testing.T) {
	got := day13.Part2(input.MainFileName)
	assert.Equal(t, "12765", got)
}
