package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/donbonifacio/codingame/blob/master/adventofcode/2023/utils"
	"github.com/samber/lo"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func main() {
	fmt.Printf("Challenge %v\n", utils.Atoi("1"))
}

type HandBid struct {
	hand     string
	numbers  []int
	handType int
	bid      int
}

const (
	HighCard int = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func readInput(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return utils.AsLines(string(data))
}

func parseHandBid(line string, _ int) HandBid {
	parts := strings.Split(line, " ")
	return HandBid{hand: parts[0], bid: utils.Atoi(parts[1])}
}

func setType(handBid HandBid, _ int) HandBid {
	handBid.handType = matchType(handBid)
	return handBid
}

func setTypeJoker(handBid HandBid, _ int) HandBid {
	handBid.handType = matchTypeJoker(handBid)
	return handBid
}

func first(m map[int][]int) []int {
	for _, v := range m {
		return v
	}
	return nil
}
func matchTypeJoker(handBid HandBid) int {
	if len(handBid.numbers) == 0 {
		handBid = setNumbersJoker(handBid, 0)
	}
	if !strings.Contains(handBid.hand, "J") {
		return matchType(handBid)
	}
	groups := lo.GroupBy(handBid.numbers, func(i int) int {
		return i
	})
	biggerGroup := -1
	target := -1
	lo.ForEach(lo.Values(groups), func(group []int, _ int) {
		if len(group) > biggerGroup && group[0] != 0 {
			biggerGroup = len(group)
			target = group[0]
		}
	})

	oldNumbers := handBid.numbers
	newNumbers := lo.ReplaceAll(handBid.numbers, 0, target)

	/*
		fmt.Println(oldNumbers)
		fmt.Println(newNumbers)
		fmt.Println(target)
	*/
	handBid.numbers = newNumbers
	typeWithJoker := matchType(handBid)
	handBid.numbers = oldNumbers
	return typeWithJoker
}

func matchType(handBid HandBid) int {
	if len(handBid.numbers) == 0 {
		handBid = setNumbers(handBid, 0)
	}
	groups := lo.GroupBy(handBid.numbers, func(i int) int {
		return i
	})
	if len(groups) == 1 {
		return FiveOfAKind
	}
	if len(groups) == 2 {
		s := len(first(groups))
		if s == 1 || s == 4 {
			return FourOfAKind
		}
		if s == 2 || s == 3 {
			return FullHouse
		}
	}
	if len(groups) == 3 {
		sizes := lo.Map(maps.Values(groups), func(group []int, _ int) int {
			return len(group)
		})
		sort.Ints(sizes)
		if sizes[len(sizes)-1] == 3 {
			return ThreeOfAKind
		}

		return TwoPair
	}
	if len(groups) == 4 {
		return OnePair
	}
	if len(groups) == 5 {
		return HighCard
	}

	return HighCard
}

func setNumbersJoker(handBid HandBid, _ int) HandBid {
	return setNumbersFunc(handBid, 0, true)
}

func setNumbers(handBid HandBid, _ int) HandBid {
	return setNumbersFunc(handBid, 0, false)
}

func setNumbersFunc(handBid HandBid, _ int, joker bool) HandBid {
	// A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2
	handBid.numbers = lo.Map([]byte(handBid.hand), func(b byte, _ int) int {
		n := -1
		if b == 'A' {
			n = 14
		} else if b == 'K' {
			n = 13
		} else if b == 'Q' {
			n = 12
		} else if b == 'J' {
			if joker {
				n = 0
			} else {
				n = 11
			}
		} else if b == 'T' {
			n = 10
		} else {
			n = utils.Atoi(string(b))
		}
		return n
	})
	return handBid
}

func compareHands(a, b HandBid) bool {
	if a.handType != b.handType {
		return a.handType < b.handType
	}
	for i := 0; i < len(a.numbers); i++ {
		c1 := a.numbers[i]
		c2 := b.numbers[i]
		if c1 != c2 {
			return c1 < c2
		}
	}
	return false
}

func part1(lines []string) int {
	handBids := lo.Map(lines, parseHandBid)
	handBids = lo.Map(handBids, setNumbers)
	handBids = lo.Map(handBids, setType)
	slices.SortFunc(handBids, compareHands)

	winnings := lo.Map(handBids, func(h HandBid, index int) int {
		fmt.Printf("Hand: %v\n", h.hand)
		return h.bid * (index + 1)
	})

	return lo.Reduce(winnings, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)
}

func part2(lines []string) int {
	handBids := lo.Map(lines, parseHandBid)
	handBids = lo.Map(handBids, setNumbersJoker)
	handBids = lo.Map(handBids, setTypeJoker)
	slices.SortFunc(handBids, compareHands)

	winnings := lo.Map(handBids, func(h HandBid, index int) int {
		fmt.Printf("Hand: %v type: %v\n", h.hand, h.handType)
		return h.bid * (index + 1)
	})

	return lo.Reduce(winnings, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)
}
