package day08

type Day08 struct {
	data []string
}

func NewDay08(data []string) *Day08 {
	return &Day08{data: data}
}

func (d *Day08) Part1() (int, error) {
	diff := 0
	for _, line := range d.data {
		count := 0
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case '"':
				continue
			case '\\':
				count++
				switch line[i+1] {
				case '"', '\\':
					i++
				default:
					i += 3
				}
			default:
				count++
			}
		}
		diff += len(line) - count
	}
	return diff, nil
}

func (d *Day08) Part2() (int, error) {
	diff := 0
	for _, line := range d.data {
		encoded := ``
		for _, c := range line {
			switch c {
			case '"':
				encoded += `\"`
			case '\\':
				encoded += `\\`
			default:
				encoded += string(c)
			}

		}
		encoded = `"` + encoded + `"`
		diff += len(encoded) - len(line)
	}
	return diff, nil
}
