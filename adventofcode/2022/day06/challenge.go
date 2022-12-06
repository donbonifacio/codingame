package main

import (
	"fmt"
	"os"
)

const challengeId = "06"

func main() {
	fmt.Printf("Challenge %v\n", challengeId)
}

func findCode(data string, markerSize int) int {
	for i := markerSize; i < len(data); i++ {
		code := data[i-markerSize : i]
		if isMarker(code, markerSize) {
			return i
		}
	}
	return 0
}

func isMarker(code string, match int) bool {
	set := map[rune]bool{}
	for _, char := range code {
		set[char] = true
	}
	return len(set) == match
}

func readInput(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}
