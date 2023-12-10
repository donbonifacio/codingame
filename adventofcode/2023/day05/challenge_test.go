package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sample1(t *testing.T) {
	assert.Equal(t, part1(readInput("sample.txt")), 35, "they should be equal")
}

func TestMain_input1(t *testing.T) {
	assert.Equal(t, part1(readInput("input.txt")), 389056265, "they should be equal")
}

func Test_sample2(t *testing.T) {
	assert.Equal(t, part2(readInput("sample.txt")), 46, "they should be equal")
}

func Test_iinput2(t *testing.T) {
	assert.Equal(t, part2(readInput("input.txt")), 46, "they should be equal")
}
