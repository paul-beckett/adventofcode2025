package day12

import (
	"aoc2025/internal/must"
	"errors"
	"strings"
	"unicode"
)

type Day12 struct {
	data []string
}

func NewDay12(data []string) *Day12 {
	return &Day12{data: data}
}

func (d *Day12) Part1() (int, error) {
	total := 0
	for _, field := range strings.FieldsFunc(d.data[0], func(r rune) bool { return !unicode.IsNumber(r) && r != '-' }) {
		total += must.ParseInt(field)
	}
	return total, nil
}

func (d *Day12) Part2() (int, error) {
	return 0, errors.ErrUnsupported
}
