package day01

import (
	"aoc2025/internal/collection"
	"math"
	"strconv"
)

type Day01 struct {
	data []string
}

func NewDay01(data []string) *Day01 {
	return &Day01{data: data}
}

func mustParseNum(s string) int {
	n, err := strconv.Atoi(s[1:])
	if err != nil {
		panic(err)
	}
	if s[0] == 'L' {
		n *= -1
	}
	return n
}

func (d *Day01) Part1() (int, error) {

	pos := 50
	f := func(s string) int {
		pos = mustParseNum(s) + pos
		return pos
	}
	nums := collection.Map(d.data, f)

	isZero := func(n int) bool {
		return n%100 == 0
	}
	nums = collection.Filter(nums, isZero)

	return len(nums), nil

}

func (d *Day01) Part2() (int, error) {
	pos := 50
	zeroCount := 0
	for _, s := range d.data {
		n := mustParseNum(s)
		zeroCount += int(math.Abs(float64(n) / 100))
		move := n % 100
		newPos := pos + move
		if newPos <= 0 || newPos >= 100 {
			if pos != 0 {
				zeroCount++
			}
		}
		pos = (100 + newPos) % 100
	}

	return zeroCount, nil
}
