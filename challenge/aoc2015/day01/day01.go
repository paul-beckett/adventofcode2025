package day01

import (
	"fmt"
	"strings"
)

type Day01 struct {
	data []string
}

func NewDay01(data []string) *Day01 {
	return &Day01{data: data}
}

func (d *Day01) Part1() (int, error) {
	line := d.data[0]
	up := strings.Count(line, "(")
	down := strings.Count(line, ")")
	return up - down, nil
}

func (d *Day01) Part2() (int, error) {
	floor := 0
	for i, c := range d.data[0] {
		if c == ')' {
			floor--
		} else if c == '(' {
			floor++
		} else {
			return -1, fmt.Errorf("invalid character '%c'", c)
		}
		if floor < 0 {
			return i + 1, nil
		}
	}
	return -1, fmt.Errorf("-1 floor not found")
}
