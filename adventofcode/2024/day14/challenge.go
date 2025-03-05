package main

import (
	"fmt"
	"os"

	"github.com/donbonifacio/codingame/blob/master/adventofcode/2023/utils"
)

func main() {
	fmt.Printf("Challenge %v\n", utils.Atoi("1"))
}

func readInput(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return utils.AsLines(string(data))
}

type Robot struct {
	p utils.Position
	s utils.Vector
}

type Matrix struct {
	m                      map[utils.Position]int
	minX, maxX, minY, maxY int
}

func part1(lines []string) int {
	result := []Robot{}
	for _, line := range lines {
		var p1, p2, v1, v2 int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &p1, &p2, &v1, &v2)
		robot := Robot{
			p: utils.Position{X: p1, Y: p2},
			s: utils.Vector{X: v1, Y: v2},
		}
		result = append(result, robot)
	}

	matrix := buildMatrix(result)
	print(matrix)

	return 0
}

func move(matrix *Matrix, robots []Robot) []Robot {
	for i := range robots {
		robots[i].p.X += robots[i].s.X

		if robots[i].p.X < matrix.minX {
			diff := matrix.minX - robots[i].p.X
			robots[i].p.X = matrix.maxX - diff
		}
		robots[i].p.Y += robots[i].s.Y
	}
	return robots
}

func buildMatrix(robots []Robot) *Matrix {
	matrix := map[utils.Position]int{}
	for _, robot := range robots {
		matrix[robot.p] = matrix[robot.p] + 1
	}
	minX, maxX := 0, 0
	minY, maxY := 0, 0

	for pos := range matrix {
		if pos.X < minX {
			minX = pos.X
		}
		if pos.X > maxX {
			maxX = pos.X
		}
		if pos.Y < minY {
			minY = pos.Y
		}
		if pos.Y > maxY {
			maxY = pos.Y
		}
	}

	return &Matrix{
		m:    matrix,
		minX: minX,
		maxX: maxX,
		minY: minY,
		maxY: maxY,
	}
}

func print(matrix *Matrix) {
	for y := matrix.minY; y <= matrix.maxY; y++ {
		for x := matrix.minX; x <= matrix.maxX; x++ {
			pos := utils.Position{X: x, Y: y}
			if matrix.m[pos] > 0 {
				fmt.Print(matrix.m[pos])
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func part2(data []string) int {
	return 0
}
