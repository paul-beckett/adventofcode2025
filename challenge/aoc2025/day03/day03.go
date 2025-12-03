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
		sum += findJoltage(nums, 2)
	}
	return sum, nil
}

func (d *Day03) Part2() (int, error) {
	sum := 0
	for _, line := range d.data {
		nums := toDigits(line)
		sum += findJoltage(nums, 12)
	}
	return sum, nil
}

func findJoltage(nums []int, size int) int {
	joltage := 0

	l := 0
	r := len(nums) - size + 1
	for r <= len(nums) {
		digit := -1
		digitI := -1
		for i := l; i < r; i++ {
			if nums[i] > digit {
				digit = nums[i]
				digitI = i
			}
		}
		l = digitI + 1
		joltage *= 10
		joltage += digit
		r++
	}
	return joltage
}

func toDigits(s string) []int {
	var nums []int
	for _, c := range strings.Split(s, "") {
		num, _ := strconv.Atoi(c)
		nums = append(nums, num)
	}
	return nums
}
