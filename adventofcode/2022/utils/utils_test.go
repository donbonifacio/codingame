package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAtoi(t *testing.T) {
	assert.Equal(t, 1, Atoi("1"))
	assert.Equal(t, 12345, Atoi("12345"))
}

func TestReadNumberFile(t *testing.T) {
	defer os.Remove("tmp.txt")

	WriteFile("tmp.txt", "1, 2,4 ")
	nums := ReadFileWithNumbers("tmp.txt")
	assert.Equal(t, nums, []int{1, 2, 4})
}
