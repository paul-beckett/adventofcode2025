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

func (d *Day10) Part2() (int, error) {
	return 0, errors.ErrUnsupported
}
