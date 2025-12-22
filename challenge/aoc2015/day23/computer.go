package day23

type instruction func(*computer, int) int

func hlf(r string) instruction {
	return func(c *computer, i int) int {
		c.registers[r] /= 2
		return i + 1
	}
}

func tpl(r string) instruction {
	return func(c *computer, i int) int {
		c.registers[r] *= 3
		return i + 1
	}
}

func inc(r string) instruction {
	return func(c *computer, i int) int {
		c.registers[r]++
		return i + 1
	}
}

func jmp(offset int) instruction {
	return func(_ *computer, i int) int {
		return i + offset
	}
}

func jie(r string, offset int) instruction {
	return func(c *computer, i int) int {
		if c.registers[r]%2 == 0 {
			return i + offset
		} else {
			return i + 1
		}
	}
}

func jio(r string, offset int) instruction {
	return func(c *computer, i int) int {
		if c.registers[r] == 1 {
			return i + offset
		} else {
			return i + 1
		}
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
