package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const challengeId = "01"

func main() {
	fmt.Printf("Challenge %v\n", challengeId)
}

func caloriesByElf(data string) []int {
	lines := strings.Split(data, "\n")
	elfs := []int{}
	sum := 0

	for _, rawLine := range lines {
		line := strings.TrimSpace(rawLine)
		if line == "" {
			elfs = append(elfs, sum)
			sum = 0
		} else {
			num, _ := strconv.Atoi(line)
			sum += num
		}
	}
	return append(elfs, sum)
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

func top3(calories []int) int {
	sort.Ints(calories)
	size := len(calories)
	sum := 0
	for _, num := range calories[size-3 : size] {
		sum += num
	}
	return sum
}

func readInput() string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(data)
}
