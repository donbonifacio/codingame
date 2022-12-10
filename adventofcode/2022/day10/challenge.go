package main

import (
	"fmt"
	"strings"

	"github.com/donbonifacio/codingame/blob/master/adventofcode/2022/utils"
)

func main() {
	fmt.Printf("Challenge %v\n", utils.Atoi("10"))
}

type CPU struct {
	X     int
	Cycle int
}

func part1(data string) (int, string) {
	lines := utils.AsLines(data)
	cpu := CPU{X: 1, Cycle: 0}
	sum := 0
	writer := new(strings.Builder)

	for _, line := range lines {
		//fmt.Printf("%v -> %v\n", line, cpu)
		parts := utils.Split(line, " ")
		if parts[0] == "noop" {
			cpu.Cycle += 1
			sum += cycleScore(&cpu)
			writeScreen(writer, &cpu)
		} else if parts[0] == "addx" {
			toAdd := utils.Atoi(parts[1])
			for i := 0; i < 2; i++ {
				cpu.Cycle += 1
				writeScreen(writer, &cpu)
				sum += cycleScore(&cpu)
			}
			cpu.X += toAdd
			//fmt.Printf(" Add %v -> %v,\n", toAdd, cpu)
		} else {
			panic("?")
		}
	}
	return sum, writer.String()
}

var targetCycles = []int{20, 60, 100, 140, 180, 220}

func cycleScore(cpu *CPU) int {
	for _, val := range targetCycles {
		if cpu.Cycle == val {
			//fmt.Printf("----%v\n", cpu)
			return cpu.Cycle * cpu.X
		}
	}
	return 0
}

func writeScreen(writer *strings.Builder, cpu *CPU) {
	if cpu.Cycle == 0 {
		return
	}
	//Cycle 200 X:38
	ref := cpu.Cycle % 40
	if ref == 0 {
		ref = 40
	}
	toWrite := "."
	if ref == cpu.X || ref == cpu.X+2 || ref == cpu.X+1 {
		toWrite = "#"
	}
	writer.WriteString(toWrite)
	if cpu.Cycle%40 == 0 {
		writer.WriteString("\n")
	}
	//fmt.Printf("Cycle %v X:%v\n%v\n", cpu.Cycle, cpu.X, writer.String())
}
