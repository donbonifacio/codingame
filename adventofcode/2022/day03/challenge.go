package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

const challengeId = "01"

func main() {
	fmt.Printf("Challenge %v\n", challengeId)
}

func totalPriorities(data string) int {
	sum := 0
	for _, line := range strings.Split(strings.TrimSpace(data), "\n") {
		inBoth := findInBoth(strings.TrimSpace(line))
		sum += score(inBoth)
	}
	return sum
}

func findInBoth(contents string) string {
	size := len(contents)
	bagSize := size / 2
	bag1 := strings.Split(contents[0:bagSize], "")
	bag2 := strings.Split(contents[bagSize:size], "")

	if len(bag1) != len(bag2) {
		panic(fmt.Sprintf("Size mismatch - %v %v", bag1, bag2))
	}

	return findInBags([]string{strings.Join(bag1, ""), strings.Join(bag2, "")})
}

func part2sum(data string) int {
	lines := strings.Split(strings.TrimSpace(data), "\n")
	sum := 0
	for i := 0; i < len(lines); i += 3 {
		inBags := findInBags([]string{
			lines[i : i+1][0],
			lines[i+1 : i+2][0],
			lines[i+2 : i+3][0],
		})
		sum += score(inBags)
	}
	return sum
}

func findInBags(bags []string) string {
	collector := map[string]int{}

	for _, bag := range bags {
		cache := map[string]bool{}
		for _, raw := range bag {
			item := string(raw)
			if _, ok := cache[item]; !ok {
				cache[item] = true
				collector[item] += 1
				if collector[item] == len(bags) {
					return item
				}
			}
		}
	}

	panic(fmt.Sprintf("Didn't find the outlier - %v", collector))
}

func score(item string) int {
	if unicode.IsUpper(rune(item[0])) {
		return int(item[0]) - int('A') + 27
	}
	return int(item[0]) - int('a') + 1
}

func readInput(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}
