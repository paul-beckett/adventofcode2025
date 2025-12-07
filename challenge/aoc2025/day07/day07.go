package day07

import (
	"strings"
)

type Day07 struct {
	data []string
}

func NewDay07(data []string) *Day07 {
	return &Day07{data: data}
}

const (
	startChar = 'S'
	splitChar = '^'
)

func (d *Day07) Part1() (int, error) {

	start := strings.Index(d.data[0], string(startChar))

	current := map[int]bool{start: true}
	splits := 0
	for _, line := range d.data[1:] {
		next := make(map[int]bool)
		for i := range current {
			if line[i] == splitChar {
				next[i-1] = true
				next[i+1] = true
				splits++
			} else {
				next[i] = true
			}
		}
		current = next
	}

	return splits, nil
}

func (d *Day07) Part2() (int, error) {
	start := strings.Index(d.data[0], string(startChar))

	current := map[int]int{start: 1}
	for _, line := range d.data[1:] {
		next := make(map[int]int)
		for i, n := range current {
			if line[i] == splitChar {
				next[i-1] += n
				next[i+1] += n
			} else {
				next[i] += n
			}
		}
		current = next
	}

	timelines := 0
	for _, n := range current {
		timelines += n
	}

	return timelines, nil
}
