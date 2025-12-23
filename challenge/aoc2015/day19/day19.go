package day19

import (
	"errors"
	"strings"
)

type Day19 struct {
	replacements map[string][]string
	start        string
}

func NewDay19(data []string) *Day19 {
	replacements := make(map[string][]string)
	for _, line := range data {
		if line == "" {
			break
		}
		fields := strings.Fields(line)
		r := fields[0]
		replacements[r] = append(replacements[r], fields[2])
	}
	return &Day19{
		replacements: replacements,
		start:        data[len(data)-1],
	}
}

func (d *Day19) Part1() (int, error) {
	possible := make(map[string]bool)

	for c, r := range d.replacements {
		i := 0
		for {
			i = nextIndex(d.start, c, i)
			if i == -1 {
				break
			}
			for _, s := range r {
				molecule := d.start[:i] + s + d.start[i+len(c):]
				possible[molecule] = true
			}
			i++
		}
	}

	return len(possible), nil
}

func (d *Day19) Part2() (int, error) {
	return 0, errors.ErrUnsupported
}

func nextIndex(s, substr string, start int) int {
	i := strings.Index(s[start:], substr)
	if i == -1 {
		return -1
	}
	return i + start
}
