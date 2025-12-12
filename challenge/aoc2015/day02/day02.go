package day02

import (
	"aoc2025/internal/must"
	"slices"
	"strings"
	"unicode"
)

type Day02 struct {
	prisms [][]int
}

func NewDay02(data []string) *Day02 {
	var prisms [][]int
	for _, line := range data {
		var prism []int
		for _, f := range strings.FieldsFunc(line, func(r rune) bool {
			return !unicode.IsNumber(r)
		}) {
			prism = append(prism, must.ParseInt(f))
		}
		prisms = append(prisms, prism)
	}
	return &Day02{prisms: prisms}
}

func (d *Day02) Part1() (int, error) {
	total := 0
	for _, prism := range d.prisms {
		total += 2*prism[0]*prism[1] + 2*prism[1]*prism[2] + 2*prism[2]*prism[0]
		p := slices.Clone(prism)
		slices.Sort(p)
		total += p[0] * p[1]
	}
	return total, nil
}

func (d *Day02) Part2() (int, error) {
	total := 0
	for _, prism := range d.prisms {
		p := slices.Clone(prism)
		slices.Sort(p)
		total += (p[0] + p[1]) * 2
		total += p[0] * p[1] * p[2]
	}
	return total, nil
}
