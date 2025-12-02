package day02

import (
	"errors"
	"strconv"
	"strings"
)

type Day02 struct {
	data []string
}

func NewDay02(data []string) *Day02 {
	return &Day02{data: data}
}

func (d *Day02) Part1() (int, error) {
	sum := 0
	ranges := strings.FieldsFunc(d.data[0], func(r rune) bool { return r == ',' })
	for _, r := range ranges {
		f := strings.Split(r, "-")
		low, _ := strconv.Atoi(f[0])
		high, _ := strconv.Atoi(f[1])
		for i := low; i <= high; i++ {
			s := strconv.Itoa(i)
			l := len(s)
			if l%2 == 0 && s[:l/2] == s[l/2:] {
				sum += i
			}
		}
	}
	return sum, nil
}

func (d *Day02) Part2() (int, error) {
	return 0, errors.ErrUnsupported
}
