package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/donbonifacio/codingame/blob/master/adventofcode/2022/utils"
)

func main() {
	fmt.Printf("Challenge %v\n", utils.Atoi("1"))
	data := utils.ReadInput("input.txt")
	tracker := part1(data)
	walkedMatrix := (*tracker.matrix)
	Print(&tracker)
	solution := main2()

	for i, _ := range solution {

		p := utils.Position{X: i.X - 1, Y: i.Y - 1}
		if walkedMatrix.Contains(p) {
			//walkedMatrix.Set(p, strings.ToUpper(string(tracker.matrix.Value(p)))[0])
		}
	}
	fmt.Printf("Matrix:\n%v\n", walkedMatrix.ToString())
}

type Tracker struct {
	matrix      *utils.ByteMatrix
	start       utils.Position
	end         utils.Position
	from        utils.Position
	curr        utils.Position
	quickest    []utils.Position
	won         bool
	deadend     bool
	bestTracker *Tracker
}

func part1(data string) Tracker {
	var tracker Tracker
	tracker.matrix = utils.AsByteMatrix(data)

	for x := 0; x < tracker.matrix.SizeX; x++ {
		for y := 0; y < tracker.matrix.SizeY; y++ {
			pos := utils.Position{X: x, Y: y}
			if tracker.matrix.Value(pos) == 'S' {
				tracker.start = pos
				tracker.curr = pos
				tracker.matrix.Set(pos, 'a')
				tracker.quickest = []utils.Position{pos}
			}
			if tracker.matrix.Value(pos) == 'E' {
				tracker.end = pos
				tracker.matrix.Set(pos, 'Z')
			}
		}
	}
	visited := map[utils.Position]bool{}
	bestTracker := findWays(tracker, &visited, 1)

	fmt.Printf("-------Best tracker: %v\n", len(bestTracker.quickest))
	return *bestTracker
}

func Print(tracker *Tracker) {
	fmt.Println()
	fmt.Printf("from: %v\n", tracker.from)
	fmt.Printf("curr: %v\n", tracker.curr)
	fmt.Printf("end: %v\n", tracker.end)
	fmt.Printf("deadend: %v\n", tracker.deadend)
	fmt.Printf("won: %v\n", tracker.won)
	fmt.Printf("quickest: %v\n", tracker.quickest)
	walkedMatrix := (*tracker.matrix)
	for _, point := range tracker.quickest {
		walkedMatrix.Set(point, strings.ToUpper(string(tracker.matrix.Value(point)))[0])
	}
	fmt.Printf("Matrix:\n%v\n", walkedMatrix.ToString())
}

func findWays(tracker Tracker, visited *map[utils.Position]bool, levels int) *Tracker {
	if levels == 150 {
		//return &tracker
	}
	//fmt.Printf("%vcurr %v - ", "-", tracker.curr)
	if tracker.curr == tracker.end {
		tracker.won = true
		fmt.Printf("%v - MATCH moves:%v\n", levels, len(tracker.quickest))
		return &tracker
	}
	possible := possibleFrom(&tracker, visited)
	if len(possible) == 0 { // dead end
		//fmt.Printf("%vdeadend\n", strings.Repeat(".", levels))
		tracker.deadend = true
		return nil
	}
	//fmt.Printf("%vpossible %v \n", strings.Repeat(".", levels), possible)
	for _, newWay := range possible {
		newTracker := tracker
		newTracker.from = newTracker.curr
		newTracker.curr = newWay
		newTracker.bestTracker = tracker.bestTracker
		newTracker.quickest = append(newTracker.quickest, newWay)
		if tracker.bestTracker != nil && len(newTracker.quickest) >= len(tracker.bestTracker.quickest) {
			fmt.Printf("%v - Doesn't beat best, skip:%v\n", levels, len(newTracker.quickest))
			continue
		}
		if false && tracker.bestTracker != nil && len(newTracker.quickest)+distance(newWay, tracker.end) >= len(tracker.bestTracker.quickest) {
			fmt.Printf("%v - Too far away, skip. Curr len:%v distance:%v\n", levels, len(newTracker.quickest), distance(newWay, tracker.end))
			continue
		}

		finalTracker := findWays(newTracker, visited, levels+1)
		if finalTracker != nil && (tracker.bestTracker == nil || len(finalTracker.quickest) < len(tracker.bestTracker.quickest)) {
			tracker.bestTracker = finalTracker
			if levels == 2 {
				//fmt.Printf("- level: %v - %v vs %v --\n", levels, len(finalTracker.quickest), len(bestTracker.quickest))
			}
		}
	}

	return tracker.bestTracker
}

func possibleFrom(tracker *Tracker, visited *map[utils.Position]bool) []utils.Position {
	possible := []utils.Position{}
	curr := tracker.curr
	currValue := tracker.matrix.Value(curr)
	utils.EachVectorXY(func(vector utils.Vector) {
		newPos := curr.Move(vector)
		if _, ok := (*visited)[newPos]; !ok {
			if tracker.matrix.Contains(newPos) {
				newValue := strings.ToLower(string(tracker.matrix.Value(newPos)))[0]
				if newValue == currValue || newValue <= currValue+1 {
					(*visited)[newPos] = true
					possible = append(possible, newPos)
				}
			}
		}
	})
	sort.Slice(possible, func(i, j int) bool {
		a := possible[i]
		b := possible[j]

		scoreA := math.Abs(float64(a.X-tracker.end.X)) + math.Abs(float64(a.Y-tracker.end.Y))
		scoreB := math.Abs(float64(b.X-tracker.end.X)) + math.Abs(float64(b.Y-tracker.end.Y))
		//fmt.Printf("%v:%v --------- %v:%v\n", a, scoreA, b, scoreB)

		return scoreA < scoreB
	})
	return possible
}

func distance(a, b utils.Position) int {
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}
