package day01

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

type Day01 struct {
	lhs, rhs []int
}

func NewDay01(data []string) *Day01 {
	var lhs, rhs []int
	for _, line := range data {
		fields := strings.Fields(line)
		l, _ := strconv.Atoi(fields[0])
		lhs = append(lhs, l)
		r, _ := strconv.Atoi(fields[1])
		rhs = append(rhs, r)
	}
	return &Day01{lhs: lhs, rhs: rhs}
}

func (d *Day01) Part1() (int, error) {
	lhs := slices.Clone(d.lhs)
	rhs := slices.Clone(d.rhs)
	diff := 0
	slices.Sort(lhs)
	slices.Sort(rhs)
	for i := 0; i < len(lhs); i++ {
		diff += int(math.Abs(float64(lhs[i] - rhs[i])))
	}

	return diff, nil
}

func (d *Day01) Part2() (int, error) {
	rhsCounts := make(map[int]int)
	for _, i := range d.rhs {
		rhsCounts[i]++
	}
	total := 0
	for _, i := range d.lhs {
		total += i * rhsCounts[i]
	}
	return total, nil
}
