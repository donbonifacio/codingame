package main

import (
	"fmt"
	"math"

	"github.com/donbonifacio/codingame/blob/master/adventofcode/2022/utils"
)

func main() {
	fmt.Printf("Challenge %v\n", utils.Atoi("1"))
}

type Sensor struct {
	pos    utils.Position
	beacon utils.Position
	area   int
}

func part1(data string, target int) int {
	lines := utils.AsLines(data)
	left := utils.Position{X: 0, Y: target}
	right := utils.Position{X: 0, Y: target}
	sensors := make([]Sensor, len(lines))
	for i, line := range lines {
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensors[i].pos.X, &sensors[i].pos.Y, &sensors[i].beacon.X, &sensors[i].beacon.Y)
		sensors[i].area = distance(sensors[i].pos, sensors[i].beacon)
		left.X, _ = utils.MinMax([]int{sensors[i].pos.X - sensors[i].area, left.X, sensors[i].beacon.X})
		_, right.X = utils.MinMax([]int{sensors[i].pos.X + sensors[i].area, left.X, sensors[i].beacon.X})
	}

	counter := 0
	fmt.Printf("Left %v Right %v\n", left, right)

	for x := left.X; x <= right.X; x++ {
		curr := utils.Position{X: x, Y: target}
		empty := false
		for _, sensor := range sensors {
			//fmt.Printf("Area %v", sensor)
			if curr == sensor.beacon {
				continue
			}
			area := distance(sensor.pos, curr)
			if curr.X == 2 {
				//fmt.Printf("%v curr:%v area:%v %v\n", sensor, curr, area, area <= sensor.area)
			}
			empty = empty || (area <= sensor.area)
		}
		if empty {
			counter++
		}
		//fmt.Printf("%v -> %v\n", curr, empty)
	}
	return counter
}

func distance(a, b utils.Position) int {
	return int(math.Abs(float64(a.X-b.X)) + math.Abs(float64(a.Y-b.Y)))
}
