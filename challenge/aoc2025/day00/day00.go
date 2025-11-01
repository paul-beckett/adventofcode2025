package day00

import (
	"strconv"
)

type Day00 struct {
	data []string
}

func NewDay00(data []string) *Day00 {
	return &Day00{data: data}
}

func (d *Day00) Part1() (int, error) {
	return strconv.Atoi(d.data[0])
}

func (d *Day00) Part2() (int, error) {
	return strconv.Atoi(d.data[1])
}
