package main

import (
	"fmt"

	"github.com/donbonifacio/codingame/blob/master/adventofcode/2022/utils"
)

func main() {
	fmt.Printf("Challenge %v %v\n", utils.Atoi("1"), utils.AsLines("a\nb"))
}

func part1(data string) int {
	matrix := utils.AsIntMatrix(data)
	fromInside := visibleFromInside(matrix)

	return (matrix.SizeY * 2) + ((matrix.SizeX - 2) * 2) + fromInside
}

func visibleFromInside(matrix *utils.IntMatrix) int {
	sum := 0
	for y := 1; y < matrix.SizeY-1; y++ {
		for x := 1; x < matrix.SizeX-1; x++ {
			pos := utils.Position{X: x, Y: y}
			visible := false
			utils.EachVectorXY(func(vector utils.Vector) {
				visible = visible || visibleFrom(matrix, matrix.Value(pos), pos, vector)
			})
			if visible {
				sum += 1
			}
			//fmt.Printf("Pos %v (%v) -> %v\n", []int{x, y}, matrix[x][y], b)
		}
	}

	return sum
}

func visibleFrom(matrix *utils.IntMatrix, value int, pos utils.Position, vector utils.Vector) bool {
	newPos := pos.Move(vector)
	//fmt.Printf("Pos: %v NewPos: %v Contains: %v\n", pos, newPos, matrix.Contains(newPos))
	if !matrix.Contains(newPos) {
		return true
	}

	//fmt.Printf("Value: %v\n", matrix.Value(newPos))
	isVisible := matrix.Value(newPos) < value
	if !isVisible {
		return false
	}

	return visibleFrom(matrix, value, newPos, vector)
}

func part2(data string) int {
	matrix := utils.AsIntMatrix(data)
	return scenicScore(matrix)
}

func scenicScore(matrix *utils.IntMatrix) int {
	max := 0
	for y := 1; y < matrix.SizeY-1; y++ {
		for x := 1; x < matrix.SizeX-1; x++ {
			pos := utils.Position{X: x, Y: y}
			value := matrix.Value(pos)
			score := 1
			utils.EachVectorXY(func(vector utils.Vector) {
				score *= directionScore(matrix, value, pos, vector, 0)
			})
			if score > max {
				max = score
			}
			//fmt.Printf("Pos %v (%v) -> %v (%v, %v, %v, %v)\n", []int{x, y}, matrix[x][y], score, a, b, c, d)
		}
	}

	return max
}

func directionScore(matrix *utils.IntMatrix, value int, pos utils.Position, vector utils.Vector, counter int) int {
	newPos := pos.Move(vector)

	if !matrix.Contains(newPos) {
		return counter
	}

	isVisible := matrix.Value(newPos) < value
	if !isVisible {
		return counter + 1
	}

	return directionScore(matrix, value, newPos, vector, counter+1)
}
