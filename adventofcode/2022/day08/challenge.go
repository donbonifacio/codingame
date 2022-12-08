package main

import (
	"fmt"

	"github.com/donbonifacio/codingame/blob/master/adventofcode/2022/utils"
)

func main() {
	fmt.Printf("Challenge %v %v\n", utils.Atoi("1"), utils.AsLines("a\nb"))
}

func part1(data string) int {
	lines := utils.AsLines(data)

	matrix := make([][]int, len(lines))
	for i, elem := range matrix {
		matrix[i] = make([]int, len(elem))
	}

	for i, line := range lines {
		digits := utils.Split(line, "")
		for _, digit := range digits {
			matrix[i] = append(matrix[i], utils.Atoi(digit))
		}
	}

	fromInside := visibleFromInside(matrix)

	return (len(lines) * 2) + ((len(lines[0]) - 2) * 2) + fromInside
}

func visibleFromInside(matrix [][]int) int {
	sum := 0
	for y := 1; y < len(matrix)-1; y++ {
		for x := 1; x < len(matrix[0])-1; x++ {
			b := visibleFrom(matrix, matrix[x][y], []int{x, y}, []int{-1, 0}) ||
				visibleFrom(matrix, matrix[x][y], []int{x, y}, []int{1, 0}) ||
				visibleFrom(matrix, matrix[x][y], []int{x, y}, []int{0, 1}) ||
				visibleFrom(matrix, matrix[x][y], []int{x, y}, []int{0, -1})
			if b {
				sum += 1
			}
			//fmt.Printf("Pos %v (%v) -> %v\n", []int{x, y}, matrix[x][y], b)
		}
	}

	return sum
}

func visibleFrom(matrix [][]int, value int, pos []int, vector []int) bool {
	newPos := []int{pos[0] + vector[0], pos[1] + vector[1]}

	if newPos[0] < 0 || newPos[1] < 0 || newPos[0] >= len(matrix) || newPos[1] >= len(matrix[0]) {
		return true
	}

	isVisible := matrix[newPos[0]][newPos[1]] < value
	if !isVisible {
		return false
	}

	return visibleFrom(matrix, value, newPos, vector)
}

func part2(data string) int {
	lines := utils.AsLines(data)

	matrix := make([][]int, len(lines))
	for i, elem := range matrix {
		matrix[i] = make([]int, len(elem))
	}

	for i, line := range lines {
		digits := utils.Split(line, "")
		for _, digit := range digits {
			matrix[i] = append(matrix[i], utils.Atoi(digit))
		}
	}

	return scenicScore(matrix)
}

func scenicScore(matrix [][]int) int {
	max := 0
	for y := 1; y < len(matrix)-1; y++ {
		for x := 1; x < len(matrix[0])-1; x++ {
			a := directionScore(matrix, matrix[x][y], []int{x, y}, []int{-1, 0}, 0)
			b := directionScore(matrix, matrix[x][y], []int{x, y}, []int{1, 0}, 0)
			c := directionScore(matrix, matrix[x][y], []int{x, y}, []int{0, 1}, 0)
			d := directionScore(matrix, matrix[x][y], []int{x, y}, []int{0, -1}, 0)

			score := a * b * c * d
			if score > max {
				max = score
			}
			//fmt.Printf("Pos %v (%v) -> %v (%v, %v, %v, %v)\n", []int{x, y}, matrix[x][y], score, a, b, c, d)
		}
	}

	return max
}

func directionScore(matrix [][]int, value int, pos []int, vector []int, counter int) int {
	newPos := []int{pos[0] + vector[0], pos[1] + vector[1]}

	if newPos[0] < 0 || newPos[1] < 0 || newPos[0] >= len(matrix) || newPos[1] >= len(matrix[0]) {
		return counter
	}

	isVisible := matrix[newPos[0]][newPos[1]] < value
	if !isVisible {
		return counter + 1
	}

	return directionScore(matrix, value, newPos, vector, counter+1)
}
