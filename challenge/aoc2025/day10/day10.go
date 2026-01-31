package day10

import (
	"aoc2025/internal/must"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Day10 struct {
	machines []machine
}

type machine struct {
	target  int
	buttons []int
}

func NewDay10(data []string) *Day10 {
	var machines []machine
	for _, line := range data {
		machines = append(machines, parseLine(line))
	}
	return &Day10{machines: machines}
}

func parseLine(line string) machine {
	fields := strings.Fields(line)
	target := fields[0]
	target = target[1 : len(target)-1]
	target = strings.ReplaceAll(target, ".", "0")
	target = strings.ReplaceAll(target, "#", "1")

	var buttons []int
	for _, button := range fields[1 : len(fields)-1] {
		button = button[1 : len(button)-1]
		a := strings.Split(strings.Repeat("0", len(target)), "")
		for _, s := range strings.Split(button, ",") {
			a[must.ParseInt(s)] = "1"
		}

		buttons = append(buttons, parseBinary(strings.Join(a, "")))
	}

	return machine{
		target:  parseBinary(target),
		buttons: buttons,
	}
}

func parseBinary(s string) int {
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func (d *Day10) Part1() (int, error) {
	total := 0
	for i, m := range d.machines {
		n, err := minMoves(m)
		if err != nil {
			return -1, fmt.Errorf("problem with machine [%d]: %w", i, err)
		}
		total += n
	}

	return total, nil
}

func minMoves(m machine) (int, error) {
	visited := make(map[int]bool)
	current := map[int]bool{
		0: true,
	}
	moves := 1
	for len(current) > 0 {
		next := make(map[int]bool)
		for c := range current {
			if visited[c] {
				continue
			}
			visited[c] = true
			for _, button := range m.buttons {
				n := c ^ button
				if n == m.target {
					return moves, nil
				}
				next[n] = true
			}
		}
		current = next
		moves++
	}
	return -1, fmt.Errorf("no moves found")
}

// Part2 solved with:
// https://www.reddit.com/r/adventofcode/comments/1pk87hl/2025_day_10_part_2_bifurcate_your_way_to_victory/
// https://topaz.github.io/paste/#XQAAAQDrBgAAAAAAAAAzHIoib6poHLpewxtGE3pTrRdzrponKxDhfDpmpp1XOH9xnlIyXvIsci+yi/TTFy44FGq6ZrL5OGunysUd322wy+hc3ZIsGd8pNfizbHiJBJTwZuKTJfFD2uUHnzBwP+u/d/PLktBiYiqXhh1rLe8pUTd4hRgQ7Y7ZnPiYgWE25rG2G/K82KYb/v3eDZYBSqI6WDTw/KZ12Dc6FqQLlurOLmsFXKRqb7yL8I8sTp9GTt2rfbMrhrR7UlhjBxofh5Ckk4hXPfRc/R87qV/BXrRJFgFbvPjBlT03fVct8umxOsqUTZ0nT7hYZl0wGUxgeOty+QYL51kUz7Jh0+LwJz28zABJLSt4UoP/08Oei2An6Y1i7Z/d7tmq1TE/qp3ZUSUTcjpJHmwOb9bGYuS9ryexTqHm6rXlEzZyiR8LjrqEDglnSy+YbNGxN2bbfvYPPco3xwCryYbgxUQ+LP53awgvEkk+We0/iyJCHhS7k3s9KLf9SkeB7/aXFoRQoHzrlkzme18oufdVmq+7hJe2xK2Z0Vyj11XfERvYggQXUIuwEbMKJWgp6jVgaoc9yXLHeaz1O+E8ECvRY5GerpLRyK0ywx4j0ItPHMWvkcySQyfoJjD+oHHVSzsrkggtm2szpnyD4QPzu0Cj4IzdkzkjA+9RYKpumjThjZIRU1tpoNVq/llLE+51NJkIUr7SRiLP5okgfiUXiJs6bd3lPfR/pwaO3iqTyVcHL/4/tOLI2Cz5lU0anKFSl9Mm6AEUTRIMfFmtD+JMwOFK226DXSLfjaopjJfI1sg4TmSJ7AjgQQOjCDKSVIvILlpaYliEsTDEg/q9vvOfpLfkVFP2WI0dt2raqrR9h44Q3NaKwKnFNmLV1yIunwwvEecBn5wyaG8Zr5JlF38jcGtGWBHYWVvmf23ByFxPZTgUDC8a1Fa967eX9K3+hB7gR8LgjOyS+DkfbLSCAqqPckeyD9PFfVH9KNAvQtTjZ/kmits4YjOhdIX/6G3UXw==
func (d *Day10) Part2() (int, error) {
	return 0, errors.ErrUnsupported
}
