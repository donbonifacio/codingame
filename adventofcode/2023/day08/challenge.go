package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/donbonifacio/codingame/blob/master/adventofcode/2023/utils"
	"github.com/samber/lo"
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

type Node struct {
	id    string
	left  string
	right string
}

type Data struct {
	directions []string
	nodes      map[string]Node
}

func parseDirections(data *Data, line string) *Data {
	data.directions = strings.Split(line, "")
	return data
}

func parseNodes(data *Data, lines []string) *Data {
	lo.ForEach(lines, func(line string, _ int) {
		if len(line) > 0 {
			var node Node
			parts := strings.Split(line, "=")
			node.id = strings.TrimSpace(parts[0])
			dirs := strings.Split(parts[1], ",")
			node.left = dirs[0][2:]
			node.right = dirs[1][1:4]
			data.nodes[node.id] = node
		}
	})
	return data
}

func part1(raw string) int {
	lines := strings.Split(raw, "\n")
	data := &Data{
		directions: []string{},
		nodes:      map[string]Node{},
	}
	data = parseDirections(data, lines[0])
	data = parseNodes(data, lines[2:])

	curr := data.nodes["AAA"]
	turns := 0
	for i := 0; curr.id != "ZZZ"; i++ {
		dir := data.directions[i%len(data.directions)]
		fmt.Printf("%v -> ", curr.id)
		if dir == "L" {
			curr = data.nodes[curr.left]
		} else {
			curr = data.nodes[curr.right]
		}
		fmt.Printf("%v\n", curr.id)
		turns += 1
		if curr.id == "ZZZ" {
			break
		}
	}
	fmt.Println(data)
	fmt.Println(turns)
	return turns
}

func part2(data string) int {
	return 0
}
