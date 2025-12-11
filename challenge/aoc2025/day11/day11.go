package day11

import (
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
	return d.countPaths("you", "out"), nil
}

func (d *Day11) dfs(src string, target string, memo map[string]int) int {
	if src == target {
		return 1
	}

	if val, ok := memo[src]; ok {
		return val
	}

	paths := 0
	for _, next := range d.graph[src] {
		paths += d.dfs(next, target, memo)
	}
	memo[src] = paths
	return paths
}

func (d *Day11) countPaths(src string, target string) int {
	return d.dfs(src, target, make(map[string]int))
}

func (d *Day11) countPathsVia(paths []string) int {
	total := 1
	for i, path := range paths[:len(paths)-1] {
		total *= d.countPaths(path, paths[i+1])
	}
	return total
}

func (d *Day11) Part2() (int, error) {
	return d.countPathsVia([]string{"svr", "dac", "fft", "out"}) + d.countPathsVia([]string{"svr", "fft", "dac", "out"}), nil
}
