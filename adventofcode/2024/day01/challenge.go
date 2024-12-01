package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

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

func extractArrays(data []string) ([]int, []int) {
	var left []int
	var right []int
	for _, line := range data {
		parts := strings.Fields(line)
		leftNum := utils.Atoi(parts[0])
		rightNum := utils.Atoi(parts[1])
		left = append(left, leftNum)
		right = append(right, rightNum)
	}
	return left, right
}

func part1(data []string) int {
	left, right := extractArrays(data)
	sort.Ints(left)
	sort.Ints(right)
	sum := 0
	for i := range left {
		diff := left[i] - right[i]
		if diff < 0 {
			diff = -diff
		}
		sum += diff
	}
	return sum
}

func part2(data []string) int {
	left, right := extractArrays(data)
	sum := 0
	for _, leftNum := range left {
		count := 0
		for _, rightNum := range right {
			if leftNum == rightNum {
				count++
			}
		}
		sum += leftNum * count
	}
	return sum
}
