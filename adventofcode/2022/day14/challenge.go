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
	return allAtRest(data, false)
}

func part2(data string) int {
	return allAtRest(data, true)
}

func allAtRest(data string, hasFloor bool) int {
	lines := utils.AsLines(data)
	var paths [][]utils.Position
	topLeft, bottomRight := utils.Position{X: 1000, Y: 0}, utils.Position{X: 0, Y: 0}
	generator := utils.Position{X: 500, Y: 0}
	for _, line := range lines {
		coords := utils.Split(line, " -> ")
		path := []utils.Position{}
		for _, coord := range coords {
			var pos utils.Position
			fmt.Sscanf(coord, "%d,%d", &pos.X, &pos.Y)
			path = append(path, pos)
			if pos.X <= topLeft.X {
				topLeft.X = pos.X
			}
			if pos.X > bottomRight.X {
				bottomRight.X = pos.X
			}
			if pos.Y > bottomRight.Y {
				bottomRight.Y = pos.Y
			}
		}
		paths = append(paths, path)

	}
	matrix := utils.BuildByteMatrix(bottomRight.Y+10, bottomRight.X*2)

	for _, path := range paths {
		for i := 0; i < len(path)-1; i++ {
			start := path[i]
			end := path[i+1]
			vector := utils.Vector{X: end.X - start.X, Y: end.Y - start.Y, Intensity: 1}
			if vector.X != 0 {
				vector.X = vector.X / int(math.Abs(float64(vector.X)))
			}
			if vector.Y != 0 {
				vector.Y = vector.Y / int(math.Abs(float64(vector.Y)))
			}
			for curr := start; curr.X != end.X || curr.Y != end.Y; curr = curr.Move(vector) {
				matrix.Set(curr, '#')
			}
			matrix.Set(end, '#')
		}
	}
	matrix.Set(generator, '.')

	if hasFloor {
		floor := utils.Position{X: 0, Y: bottomRight.Y + 2}
		for i := 0; i < matrix.SizeX; i++ {
			floor.X = i
			matrix.Set(floor, '#')
		}
	}

	show(matrix, topLeft, bottomRight, generator)

	running := true
	counter := 0
	down := utils.Vector{X: 0, Y: 1, Intensity: 1}
	downLeft := utils.Vector{X: -1, Y: 1, Intensity: 1}
	downRight := utils.Vector{X: 1, Y: 1, Intensity: 1}
	sand := generator
	for running {
		if matrix.Value(sand) == 'o' {
			running = false
			break
		}
		for true {
			newSand := sand.Move(down)

			if !matrix.Contains(newSand) {
				running = false
				break
			}
			if matrix.Value(newSand) != '.' {
				newSand = sand.Move(downLeft)
				if !matrix.Contains(newSand) || matrix.Value(newSand) != '.' {
					newSand = sand.Move(downRight)
					if !matrix.Contains(newSand) || matrix.Value(newSand) != '.' {
						matrix.Set(sand, 'o')
						sand = generator
						counter++
						break
					}
				}
			}
			show(matrix, topLeft, bottomRight, newSand)
			sand = newSand
		}
	}

	return counter
}

func show(matrix *utils.ByteMatrix, topLeft utils.Position, bottomRight utils.Position, sand utils.Position) {
	return
	smallMatrix := (*matrix)
	//fmt.Printf("%v %v:%v -  %v %v:%v", len(smallMatrix.Data), topLeft.Y, bottomRight.Y, len(smallMatrix.Data[0]), topLeft.X, bottomRight.X)
	newData := make([][]byte, len(matrix.Data))
	copy(newData, matrix.Data[topLeft.Y:bottomRight.Y+5])
	smallMatrix.SizeY = len(newData)
	for i := 0; i < len(newData); i++ {
		newData[i] = make([]byte, len(matrix.Data[i]))
		copy(newData[i], matrix.Data[i][topLeft.X-5:bottomRight.X+5])
		smallMatrix.SizeX = len(newData[i])
	}
	smallMatrix.Data = newData
	newSand := utils.Position{X: sand.X - topLeft.X + 5, Y: sand.Y}
	smallMatrix.Set(newSand, 'o')
	fmt.Printf("\n----%v Sand: %v -> %v", topLeft, sand, newSand)
	fmt.Printf("\n%v\n", smallMatrix.ToString())
}
