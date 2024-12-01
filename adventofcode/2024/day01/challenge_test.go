package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sample1(t *testing.T) {
	assert.Equal(t, part1(readInput("sample.txt")), 11, "they should be equal")
}

func TestMain_input1(t *testing.T) {
	assert.Equal(t, part1(readInput("input.txt")), 2057374, "they should be equal")
}

func Test_sample2(t *testing.T) {
	assert.Equal(t, part2(readInput("sample.txt")), 31, "they should be equal")
}

func Test_input2(t *testing.T) {
	assert.Equal(t, part2(readInput("input.txt")), 23177084, "they should be equal")
}
