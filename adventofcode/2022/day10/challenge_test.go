package main

import (
	"strings"
	"testing"

	"github.com/donbonifacio/codingame/blob/master/adventofcode/2022/utils"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	main()
	assert.Equal(t, 123, 123, "they should be equal")
}

func TestMod(t *testing.T) {
	assert.Equal(t, 0, 20%20)
	assert.Equal(t, 0, 40%20)
	assert.Equal(t, 10, 30%20)
	assert.Equal(t, 5, 5%20)
	assert.Equal(t, 0, 200%20)
}

func TestSamplePart1(t *testing.T) {
	data := utils.ReadInput("sample.txt")
	cpu := part1(data)
	assert.Equal(t, 13140, cpu.TotalSignalStrength)
}

func TestSamplePart2(t *testing.T) {
	data := utils.ReadInput("sample.txt")
	cpu := part1(data)
	expected := []string{
		"##..##..##..##..##..##..##..##..##..##..",
		"###...###...###...###...###...###...###.",
		"####....####....####....####....####....",
		"#####.....#####.....#####.....#####.....",
		"######......######......######......####",
		"#######.......#######.......#######.....",
	}
	assert.Equal(t, strings.Join(expected, "\n"), strings.TrimSpace(cpu.Screen.String()))
}

func TestInputPart1(t *testing.T) {
	data := utils.ReadInput("input.txt")
	cpu := part1(data)
	assert.Equal(t, 14160, cpu.TotalSignalStrength)
}

func TestInputPart2(t *testing.T) {
	data := utils.ReadInput("input.txt")
	cpu := part1(data)
	// RJERPEFC
	expected := []string{
		"###....##.####.###..###..####.####..##..",
		"#..#....#.#....#..#.#..#.#....#....#..#.",
		"#..#....#.###..#..#.#..#.###..###..#....",
		"###.....#.#....###..###..#....#....#....",
		"#.#..#..#.#....#.#..#....#....#....#..#.",
		"#..#..##..####.#..#.#....####.#.....##..",
	}
	assert.Equal(t, strings.Join(expected, "\n"), strings.TrimSpace(cpu.Screen.String()))
}
