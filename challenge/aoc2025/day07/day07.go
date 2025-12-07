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
	splits, _ := d.tachyons()
	return splits, nil
}

func (d *Day07) Part2() (int, error) {
	_, timelines := d.tachyons()
	return timelines, nil
}

// tachyons returns the number of splits and the total timelines
func (d *Day07) tachyons() (int, int) {
	start := strings.Index(d.data[0], string(startChar))
	current := map[int]int{start: 1}
	splits := 0
	for _, line := range d.data[1:] {
		next := make(map[int]int)
		for i, n := range current {
			if line[i] == splitChar {
				next[i-1] += n
				next[i+1] += n
				splits++
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
	return splits, timelines
}
