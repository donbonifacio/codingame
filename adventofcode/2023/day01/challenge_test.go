package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	assert.Equal(t, part1(readInput("sample.txt")), 142, "they should be equal")
}

func TestMain_sample(t *testing.T) {
	assert.Equal(t, part1(readInput("input.txt")), 53334, "they should be equal")
}
