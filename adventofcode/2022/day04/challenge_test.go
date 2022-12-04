package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	main()
	assert.Equal(t, 123, 123, "they should be equal")
}

func TestSamplePart1(t *testing.T) {
	data := readInput("sample.txt")
	assert.Equal(t, 2, part1(data))
}

func TestInputPart1(t *testing.T) {
	data := readInput("input.txt")
	assert.Equal(t, 651, part1(data))
}

func TestOverlaps(t *testing.T) {
	assert.Equal(t, true, overlaps("1-4", "2-2"))
	assert.Equal(t, false, overlaps("1-4", "5-5"))
}

func TestPartialOverlaps(t *testing.T) {
	assert.Equal(t, false, partialOverlaps("2-4", "6-8"))
	assert.Equal(t, true, partialOverlaps("5-7", "7-9"))
}

func TestSamplePart2(t *testing.T) {
	data := readInput("sample.txt")
	assert.Equal(t, 4, part2(data))
}

func TestInputPart2(t *testing.T) {
	data := readInput("input.txt")
	assert.Equal(t, 2, part2(data))
}
