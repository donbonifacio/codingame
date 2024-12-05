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

type Rule struct {
	First int
	Last  int
}

type Update struct {
	PageNumbers []int
}

type Input struct {
	Rules          []Rule
	RulesByPage    map[int][]int
	Updates        []Update
	CorrectUpdates []Update
}

func readInput(filename string) Input {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	raw := strings.Split(string(data), "\n\n")
	rawRules := raw[0]
	rawUpdates := raw[1]

	rules := []Rule{}
	for _, line := range utils.AsLines(rawRules) {
		parts := strings.Split(string(line), "|")
		rules = append(rules, Rule{First: utils.Atoi(parts[0]), Last: utils.Atoi(parts[1])})
	}

	updates := []Update{}
	for _, line := range utils.AsLines(rawUpdates) {
		parts := strings.Split(string(line), ",")
		pageNumbers := []int{}
		for _, part := range parts {
			pageNumbers = append(pageNumbers, utils.Atoi(part))
		}
		updates = append(updates, Update{PageNumbers: pageNumbers})
	}

	rulesByPage := lo.GroupBy(rules, func(rule Rule) int { return rule.First })
	rulesByPageRaw := map[int][]int{}
	for key, value := range rulesByPage {
		rulesByPageRaw[key] = lo.Map(value, func(rule Rule, _ int) int { return rule.Last })
	}

	return Input{
		Rules:       rules,
		RulesByPage: rulesByPageRaw,
		Updates:     updates,
	}
}

func findCorrectUpdates(input Input) Input {
	correct := []Update{}
	for _, update := range input.Updates {
		valid := false
		for index, pageNumber := range update.PageNumbers {
			fmt.Println(update.PageNumbers, update.PageNumbers[index+1:], index+1)
			invalid := lo.Intersect(update.PageNumbers[index+1:], input.RulesByPage[pageNumber])
			if len(invalid) == 0 && index < len(update.PageNumbers)-1 {
				valid = true
			}
		}
		if valid {
			correct = append(correct, update)
		}
	}
	input.CorrectUpdates = correct
	return input
}

func part1(input Input) int {
	input = findCorrectUpdates(input)
	fmt.Println(input.CorrectUpdates)
	return len(input.CorrectUpdates)
}

func part2(input Input) int {
	return 0
}
