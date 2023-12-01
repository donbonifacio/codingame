package main

import (
	"fmt"
	"os"
	"strings"
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
	//fmt.Println(values)
	return lo.Reduce(values, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)
}

func part2(data []string) int {
	dict := [][]string{
		[]string{"seven", "7"},
		[]string{"nine", "9"},
		[]string{"one", "1"},
		[]string{"eight", "8"},
		[]string{"two", "2"},
		[]string{"three", "3"},
		[]string{"four", "4"},
		[]string{"five", "5"},
		[]string{"six", "6"},
	}
	transformed := lo.Map(data, func(line string, index int) string {
		newLine := line
		for i := 0; i < len(newLine); i++ {
			rest := newLine[i:]
			for _, v := range dict {
				if strings.HasPrefix(rest, v[0]) {
					newLine = strings.Replace(newLine, v[0], v[1], 1)
					//fmt.Printf("restt: %v, prefix: %v %v -> %v \n", rest, v[0], v[1], newLine)
				}
			}
		}

		return newLine
	})
	//fmt.Println(transformed)
	return part1(transformed)
}
