package day17

import (
	"aoc2025/internal/math"
	"crypto/md5"
	"fmt"
	"io"
)

type Day17 struct {
	passcode string
}

func NewDay17(data []string) *Day17 {
	return &Day17{passcode: data[0]}
}

func (d *Day17) Part1() (string, error) {
	paths := findPaths(d.passcode)
	if len(paths) == 0 {
		return "", fmt.Errorf("not found")
	}
	return paths[0], nil
}

type key struct {
	pos  math.Vector2
	path string
}

var (
	directions = []math.Vector2{{X: 0, Y: -1}, {X: 0, Y: 1}, {X: -1, Y: 0}, {X: 1, Y: 0}}
	path       = []string{"U", "D", "L", "R"}
)

func findPaths(passcode string) []string {
	start := math.Vector2{X: 0, Y: 0}
	target := math.Vector2{X: 3, Y: 3}

	var queue []key
	queue = append(queue, key{pos: start, path: ""})

	var paths []string

	h := md5.New()
	for len(queue) > 0 {
		var nextLevel []key
		for _, p := range queue {
			if p.pos == target {
				paths = append(paths, p.path)
				continue
			}
			h.Reset()
			io.WriteString(h, passcode+p.path)
			s := fmt.Sprintf("%x", h.Sum(nil))
			for i, c := range s[:4] {
				if c < 'b' || c > 'f' {
					continue
				}
				next := p.pos.Add(directions[i])
				if next.X < 0 || next.X > 3 || next.Y < 0 || next.Y > 3 {
					continue
				}
				nextLevel = append(nextLevel, key{pos: next, path: p.path + path[i]})
			}
		}
		queue = nextLevel
	}
	return paths
}

func (d *Day17) Part2() (int, error) {
	paths := findPaths(d.passcode)
	if len(paths) == 0 {
		return 0, fmt.Errorf("not found")
	}
	return len(paths[len(paths)-1]), nil
}
