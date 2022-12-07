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
	fs := part1(data)
	assert.Equal(t, 48381165, fs.root.size)
	assert.Equal(t, 95437, sumLargeDirs(&fs.root))
}

func TestInputPart1(t *testing.T) {
	data := readInput("input.txt")
	fs := part1(data)
	assert.Equal(t, 1307902, sumLargeDirs(&fs.root))
}

func TestSamplePart2(t *testing.T) {
	data := readInput("sample.txt")
	fs := part1(data)
	assert.Equal(t, 24933642, smallestToDelete(fs).size)
}

func TestInputPart2(t *testing.T) {
	data := readInput("input.txt")
	fs := part1(data)
	assert.Equal(t, 7068748, smallestToDelete(fs).size)
}
