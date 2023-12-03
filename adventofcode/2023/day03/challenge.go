package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/donbonifacio/codingame/blob/master/adventofcode/2023/utils"
	"github.com/samber/lo"
)

func main() {
	fmt.Printf("Challenge %v\n", utils.Atoi("1"))
}

func readInput(filename string) *utils.ByteMatrix {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return utils.AsByteMatrix(string(data))
}

func isPart(matrix *utils.ByteMatrix, number int, start utils.Position, end utils.Position) bool {
	// top line
	iter := utils.Position{X: start.X - 1, Y: start.Y - 1}
	right := utils.GetVectorForDir("R")
	for i := start.X - 1; i <= end.X+1; i++ {
		if !matrix.Contains(iter) {
			iter = iter.Move(right)
			continue
		}
		value := string(matrix.Value(iter))
		if value != "." && !IsNumber(value) {
			//fmt.Printf("  %v is a symbol\n", value)
			return true
		}
		iter = iter.Move(right)
	}

	// bottom line
	iter = utils.Position{X: start.X - 1, Y: start.Y + 1}
	for i := start.X - 1; i <= end.X+1; i++ {
		//fmt.Printf("iter: %v dir: %v\n", iter, right)
		if !matrix.Contains(iter) {
			iter = iter.Move(right)
			continue
		}
		value := string(matrix.Value(iter))
		if value != "." && !IsNumber(value) {
			//fmt.Printf("  %v is a symbol\n", value)
			return true
		}
		iter = iter.Move(right)
	}

	// sides
	side1 := utils.Position{X: start.X - 1, Y: start.Y}
	if matrix.Contains(side1) {
		value := string(matrix.Value(side1))
		if value != "." && !IsNumber(value) {
			return true
		}
	}
	side2 := utils.Position{X: end.X + 1, Y: end.Y}
	if matrix.Contains(side2) {
		value := string(matrix.Value(side2))
		if value != "." && !IsNumber(value) {
			return true
		}
	}

	return false
}

func part1(matrix *utils.ByteMatrix) int {
	numbers := []int{}
	for y := 0; y < matrix.SizeY; y++ {
		rawNumber := ""
		startPos := utils.Position{X: 0, Y: 0}
		endPos := utils.Position{X: 0, Y: 0}
		for x := 0; x < matrix.SizeX; x++ {
			pos := utils.Position{X: x, Y: y}
			value := string(matrix.Value(pos))
			if IsNumber(value) {
				if rawNumber == "" {
					startPos = pos
				}
				rawNumber = rawNumber + value
			}
			if !IsNumber(value) && len(rawNumber) > 0 || len(rawNumber) > 0 && x+1 >= matrix.SizeX {
				number := utils.Atoi(rawNumber)
				rawNumber = ""
				endPos = utils.Position{X: x - 1, Y: y}
				//fmt.Printf("%v -> %v : %v\n", number, startPos, endPos)
				if isPart(matrix, number, startPos, endPos) {
					numbers = append(numbers, number)
				} else {
					//fmt.Printf("Not part: %v\n", number)
				}
				//return -1
			}

		}
	}
	//fmt.Println(numbers)

	return lo.Reduce(numbers, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)
}

func IsNumber(raw string) bool {
	if _, err := strconv.Atoi(raw); err != nil {
		return false
	}

	return true
}

func part2(data []string) int {
	return 0
}
