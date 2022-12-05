package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	main()
	assert.Equal(t, 123, 123, "they should be equal")
}

func Test_load(t *testing.T) {
	data := readInput("sample.txt")
	cargo := load(data)
	assert.Equal(t, 3, len(cargo.cranes))
	assert.Equal(t, 4, len(cargo.operations))
	assert.Equal(t, Crane{[]string{"M", "C", "D"}}, cargo.cranes[2])
	assert.Equal(t, Operation{2, 2, 1}, cargo.operations[2])
}

func Test_SamplePart1(t *testing.T) {
	top := run(readInput("sample.txt"), runOperationsPart1)
	assert.Equal(t, "CMZ", top)
}

func Test_InputPart1(t *testing.T) {
	top := run(readInput("input.txt"), runOperationsPart1)
	assert.Equal(t, "PTWLTDSJV", top)
}

func Test_SamplePart2(t *testing.T) {
	top := run(readInput("sample.txt"), runOperationsPart2)
	assert.Equal(t, "MCD", top)
}

func Test_InputPart2(t *testing.T) {
	top := run(readInput("input.txt"), runOperationsPart2)
	assert.Equal(t, "WZMFVGGZP", top)
}
