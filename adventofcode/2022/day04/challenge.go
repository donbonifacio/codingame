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
	lines := strings.Split(strings.TrimSpace(data), "\n")
	sum := 0
	for _, line := range lines {
		pairs := strings.Split(strings.TrimSpace(line), ",")
		if overlaps(pairs[0], pairs[1]) || overlaps(pairs[1], pairs[0]) {
			//fmt.Printf("Overlaps: %v\n", pairs)
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
