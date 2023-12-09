package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/donbonifacio/codingame/blob/master/adventofcode/2023/utils"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

func main() {
	fmt.Printf("Challenge %v\n", utils.Atoi("1"))
}

type RangeMap struct {
	sourceStart      int
	destinationStart int
	rangeLength      int
}

func (rm *RangeMap) Map(value int) int {
	//fmt.Printf("---- Here %v %v from %v to %v\n", rm, value, rm.sourceStart, rm.sourceStart+rm.rangeLength)
	if value >= rm.sourceStart && value <= rm.sourceStart+rm.rangeLength {
		match := rm.destinationStart + value - rm.sourceStart
		//fmt.Printf("------> %v %v -> %v\n", rm, value, match)
		return match
	}
	return -1
}

type SourceDestinationMap struct {
	source      string
	destination string
	rangeMaps   []RangeMap
}

type Data struct {
	seeds []int
	maps  map[string]SourceDestinationMap
}

func readInput(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func parseSeeds(data *Data, line string) *Data {
	raw := strings.Split(line, ":")
	rawSeeds := strings.Split(strings.TrimSpace(raw[1]), " ")
	data.seeds = lo.Map(rawSeeds, func(item string, _ int) int {
		return utils.Atoi(item)
	})
	return data
}

func parseMap(data *Data, raw string) *Data {
	lines := strings.Split(raw, "\n")
	array := regexp.MustCompile("[-, ]+").Split(lines[0], -1)

	if data.maps == nil {
		data.maps = map[string]SourceDestinationMap{}
	}

	source := array[0]
	destination := array[2]
	rangeMaps := []RangeMap{}

	lo.ForEach(lines[1:], func(line string, _ int) {
		if line != "" {
			raw := strings.Split(line, " ")
			rangeMaps = append(rangeMaps, RangeMap{
				destinationStart: utils.Atoi(raw[0]),
				sourceStart:      utils.Atoi(raw[1]),
				rangeLength:      utils.Atoi(raw[2]),
			})
		}
	})

	data.maps[source] = SourceDestinationMap{
		source:      source,
		destination: destination,
		rangeMaps:   rangeMaps,
	}

	return data
}

func printData(data *Data) {
	fmt.Printf("Seeds: %v\n", data.seeds)
	for _, v := range data.maps {
		fmt.Printf("%v to %v\n", v.source, v.destination)
		for _, rm := range v.rangeMaps {
			fmt.Printf("%v %v %v\n", rm.destinationStart, rm.sourceStart, rm.rangeLength)
		}
		fmt.Println()
	}
	fmt.Println()
}

func part1(raw string) int {
	lines := strings.Split(raw, "\n\n")
	data := &Data{}
	data = parseSeeds(data, lines[0])
	lo.ForEach(lines[1:], func(m string, _ int) {
		parseMap(data, m)
	})

	flow := []string{"seed", "soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}
	locations := []int{}

	lo.ForEach(data.seeds, func(seed int, _ int) {
		fmt.Printf("seed: %v ", seed)
		nextVal := seed
		for _, curr := range flow {
			mapper := data.maps[curr]
			matches := lo.Map(mapper.rangeMaps, func(rm RangeMap, _ int) int {
				return rm.Map(nextVal)
			})
			luckies := lo.Filter(matches, func(val int, _ int) bool {
				return val >= 0
			})
			if len(luckies) > 0 {
				nextVal = luckies[0]
			}
			fmt.Printf("%v: %v ", mapper.destination, nextVal)
			if mapper.destination == "location" {
				locations = append(locations, nextVal)
			}
		}
		fmt.Println()
	})

	slices.Sort(locations)
	fmt.Printf("locations: %v", locations[0])
	//printData(data)
	return locations[0]
}

func part2(data string) int {
	return 0
}
