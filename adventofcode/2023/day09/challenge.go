package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/donbonifacio/codingame/blob/master/adventofcode/2023/utils"
	"github.com/samber/lo"
)

func main() {
	fmt.Printf("\nFinal: %v\n", part1(readInput("input.txt")))
}

func readInput(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return utils.AsLines(string(data))
}

type Line struct {
	numbers    []int
	iterations [][]int
}

func parseData(data []string) []Line {
	return lo.Map(data, func(raw string, _ int) Line {
		numbers := lo.Map(strings.Split(raw, " "), func(n string, _ int) int {
			return utils.Atoi(n)
		})
		line := Line{numbers: numbers, iterations: [][]int{}}
		return line
	})
}

func generateIterations(lines []Line) []Line {
	return lo.Map(lines, func(line Line, _ int) Line {
		numbers := line.numbers
		line.iterations = append(line.iterations, numbers)
		for true {
			iteration := []int{}
			for i := 0; i < len(numbers)-1; i++ {
				a := numbers[i]
				b := numbers[i+1]
				iteration = append(iteration, b-a)
			}
			line.iterations = append(line.iterations, iteration)
			uniqueValues := len(lo.FindUniques(iteration))
			if utils.Sum(iteration) == 0 && uniqueValues == 1 {
				break
			}
			numbers = iteration
		}
		return line
	})
}

func extrapolate(lines []Line) []Line {
	lo.Map(lines, func(line Line, _ int) Line {
		for i := len(line.iterations) - 2; i >= 0; i-- {
			iteration := line.iterations[i]
			previousIteration := line.iterations[i+1]
			value := iteration[len(iteration)-1] + previousIteration[len(previousIteration)-1]
			newIteration := append(line.iterations[i], value)
			line.iterations[i] = newIteration
		}
		return line
	})
	return lines
}

func extrapolateLeft(lines []Line) []Line {
	lo.Map(lines, func(line Line, _ int) Line {
		for i := len(line.iterations) - 2; i >= 0; i-- {
			iteration := line.iterations[i]
			previousIteration := line.iterations[i+1]
			value := iteration[0] - previousIteration[0]
			newIteration := append([]int{value}, line.iterations[i]...)
			line.iterations[i] = newIteration
		}
		return line
	})
	return lines
}

func printIterations(lines []Line) {
	for _, line := range lines {
		fmt.Println("--")
		for i, iteration := range line.iterations {
			for s := 0; s < i+1; s++ {
				fmt.Printf(" ")
			}
			fmt.Println(iteration)
		}
	}
}

func part1(data []string) int {
	lines := parseData(data)
	lines = generateIterations(lines)
	lines = extrapolate(lines)
	extrapolations := lo.Map(lines, func(line Line, _ int) int {
		size := len(line.iterations[0])
		fmt.Printf("%v ", line.iterations[0][size-1])
		return line.iterations[0][size-1]
	})
	printIterations(lines)
	return utils.Sum(extrapolations)
}

func part2(data []string) int {
	lines := parseData(data)
	lines = generateIterations(lines)
	lines = extrapolateLeft(lines)
	extrapolations := lo.Map(lines, func(line Line, _ int) int {
		fmt.Printf("%v ", line.iterations[0][0])
		return line.iterations[0][0]
	})
	printIterations(lines)
	return utils.Sum(extrapolations)
}
