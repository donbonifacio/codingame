package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sample1(t *testing.T) {
	races := []Race{
		{time: 7, distance: 9, distances: []int{}},
		{time: 15, distance: 40, distances: []int{}},
		{time: 30, distance: 200, distances: []int{}},
	}
	assert.Equal(t, part1(races), 288, "they should be equal")
}

func TestMain_input1(t *testing.T) {
	races := []Race{
		{time: 44, distance: 277, distances: []int{}},
		{time: 89, distance: 1136, distances: []int{}},
		{time: 96, distance: 1890, distances: []int{}},
		{time: 91, distance: 1768, distances: []int{}},
	}
	assert.Equal(t, part1(races), 2344708, "they should be equal")
}
