package day20

import (
	"aoc2025/internal/must"
	"fmt"
	"math"
	"slices"
	"strings"
)

type Day20 struct {
	ipRanges []ipRange
}

type ipRange struct {
	min, max int
}

var byMin = func(a, b ipRange) int {
	return a.min - b.min
}

func mergeRanges(ranges []ipRange) []ipRange {
	merged := []ipRange{ranges[0]}

	for _, r := range ranges {
		last := len(merged) - 1
		if r.min <= merged[last].max {
			merged[last].max = max(r.max, merged[last].max)
		} else {
			merged = append(merged, r)
		}
	}
	return merged
}

func NewDay20(data []string) *Day20 {
	var ipRanges []ipRange
	for _, line := range data {
		fields := strings.Split(line, "-")
		ipRanges = append(ipRanges, ipRange{min: must.ParseInt(fields[0]), max: must.ParseInt(fields[1])})

	}
	return &Day20{ipRanges: ipRanges}
}

func (d *Day20) Part1() (int, error) {
	ipRanges := slices.Clone(d.ipRanges)
	slices.SortFunc(ipRanges, byMin)
	merged := mergeRanges(ipRanges)

	current := 0
	for _, r := range merged {
		if current < r.min {
			return current, nil
		}
		current = r.max + 1
	}

	return 0, fmt.Errorf("not found")
}

func (d *Day20) Part2() (int, error) {
	return d.Part2WithMax(math.MaxUint32)
}

func (d *Day20) Part2WithMax(n int) (int, error) {
	ipRanges := slices.Clone(d.ipRanges)
	slices.SortFunc(ipRanges, byMin)
	merged := mergeRanges(ipRanges)

	totalValid := 0
	current := 0
	for _, r := range merged {
		if current < r.min {
			totalValid += r.min - current
		}
		current = r.max + 1
	}
	if current <= n {
		totalValid += n - current + 1
	}
	return totalValid, nil
}
