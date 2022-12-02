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
	score := sumScore(input, part1calculator)
	assert.Equal(t, 15, score)
}

func TestInput(t *testing.T) {
	input := readInput("input.txt")
	score := sumScore(input, part1calculator)
	assert.Equal(t, 9177, score)
}

func TestCalculatePlayerScore(t *testing.T) {
	assert.Equal(t, 6+1, part1calculator("C", "X"))
	assert.Equal(t, 6+2, part1calculator("A", "Y"))
	assert.Equal(t, 6+3, part1calculator("B", "Z"))

	assert.Equal(t, 1, part1calculator("B", "X"))
	assert.Equal(t, 6, part1calculator("C", "Z"))

	assert.Equal(t, 2, part1calculator("C", "Y"))
	assert.Equal(t, 3, part1calculator("A", "Z"))

	assert.Equal(t, 3+1, part1calculator("A", "X"))
	assert.Equal(t, 3+2, part1calculator("B", "Y"))
	assert.Equal(t, 3+3, part1calculator("C", "Z"))
}

func TestSamplePart2(t *testing.T) {
	input := readInput("sample.txt")
	score := sumScore(input, part2calculator)
	assert.Equal(t, 12, score)
}

func TestInputPart2(t *testing.T) {
	input := readInput("input.txt")
	score := sumScore(input, part2calculator)
	assert.Equal(t, 12111, score)
}

func TestCalculatePlayerScorePart2(t *testing.T) {
	assert.Equal(t, 4, part2calculator("A", "Y"))
	assert.Equal(t, 1, part2calculator("B", "X"))
	assert.Equal(t, 7, part2calculator("C", "Z"))
}
