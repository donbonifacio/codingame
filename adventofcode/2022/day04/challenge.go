package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const challengeId = "01"

func main() {
	fmt.Printf("Challenge %v\n", challengeId)
}

func part1(data string) int {
	return run(data, overlaps)
}

func part2(data string) int {
	return run(data, partialOverlaps)
}

func run(data string, test func(string, string) bool) int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	sum := 0
	for _, line := range lines {
		pairs := strings.Split(strings.TrimSpace(line), ",")
		if test(pairs[0], pairs[1]) || test(pairs[1], pairs[0]) {
			sum += 1
		}
	}
	return sum
}

func overlaps(pair1 string, pair2 string) bool {
	a1, a2 := parseVector(pair1)
	b1, b2 := parseVector(pair2)

	if a1 <= b1 && a2 >= b2 {
		return true
	}

	return false
}

func partialOverlaps(pair1 string, pair2 string) bool {
	a1, a2 := parseVector(pair1)
	b1, b2 := parseVector(pair2)

	if b1 >= a1 && b1 <= a2 {
		return true
	}
	if b2 >= a1 && b2 <= a2 {
		return true
	}

	return false
}

func parseVector(pair string) (int, int) {
	vector := strings.Split(strings.TrimSpace(pair), "-")

	x, err := strconv.Atoi(vector[0])
	if err != nil {
		panic(err)
	}

	y, err := strconv.Atoi(vector[1])
	if err != nil {
		panic(err)
	}
	return x, y
}

func readInput(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}
