package day01

import (
	"aoc2025/internal/math"
	"aoc2025/internal/must"
	"fmt"
	math2 "math"
	"strings"
	"unicode"
)

type Day01 struct {
	data []string
}

var directions = []math.Vector2{
	{0, -1}, {1, 0}, {0, 1}, {-1, 0},
}

func NewDay01(data []string) *Day01 {
	return &Day01{data: data}
}

func (d *Day01) Part1() (int, error) {
	pos := math.Vector2{X: 0, Y: 0}
	instructions := strings.FieldsFunc(d.data[0], func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
	dir := 0
	for _, instr := range instructions {
		turn := instr[0]
		switch turn {
		case 'R':
			dir++
		case 'L':
			dir--
		default:
			return -1, fmt.Errorf("unknown dir")
		}
		dir = (len(directions) + dir) % len(directions)
		scale := must.ParseInt(instr[1:])

		move := directions[dir]
		move.X *= scale
		move.Y *= scale

		pos = pos.Add(move)
	}
	return int(math2.Abs(float64(pos.X)) + math2.Abs(float64(pos.Y))), nil
}

func (d *Day01) Part2() (int, error) {
	pos := math.Vector2{X: 0, Y: 0}
	instructions := strings.FieldsFunc(d.data[0], func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
	dir := 0
	visited := make(map[math.Vector2]bool)
	for _, instr := range instructions {
		turn := instr[0]
		switch turn {
		case 'R':
			dir++
		case 'L':
			dir--
		default:
			return -1, fmt.Errorf("unknown dir")
		}
		dir = (len(directions) + dir) % len(directions)
		scale := must.ParseInt(instr[1:])
		for range scale {
			if visited[pos] {
				break
			} else {
				visited[pos] = true
			}
			pos = pos.Add(directions[dir])
		}
	}
	return int(math2.Abs(float64(pos.X)) + math2.Abs(float64(pos.Y))), nil
}
