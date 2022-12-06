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
	assert.Equal(t, 7, findCode("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 4))
}

func TestInputPart1(t *testing.T) {
	data := readInput("input.txt")
	assert.Equal(t, 1912, findCode(data, 4))
}

func TestIsMarker(t *testing.T) {
	assert.Equal(t, true, isMarker("abcd", 4))
	assert.Equal(t, false, isMarker("abca", 4))
}

func TestSamplePart2(t *testing.T) {
	assert.Equal(t, 19, findCode("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14))
}

func TestInputPart2(t *testing.T) {
	data := readInput("input.txt")
	assert.Equal(t, 2122, findCode(data, 14))
}
