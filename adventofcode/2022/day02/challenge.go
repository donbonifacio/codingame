package main

import (
	"fmt"
	"os"
	"strings"
)

const challengeId = "02"

func main() {
	fmt.Printf("Challenge %v\n", challengeId)
}

func sumScore(data string, calculator func(string, string) int) int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	sum := 0
	for _, line := range lines {
		plays := strings.Split(strings.TrimSpace(line), " ")
		sum += calculator(plays[0], plays[1])
	}
	return sum
}

var resultMapPart1 = map[string]int{
	"AX": 3 + 1,
	"BX": 0 + 1,
	"CX": 6 + 1,

	"AY": 6 + 2,
	"BY": 3 + 2,
	"CY": 0 + 2,

	"AZ": 0 + 3,
	"BZ": 6 + 3,
	"CZ": 3 + 3,
}

func part1calculator(player1 string, player2 string) int {
	return resultMapPart1[fmt.Sprintf("%v%v", player1, player2)]
}

var resultMapPart2 = map[string]string{
	"AX": "Z",
	"BX": "X",
	"CX": "Y",

	"AY": "X",
	"BY": "Y",
	"CY": "Z",

	"AZ": "Y",
	"BZ": "Z",
	"CZ": "X",
}

func part2calculator(player1 string, player2 string) int {
	newPlayer2 := resultMapPart2[fmt.Sprintf("%v%v", player1, player2)]
	return part1calculator(player1, newPlayer2)
}

func readInput(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}
