package day18

import (
	"aoc2025/internal/math"
)

type Day18 struct {
	grid grid
}

type grid struct {
	g    map[math.Vector2]bool
	size int
}

func (g *grid) next() {
	next := make(map[math.Vector2]bool)
	for p := range g.g {
		next[p] = true
		for _, dir := range directions {
			l := p.Add(dir)
			if l.X >= 0 && l.X < g.size && l.Y >= 0 && l.Y < g.size {
				next[l] = g.g[l]
			}
		}
	}

	for p, state := range next {
		count := 0
		for _, dir := range directions {
			if g.g[p.Add(dir)] {
				count++
			}
		}
		newState := (state && (count == 2 || count == 3)) || (!state && count == 3)
		if !newState {
			delete(next, p)
		} else {
			next[p] = true
		}
	}
	g.g = next
}

var directions = []math.Vector2{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

func NewDay18(data []string) *Day18 {
	g := make(map[math.Vector2]bool)
	for y, row := range data {
		for x, c := range row {
			if c == '#' {
				g[math.Vector2{X: x, Y: y}] = true
			}
		}
	}
	return &Day18{grid: grid{g: g, size: len(data)}}
}

func (d *Day18) Part1() (int, error) {
	return d.Part1WithSteps(100)
}

func (d *Day18) Part1WithSteps(steps int) (int, error) {
	current := d.grid
	for range steps {
		current.next()
	}
	return len(current.g), nil
}

func (d *Day18) Part2WithSteps(steps int) (int, error) {
	current := d.grid
	alwaysOn := []math.Vector2{
		{0, 0}, {X: current.size - 1, Y: 0},
		{0, current.size - 1}, {current.size - 1, current.size - 1},
	}
	for _, l := range alwaysOn {
		current.g[l] = true
	}
	for range steps {
		current.next()
		for _, l := range alwaysOn {
			current.g[l] = true
		}
	}
	return len(current.g), nil
}

func (d *Day18) Part2() (int, error) {
	return d.Part2WithSteps(100)
}
