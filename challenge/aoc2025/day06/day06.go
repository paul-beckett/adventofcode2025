package day06

import (
	"aoc2025/internal/must"
	"strings"
)

type Day06 struct {
	data []string
}

func NewDay06(data []string) *Day06 {
	return &Day06{data: data}
}

var operations = map[string]func(a, b int) int{
	"+": func(a, b int) int {
		return a + b
	},
	"*": func(a, b int) int {
		return a * b
	},
}

func (d *Day06) Part1() (int, error) {
	operationRow := len(d.data) - 1
	ops := strings.Fields(d.data[operationRow])
	var nums [][]int
	for _, line := range d.data[:operationRow] {
		var row []int
		for _, n := range strings.Fields(line) {
			row = append(row, must.ParseInt(n))
		}
		nums = append(nums, row)
	}

	total := 0
	for x, opCode := range ops {
		op := operations[opCode]
		n := nums[0][x]
		for y := range len(nums) - 1 {
			n = op(n, nums[y+1][x])
		}
		total += n
	}

	return total, nil
}

func (d *Day06) Part2() (int, error) {
	operationRow := len(d.data) - 1
	total := 0
	sum := 0
	var op func(a, b int) int
	for x := range d.data[0] {
		opCode := d.data[operationRow][x]
		if opCode != ' ' {
			op = operations[string(opCode)]
			total += sum
			if opCode == '+' {
				sum = 0
			} else {
				sum = 1
			}
		}
		num := ""
		for y := range operationRow {
			c := d.data[y][x]
			if c != ' ' {
				num += string(c)
			}
		}
		if num == "" {
			continue
		}
		sum = op(sum, must.ParseInt(num))
	}
	total += sum
	return total, nil
}
