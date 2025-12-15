package day12

import (
	"encoding/json"
)

type Day12 struct {
	data []string
}

func NewDay12(data []string) *Day12 {
	return &Day12{data: data}
}

func (d *Day12) Part1() (int, error) {
	var doc any
	err := json.Unmarshal([]byte(d.data[0]), &doc)
	if err != nil {
		return -1, err
	}
	return sum(doc, func(v any) bool { return false }), nil
}

func (d *Day12) Part2() (int, error) {
	var doc any
	err := json.Unmarshal([]byte(d.data[0]), &doc)
	if err != nil {
		return -1, err
	}
	return sum(doc, func(v any) bool { return v == "red" }), nil
}

func sum(doc any, f func(v any) bool) int {
	total := 0
	switch d := doc.(type) {
	case float64:
		total += int(d)
	case []any:
		for _, v := range d {
			total += sum(v, f)
		}
	case map[string]any:
		for _, v := range d {
			if f(v) {
				return 0
			}
			total += sum(v, f)
		}
	}
	return total
}
