package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/donbonifacio/codingame/blob/master/adventofcode/2022/utils"
)

func main() {
	fmt.Printf("Challenge %v\n", utils.Atoi("1"))
}

type Monkey struct {
	id    int
	items []int
	op    func(int) int
	test  func(int) int
}

func part1(data string) int {
	return monkeyBusiness(data, 20, func(v int) int { return v / 3 })
}
func part2(data string) int {
	return 0
}

func monkeyBusiness(data string, nRounds int, attenuation func(int) int) int {
	blocks := strings.Split(strings.TrimSpace(data), "\n\n")
	monkeys := map[int]*Monkey{}
	for _, block := range blocks {
		monkey := scanMonkey(block)
		monkeys[monkey.id] = monkey
	}
	return process(monkeys, nRounds, attenuation)
}

func scanMonkey(data string) *Monkey {
	var monkey Monkey
	var items, op string
	var value, divisible, trueNext, falseNext int
	data2 := strings.NewReplacer(", ", ",", "* old", "^ 2").Replace(data)
	fmt.Sscanf(data2, utils.ReadInput("monkeyTemplate.txt"), &monkey.id, &items, &op, &value, &divisible, &trueNext, &falseNext)

	json.Unmarshal([]byte("["+items+"]"), &monkey.items)

	monkey.op = map[string]func(int) int{
		"+": func(old int) int { return old * value },
		"*": func(old int) int { return old * value },
		"^": func(old int) int { return old * old },
	}[op]

	monkey.test = func(old int) int {
		if old%divisible == 0 {
			return trueNext
		}
		return falseNext
	}

	return &monkey
}

func process(monkeys map[int]*Monkey, nRounds int, attenuation func(int) int) int {
	inspected := make([]int, len(monkeys))
	for round := 0; round < nRounds; round++ {
		for monkeyId := 0; monkeyId < len(monkeys); monkeyId++ {
			monkey := monkeys[monkeyId]
			for itemIdx := 0; itemIdx < len(monkey.items); itemIdx++ {
				item := monkey.items[itemIdx]
				worryLevel := attenuation(monkey.op(item))
				nextMonkey := monkeys[monkey.test(worryLevel)]
				nextMonkey.items = append(nextMonkey.items, worryLevel)
				inspected[monkeyId]++
			}
			monkey.items = nil
		}
	}
	sort.Ints(inspected)
	score := inspected[len(inspected)-1] * inspected[len(inspected)-2]
	//fmt.Printf("Score: %v, Inspects: %v\n", score, inspects)
	return score
}
