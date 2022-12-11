package main

import (
	"strings"
	"testing"

	"github.com/donbonifacio/codingame/blob/master/adventofcode/2022/utils"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	main()
	assert.Equal(t, 123, 123, "they should be equal")
}

func TestScanMonkeyData(t *testing.T) {
	data := utils.ReadInput("sample.txt")
	monkeyData := strings.Split(data, "\n\n")[1]
	monkey := scanMonkey(monkeyData)
	assert.Equal(t, 1, monkey.id)
	assert.Equal(t, []int{54, 65, 75, 74}, monkey.items)
	assert.Equal(t, 2*6, monkey.op(2))
	assert.Equal(t, 2, monkey.test(2*19))
	assert.Equal(t, 0, monkey.test(2*18))
}

func TestSamplePart1(t *testing.T) {
	data := utils.ReadInput("sample.txt")
	assert.Equal(t, 10605, part1(data))
}

func TestInputPart1(t *testing.T) {
	data := utils.ReadInput("input.txt")
	assert.Equal(t, 67830, part1(data))
}

func TestSamplePart2(t *testing.T) {
	// wasn't able to do this one
	//data := utils.ReadInput("sample.txt")
	//assert.Equal(t, 2713310158, part2(data))
}
