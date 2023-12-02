package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/donbonifacio/codingame/blob/master/adventofcode/2023/utils"
	"github.com/samber/lo"
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
	luckies := []int{}
	for i, line := range data {
		raw := strings.TrimSpace(strings.Split(line, ":")[1])
		cubes := strings.Split(raw, ";")
		lucky := true
		for _, play := range cubes {
			cubes := strings.Split(play, ",")
			for _, cube := range cubes {
				colors := strings.Split(strings.TrimSpace(cube), " ")
				amount := utils.Atoi(colors[0])
				color := colors[1]
				if color == "red" && amount > 12 {
					lucky = false
				} else if color == "green" && amount > 13 {
					lucky = false
				} else if color == "blue" && amount > 14 {
					lucky = false
				}
			}
		}
		if lucky {
			//fmt.Printf("Game %v\n", i+1)
			luckies = append(luckies, i+1)
		}
	}
	return lo.Reduce(luckies, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)
}

func part2(data []string) int {
	return 0
}
