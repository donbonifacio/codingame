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

func TestAsLines(t *testing.T) {
	assert.Equal(t, []string{"1", "2", "3"}, AsLines("1\n 2\n3\n\n"))
}

func TestEachString(t *testing.T) {
	data := []string{"a", "b", "c"}
	output := ""
	EachString(data, func(str string) {
		output = output + str
	})
	assert.Equal(t, "abc", output)
}

func TestIsEmpty(t *testing.T) {
	assert.Equal(t, true, IsEmpty(""))
	assert.Equal(t, false, IsEmpty("ddd"))
}

func TestMinMax(t *testing.T) {
	min, max := MinMax([]int{1, 3, 3, 0, 10, 5})
	assert.Equal(t, 0, min)
	assert.Equal(t, 10, max)
}
