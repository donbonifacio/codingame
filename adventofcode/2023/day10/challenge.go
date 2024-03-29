package main

import (
	"fmt"
	"os"

	"github.com/donbonifacio/codingame/blob/master/adventofcode/2023/utils"
)

func main() {
	fmt.Printf("Challenge %v\n", utils.Atoi("1"))
}

func readInput(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}

type Loop struct {
	matrix      *utils.ByteMatrix
	distances   *utils.IntMatrix
	maxDistance int
	start       utils.Position
}

func part1(data string) int {
	matrix := utils.AsByteMatrix(data)
	loop := Loop{matrix: matrix, distances: utils.BuildIntMatrix(matrix.SizeX, matrix.SizeY)}
	loop = processStart(loop)
	loop = processDistances(loop, loop.start, 0)
	p(loop)
	return loop.maxDistance
}

func dirValue(loop Loop, pos utils.Position) byte {
	if loop.matrix.Contains(pos) {
		return loop.matrix.Value(pos)
	}
	return '?'
}

func processStart(loop Loop) Loop {
	for x := 0; x < loop.matrix.SizeX; x++ {
		for y := 0; y < loop.matrix.SizeY; y++ {
			pos := utils.Position{X: x, Y: y}
			value := loop.matrix.Value(pos)
			if value == 'S' {
				loop.start = pos
				n := dirValue(loop, utils.Position{X: pos.X, Y: pos.Y - 1})
				s := dirValue(loop, utils.Position{X: pos.X, Y: pos.Y + 1})
				e := dirValue(loop, utils.Position{X: pos.X + 1, Y: pos.Y})
				w := dirValue(loop, utils.Position{X: pos.X - 1, Y: pos.Y})
				curr := byte('0')
				if n == '|' && s == '|' {
					curr = '|'
				}
				if w == '-' && e == '-' {
					curr = '-'
				}
				if s == '|' && e == '-' {
					curr = 'F'
				}
				if s == '|' && e == 'J' {
					curr = 'F'
				}
				if s == 'J' && e == '7' {
					curr = 'F'
				}
				if curr == '0' {
					panic(fmt.Sprintf("n:%v s:%v e:%v :w:%v", string(n), string(s), string(e), string(w)))
				}
				loop.matrix.Set(loop.start, curr)
			}
		}
	}
	return loop
}

/*
| is a vertical pipe connecting north and south.
- is a horizontal pipe connecting east and west.
L is a 90-degree bend connecting north and east.
J is a 90-degree bend connecting north and west.
7 is a 90-degree bend connecting south and west.
F is a 90-degree bend connecting south and east.
*/

type node struct {
	pos    utils.Position
	weight int
}

func processDistances(loop Loop, curr utils.Position, step int) Loop {
	connections := map[byte][]utils.Vector{
		'|': {{X: 0, Y: 1}, {X: 0, Y: -1}},
		'-': {{X: -1, Y: 0}, {X: 1, Y: 0}},
		'L': {{X: 0, Y: -1}, {X: 1, Y: 0}},
		'J': {{X: 0, Y: -1}, {X: -1, Y: 0}},
		'7': {{X: 0, Y: 1}, {X: -1, Y: 0}},
		'F': {{X: 0, Y: 1}, {X: 1, Y: 0}},
	}

	queue := []node{{pos: curr, weight: 0}}
	i := 0

	for true {
		curr := queue[0]
		queue = queue[1:]
		currValue := loop.matrix.Value(curr.pos)
		directions := connections[currValue]
		step++
		for _, dir := range directions {
			newPos := curr.pos.Move(dir)
			newPosDistance := loop.distances.Value(newPos)
			if newPos != loop.start && newPosDistance == 0 {
				//fmt.Printf("%v: Processing curr:%v newPos:%v distance:%v step:%v queue:%v\n", i, curr, newPos, newPosDistance, step, queue)
				weight := curr.weight + 1
				loop.distances.Set(newPos, weight)
				queue = append(queue, node{pos: newPos, weight: weight})
				if loop.maxDistance < weight {
					loop.maxDistance = weight
				}
			}
		}
		i++
		if len(queue) == 0 {
			break
		}
		if i == 2 {
			//break
		}
	}

	return loop
}

func p(loop Loop) {
	fmt.Printf("--%v maxDistance:%v\n", loop.start, loop.maxDistance)
	fmt.Println(loop.matrix.ToString())
	fmt.Println(loop.distances.ToString())
}

func part2(data []string) int {
	return 0
}
