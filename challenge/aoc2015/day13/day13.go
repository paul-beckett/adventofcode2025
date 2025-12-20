package day13

import (
	"aoc2025/internal/math"
	"aoc2025/internal/must"
	"maps"
	math2 "math"
	"slices"
	"strings"
)

type Day13 struct {
	data []string
}

func NewDay13(data []string) *Day13 {
	return &Day13{data: data}
}

func (d *Day13) parsePotentialHappiness() map[string]map[string]int {
	names := make(map[string]map[string]int)

	for _, line := range d.data {
		fields := strings.Fields(line)

		p1 := fields[0]
		p2, _ := strings.CutSuffix(fields[10], ".")
		v := must.ParseInt(fields[3])
		if fields[2] == "lose" {
			v = -v
		}

		name, exists := names[p1]
		if !exists {
			name = make(map[string]int)
			names[p1] = name
		}
		name[p2] = v
	}
	return names
}

func (d *Day13) Part1() (int, error) {
	return d.solve(nil), nil
}

func (d *Day13) Part2() (int, error) {
	return d.solve([]string{"yourself"}), nil
}

func (d *Day13) solve(extra []string) int {
	potentialHappiness := d.parsePotentialHappiness()
	best := math2.MinInt
	people := slices.Collect(maps.Keys(potentialHappiness))
	people = append(people, extra...)
	for p := range math.AllPermutations(people) {
		change := 0
		for i, name := range p {
			var name2 string
			if i == 0 {
				name2 = p[len(p)-1]
			} else {
				name2 = p[i-1]
			}
			if slices.Contains(extra, name) || slices.Contains(extra, name2) {
				continue
			}

			change += potentialHappiness[name][name2]
			change += potentialHappiness[name2][name]
		}
		best = max(best, change)
	}
	return best
}
