package day06

import (
	"aoc2025/internal/must"
	"strings"
	"unicode"
)

type Day06 struct {
	data []string
}

func NewDay06(data []string) *Day06 {
	return &Day06{data: data}
}

func (d *Day06) Part1() (int, error) {
	instructions := map[string]func(int) int{
		"turn on": func(b int) int {
			return 1
		},
		"turn off": func(b int) int {
			return 0
		},
		"toggle": func(b int) int {
			if b == 0 {
				return 1
			} else {
				return 0
			}
		},
	}
	return d.solve(instructions), nil
}

func (d *Day06) Part2() (int, error) {
	instructions := map[string]func(int) int{
		"turn on": func(b int) int {
			return b + 1
		},
		"turn off": func(b int) int {
			if b > 0 {
				return b - 1
			} else {
				return 0
			}
		},
		"toggle": func(b int) int {
			return b + 2
		},
	}
	return d.solve(instructions), nil
}

func (d *Day06) solve(instructions map[string]func(int) int) int {
	var lights [1000][1000]int
	for _, line := range d.data {
		var instruction func(int) int
		for name, instr := range instructions {
			if strings.HasPrefix(line, name) {
				instruction = instr
				break
			}
		}
		fields := strings.FieldsFunc(line, func(r rune) bool {
			return !unicode.IsNumber(r)
		})
		startX := must.ParseInt(fields[0])
		endX := must.ParseInt(fields[2])
		startY := must.ParseInt(fields[1])
		endY := must.ParseInt(fields[3])

		for x := startX; x <= endX; x++ {
			for y := startY; y <= endY; y++ {
				lights[y][x] = instruction(lights[y][x])
			}
		}
	}
	count := 0
	for _, row := range lights {
		for _, light := range row {
			count += light

		}
	}
	return count
}
