package day18

import (
	"aoc2025/internal/math"
	"maps"
)

type Day18 struct {
	tiles map[math.Vector2]bool
	width int
}

func NewDay18(data []string) *Day18 {
	tiles := make(map[math.Vector2]bool)
	for i, c := range data[0] {
		v := math.Vector2{X: i, Y: 0}
		tiles[v] = c == '^'
	}
	return &Day18{tiles: tiles, width: len(data[0])}
}
func (d *Day18) Part1() (int, error) {
	return d.Part1WithRows(40)
}

func (d *Day18) Part1WithRows(n int) (int, error) {
	tiles := maps.Clone(d.tiles)
	width := d.width

	directions := []math.Vector2{{X: -1, Y: -1}, {X: 0, Y: -1}, {X: 1, Y: -1}}

	for y := range n {
		if y == 0 {
			continue
		}
		for x := range width {
			pos := math.Vector2{X: x, Y: y}
			var traps [3]bool
			for i, dir := range directions {
				traps[i] = tiles[pos.Add(dir)]
			}
			trap := (traps[0] && traps[1] && !traps[2]) || (!traps[0] && traps[1] && traps[2]) || (traps[0] && !traps[1] && !traps[2]) || (!traps[0] && !traps[1] && traps[2])
			tiles[pos] = trap
		}
	}
	safe := 0
	for _, t := range tiles {
		if !t {
			safe++
		}
	}
	return safe, nil
}

func (d *Day18) Part2() (int, error) {
	return d.Part1WithRows(400_000)
}
