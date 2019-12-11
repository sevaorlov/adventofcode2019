package day11_test

import (
	"adventofcode2019/day11"
	"adventofcode2019/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	got := day11.Part1(input.MainFileName)
	assert.Equal(t, "2322", got)
}
