package main

import (
	"fmt"

	"github.com/donbonifacio/codingame/blob/master/adventofcode/2022/utils"
)

func main() {
	fmt.Printf("Challenge %v\n", utils.Atoi("1"))
}

func part1(data string) int {
	return snake(data, 2)
}

func part2(data string) int {
	return snake(data, 10)
}

func snake(data string, size int) int {
	lines := utils.AsLines(data)
	snake := make([]utils.Position, size)
	for i := 0; i < len(snake); i++ {
		snake[i] = utils.Position{X: 0, Y: 0}
	}
	head := len(snake) - 1

	visited := map[utils.Position]bool{}

	for _, line := range lines {
		vector := parseVector(line)

		for i := 0; i < vector.Intensity; i++ {
			snake[head] = snake[head].Move(vector)
			for k := head - 1; k >= 0; k-- {
				isAdjacent := snake[k].IsAdjacent(snake[k+1])
				if !isAdjacent {
					snake[k] = nextPos(snake[k], snake[k+1])
				}
			}
			visited[snake[0]] = true
		}
	}
	return len(visited)
}

func parseVector(data string) utils.Vector {
	parts := utils.Split(data, " ")

	vector := utils.GetVectorForDir(parts[0])
	vector.Intensity = utils.Atoi(parts[1])

	return vector
}

func nextPos(pos utils.Position, target utils.Position) utils.Position {
	lucky := utils.Position{X: -999999, Y: -99999}
	utils.EachVector(func(vector utils.Vector) {
		try := pos.Move(vector)
		if try.IsSideBySide(target) {
			lucky = try
		}
	})
	if lucky.X == -999999 {
		utils.EachVector(func(vector utils.Vector) {
			try := pos.Move(vector)
			if try.IsAdjacent(target) {
				lucky = try
			}
		})
	}
	return lucky
}
