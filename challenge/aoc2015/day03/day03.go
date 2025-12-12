package day03

import (
	"aoc2025/internal/math"
)

var deltas = map[rune]math.Vector2{
	'>': {X: 1, Y: 0},
	'v': {X: 0, Y: 1},
	'<': {X: -1, Y: 0},
	'^': {X: 0, Y: -1},
}

type Day03 struct {
	data []string
}

func NewDay03(data []string) *Day03 {
	return &Day03{data: data}
}

func (d *Day03) Part1() (int, error) {
	var position math.Vector2
	visited := map[math.Vector2]bool{
		position: true,
	}
	for _, move := range d.data[0] {
		position = position.Add(deltas[move])
		visited[position] = true
	}
	return len(visited), nil
}

func (d *Day03) Part2() (int, error) {
	var santa math.Vector2
	var robot math.Vector2
	visited := map[math.Vector2]bool{
		santa: true,
		robot: true,
	}
	moves := d.data[0]
	for i := 0; i < len(moves); i += 2 {
		santa = santa.Add(deltas[rune(moves[i])])
		visited[santa] = true
		if i < len(moves)-1 {
			robot = robot.Add(deltas[rune(moves[i+1])])
			visited[robot] = true
		}
	}

	return len(visited), nil
}
