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

type Card struct {
	id      int
	winning []int
	numbers []int
	matches []int
	worth   int
}

func readInput(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return utils.AsLines(string(data))
}

func parseCard(raw string, index int) Card {
	card := Card{id: index + 1}
	gameInfo := strings.Split(raw, ":")
	parts := strings.Split(gameInfo[1], "|")

	rawWinning := lo.Filter(strings.Split(parts[0], " "), func(number string, _ int) bool { return len(number) > 0 })
	card.winning = lo.Map(rawWinning, func(number string, _ int) int { return utils.Atoi(number) })

	rawNumbers := lo.Filter(strings.Split(parts[1], " "), func(number string, _ int) bool { return len(number) > 0 })
	card.numbers = lo.Map(rawNumbers, func(number string, _ int) int { return utils.Atoi(number) })

	return card
}

func loadMatches(card Card, _ int) Card {
	card.matches = lo.Filter(card.numbers, func(number int, _ int) bool {
		return lo.Contains(card.winning, number)
	})
	return card
}

func calculateWorth(card Card, _ int) Card {
	if len(card.matches) == 0 {
		card.worth = 0
		return card
	}

	if len(card.matches) == 1 {
		card.worth = 1
		return card
	}

	card.worth = lo.Reduce(card.matches[1:], func(agg int, item int, _ int) int {
		return agg * 2
	}, 1)
	return card
}

func part1(data []string) int {
	cards := lo.Map(data, parseCard)
	cards = lo.Map(cards, loadMatches)
	cards = lo.Map(cards, calculateWorth)

	lo.ForEach(cards, func(card Card, _ int) {
		fmt.Printf("Card %v: matches:%v worth:%v\n", card.id, card.matches, card.worth)
	})

	return lo.Reduce(cards, func(agg int, card Card, _ int) int {
		return agg + card.worth
	}, 0)
}

func part2(data []string) int {
	return 0
}
