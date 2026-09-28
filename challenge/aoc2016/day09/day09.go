package day09

import (
	"aoc2025/internal/must"
	"strings"
)

type Day09 struct {
	data []string
}

func NewDay09(data []string) *Day09 {
	return &Day09{data: data}
}

func decompress(s string, recurse bool) int {
	count := 0
	start := 0
	i := 0
	for i < len(s) {
		if s[i] == '(' {
			count += i - start
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
			if recurse {
				count += decompress(s[j+1:j+1+chars], true) * repeat
			} else {
				count += repeat * chars
			}
			i = j + 1 + chars
			start = i
		} else {
			i++
		}
	}
	if start < len(s) {
		count += len(s) - start
	}
	return count
}

func (d *Day09) Part1() (int, error) {
	s := decompress(d.data[0], false)
	return s, nil
}

func (d *Day09) Part2() (int, error) {
	s := decompress(d.data[0], true)
	return s, nil
}
