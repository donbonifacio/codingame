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

func TestPart1(t *testing.T) {
	data := utils.ReadInput("sample.txt")
	assert.Equal(t, 16+5, part1(data))
}

func TestPartInput1(t *testing.T) {
	data := utils.ReadInput("input.txt")
	assert.Equal(t, 1736, part1(data))
}

func TestVisibleFrom(t *testing.T) {
	matrix := utils.AsIntMatrix("29512\n25512\n23512")
	position := utils.Position{X: 1, Y: 1}
	value := matrix.Value(position)

	assert.Equal(t, false, visibleFrom(matrix, value, position, utils.Vector{X: 0, Y: -1}))
	assert.Equal(t, true, visibleFrom(matrix, value, position, utils.Vector{X: -1, Y: 0}))
	assert.Equal(t, true, visibleFrom(matrix, value, position, utils.Vector{X: 0, Y: 1}))
	assert.Equal(t, false, visibleFrom(matrix, value, position, utils.Vector{X: 1, Y: 0}))
}

func TestPart2(t *testing.T) {
	data := utils.ReadInput("sample.txt")
	assert.Equal(t, 8, part2(data))
}

func TestPartInput2(t *testing.T) {
	data := utils.ReadInput("input.txt")
	assert.Equal(t, 268800, part2(data))
}
