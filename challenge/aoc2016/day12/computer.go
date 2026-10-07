package day12

import "strconv"

type instruction func(*computer, int) int

func valueOrRegisterValue(c *computer, v string) int {
	if n, err := strconv.Atoi(v); err == nil {
		return n
	}
	return c.registers[v]
}

func cpy(v string, target string) instruction {
	return func(c *computer, i int) int {
		c.registers[target] = valueOrRegisterValue(c, v)
		return i + 1
	}
}

func inc(r string) instruction {
	return func(c *computer, i int) int {
		c.registers[r]++
		return i + 1
	}
}

func dec(r string) instruction {
	return func(c *computer, i int) int {
		c.registers[r]--
		return i + 1
	}
}

func jnz(v string, offset int) instruction {
	return func(c *computer, i int) int {
		if valueOrRegisterValue(c, v) == 0 {
			return i + 1
		}
		return i + offset
	}
}

type computer struct {
	registers    map[string]int
	instructions []instruction
}

func newComputer(instructions []instruction) *computer {
	return &computer{
		registers:    make(map[string]int),
		instructions: instructions,
	}
}
