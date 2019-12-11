package day1_test

import (
	"adventofcode2019/day1"
	"adventofcode2019/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart2(t *testing.T) {
	got := day1.Part2(input.MainFileName)
	assert.Equal(t, "5094261", got)
}
