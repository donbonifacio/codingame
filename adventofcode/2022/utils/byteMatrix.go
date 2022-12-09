package utils

import (
	"fmt"
	"strings"
)

type ByteMatrix struct {
	Data  [][]byte
	SizeX int
	SizeY int
}

func (matrix *ByteMatrix) Set(pos Position, value byte) {
	matrix.Data[pos.Y][pos.X] = value
}

func (matrix *ByteMatrix) Contains(pos Position) bool {
	return pos.X >= 0 && pos.Y >= 0 && pos.X < matrix.SizeX && pos.Y < matrix.SizeY
}

func (matrix *ByteMatrix) ToString() string {
	writer := new(strings.Builder)
	for y := 0; y < matrix.SizeY; y++ {
		for x := 0; x < matrix.SizeX; x++ {
			writer.WriteString(string(matrix.Value(Position{X: x, Y: y})))
		}
		writer.WriteString("\n")
	}
	return strings.TrimSpace(writer.String())
}

func (matrix *ByteMatrix) Value(pos Position) byte {
	if !matrix.Contains(pos) {
		panic(fmt.Sprintf("Matrix doesn't contain position %v - %v", pos, matrix))
	}
	return matrix.Data[pos.Y][pos.X]
}

func BuildByteMatrix(sy int, sx int) *ByteMatrix {
	matrix := ByteMatrix{SizeX: sx, SizeY: sy}
	matrix.Data = make([][]byte, sy)
	for i := range matrix.Data {
		matrix.Data[i] = make([]byte, sx)
		for j := 0; j < sx; j++ {
			matrix.Data[i][j] = '.'
		}
	}
	return &matrix
}

func AsByteMatrix(data string) *ByteMatrix {
	lines := AsLines(data)

	matrix := BuildByteMatrix(len(lines), len(lines[0]))

	for i, line := range lines {
		digits := Split(line, "")
		for c, b := range digits {
			matrix.Set(Position{c, i}, b[0])
		}
	}

	return matrix
}
