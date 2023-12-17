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

func checkNumber(gears map[utils.Position][]utils.Position, matrix *utils.ByteMatrix, start utils.Position, end utils.Position, lineAdd int) bool {
	iter := utils.Position{X: start.X - 1, Y: start.Y + lineAdd}
	right := utils.GetVectorForDir("R")
	symbol := false
	for i := start.X - 1; i <= end.X+1; i++ {
		if !matrix.Contains(iter) {
			iter = iter.Move(right)
			continue
		}
		value := string(matrix.Value(iter))
		if value != "." && !IsNumber(value) {
			symbol = true
			if value == "*" {
			}
		}
		iter = iter.Move(right)
	}
	return symbol
}

func isPart(gears map[utils.Position][]utils.Position, matrix *utils.ByteMatrix, number int, start utils.Position, end utils.Position) bool {
	if checkNumber(gears, matrix, start, end, -1) {
		return true
	}
	if checkNumber(gears, matrix, start, end, 1) {
		return true
	}

	if checkNumber(gears, matrix, start, end, 0) {
		return true
	}

	return false
}

func part1(matrix *utils.ByteMatrix) int {
	gears := map[utils.Position][]utils.Position{}
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
				if isPart(gears, matrix, number, startPos, endPos) {
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
