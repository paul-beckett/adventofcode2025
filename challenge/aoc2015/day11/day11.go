package day11

type Day11 struct {
	data []string
}

func NewDay11(data []string) *Day11 {
	return &Day11{data: data}
}

func (d *Day11) Part1() (string, error) {
	return nextValidPassword(d.data[0]), nil
}

func (d *Day11) Part2() (string, error) {
	s := nextValidPassword(d.data[0])
	return nextValidPassword(s), nil
}

func nextValidPassword(s string) string {
	var password []int
	for _, c := range s {
		password = append(password, int(c-'a'))
	}
	for {
		next(password)
		if isValid(password) {
			break
		}
	}

	var runes []rune
	for _, c := range password {
		runes = append(runes, rune(c+'a'))
	}
	return string(runes)
}

func next(n []int) {
	for i := len(n) - 1; i >= 0; i-- {
		if n[i] < 25 {
			n[i]++
			return
		}
		n[i] = 0
	}
}

func isValid(n []int) bool {
	firstPair := -1
	pairsFound := false
	straightFound := false
	for i, v := range n {
		if v == 'i'-'a' || v == 'o'-'a' || v == 'l'-'a' {
			return false
		}
		if i > 0 && v == n[i-1] {
			if firstPair == -1 {
				firstPair = v
			} else if v != firstPair {
				pairsFound = true
			}
		}
		if i > 1 && !straightFound {
			straightFound = v-n[i-1] == 1 && v-n[i-2] == 2
		}
	}
	return pairsFound && straightFound
}
