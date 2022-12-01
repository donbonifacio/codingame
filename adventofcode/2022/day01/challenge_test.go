package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const sampleData = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

func TestMain(t *testing.T) {
	main()
	assert.Equal(t, 123, 123, "they should be equal")
}

func TestSample(t *testing.T) {
	calories := caloriesByElf(sampleData)
	assert.Equal(t, calories, []int{6000, 4000, 11000, 24000, 10000})

	_, max := MinMax(calories)
	assert.Equal(t, max, 24000)
}

func TestInput(t *testing.T) {
	calories := caloriesByElf(readInput())
	_, max := MinMax(calories)
	assert.Equal(t, max, 71934)
}

func TestSamplePart2(t *testing.T) {
	calories := caloriesByElf(sampleData)
	assert.Equal(t, top3(calories), 45000)
}

func TestInputPart2(t *testing.T) {
	calories := caloriesByElf(readInput())
	assert.Equal(t, top3(calories), 211447)
}
