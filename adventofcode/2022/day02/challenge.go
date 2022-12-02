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

func calculateScore(data string) int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	sum := 0
	for _, line := range lines {
		plays := strings.Split(strings.TrimSpace(line), " ")
		player1 := plays[0]
		player2 := plays[1]
		sum += calculatePlayScore(player1, player2)
	}
	return sum
}

var playScores = map[string]int{
	"A": 1, // R
	"B": 2, // P
	"C": 3, // S

	"X": 1, // R
	"Y": 2, // P
	"Z": 3, // S
}

func calculatePlayScore(player1 string, player2 string) int {
	playScore1 := playScores[player1]
	playScore2 := playScores[player2]

	resultScore := 0
	if playScore1 == playScore2 {
		resultScore = 3
	} else if player1 == "A" && player2 == "Y" {
		resultScore = 6
	} else if player1 == "B" && player2 == "Z" {
		resultScore = 6
	} else if player1 == "C" && player2 == "X" {
		resultScore = 6
	}

	return resultScore + playScore2
}

func readInput(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}
