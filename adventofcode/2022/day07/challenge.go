package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const challengeId = "01"

func main() {
	fmt.Printf("Challenge %v\n", challengeId)
}

type Item struct {
	name     string
	dir      bool
	size     int
	children []Item
	parent   *Item
}

type FileSystem struct {
	root    Item
	current *Item
}

func part1(data string) *FileSystem {
	fs := FileSystem{}
	lines := strings.Split(data, "\n")

	cmd := []string{lines[0]}

	for _, line := range lines[1:len(lines)] {
		if strings.HasPrefix(line, "$") {
			processCmd(&fs, cmd)
			cmd = []string{line}
		} else {
			cmd = append(cmd, line)
		}
	}
	processCmd(&fs, cmd)

	return &fs
}

func processCmd(fs *FileSystem, cmdAndOutput []string) {
	raw := strings.Split(cmdAndOutput[0], " ")
	cmd := raw[1]
	arg := ""
	output := cmdAndOutput[1:len(cmdAndOutput)]

	if len(raw) == 3 {
		arg = raw[2]
	}
	//fmt.Printf("\nCmd: %v %v\n", cmd, arg)

	if cmd == "cd" {
		if arg == "/" {
			fs.root = Item{name: "/", dir: true, size: -1, children: []Item{}, parent: nil}
			fs.current = &fs.root
		} else if arg == ".." {
			fs.current = fs.current.parent
		} else {
			for i := 0; i < len(fs.current.children); i++ {
				item := &fs.current.children[i]
				if item.dir && item.name == arg {
					fs.current = item
				}
			}
		}
	}
	if cmd == "ls" {
		addItems(fs, output)
	}
	//fmt.Printf("fs[%v]: %v items, size: %v\n", fs.current.name, len(fs.current.children), fs.current.size)
}

func addItems(fs *FileSystem, lines []string) {
	sum := 0
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		var item Item
		parts := strings.Split(line, " ")
		if parts[0] == "dir" {
			item = Item{name: parts[1], dir: true, size: -1, children: []Item{}, parent: fs.current}
		} else {
			item = Item{name: parts[1], dir: false, size: atoi(parts[0]), children: []Item{}, parent: fs.current}
			sum += item.size
		}
		fs.current.children = append(fs.current.children, item)
	}
	fs.current.size = sum

	//fmt.Printf("\n-------Calculating size for %v\n", fs.current.name)
	for item := fs.current.parent; item != nil; item = item.parent {
		//fmt.Printf("\n-------item  %v size %v\n", item.name, item.size)
		sum = 0
		for _, child := range item.children {
			sum += child.size
		}
		item.size = sum
	}
	//fmt.Printf("\n-------end Calculating size for %v : %v\n", fs.current.name, fs.current.size)
}

func sumLargeDirs(root *Item) int {
	sum := 0
	for i := 0; i < len(root.children); i++ {
		item := &root.children[i]
		//fmt.Printf("size: %v - %v :%v\n", item.size, item.name, item.dir)
		if item.dir && item.size <= 100000 {
			sum += item.size
		}
		if item.dir {
			sum += sumLargeDirs(item)
		}
	}
	return sum
}

func smallestToDelete(fs *FileSystem) *Item {
	target := 30000000 - (70000000 - fs.root.size)
	//fmt.Printf("Target: %v\n", target)

	dirs := []*Item{}
	dirs = append(dirs, loadDirs(&fs.root)...)
	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].size < dirs[j].size
	})

	for _, elem := range dirs {
		if elem.size > target {
			return elem
		}
	}

	panic("?")
}

func loadDirs(item *Item) []*Item {
	dirs := []*Item{}
	for i := 0; i < len(item.children); i++ {
		child := &item.children[i]
		if child.dir {
			dirs = append(dirs, child)
			dirs = append(dirs, loadDirs(child)...)
		}

	}
	return dirs
}

func readInput(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func atoi(raw string) int {
	val, err := strconv.Atoi(raw)
	if err != nil {
		panic(err)
	}
	return val
}
