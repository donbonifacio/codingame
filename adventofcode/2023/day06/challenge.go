package main

import (
	"fmt"

	"github.com/donbonifacio/codingame/blob/master/adventofcode/2023/utils"
	"github.com/samber/lo"
)

func main() {
	fmt.Printf("Challenge %v\n", utils.Atoi("1"))
}

type Race struct {
	time      int
	distance  int
	distances []int
	records   int
}

func findRecords(race Race, _ int) int {
	for i := 0; i <= race.time; i++ {
		hold := i
		moving := race.time - i
		velocity := hold
		distance := moving * velocity
		//race.distances = append(race.distances, distance)
		if distance > race.distance {
			race.records += 1
		}
	}

	return race.records
}

func part1(races []Race) int {
	records := lo.Map(races, findRecords)
	return lo.Reduce(records, func(agg int, item int, _ int) int {
		return agg * item
	}, 1)
}

func part2(races []Race) int {
	return part1(races)
}
