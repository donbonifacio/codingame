package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/donbonifacio/codingame/blob/master/adventofcode/2022/utils"
)

func main() {
	fmt.Printf("Challenge %v\n", utils.Atoi("1"))
}

type Monkey struct {
	id       int
	items    []int
	op       func(int) int
	test     func(int) bool
	next     map[bool]int
	inspects int
}

func part1(data string) int {
	return monkeyBusiness(data, 20, 3)
}
func part2(data string) int {
	return 0
}

func monkeyBusiness(data string, nRounds int, attenuation int) int {
	lines := utils.AsLines(data)
	monkeyData := []string{}
	monkeys := map[int]*Monkey{}
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			monkey := buildMonkey(monkeyData)
			monkeys[monkey.id] = monkey
			monkeyData = []string{}
		} else {
			monkeyData = append(monkeyData, strings.TrimSpace(line))
		}
	}
	monkey := buildMonkey(monkeyData)
	monkeys[monkey.id] = monkey
	return process(monkeys, nRounds, attenuation)
}

func process(monkeys map[int]*Monkey, nRounds int, attenuation int) int {
	for round := 0; round < nRounds; round++ {
		//fmt.Printf("Round: %v\n", round)
		for monkeyId := 0; monkeyId < len(monkeys); monkeyId++ {
			monkey := monkeys[monkeyId]
			for itemIdx := 0; itemIdx < len(monkey.items); itemIdx++ {
				item := monkey.items[itemIdx]
				worryLevel := monkey.op(item) / attenuation
				test := monkey.test(worryLevel)
				nextMonkey := monkeys[monkey.next[test]]
				nextMonkey.items = append(nextMonkey.items, worryLevel)
				monkey.inspects += 1
				//fmt.Printf(" Monkey %v: worry:%v op:%v next:%v\n", monkeyId, item, worryLevel, nextMonkey.id)
			}
			monkey.items = []int{}
		}
		for i := 0; i < len(monkeys); i++ {
			monkey := monkeys[i]
			worry := []int{}
			for _, item := range monkey.items {
				worry = append(worry, monkey.op(item))
			}
			//fmt.Printf(" Monkey %v: %v - inspects: %v\n", i, worry, monkey.inspects)
		}
	}
	inspects := []int{}
	for i := 0; i < len(monkeys); i++ {
		inspects = append(inspects, monkeys[i].inspects)
	}
	sort.Ints(inspects)
	score := inspects[len(inspects)-1] * inspects[len(inspects)-2]
	//fmt.Printf("Score: %v, Inspects: %v\n", score, inspects)
	return score
}

var numberRegex = regexp.MustCompile(`(\d+)`)

func buildMonkey(data []string) *Monkey {
	var monkey Monkey
	monkey.inspects = 0

	rawId := numberRegex.FindAllStringSubmatch(data[0], -1)
	monkey.id = utils.Atoi(rawId[0][0])

	rawStartingItems := numberRegex.FindAllStringSubmatch(data[1], -1)
	for i := 0; i < len(rawStartingItems); i++ {
		monkey.items = append(monkey.items, utils.Atoi(rawStartingItems[i][0]))
	}

	rawOperation := utils.Split(data[2], " ")
	operator := rawOperation[4]
	opValue := rawOperation[5]
	monkey.op = func(old int) int {
		a := old
		b := 0
		if opValue == "old" {
			b = a
		} else {
			b = utils.Atoi(opValue)
		}
		if operator == "*" {
			return a * b
		} else if operator == "+" {
			return a + b
		} else {
			panic(fmt.Sprintf("Don't know how to handle '%v'", operator))
		}
	}

	rawTest := utils.Split(data[3], " ")
	if rawTest[1] != "divisible" {
		panic(fmt.Sprintf("Don't know how to handle '%v'", rawTest))
	}
	testValue := rawTest[3]
	monkey.test = func(item int) bool {
		return item%utils.Atoi(testValue) == 0
	}

	monkey.next = map[bool]int{}
	ifTrue := utils.Split(data[4], " ")
	monkey.next[true] = utils.Atoi(ifTrue[5])
	ifFalse := utils.Split(data[5], " ")
	monkey.next[false] = utils.Atoi(ifFalse[5])
	if ifTrue[1] != "true:" || ifFalse[1] != "false:" {
		panic(fmt.Sprintf("True:%v False:%v", ifTrue, ifFalse))
	}

	return &monkey
}
