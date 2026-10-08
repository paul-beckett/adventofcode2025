package day18

import "slices"

type Day18 struct {
	tiles []bool
}

func NewDay18(data []string) *Day18 {
	tiles := make([]bool, len(data[0]))
	for i, c := range data[0] {
		tiles[i] = c == '^'
	}
	return &Day18{tiles: tiles}
}
func (d *Day18) Part1() (int, error) {
	return d.Part1WithRows(40)
}

func (d *Day18) Part1WithRows(n int) (int, error) {
	return resolve(d.tiles, n), nil
}

func resolve(start []bool, n int) int {
	safe := 0
	prev := start
	for _, b := range prev {
		if !b {
			safe++
		}
	}
	for range n - 1 {
		next := slices.Clone(prev)
		for x := range len(prev) {
			left := x > 0 && prev[x-1]
			centre := prev[x]
			right := x < len(prev)-1 && prev[x+1]

			trap := (left && centre && !right) || (!left && centre && right) || (left && !centre && !right) || (!left && !centre && right)
			if !trap {
				safe++
			}
			next[x] = trap
		}
		prev = next
	}
	return safe
}

func (d *Day18) Part2() (int, error) {
	return resolve(d.tiles, 400_000), nil
}
