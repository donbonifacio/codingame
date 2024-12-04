package main

import (
	"fmt"
	"os"

	"github.com/donbonifacio/codingame/blob/master/adventofcode/2023/utils"
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

func part1(data *utils.ByteMatrix) int {
	fmt.Println(data.ToString())
	sum := 0
	for y := 0; y < data.SizeY; y++ {
		for x := 0; x < data.SizeX; x++ {
			pos := utils.Position{X: x, Y: y}
			for _, dir := range utils.VectorsAll {
				sum += hasWord(data, "XMAS", dir, pos)
			}
		}
	}
	return sum
}

func hasWord(data *utils.ByteMatrix, word string, dir utils.Vector, pos utils.Position) int {
	if !data.Contains(pos) {
		return 0
	}
	curr := data.Value(pos)
	if len(word) == 1 && word[0] == curr {
		return 1
	}
	if word[0] != curr {
		return 0
	}
	return hasWord(data, word[1:], dir, pos.Move(dir))
}

func part2(data *utils.ByteMatrix) int {
	fmt.Println(data.ToString())
	sum := 0
	for y := 0; y < data.SizeY; y++ {
		for x := 0; x < data.SizeX; x++ {
			pos := utils.Position{X: x, Y: y}
			curr := data.Value(pos)
			if curr == 'A' {
				topLeft := pos.Move(utils.Vector{X: -1, Y: -1})
				bottomLeft := pos.Move(utils.Vector{X: -1, Y: 1})
				diagonal1 := utils.Vector{X: 1, Y: 1}
				diagonal2 := utils.Vector{X: 1, Y: -1}

				hasDiagonal1 := hasWord(data, "MAS", diagonal1, topLeft) == 1 ||
					hasWord(data, "SAM", diagonal1, topLeft) == 1
				hasDiagonal2 := hasWord(data, "MAS", diagonal2, bottomLeft) == 1 ||
					hasWord(data, "SAM", diagonal2, bottomLeft) == 1

				if hasDiagonal1 && hasDiagonal2 {
					sum += 1
				}
			}
		}
	}
	return sum
}
