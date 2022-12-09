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
	assert.Equal(t, 13, part1(data))
}

func Test1Part1(t *testing.T) {
	data := utils.ReadInput("test1.txt")
	assert.Equal(t, 1, part1(data))
}

func Test2Part1(t *testing.T) {
	data := utils.ReadInput("test2.txt")
	assert.Equal(t, 3, part1(data))
}

func TestFirstPart1(t *testing.T) {
	data := utils.ReadInput("first.txt")
	assert.Equal(t, 10, part1(data))
}

func TestInputPart1(t *testing.T) {
	data := utils.ReadInput("input.txt")
	assert.Equal(t, 5960, part1(data))
}

func TestSamplePart2(t *testing.T) {
	data := utils.ReadInput("sample.txt")
	assert.Equal(t, 1, part2(data))
}

func TestInputPart2(t *testing.T) {
	data := utils.ReadInput("input.txt")
	assert.Equal(t, 2327, part2(data))
}
