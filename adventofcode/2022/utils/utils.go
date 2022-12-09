package utils

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {}

func Atoi(raw string) int {
	val, err := strconv.Atoi(raw)
	if err != nil {
		panic(err)
	}
	return val
}

func AsLines(data string) []string {
	return Split(data, "\n")
}

func Split(data string, separator string) []string {
	parts := strings.Split(strings.TrimSpace(data), separator)
	for i := 0; i < len(parts); i++ {
		parts[i] = strings.TrimSpace(parts[i])
	}
	return parts
}

func IsEmpty(data string) bool {
	return strings.TrimSpace(data) == ""
}

func EachString(data []string, op func(str string)) {
	for i := 0; i < len(data); i++ {
		item := data[i]
		op(item)
	}
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

type IntMatrix struct {
	data  [][]int
	SizeX int
	SizeY int
}

type Position struct {
	X int
	Y int
}

type Vector struct {
	X         int
	Y         int
	Intensity int
}

func BuildVector(x int, y int) Vector {
	return Vector{X: x, Y: y, Intensity: 1}
}

func (pos *Position) Move(vector Vector) Position {
	return Position{pos.X + vector.X, pos.Y + vector.Y}
}

func (pos *Position) IsAdjacent(other Position) bool {
	return math.Abs(float64(pos.X-other.X)) <= 1 && math.Abs(float64(pos.Y-other.Y)) <= 1
}

func (pos *Position) IsSideBySide(other Position) bool {
	return (pos.X == other.X && math.Abs(float64(pos.Y-other.Y)) <= 1 ||
		pos.Y == other.Y && math.Abs(float64(pos.X-other.X)) <= 1) &&
		!(pos.X == other.X && pos.Y == other.Y)
}

var VectorsXY = []Vector{
	BuildVector(-1, 0),
	BuildVector(1, 0),
	BuildVector(0, 1),
	BuildVector(0, -1),
}

var vectorByDir = map[string]Vector{
	"R": BuildVector(1, 0),
	"L": BuildVector(-1, 0),
	"U": BuildVector(0, -1),
	"D": BuildVector(0, 1),
}

func GetVectorForDir(dir string) Vector {
	return vectorByDir[dir]
}

var VectorsAll = []Vector{
	BuildVector(-1, 0),
	BuildVector(1, 0),
	BuildVector(0, 1),
	BuildVector(0, -1),

	BuildVector(-1, -1),
	BuildVector(1, 1),
	BuildVector(-1, 1),
	BuildVector(1, -1),
}

func EachVector(op func(vector Vector)) {
	for _, v := range VectorsAll {
		op(v)
	}
}

func EachVectorXY(op func(vector Vector)) {
	for _, v := range VectorsXY {
		op(v)
	}
}

func (matrix *IntMatrix) Set(pos Position, value int) {
	matrix.data[pos.Y][pos.X] = value
}

func (matrix *IntMatrix) Each(op func(pos Position, value int)) {
	for y := 0; y < matrix.SizeY; y++ {
		for x := 0; x < matrix.SizeX; x++ {
			pos := Position{X: x, Y: y}
			op(pos, matrix.Value(pos))
		}
	}
}

func (matrix *IntMatrix) EachInner(op func(pos Position, value int)) {
	matrix.Each(func(pos Position, value int) {
		if pos.X > 0 && pos.Y > 0 && pos.X < matrix.SizeX-1 && pos.Y < matrix.SizeY-1 {
			op(pos, value)
		}
	})
}

func (matrix *IntMatrix) Contains(pos Position) bool {
	return pos.X >= 0 && pos.Y >= 0 && pos.X < matrix.SizeX && pos.Y < matrix.SizeY
}

func (matrix *IntMatrix) Value(pos Position) int {
	if !matrix.Contains(pos) {
		panic(fmt.Sprintf("Matrix doesn't contain position %v - %v", pos, matrix))
	}
	return matrix.data[pos.Y][pos.X]
}

func BuildIntMatrix(sy int, sx int) *IntMatrix {
	matrix := IntMatrix{SizeX: sx, SizeY: sy}
	matrix.data = make([][]int, sy)
	for i := range matrix.data {
		matrix.data[i] = make([]int, sx)
		for j := 0; j < sx; j++ {
			matrix.data[i][j] = 0
		}
	}
	return &matrix
}

func AsIntMatrix(data string) *IntMatrix {
	lines := AsLines(data)

	matrix := BuildIntMatrix(len(lines), len(lines[0]))

	for i, line := range lines {
		digits := Split(line, "")
		for c, digit := range digits {
			matrix.Set(Position{c, i}, Atoi(digit))
		}
	}

	return matrix
}

func ReadInput(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func ReadFileWithNumbers(fileName string) []int {
	raw, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	rawString := string(raw)
	parts := strings.Split(rawString, ",")
	numbers := []int{}

	for i := 0; i < len(parts); i++ {
		part := strings.TrimSpace(parts[i])
		num, err := strconv.Atoi(part)

		if err != nil {
			panic(fmt.Sprintf("Error parsing number '%v'", num))
		}

		numbers = append(numbers, num)
	}

	return numbers
}

func WriteFile(fileName string, content string) {
	f, err := os.Create(fileName)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	_, err2 := f.WriteString(content)

	if err2 != nil {
		panic(err)
	}
}
