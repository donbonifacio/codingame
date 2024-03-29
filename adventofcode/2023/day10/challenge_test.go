package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sample1(t *testing.T) {
	assert.Equal(t, part1(readInput("sample.txt")), 4, "they should be equal")
}

func Test_sample2(t *testing.T) {
	assert.Equal(t, part1(readInput("sample2.txt")), 8, "they should be equal")
}

func TestMain_input1(t *testing.T) {
	assert.Equal(t, part1(readInput("input.txt")), 6838, "they should be equal")
}
