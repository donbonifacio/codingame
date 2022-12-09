package main

import (
	"fmt"
	"math"

	"github.com/donbonifacio/codingame/blob/master/adventofcode/2022/utils"
)

func main() {
	fmt.Printf("Challenge %v\n", utils.Atoi("1"))
}

func part1(data string) int {
	lines := utils.AsLines(data)
	head := utils.Position{X: 0, Y: 0}
	tail := utils.Position{X: 0, Y: 0}
	lastHeadPosition := head
	visited := map[utils.Position]bool{}

	for _, line := range lines {
		vector := parseVector(line)

		for i := 0; i < vector.Intensity; i++ {
			head = head.Move(vector)
			isAdjacent := adjacent(head, tail)
			if !isAdjacent {
				tail = lastHeadPosition
			}
			lastHeadPosition = head
			visited[tail] = true
		}
	}
	return len(visited)
}

func adjacent(a utils.Position, b utils.Position) bool {
	return math.Abs(float64(a.X-b.X)) <= 1 && math.Abs(float64(a.Y-b.Y)) <= 1
}

func sideBySide(a utils.Position, b utils.Position) bool {
	return (a.X == b.X && math.Abs(float64(a.Y-b.Y)) <= 1 ||
		a.Y == b.Y && math.Abs(float64(a.X-b.X)) <= 1) && !(a.X == b.X && a.Y == b.Y)
}

func parseVector(data string) utils.Vector {
	parts := utils.Split(data, " ")
	intensity := utils.Atoi(parts[1])

	var vector utils.Vector
	if parts[0] == "R" {
		vector = utils.BuildVector(1, 0)
	} else if parts[0] == "L" {
		vector = utils.BuildVector(-1, 0)
	} else if parts[0] == "U" {
		vector = utils.BuildVector(0, -1)
	} else if parts[0] == "D" {
		vector = utils.BuildVector(0, 1)
	} else {
		panic("?")
	}
	vector.Intensity = intensity
	return vector
}

func part2(data string) int {
	lines := utils.AsLines(data)
	snake := make([]utils.Position, 10)
	for i := 0; i < len(snake); i++ {
		snake[i] = utils.Position{X: 0, Y: 0}
	}

	visited := map[utils.Position]bool{}

	for _, line := range lines {
		vector := parseVector(line)

		for i := 0; i < vector.Intensity; i++ {

			//fmt.Printf("Start : %v\n", snake)
			snake[9] = snake[9].Move(vector)
			//fmt.Printf("--curr: %v %v\n", 9, snake[9])
			for k := 8; k >= 0; k-- {
				isAdjacent := adjacent(snake[k], snake[k+1])
				//fmt.Printf("--curr: %v %v\n", k, snake[k])
				if !isAdjacent {
					next := nextPos(snake[k], snake[k+1])
					//fmt.Printf("-----next: %v\n", next)
					snake[k] = next
				}
			}
			visited[snake[0]] = true
		}
	}
	return len(visited)
}

func nextPos(pos utils.Position, target utils.Position) utils.Position {
	lucky := utils.Position{X: -999999, Y: -99999}
	utils.EachVector(func(vector utils.Vector) {
		try := pos.Move(vector)
		if sideBySide(try, target) {
			lucky = try
		}
	})
	if lucky.X == -999999 {
		utils.EachVector(func(vector utils.Vector) {
			try := pos.Move(vector)
			if adjacent(try, target) {
				lucky = try
			}
		})
	}
	return lucky
}
