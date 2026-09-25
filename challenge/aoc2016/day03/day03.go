package day03

import (
	"aoc2025/internal/must"
	"sort"
	"strings"
)

type Day03 struct {
	data []string
}

func NewDay03(data []string) *Day03 {
	return &Day03{data: data}
}

func (d *Day03) Part1() (int, error) {
	count := 0
	for _, l := range d.data {
		var sides []int
		for _, f := range strings.Fields(l) {
			sides = append(sides, must.ParseInt(f))
		}
		sort.Ints(sides)
		if sides[0]+sides[1] > sides[2] {
			count++
		}
	}
	return count, nil
}

func (d *Day03) Part2() (int, error) {
	count := 0
	var lines [][]int
	for _, l := range d.data {
		var line []int
		for _, f := range strings.Fields(l) {
			line = append(line, must.ParseInt(f))
		}
		lines = append(lines, line)
	}

	for x := 0; x < len(lines[0]); x++ {
		for y := 0; y < len(lines); y = y + 3 {
			sides := []int{lines[y][x], lines[y+1][x], lines[y+2][x]}
			sort.Ints(sides)
			if sides[0]+sides[1] > sides[2] {
				count++
			}
		}
	}

	return count, nil
}
