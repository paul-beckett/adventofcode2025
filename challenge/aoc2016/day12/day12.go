package day12

import (
	"aoc2025/internal/must"
	"strings"
)

type Day12 struct {
	computer *computer
}

var rules = map[string]func(...string) instruction{
	"cpy": func(a ...string) instruction {
		return cpy(a[0], a[1])
	},
	"inc": func(a ...string) instruction {
		return inc(a[0])
	},
	"dec": func(a ...string) instruction {
		return dec(a[0])
	},
	"jnz": func(a ...string) instruction {
		return jnz(a[0], must.ParseInt(a[1]))
	},
}

func NewDay12(data []string) *Day12 {
	var instructions []instruction
	for _, line := range data {
		fields := strings.Fields(line)
		instructions = append(instructions, rules[fields[0]](fields[1:]...))
	}
	return &Day12{computer: newComputer(instructions)}
}

func (d *Day12) Part1() (int, error) {
	c := d.computer
	runProgram(c)
	return c.registers["a"], nil
}

func (d *Day12) Part2() (int, error) {
	c := d.computer
	c.registers = map[string]int{"c": 1}
	runProgram(c)
	return c.registers["a"], nil
}

func runProgram(c *computer) {
	i := 0
	for i < len(c.instructions) {
		instr := c.instructions[i]
		i = instr(c, i)
	}
}
