package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sample1(t *testing.T) {
	assert.Equal(t, part1(readInput("sample.txt")), 2, "they should be equal")
}

func TestMain_input1(t *testing.T) {
	assert.Equal(t, part1(readInput("input.txt")), 524, "they should be equal")
}

func Test_sample2(t *testing.T) {
	assert.Equal(t, part2(readInput("sample.txt")), 14, "they should be equal")
}

func TestMain_input2(t *testing.T) {
	assert.Equal(t, part2(readInput("input.txt")), 524, "they should be equal")
}
func Test_sample_withoutIndex(t *testing.T) {
	assert.Equal(t, withoutIndex([]string{"1", "2", "3"}, 1), []string{"1", "3"}, "they should be equal")
	assert.Equal(t, withoutIndex([]string{"4", "5", "6"}, 0), []string{"5", "6"}, "they should be equal")
	assert.Equal(t, withoutIndex([]string{"7", "8", "9"}, 2), []string{"7", "8"}, "they should be equal")
}
