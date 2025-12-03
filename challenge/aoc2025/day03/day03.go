package day03

import (
	"errors"
	"fmt"
	"slices"
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
		nums := toInts(line)
		sum += joltage(nums)
	}
	return sum, nil
}

func (d *Day03) Part2() (int, error) {
	return 0, errors.ErrUnsupported
}

func joltage(nums []int) int {
	maxL := slices.Clone(nums)
	for i := 1; i < len(nums); i++ {
		maxL[i] = max(maxL[i], max(maxL[i-1]))
	}

	maxR := slices.Clone(nums)
	for i := len(nums) - 2; i >= 0; i-- {
		maxR[i] = max(maxR[i], max(maxR[i+1]))
	}
	fmt.Println("maxL:", maxL)
	fmt.Println("maxR:", maxR)

	best := 0
	for i := 0; i < len(nums)-1; i++ {
		best = max(best, maxL[i]*10+maxR[i+1])
	}
	fmt.Println("best:", best)
	return best
}

func toInts(s string) []int {
	var nums []int
	for _, c := range strings.Split(s, "") {
		num, _ := strconv.Atoi(c)
		nums = append(nums, num)
	}
	return nums
}
