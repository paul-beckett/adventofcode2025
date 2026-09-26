package day06

import "math"

type Day06 struct {
	data []string
}

func NewDay06(data []string) *Day06 {
	return &Day06{data: data}
}

func (d *Day06) Part1() (string, error) {
	var msg []rune
	for _, count := range d.toLetterCounts() {
		var maxC rune
		var maxN int

		for c, n := range count {
			if n > maxN {
				maxN = n
				maxC = c
			}

		}
		msg = append(msg, maxC)
	}
	return string(msg), nil
}

func (d *Day06) toLetterCounts() []map[rune]int {
	var letterCounts []map[rune]int
	for _, line := range d.data {
		for i, c := range line {
			if i >= len(letterCounts) {
				letterCounts = append(letterCounts, make(map[rune]int))
			}
			letterCounts[i][c]++
		}
	}
	return letterCounts
}

func (d *Day06) Part2() (string, error) {
	var msg []rune
	for _, count := range d.toLetterCounts() {
		var minC rune
		minN := math.MaxInt

		for c, n := range count {
			if n < minN {
				minN = n
				minC = c
			}

		}
		msg = append(msg, minC)
	}
	return string(msg), nil
}
