package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"

	"github.com/donbonifacio/codingame/blob/master/adventofcode/2023/utils"
	"github.com/samber/lo"
)

func main() {
	fmt.Printf("Challenge %v\n", utils.Atoi("1"))
	part2(readInput("input.txt"))
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
	cache map[string]int
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

func parseSeeds2(data *Data, line string) *Data {
	raw := strings.Split(line, ":")
	rawSeeds := strings.Split(strings.TrimSpace(raw[1]), " ")
	intervals := lo.Map(rawSeeds, func(item string, _ int) int {
		return utils.Atoi(item)
	})
	seeds := []int{}
	for i := 0; i < len(intervals); i += 2 {
		for s := intervals[i]; s <= intervals[i]+intervals[i+1]; s++ {
			seeds = append(seeds, s)
		}
		break // TODO
	}
	data.seeds = seeds
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
		fmt.Printf("Cache: %v\n", data.cache)
	}
	fmt.Println()
}

func processMap(data *Data, curr string, nextVal int) int {
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
	return nextVal
	//fmt.Printf("%v: %v ", mapper.destination, nextVal)
}

func locationFor(data *Data, seed int) int {
	flow := []string{"seed", "soil", "fertilizer", "water", "light", "temperature", "humidity", "location"}
	return locationForStep(data, flow, seed)
}

func locationForStep(data *Data, flow []string, nextVal int) int {
	curr := flow[0]
	cacheKey := fmt.Sprintf("%v-%v", curr, nextVal)
	if location, ok := data.cache[cacheKey]; ok {
		return location
	}
	mapper := data.maps[curr]
	nextVal = processMap(data, curr, nextVal)
	if mapper.destination == "location" {
		return nextVal
	}

	location := locationForStep(data, flow[1:], nextVal)
	data.cache[cacheKey] = location
	return location
}

func process(data *Data) int {
	bestLocation := math.MaxInt32

	lo.ForEach(data.seeds, func(seed int, idx int) {
		location := locationFor(data, seed)
		if location < bestLocation {
			bestLocation = location
		}
	})

	//printData(data)
	return bestLocation
}

func part1(raw string) int {
	lines := strings.Split(raw, "\n\n")
	data := &Data{}
	data = parseSeeds(data, lines[0])
	lo.ForEach(lines[1:], func(m string, _ int) {
		parseMap(data, m)
	})

	return process(data)
}

func part2(raw string) int {
	lines := strings.Split(raw, "\n\n")
	data := &Data{cache: map[string]int{}}
	data = parseSeeds2(data, lines[0])
	lo.ForEach(lines[1:], func(m string, _ int) {
		parseMap(data, m)
	})

	return process(data)
}
