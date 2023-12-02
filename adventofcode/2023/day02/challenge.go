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

func process(data []string) (int, int) {
	luckies := []int{}
	powers := []int{}
	for i, line := range data {
		raw := strings.TrimSpace(strings.Split(line, ":")[1])
		cubes := strings.Split(raw, ";")
		lucky := true
		max := map[string]int{
			"blue":  0,
			"green": 0,
			"red":   0,
		}
		for _, play := range cubes {
			cubes := strings.Split(play, ",")
			for _, cube := range cubes {
				colors := strings.Split(strings.TrimSpace(cube), " ")
				amount := utils.Atoi(colors[0])
				color := colors[1]
				if amount > max[color] {
					max[color] = amount
				}
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
		power := 1
		for _, v := range max {
			power = power * v
		}
		powers = append(powers, power)
		//fmt.Printf("Game %v, %v\n", i+1, power)
	}
	luckiesSum := lo.Reduce(luckies, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)
	powersSum := lo.Reduce(powers, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)

	return luckiesSum, powersSum
}

func part1(data []string) int {
	luckies, _ := process(data)
	return luckies
}

func part2(data []string) int {
	_, power := process(data)
	return power
	return 0
}
