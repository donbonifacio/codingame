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
		if isSafe(report, 0) {
			count++
		}

	}
	return count
}

func isSafe(report []string, allowedBadLevels int) bool {
	fmt.Print("Report:", report)
	prev := report[0]
	firstDiff := false
	dir := 1
	badLevels := 0
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
			if allowedBadLevels > 0 {
				fmt.Printf("--- %v - %v ----\n", report[:i], report[i+1:])
				newReport := append([]int{}, report[:i]..., report[i+1:]...)
				fmt.Printf("--- %v - %v - %v ----\n", report[:i], report[i+1:], newReport)
				newReport2 := append(report[:i-1], report[i:]...)
				fmt.Printf("--- newReport: %v newReport2: %v index: %v ----\n", newReport, newReport2, i)
				return isSafe(newReport, 0) || isSafe(newReport2, 0)
			}
			return false
		}
		if diff < 0 {
			diff = -diff
		}
		if diff < 1 || diff > 3 {
			badLevels++
			if allowedBadLevels > 0 {
				newReport := append(report[:i], report[i+1:]...)
				return isSafe(newReport, 0)
			}
			fmt.Println("false by levels ", diff)
			return false
		}
		prev = curr
	}

	fmt.Println("true", badLevels, badLevels <= allowedBadLevels)
	return badLevels <= allowedBadLevels
}

func part2(data []string) int {
	count := 0
	for _, line := range data {
		report := strings.Fields(line)
		if isSafe(report, 1) {
			count++
		}

	}
	return count
}
