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
	idRanges := d.parseRanges()
	fresh := 0
	for _, line := range d.data[len(idRanges)+1:] {
		id := mustParseInt(line)
		for _, r := range idRanges {
			if id >= r.min && id <= r.max {
				fresh++
				break
			}
		}
	}
	return fresh, nil
}

func (d *Day05) Part2() (int, error) {
	idRanges := d.parseRanges()
	slices.SortFunc(idRanges, byMin)

	total := 0
	for _, m := range mergeRanges(idRanges) {
		total += m.max - m.min + 1
	}
	return total, nil
}

func mustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func mergeRanges(ranges []idRange) []idRange {
	merged := []idRange{ranges[0]}

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

func (d *Day05) parseRanges() []idRange {
	var idRanges []idRange
	for _, line := range d.data {
		if len(line) == 0 {
			break
		}
		fields := strings.Split(line, "-")
		idRanges = append(idRanges, idRange{min: mustParseInt(fields[0]), max: mustParseInt(fields[1])})
	}
	return idRanges
}
