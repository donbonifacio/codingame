package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sample1(t *testing.T) {
	assert.Equal(t, part1(readInput("sample.txt")), 161, "they should be equal")
}

func TestMain_input1(t *testing.T) {
	assert.Equal(t, part1(readInput("input.txt")), 166357705, "they should be equal")
}

func Test_sample2(t *testing.T) {
	assert.Equal(t, part2(readInput("sample2.txt")), 48, "they should be equal")
}

func Test_input2(t *testing.T) {
	assert.Equal(t, part2(readInput("input.txt")), 48, "they should be equal")
}
