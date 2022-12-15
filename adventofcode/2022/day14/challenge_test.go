package main

import (
	"testing"

	"github.com/donbonifacio/codingame/blob/master/adventofcode/2022/utils"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	main()
	assert.Equal(t, 123, 123, "they should be equal")
}

func TestSamplePart1(t *testing.T) {
	data := utils.ReadInput("sample.txt")
	assert.Equal(t, 24, part1(data))
}

func TestInputPart1(t *testing.T) {
	data := utils.ReadInput("input.txt")
	assert.Equal(t, 825, part1(data))
}

func TestSamplePart2(t *testing.T) {
	data := utils.ReadInput("sample.txt")
	assert.Equal(t, 93, part2(data))
}

func TestInputPart2(t *testing.T) {
	data := utils.ReadInput("input.txt")
	assert.Equal(t, 26729, part2(data))
}
