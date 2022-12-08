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

func TestAsIntMatrix(t *testing.T) {
	matrix := AsIntMatrix("12300\n45600\n78900")
	assert.Equal(t, [][]int{{1, 2, 3, 0, 0}, {4, 5, 6, 0, 0}, {7, 8, 9, 0, 0}}, matrix.data)
	assert.Equal(t, 3, matrix.SizeY)
	assert.Equal(t, 5, matrix.SizeX)
	assert.Equal(t, true, matrix.Contains(Position{X: 0, Y: 1}))
	assert.Equal(t, false, matrix.Contains(Position{X: 0, Y: -1}))
	assert.Equal(t, true, matrix.Contains(Position{X: 4, Y: 0}))
	assert.Equal(t, 0, matrix.Value(Position{X: 4, Y: 0}))
}
