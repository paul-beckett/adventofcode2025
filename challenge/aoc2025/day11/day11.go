package day11

import (
	"errors"
	"strings"
	"unicode"
)

type Day11 struct {
	graph map[string][]string
}

func NewDay11(data []string) *Day11 {
	graph := make(map[string][]string)
	for _, line := range data {
		fields := strings.FieldsFunc(line, func(r rune) bool {
			return !unicode.IsLetter(r)
		})
		graph[fields[0]] = fields[1:]
	}

	return &Day11{graph: graph}
}

func (d *Day11) Part1() (int, error) {
	return d.countPaths("you", "out", make(map[string]bool)), nil
}

func (d *Day11) countPaths(src string, target string, visited map[string]bool) int {

	if src == target {
		return 1
	}

	visited[src] = true
	paths := 0
	for _, next := range d.graph[src] {
		if !visited[next] {
			paths += d.countPaths(next, target, visited)
		}
	}
	visited[src] = false
	return paths
}

func (d *Day11) Part2() (int, error) {
	return 0, errors.ErrUnsupported
}
