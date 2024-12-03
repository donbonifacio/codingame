package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/donbonifacio/codingame/blob/master/adventofcode/2023/utils"
)

func main() {
	fmt.Printf("Challenge %v\n", utils.Atoi("1"))
}

func readInput(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func part1(data string) int {
	pattern := `mul\((\d+),(\d+)\)`
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		panic("regex")
	}

	matches := re.FindAllStringSubmatch(data, -1)
	fmt.Println("Matches found:")
	sum := 0
	for _, match := range matches {
		fmt.Println(match)
		sum += utils.Atoi(match[1]) * utils.Atoi(match[2])
	}
	return sum
}

func part2(data string) int {
	pattern := `don't\(\).*(mul\((\d+),(\d+)\))+.*(do\(\))+?`
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		panic("regex")
	}

	parts := utils.Split(data, "\n")
	sum := 0
	for _, rawpart := range parts {
		part := rawpart
		matches := re.FindAllStringSubmatch(part, -1)
		fmt.Println("Matches found:")
		for _, match := range matches {
			fmt.Println(match)
			//sum += utils.Atoi(match[1]) * utils.Atoi(match[2])
			part = strings.Replace(part, match[0], "", -1)
		}
		fmt.Println("Processed", part)
		sum += part1(part)
	}
	return sum
}
