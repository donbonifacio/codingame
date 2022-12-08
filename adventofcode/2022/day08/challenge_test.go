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
	matrix := [][]int{
		[]int{2, 9, 5, 1, 2},
		[]int{2, 5, 5, 1, 2},
		[]int{2, 3, 5, 1, 2},
	}
	assert.Equal(t, false, visibleFrom(matrix, 5, []int{1, 1}, []int{-1, 0}))
	assert.Equal(t, true, visibleFrom(matrix, 5, []int{1, 1}, []int{0, -1}))
	assert.Equal(t, true, visibleFrom(matrix, 5, []int{1, 1}, []int{1, 0}))
	assert.Equal(t, false, visibleFrom(matrix, 5, []int{1, 1}, []int{0, 1}))
}

func TestPart2(t *testing.T) {
	data := utils.ReadInput("sample.txt")
	assert.Equal(t, 8, part2(data))
}

func TestPartInput2(t *testing.T) {
	data := utils.ReadInput("input.txt")
	assert.Equal(t, 268800, part2(data))
}
