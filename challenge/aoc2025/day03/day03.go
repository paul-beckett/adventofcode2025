package day03

import (
	"strconv"
	"strings"
)

type Day03 struct {
	data []string
}

func NewDay03(data []string) *Day03 {
	return &Day03{data: data}
}

func (d *Day03) Part1() (int, error) {
	sum := 0
	for _, line := range d.data {
		nums := toDigits(line)
		sum += joltage(nums, 2)
	}
	return sum, nil
}

func (d *Day03) Part2() (int, error) {
	sum := 0
	for _, line := range d.data {
		nums := toDigits(line)
		sum += joltage(nums, 12)
	}
	return sum, nil
}

func joltage(nums []int, size int) int {
	j := 0
	start := 0
	for size > 0 {
		digit := -1
		digitI := -1
		for i := start; i <= len(nums)-size; i++ {
			if nums[i] > digit {
				digit = nums[i]
				digitI = i
			}
		}
		start = digitI + 1
		j *= 10
		j += digit
		size--
	}
	return j
}

func toDigits(s string) []int {
	var nums []int
	for _, c := range strings.Split(s, "") {
		num, _ := strconv.Atoi(c)
		nums = append(nums, num)
	}
	return nums
}
