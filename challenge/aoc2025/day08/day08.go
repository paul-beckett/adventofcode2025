package day08

import (
	"aoc2025/internal/math"
	"aoc2025/internal/must"
	"fmt"
	"maps"
	"slices"
	"strings"
)

type Day08 struct {
	points []math.Vector3
}

func NewDay08(data []string) *Day08 {
	var points []math.Vector3
	for _, line := range data {
		fields := strings.FieldsFunc(line, func(r rune) bool {
			return r == ','
		})
		point := math.Vector3{
			X: must.ParseInt(fields[0]),
			Y: must.ParseInt(fields[1]),
			Z: must.ParseInt(fields[2]),
		}
		points = append(points, point)
	}

	return &Day08{points: points}
}

func (d *Day08) Part1() (int, error) {
	return d.Part1WithLimit(1000)
}

type pointPair struct {
	start, end math.Vector3
	dst        int
}

func newPointPair(start, end math.Vector3) *pointPair {
	return &pointPair{
		start: start,
		end:   end,
		dst:   start.DstSquared(end),
	}
}

func (d *Day08) Part1WithLimit(limit int) (int, error) {
	neighbours := make(map[math.Vector3][]math.Vector3)

	var pairs []pointPair
	for i, start := range d.points[:len(d.points)-1] {
		for _, end := range d.points[i+1:] {
			pairs = append(pairs, *newPointPair(start, end))
		}
	}
	slices.SortFunc(pairs, func(a, b pointPair) int {
		return a.dst - b.dst
	})

	for _, pair := range pairs[:limit] {
		neighbours[pair.start] = append(neighbours[pair.start], pair.end)
		neighbours[pair.end] = append(neighbours[pair.end], pair.start)
	}

	visited := make(map[math.Vector3]bool)
	var circuits [][]math.Vector3
	for _, point := range d.points {
		queue := []math.Vector3{point}
		var circuit []math.Vector3
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			if visited[current] {
				continue
			}
			visited[current] = true
			circuit = append(circuit, current)
			queue = append(queue, neighbours[current]...)
		}
		if len(circuit) > 0 {
			circuits = append(circuits, circuit)
		}
	}
	slices.SortFunc(circuits, func(a, b []math.Vector3) int {
		return len(b) - len(a)
	})

	total := 1
	for _, circuit := range circuits[:3] {
		total *= len(circuit)
	}
	return total, nil
}

func (d *Day08) Part2() (int, error) {
	var circuits []map[math.Vector3]bool
	for _, point := range d.points {
		circuits = append(circuits, map[math.Vector3]bool{point: true})
	}

	var pairs []pointPair
	for i, start := range d.points[:len(d.points)-1] {
		for _, end := range d.points[i+1:] {
			pairs = append(pairs, *newPointPair(start, end))
		}
	}
	slices.SortFunc(pairs, func(a, b pointPair) int {
		return a.dst - b.dst
	})

	for _, pair := range pairs {
		i := slices.IndexFunc(circuits, func(m map[math.Vector3]bool) bool {
			return m[pair.start]
		})
		j := slices.IndexFunc(circuits, func(m map[math.Vector3]bool) bool {
			return m[pair.end]
		})
		if i == j {
			continue
		}
		maps.Copy(circuits[i], circuits[j])
		circuits = slices.Delete(circuits, j, j+1)
		if len(circuits) == 1 {
			return pair.start.X * pair.end.X, nil
		}
	}

	return 0, fmt.Errorf("no solution found")
}
