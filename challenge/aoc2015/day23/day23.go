package day23

import (
	"aoc2025/internal/must"
	"strings"
)

type Day23 struct {
	computer *computer
}

var rules = map[string]func(...string) instruction{
	"hlf": func(a ...string) instruction {
		return hlf(a[0])
	},
	"tpl": func(a ...string) instruction {
		return tpl(a[0])
	},
	"inc": func(a ...string) instruction {
		return inc(a[0])
	},
	"jmp": func(a ...string) instruction {
		return jmp(must.ParseInt(a[0]))
	},
	"jie": func(a ...string) instruction {
		r, _ := strings.CutSuffix(a[0], ",")
		return jie(r, must.ParseInt(a[1]))
	},
	"jio": func(a ...string) instruction {
		r, _ := strings.CutSuffix(a[0], ",")
		return jio(r, must.ParseInt(a[1]))
	},
}

func NewDay23(data []string) *Day23 {
	var instructions []instruction
	for _, line := range data {
		fields := strings.Fields(line)
		instructions = append(instructions, rules[fields[0]](fields[1:]...))
	}
	return &Day23{computer: newComputer(instructions)}
}

func (d *Day23) Part1() (int, error) {
	return d.Part1WithRegister("b")
}

func (d *Day23) Part1WithRegister(r string) (int, error) {
	c := d.computer
	runProgram(c)
	return c.registers[r], nil
}

func (d *Day23) Part2() (int, error) {
	c := d.computer
	c.registers = map[string]int{
		"a": 1,
	}
	runProgram(c)
	return c.registers["b"], nil
}

func runProgram(c *computer) {
	i := 0
	for i < len(c.instructions) {
		instr := c.instructions[i]
		i = instr(c, i)
	}
}
