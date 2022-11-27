package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	main()
	assert.Equal(t, 123, 123, "they should be equal")
}

func TestReadNumberFile(t *testing.T) {
	defer os.Remove("tmp.txt")

	writeFile("tmp.txt", "1, 2,4 ")
	nums := readFileWithNumbers("tmp.txt")
	assert.Equal(t, nums, []int{1, 2, 4})
}
