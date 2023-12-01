package main

import (
	"fmt"
	"os"
	"unicode"

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
	values := lo.Map(data, func(line string, index int) int {
		first, last := "", ""
		for _, c := range line {
			if unicode.IsDigit(c) {
				if first == "" {
					first = string(c)
				}
				last = string(c)
			}
		}
		return utils.Atoi(first + last)
	})
	return lo.Reduce(values, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)
}
