package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const challengeId = "01"

func main() {
	fmt.Printf("Challenge %v\n", challengeId)
}

type Operation struct {
	quantity int
	from     int
	to       int
}

type Crane struct {
	stack []string
}

type Cargo struct {
	cranes     map[int]Crane
	operations []Operation
}

func load(data string) *Cargo {
	cargo := Cargo{
		cranes:     map[int]Crane{},
		operations: []Operation{},
	}

	lines := strings.Split(data, "\n")
	separator := 0
	for index, val := range lines {
		if strings.TrimSpace(val) == "" {
			separator = index
			break
		}
	}

	loadCranes(&cargo, lines[0:separator])
	loadOperations(&cargo, lines[separator+1:len(lines)])

	return &cargo
}

func loadCranes(cargo *Cargo, raw []string) {
	craneNumbers := strings.Split(strings.TrimSpace(raw[len(raw)-1]), " ")
	nCranes := atoi(craneNumbers[len(craneNumbers)-1])
	//fmt.Printf("%v\n", nCranes)
	for i := 1; i < nCranes+1; i++ {
		cargo.cranes[i] = Crane{[]string{}}
	}

	for i := len(raw) - 2; i >= 0; i-- {
		line := strings.Split(raw[i], "")

		for c, inc := 1, 1; c < nCranes+1; c, inc = c+1, inc+4 {
			if crane, ok := cargo.cranes[c]; ok {
				if strings.TrimSpace(line[inc]) != "" {
					crane.stack = append(crane.stack, line[inc])
					cargo.cranes[c] = crane
				}
			}
		}
	}
	//fmt.Printf("Cargo: %v\n", cargo)
}

func loadOperations(cargo *Cargo, rawOperations []string) {
	re := regexp.MustCompile(`(\d+)`)

	for _, rawOperation := range rawOperations {
		if rawOperation == "" {
			continue
		}

		//fmt.Printf("rawOperation: %v\n", rawOperation)
		matches := re.FindAllStringSubmatch(rawOperation, -1)
		//fmt.Printf("matches: %v\n", matches)

		operation := Operation{atoi(matches[0][0]), atoi(matches[1][0]), atoi(matches[2][0])}
		cargo.operations = append(cargo.operations, operation)

		//fmt.Printf("Operations:: %v\n", cargo)
	}
}

func run(data string, runOperations func(*Cargo)) string {
	cargo := load(data)
	top := []string{}

	runOperations(cargo)
	//fmt.Printf("Cargo:: %v\n", cargo)

	for i := 1; i < len(cargo.cranes)+1; i++ {
		crane := cargo.cranes[i]
		//fmt.Printf("-Crane:: %v\n", crane)
		top = append(top, crane.stack[len(crane.stack)-1])
	}
	return strings.Join(top, "")
}

func runOperationsPart1(cargo *Cargo) {
	for _, operation := range cargo.operations {
		from := cargo.cranes[operation.from]
		to := cargo.cranes[operation.to]
		for i := 0; i < operation.quantity; i++ {
			element := from.stack[len(from.stack)-1]
			from.stack = from.stack[:len(from.stack)-1]
			to.stack = append(to.stack, element)
		}
		cargo.cranes[operation.from] = from
		cargo.cranes[operation.to] = to
	}
}

func runOperationsPart2(cargo *Cargo) {
	for _, operation := range cargo.operations {
		from := cargo.cranes[operation.from]
		fromSize := len(from.stack)
		to := cargo.cranes[operation.to]

		//fmt.Printf("%v - %v\n", from.stack, operation.quantity)
		elements := from.stack[fromSize-operation.quantity : fromSize]
		from.stack = from.stack[:fromSize-operation.quantity]
		to.stack = append(to.stack, elements...)
		//fmt.Printf("%v -> %v\n", from.stack, to.stack)

		cargo.cranes[operation.from] = from
		cargo.cranes[operation.to] = to
	}
}

func atoi(raw string) int {
	val, err := strconv.Atoi(raw)
	if err != nil {
		panic(err)
	}
	return val
}

func readInput(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(data)
}
