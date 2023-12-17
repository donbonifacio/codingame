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
	end   int
}

type Data struct {
	directions []string
	nodes      map[string]Node
	start      []string
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

func parseStart(data *Data) *Data {
	for _, node := range data.nodes {
		if strings.HasSuffix(node.id, "A") {
			data.start = append(data.start, node.id)
		}
	}
	return data
}

func part2(raw string) int {
	lines := strings.Split(raw, "\n")
	data := &Data{
		directions: []string{},
		nodes:      map[string]Node{},
		start:      []string{},
	}
	data = parseDirections(data, lines[0])
	data = parseNodes(data, lines[2:])
	data = parseStart(data)

	curr := lo.Map(data.start, func(id string, _ int) Node {
		node := data.nodes[id]
		return node
	})
	//curr = curr[0:1]
	turns := 0
	for i := 0; true; i++ {
		dir := data.directions[i%len(data.directions)]
		curr = lo.Map(curr, func(node Node, _ int) Node {
			var newNode Node
			if dir == "L" {
				newNode = data.nodes[node.left]
			} else {
				newNode = data.nodes[node.right]
			}
			if node.end != 0 {
				newNode.end = node.end
			}
			if newNode.end == 0 && strings.HasSuffix(newNode.id, "Z") {
				newNode.end = turns
				fmt.Println(newNode)
			}
			return newNode
		})
		turns += 1
		ends := lo.Reduce(curr, func(agg []int, item Node, _ int) []int {
			if item.end != 0 {
				return append(agg, item.end)
			}
			return agg
		}, []int{})

		//fmt.Println(ends)
		if len(ends) == len(curr) {
			fmt.Println(ends)
			fmt.Println(utils.LCM(ends[0], ends[1:]))
			break
		}
		if i > 1000000 {
			break
		}
	}

	return turns
}
