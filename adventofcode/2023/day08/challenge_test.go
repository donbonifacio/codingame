package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sample1(t *testing.T) {
	assert.Equal(t, part1(readInput("sample.txt")), 2, "they should be equal")
}

func Test_sample2(t *testing.T) {
	assert.Equal(t, part1(readInput("sample2.txt")), 6, "they should be equal")
}

func TestMain_input1(t *testing.T) {
	assert.Equal(t, part1(readInput("input.txt")), 22199, "they should be equal")
}
