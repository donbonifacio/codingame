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
		safe := isSafe(report, 0)
		fmt.Printf("____ %v -> %v", report, safe)
		if safe {
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
			fmt.Println("false by dir ", diff, dir, i)
			if allowedBadLevels > 0 {
				newReport := withoutIndex(report, i)
				newReport2 := withoutIndex(report, i-1)
				newReport3 := withoutIndex(report, i-2)
				return isSafe(newReport, 0) || isSafe(newReport2, 0) || isSafe(newReport3, 0)
			}
			return false
		}
		if diff < 0 {
			diff = -diff
		}
		if diff < 1 || diff > 3 {
			badLevels++
			if allowedBadLevels > 0 {
				newReport := withoutIndex(report, i)
				newReport2 := withoutIndex(report, i-1)
				newReport3 := withoutIndex(report, i-2)
				return isSafe(newReport, 0) || isSafe(newReport2, 0) || isSafe(newReport3, 0)
			}
			fmt.Println("false by levels ", diff)
			return false
		}
		prev = curr
	}

	fmt.Println("true", badLevels, badLevels <= allowedBadLevels)
	return badLevels <= allowedBadLevels
}

func withoutIndex(arr []string, index int) []string {
	if index < 0 || index >= len(arr) {
		return arr
	}
	newArr := []string{}
	newArr = append(newArr, arr[:index]...)
	newArr = append(newArr, arr[index+1:]...)
	return newArr
}

func part2(data []string) int {
	count := 0
	for _, line := range data {
		report := strings.Fields(line)
		safe := isSafe(report, 1)
		fmt.Printf("____ %v -> %v\n", report, safe)
		if safe {
			count++
		}

	}
	return count
}
