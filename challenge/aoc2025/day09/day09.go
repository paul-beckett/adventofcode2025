package day09

import (
	"aoc2025/internal/math"
	"aoc2025/internal/must"
	"errors"
	"strings"
)

type Day09 struct {
	points []math.Vector2
}

func NewDay09(data []string) *Day09 {
	var points []math.Vector2
	for _, line := range data {
		fields := strings.FieldsFunc(line, func(r rune) bool {
			return r == ','
		})
		point := math.Vector2{
			X: must.ParseInt(fields[0]),
			Y: must.ParseInt(fields[1]),
		}
		points = append(points, point)
	}

	return &Day09{points: points}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (d *Day09) Part1() (int, error) {
	maxArea := 0
	for i, start := range d.points[:len(d.points)-1] {
		for _, end := range d.points[i+1:] {
			dx := abs(start.X-end.X) + 1
			dy := abs(start.Y-end.Y) + 1
			area := dx * dy
			maxArea = max(maxArea, area)
		}
	}
	return maxArea, nil
}

func (d *Day09) Part2() (int, error) {
	return 0, errors.ErrUnsupported
}
