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
				sum += hasWord(data, "XMAS", dir, pos, "")
			}
		}
	}
	return sum
}

func hasWord(data *utils.ByteMatrix, word string, dir utils.Vector, pos utils.Position, space string) int {
	curr := data.Value(pos)
	if len(word) == 1 && word[0] == curr {
		return 1
	}
	if word[0] != curr {
		return 0
	}
	newWord := word[1:]
	newPos := pos.Move(dir)
	if data.Contains(newPos) {
		return hasWord(data, newWord, dir, newPos, space+" ")
	}
	return 0
}

func part2(data utils.ByteMatrix) int {
	return 0
}
