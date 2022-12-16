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
	assert.Equal(t, 26, part1(data, 10))
}

func TestInputPart1(t *testing.T) {
	data := utils.ReadInput("input.txt")
	assert.Equal(t, 26, part1(data, 2000000))
}
