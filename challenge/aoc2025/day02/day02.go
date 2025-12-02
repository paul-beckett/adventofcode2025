package day02

import (
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
	return d.sumInvalid(func(s string) bool {
		return len(s)%2 == 0 && isRepeating(s, len(s)/2)
	}), nil
}

func (d *Day02) Part2() (int, error) {
	return d.sumInvalid(isInvalid), nil
}

type Range struct {
	low  int
	high int
}

func parseRanges(s string) []Range {
	var ranges []Range
	fields := strings.FieldsFunc(s, func(r rune) bool { return r == ',' })
	for _, r := range fields {
		f := strings.Split(r, "-")
		low, _ := strconv.Atoi(f[0])
		high, _ := strconv.Atoi(f[1])
		ranges = append(ranges, Range{
			low:  low,
			high: high,
		})
	}
	return ranges
}

func isInvalid(s string) bool {
	l := len(s)
	for n := 1; n <= l/2; n++ {
		if isRepeating(s, n) {
			return true
		}
	}
	return false
}

func isRepeating(s string, l int) bool {
	if len(s)%l != 0 {
		return false
	}
	for i := 0; i < len(s)-l; i += l {
		if s[i:i+l] != s[i+l:i+l+l] {
			return false
		}
	}
	return true
}

func (d *Day02) sumInvalid(f func(string) bool) int {
	sum := 0
	ranges := parseRanges(d.data[0])
	for _, r := range ranges {
		for i := r.low; i <= r.high; i++ {
			if f(strconv.Itoa(i)) {
				sum += i
			}
		}
	}
	return sum
}
