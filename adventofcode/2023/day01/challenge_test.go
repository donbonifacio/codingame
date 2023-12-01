package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	assert.Equal(t, part1(readInput("sample.txt")), 142, "they should be equal")
}

func TestMain_sample(t *testing.T) {
	assert.Equal(t, part1(readInput("input.txt")), 53334, "they should be equal")
}

func TestMain_sample2(t *testing.T) {
	assert.Equal(t, part2(readInput("sample2.txt")), 281, "they should be equal")
}

func TestMain_input2(t *testing.T) {
	assert.Equal(t, part2(readInput("input2.txt")), 52834, "they should be equal")
}

func Test_values(t *testing.T) {
	assert.Equal(t, part2([]string{"oneight1"}), 11, "they should be equal")
	assert.Equal(t, part2([]string{"eightwo1"}), 81, "they should be equal")
	assert.Equal(t, part2([]string{"twone1"}), 21, "they should be equal")
	assert.Equal(t, part2([]string{"sevenine1"}), 71, "they should be equal")
}
