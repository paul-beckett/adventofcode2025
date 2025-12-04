package day04

import (
	"aoc2025/internal/math"
)

type Day04 struct {
	data []string
}

func NewDay04(data []string) *Day04 {
	return &Day04{data: data}
}

type vectorSet = map[math.Vector2]bool

func (d *Day04) parseInput() map[math.Vector2]vectorSet {
	positions := make(map[math.Vector2]vectorSet)
	for y, line := range d.data {
		for x, c := range line {
			if c == '@' {
				v := math.Vector2{X: x, Y: y}
				positions[v] = make(vectorSet)
			}
		}
	}
	for v, n := range positions {
		for _, dir := range directions {
			neighbour := v.Add(dir)
			if _, exists := positions[neighbour]; exists {
				n[neighbour] = true
			}
		}
	}
	return positions
}

var directions = []math.Vector2{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

func (d *Day04) Part1() (int, error) {
	positions := d.parseInput()
	rolls := 0
	for _, n := range positions {
		if len(n) < 4 {
			rolls++
		}
	}
	return rolls, nil
}

func (d *Day04) Part2() (int, error) {
	positions := d.parseInput()
	total := 0
	for {
		removed := false
		for pos, neighbours := range positions {
			if len(neighbours) < 4 {
				removed = true
				total++
				for n := range neighbours {
					delete(positions[n], pos)
				}
				delete(positions, pos)
			}
		}
		if !removed {
			break
		}
	}
	return total, nil
}
