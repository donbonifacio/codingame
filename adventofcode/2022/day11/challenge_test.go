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

func TestBuildMonkeyData(t *testing.T) {
	data := []string{
		"Monkey 79:",
		"Starting items: 75, 98",
		"Operation: new = old * 19",
		"Test: divisible by 23",
		"If true: throw to monkey 2",
		"If false: throw to monkey 3",
	}
	monkey := buildMonkey(data)
	assert.Equal(t, 79, monkey.id)
	assert.Equal(t, []int{75, 98}, monkey.items)
	assert.Equal(t, 2*19, monkey.op(2))
	assert.Equal(t, true, monkey.test(2*23))
	assert.Equal(t, false, monkey.test(2*24))
	assert.Equal(t, 2, monkey.next[true])
	assert.Equal(t, 3, monkey.next[false])
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
