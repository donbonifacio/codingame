package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sample1(t *testing.T) {
	assert.Equal(t, part1(readInput("sample.txt")), 13, "they should be equal")
}

func TestMain_input1(t *testing.T) {
	assert.Equal(t, part1(readInput("input.txt")), 23028, "they should be equal")
}

func Test_sample2(t *testing.T) {
	assert.Equal(t, part2(readInput("sample.txt")), 30, "they should be equal")
}

func TestMain_input2(t *testing.T) {
	assert.Equal(t, part2(readInput("input.txt")), 9236992, "they should be equal")
}
