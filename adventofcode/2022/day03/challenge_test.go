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
	data := readInput("sample.txt")
	assert.Equal(t, 157, totalPriorities(data))
}

func TestInput(t *testing.T) {
	data := readInput("input.txt")
	assert.Equal(t, 7908, totalPriorities(data))
}

func Test_findInBoth(t *testing.T) {
	assert.Equal(t, "a", findInBoth("abca"))
	assert.Equal(t, "b", findInBoth("aaabcxybzz"))
}

func Test_score(t *testing.T) {
	assert.Equal(t, 1, score("a"))
	assert.Equal(t, 26, score("z"))

	assert.Equal(t, 27, score("A"))
	assert.Equal(t, 52, score("Z"))
}

func Test_findInBags(t *testing.T) {
	assert.Equal(t, "a", findInBags([]string{"aabc", "xyza"}))
}

func TestSample_part2(t *testing.T) {
	data := readInput("sample.txt")
	assert.Equal(t, 70, part2sum(data))
}

func TestInput_part2(t *testing.T) {
	data := readInput("input.txt")
	assert.Equal(t, 2838, part2sum(data))
}
