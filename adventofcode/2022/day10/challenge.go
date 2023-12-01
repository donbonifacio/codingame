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
	X                   int
	Cycle               int
	Screen              *strings.Builder
	TotalSignalStrength int
}

func NewCpu() *CPU {
	return &CPU{X: 1, Cycle: 0, TotalSignalStrength: 0, Screen: new(strings.Builder)}
}

func (cpu *CPU) writeScreen() {
	if cpu.Cycle == 0 {
		return
	}
	ref := cpu.Cycle % 40
	if ref == 0 {
		ref = 40
	}
	toWrite := "."
	if ref == cpu.X || ref == cpu.X+2 || ref == cpu.X+1 {
		toWrite = "#"
	}
	cpu.Screen.WriteString(toWrite)
	if cpu.Cycle%40 == 0 {
		cpu.Screen.WriteString("\n")
	}
}

func cycleScore(cpu *CPU) int {
	var targetCycles = map[int]bool{20: true, 60: true, 100: true, 140: true, 180: true, 220: true}

	if _, ok := targetCycles[cpu.Cycle]; ok {
		return cpu.Cycle * cpu.X
	}
	return 0
}

func (cpu *CPU) ProcessCycle() {
	cpu.Cycle += 1
	cpu.TotalSignalStrength += cycleScore(cpu)
	cpu.writeScreen()
}

func part1(data string) *CPU {
	lines := utils.AsLines(data)
	cpu := NewCpu()

	for _, line := range lines {
		parts := utils.Split(line, " ")
		if op, ok := ops[parts[0]]; ok {
			op(cpu, parts[1:])
		} else {
			panic(fmt.Sprintf("Don't know how to process '%v'", line))
		}
	}
	return cpu
}

var ops = map[string]func(*CPU, []string){
	"noop": func(cpu *CPU, args []string) {
		cpu.ProcessCycle()
	},

	"addx": func(cpu *CPU, args []string) {
		for i := 0; i < 2; i++ {
			cpu.ProcessCycle()
		}
		cpu.X += utils.Atoi(args[0])
	},
}
