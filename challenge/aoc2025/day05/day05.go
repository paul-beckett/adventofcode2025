package day05

import (
	"slices"
	"strconv"
	"strings"
)

type Day05 struct {
	data []string
}

func NewDay05(data []string) *Day05 {
	return &Day05{data: data}
}

type idRange struct {
	min, max int
}

var byMin = func(a, b idRange) int {
	return a.min - b.min
}

func (d *Day05) Part1() (int, error) {

	var idRanges []idRange
	var i int
	for i = 0; i < len(d.data); i++ {
		line := d.data[i]
		if len(line) == 0 {
			break
		}
		fields := strings.Split(line, "-")
		idRanges = append(idRanges, idRange{min: mustParseInt(fields[0]), max: mustParseInt(fields[1])})
	}
	fresh := 0
	for i = i + 1; i < len(d.data); i++ {
		id := mustParseInt(d.data[i])
		for _, r := range idRanges {
			if id >= r.min && id <= r.max {
				fresh++
				break
			}
		}
	}
	return fresh, nil
}

func mustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func (d *Day05) Part2() (int, error) {
	var idRanges []idRange
	for _, line := range d.data {
		if len(line) == 0 {
			break
		}
		fields := strings.Split(line, "-")
		idRanges = append(idRanges, idRange{min: mustParseInt(fields[0]), max: mustParseInt(fields[1])})
	}
	slices.SortFunc(idRanges, byMin)

	total := 0
	for _, m := range merge(idRanges) {
		total += m.max - m.min + 1
	}
	return total, nil
}

func merge(ranges []idRange) []idRange {
	merged := []idRange{ranges[0]}

	for _, r := range ranges {
		slices.SortFunc(merged, byMin)
		last := len(merged) - 1
		if r.min <= merged[last].max {
			merged[last].max = max(r.max, merged[last].max)
		} else {
			merged = append(merged, r)
		}
	}
	return merged
}
