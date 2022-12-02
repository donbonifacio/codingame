package main

import (
	"fmt"
	"os"
	"strings"
)

const challengeId = "01"

func main() {
	fmt.Printf("Challenge %v\n", challengeId)
}

func sumScore(data string, calculator func(string, string) int) int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	sum := 0
	for _, line := range lines {
		plays := strings.Split(strings.TrimSpace(line), " ")
		player1 := plays[0]
		player2 := plays[1]
		sum += calculator(player1, player2)
	}
	return sum
}

var playScores = map[string]int{
	"X": 1, // R
	"Y": 2, // P
	"Z": 3, // S
}
var resultMapPart1 = map[string]int{
	"AX": 3,
	"BX": 0,
	"CX": 6,

	"AY": 6,
	"BY": 3,
	"CY": 0,

	"AZ": 0,
	"BZ": 6,
	"CZ": 3,
}

func part1calculator(player1 string, player2 string) int {
	playScore2 := playScores[player2]
	resultScore := resultMapPart1[fmt.Sprintf("%v%v", player1, player2)]

	return resultScore + playScore2
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
