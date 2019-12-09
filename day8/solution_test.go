package day8_test

import (
	"adventofcode2019/day8"
	"adventofcode2019/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {
	got := day8.Part1(input.MainFileName)
	assert.Equal(t, "2048", got)
}
