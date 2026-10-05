package day10

import (
	"aoc2025/internal/must"
	"fmt"
	"maps"
	"slices"
	"strings"
)

type Day10 struct {
	robots map[string][]int
	rules  map[string][]string
}

func NewDay10(data []string) *Day10 {
	robots := make(map[string][]int)
	rules := make(map[string][]string)

	for _, line := range data {
		fields := strings.Fields(line)
		switch fields[0] {
		case "value":
			id := fields[5]
			robots[id] = append(robots[id], must.ParseInt(fields[1]))
		case "bot":
			id := fields[1]
			robots[id] = robots[id]
			low := fields[6]
			if fields[5] == "output" {
				low = fields[5] + low
			}
			high := fields[11]
			if fields[10] == "output" {
				high = fields[10] + high
			}
			rules[id] = []string{low, high}
		}
	}

	return &Day10{
		robots: robots,
		rules:  rules,
	}
}

func (d *Day10) Part1() (string, error) {
	return d.Part1WithComparing(17, 61)
}
func (d *Day10) Part1WithComparing(low, high int) (string, error) {
	robots := maps.Clone(d.robots)
	resolve(robots, d.rules)
	for r, chips := range robots {
		if slices.Contains(chips, low) && slices.Contains(chips, high) {
			return r, nil
		}
	}
	return "", fmt.Errorf("not found")
}

func resolve(robots map[string][]int, rules map[string][]string) {
	visited := make(map[string]bool)

	var queue []string
	for r, chips := range robots {
		if len(chips) == 2 {
			queue = append(queue, r)
		}
	}

	for len(queue) > 0 {
		r := queue[0]
		queue = queue[1:]
		if visited[r] {
			continue
		}
		visited[r] = true
		slices.Sort(robots[r])
		for i, target := range rules[r] {
			robots[target] = append(robots[target], robots[r][i])
			if len(robots[target]) == 2 {
				queue = append(queue, target)
			}
		}
	}
}

func (d *Day10) Part2() (int, error) {
	robots := maps.Clone(d.robots)
	for _, o := range []string{"output0", "output1", "output2"} {
		fmt.Println(robots[o])
	}

	resolve(robots, d.rules)
	for _, o := range []string{"output0", "output1", "output2"} {
		fmt.Println(robots[o])
	}

	return robots["output0"][0] * robots["output1"][0] * robots["output2"][0], nil
}
