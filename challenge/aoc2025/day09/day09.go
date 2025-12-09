package day09

import (
	"aoc2025/internal/math"
	"aoc2025/internal/must"
	"fmt"
	"slices"
	"strings"
)

type pointPair struct {
	start, end math.Vector2
	area       int
}

func newPointPair(start, end math.Vector2) *pointPair {
	dx := abs(start.X-end.X) + 1
	dy := abs(start.Y-end.Y) + 1

	return &pointPair{
		start: start,
		end:   end,
		area:  dx * dy,
	}
}

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
	var lines []pointPair
	for i, point := range d.points {
		var line pointPair
		if i < len(d.points)-1 {
			line = *newPointPair(point, d.points[i+1])
		} else {
			line = *newPointPair(point, d.points[0])
		}
		lines = append(lines, line)
	}

	var pairs []pointPair
	for i, start := range d.points[:len(d.points)-1] {
		for _, end := range d.points[i+1:] {
			pairs = append(pairs, *newPointPair(start, end))
		}
	}
	slices.SortFunc(pairs, func(a, b pointPair) int {
		return b.area - a.area
	})

	for _, pair := range pairs {
		crosses := false
		for _, line := range lines {
			if intersects(pair, line) {
				crosses = true
				break
			}
		}
		if !crosses {
			return pair.area, nil
		}
	}

	return -1, fmt.Errorf("not found")
}

func intersects(rect pointPair, line pointPair) bool {
	if rect.start == line.start || rect.start == line.end || rect.end == line.start || rect.end == line.end {
		return false
	}
	minX := min(rect.start.X, rect.end.X)
	maxX := max(rect.start.X, rect.end.X)
	minY := min(rect.start.Y, rect.end.Y)
	maxY := max(rect.start.Y, rect.end.Y)

	lineMinX := min(line.start.X, line.end.X)
	lineMaxX := max(line.start.X, line.end.X)
	lineMinY := min(line.start.Y, line.end.Y)
	lineMaxY := max(line.start.Y, line.end.Y)

	if lineMinX >= maxX || lineMaxX <= minX || lineMinY >= maxY || lineMaxY <= minY {
		return false
	}

	return true
}
