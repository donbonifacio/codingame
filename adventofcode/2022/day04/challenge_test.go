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
	assert.Equal(t, 2, part1(data))
}

func TestOverlaps(t *testing.T) {
	assert.Equal(t, true, overlaps("1-4", "2-2"))
	assert.Equal(t, false, overlaps("1-4", "5-5"))
}
