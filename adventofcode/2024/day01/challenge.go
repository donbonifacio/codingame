package main

import (
	"fmt"
	"os"

	"github.com/donbonifacio/codingame/blob/master/adventofcode/2023/utils"
)

func main() {
	data := readInput("input.txt")
	result := part1(data)
	fmt.Printf("Challenge result: %v\n", result)
}

func readInput(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return utils.AsLines(string(data))
}

func part1(data []string) int {
	return 0
}

func part2(data []string) int {
	return 0
}
