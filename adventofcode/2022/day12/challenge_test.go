package main

import (
	"testing"

	"github.com/donbonifacio/codingame/blob/master/adventofcode/2022/utils"
	"github.com/stretchr/testify/assert"
)

func TestSamplePart1(t *testing.T) {
	data := utils.ReadInput("sample.txt")
	tracker := part1(data)
	Print(&tracker)
	assert.Equal(t, 311, len(tracker.quickest)-1)
}

func TestSample2Part1(t *testing.T) {
	data := utils.ReadInput("sample2.txt")
	tracker := part1(data)
	Print(&tracker)
	assert.Equal(t, 2, len(tracker.quickest)-1)
}

func TestSample3Part1(t *testing.T) {
	data := utils.ReadInput("sample3.txt")
	tracker := part1(data)
	Print(&tracker)
	assert.Equal(t, 50, len(tracker.quickest)+1)
}

func TestSample4Part1(t *testing.T) {
	data := utils.ReadInput("sample.txt")
	tracker := part1(data)
	Print(&tracker)
	assert.Equal(t, 31, len(tracker.quickest)+1)
}

func TestPossibleFrom(t *testing.T) {
	data := utils.ReadInput("sample.txt")
	tracker := part1(data)
	assert.Equal(t, []utils.Position{{X: 1, Y: 0}, {X: 0, Y: 1}}, possibleFrom(&tracker, &map[utils.Position]bool{}))
}

func TestInputPart1(t *testing.T) {
	data := utils.ReadInput("input.txt")
	tracker := part1(data)
	Print(&tracker)
	assert.Equal(t, 311, len(tracker.quickest)-1)
}
