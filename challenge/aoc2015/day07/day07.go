package day07

import (
	"aoc2025/internal/must"
	"maps"
	"strconv"
	"strings"
)

type Day07 struct {
	wires map[string]string
}

func NewDay07(data []string) *Day07 {
	wires := make(map[string]string)

	for _, line := range data {
		fields := strings.Split(line, " -> ")
		wires[fields[1]] = fields[0]
	}
	return &Day07{wires: wires}
}

func (d *Day07) Part1() (int, error) {
	result := resolve("a", d.wires, make(map[string]uint16))
	return int(result), nil
}

func (d *Day07) Part2() (int, error) {
	wires := maps.Clone(d.wires)
	wires["b"] = strconv.Itoa(int(resolve("a", wires, make(map[string]uint16))))
	result := resolve("a", wires, make(map[string]uint16))
	return int(result), nil
}

func resolve(wire string, wires map[string]string, memo map[string]uint16) uint16 {
	val, resolved := memo[wire]
	if resolved {
		return val
	}
	if v, err := strconv.Atoi(wire); err == nil {
		return uint16(v)
	}
	parts := strings.Fields(wires[wire])
	if len(parts) == 1 { //value
		val = resolve(parts[0], wires, memo)
	} else if len(parts) == 2 {
		val = ^resolve(parts[1], wires, memo)
	} else {
		switch parts[1] {
		case "AND":
			val = resolve(parts[0], wires, memo) & resolve(parts[2], wires, memo)
		case "OR":
			val = resolve(parts[0], wires, memo) | resolve(parts[2], wires, memo)
		case "LSHIFT":
			val = resolve(parts[0], wires, memo) << must.ParseInt(parts[2])
		case "RSHIFT":
			val = resolve(parts[0], wires, memo) >> must.ParseInt(parts[2])
		}
	}
	memo[wire] = val
	return val
}
