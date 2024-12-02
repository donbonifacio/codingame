package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/donbonifacio/codingame/blob/master/adventofcode/2023/utils"
)

func main() {
	fmt.Printf("Challenge %v\n", utils.Atoi("1"))
}

func readInput(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return utils.AsLines(string(data))
}

func part1(data []string) int {
	count := 0
	for _, line := range data {
		report := strings.Fields(line)
		if isSafe(report) {
			count++
		}

	}
	return count
}

func isSafe(report []string) bool {
	fmt.Print("Report:", report)
	prev := report[0]
	firstDiff := false
	dir := 1
	for i := 1; i < len(report); i++ {
		curr := report[i]
		diff := utils.Atoi(prev) - utils.Atoi(curr)
		if firstDiff == false {
			firstDiff = true
			if diff < 0 {
				dir = -1
			}
		}
		//fmt.Printf("::: %s - %s = %d :", curr, prev, diff)
		if diff < 0 && dir > 0 || diff > 0 && dir < 0 {
			fmt.Println("false by dir ", diff, dir)
			return false
		}
		if diff < 0 {
			diff = -diff
		}
		if diff < 1 || diff > 3 {
			fmt.Println("false by diff ", diff, dir)
			return false
		}
		prev = curr
	}

	fmt.Println("true", dir)
	return true
}

func part2(data []string) int {
	return 0
}
