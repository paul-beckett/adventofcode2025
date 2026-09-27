package day09

import (
	"aoc2025/internal/must"
	"errors"
	"strings"
)

type Day09 struct {
	data []string
}

func NewDay09(data []string) *Day09 {
	return &Day09{data: data}
}

func decompress(s string) string {
	var d []string
	start := 0
	i := 0
	for i < len(s) {
		if s[i] == '(' {
			d = append(d, s[start:i])
			j := i
			for j < len(s) {
				if s[j] == ')' {
					break
				}
				j++
			}
			nums := strings.Split(s[i+1:j], "x")
			chars := must.ParseInt(nums[0])
			repeat := must.ParseInt(nums[1])
			d = append(d, strings.Repeat(s[j+1:j+1+chars], repeat))
			i = j + 1 + chars
			start = i
		} else {
			i++
		}
	}
	if start < len(s) {
		d = append(d, s[start:])
	}
	return strings.Join(d, "")
}

func (d *Day09) Part1() (int, error) {
	s := decompress(d.data[0])
	return len(s), nil
}

func (d *Day09) Part2() (int, error) {
	return 0, errors.ErrUnsupported
}
