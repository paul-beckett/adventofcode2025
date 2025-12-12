package day12

import (
	"aoc2025/internal/must"
	"errors"
	"strings"
)

type Day12 struct {
	data []string
}

func NewDay12(data []string) *Day12 {
	return &Day12{data: data}
}

func (d *Day12) Part1() (int, error) {
	total := 0
	for _, line := range d.data {
		if !strings.Contains(line, "x") {
			continue
		}
		parts := strings.Split(line, ": ")
		size := strings.Split(parts[0], "x")
		area := must.ParseInt(size[0]) / 3 * must.ParseInt(size[1]) / 3
		count := 0
		for _, n := range strings.Fields(parts[1]) {
			count += must.ParseInt(n)
		}
		if area >= count {
			total++
		}
	}
	return total, nil
}

func (d *Day12) Part2() (int, error) {
	return 0, errors.ErrUnsupported
}
