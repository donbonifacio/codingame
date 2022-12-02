package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	main()
	assert.Equal(t, 123, 123, "they should be equal")
}

func TestSample(t *testing.T) {
	input := readInput("sample.txt")
	score := calculateScore(input)
	assert.Equal(t, 15, score)
}

func TestInput(t *testing.T) {
	input := readInput("input.txt")
	score := calculateScore(input)
	assert.Equal(t, 9177, score)
}

func TestCalculatePlayerScore(t *testing.T) {
	assert.Equal(t, 6+1, calculatePlayScore("C", "X"))
	assert.Equal(t, 6+2, calculatePlayScore("A", "Y"))
	assert.Equal(t, 6+3, calculatePlayScore("B", "Z"))

	assert.Equal(t, 1, calculatePlayScore("B", "X"))
	assert.Equal(t, 6, calculatePlayScore("C", "Z"))

	assert.Equal(t, 2, calculatePlayScore("C", "Y"))
	assert.Equal(t, 3, calculatePlayScore("A", "Z"))

	assert.Equal(t, 3+1, calculatePlayScore("A", "X"))
	assert.Equal(t, 3+2, calculatePlayScore("B", "Y"))
	assert.Equal(t, 3+3, calculatePlayScore("C", "Z"))
}
