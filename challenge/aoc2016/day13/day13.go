package day13

import (
	"aoc2025/internal/math"
	"aoc2025/internal/must"
)

type Day13 struct {
	favourite int
}

var directions = []math.Vector2{
	{0, -1}, {1, 0}, {0, 1}, {-1, 0},
}

func NewDay13(data []string) *Day13 {
	return &Day13{favourite: must.ParseInt(data[0])}
}

func (d *Day13) Part1() (int, error) {
	return d.Part1WithTarget(math.Vector2{X: 31, Y: 39})
}

func (d *Day13) Part1WithTarget(loc math.Vector2) (int, error) {
	dist, _ := bfs(
		func(pos math.Vector2, _ int) bool { return pos == loc },
		func(v math.Vector2) bool { return isWall(v, d.favourite) },
	)
	return dist, nil
}

func countOnes(n int) int {
	count := 0
	for n > 0 {
		count = count + 1
		n = n & (n - 1)
	}
	return count
}

func isWall(loc math.Vector2, n int) bool {
	x := loc.X
	y := loc.Y
	total := x*x + 3*x + 2*x*y + y + y*y
	total += n

	return countOnes(total)%2 != 0
}

func bfs(targetF func(pos math.Vector2, steps int) bool, wallF func(math.Vector2) bool) (int, map[math.Vector2]bool) {
	visited := make(map[math.Vector2]bool)
	var queue []math.Vector2

	start := math.Vector2{X: 1, Y: 1}
	queue = append(queue, start)
	visited[start] = true
	distance := 0

	for len(queue) > 0 {
		var nextLevel []math.Vector2
		for _, p := range queue {
			if targetF(p, distance) {
				return distance, visited
			}
			for _, d := range directions {
				next := p.Add(d)
				if next.X < 0 || next.Y < 0 || visited[next] || wallF(next) {
					continue
				}
				visited[next] = true
				nextLevel = append(nextLevel, next)
			}
		}
		distance++
		queue = nextLevel
	}
	return -1, nil
}

func (d *Day13) Part2() (int, error) {
	_, visited := bfs(
		func(_ math.Vector2, steps int) bool { return steps >= 50 },
		func(v math.Vector2) bool { return isWall(v, d.favourite) },
	)
	return len(visited), nil
}
