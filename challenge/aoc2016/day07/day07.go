package day07

import (
	"strings"
)

type Day07 struct {
	data []string
}

func NewDay07(data []string) *Day07 {
	return &Day07{data: data}
}

func isABBA(s string) bool {
	s = strings.ReplaceAll(s, "]", "[")
	var hasABBA bool
	for i, c := range strings.Split(s, "[") {
		abba := containsABBA(c)
		if !abba {
			continue
		}
		if i%2 == 0 {
			hasABBA = abba
		} else {
			return false
		}
	}
	return hasABBA
}

func containsABBA(s string) bool {
	if len(s) < 4 {
		return false
	}
	for i := 0; i <= len(s)-4; i++ {
		if s[i] != s[i+1] && s[i] == s[i+3] && s[i+1] == s[i+2] {
			return true
		}
	}
	return false
}

func (d *Day07) Part1() (int, error) {
	count := 0
	for _, l := range d.data {
		if isABBA(l) {
			count++
		}
	}
	return count, nil
}

func supportsSSL(s string) bool {
	s = strings.ReplaceAll(s, "]", "[")
	supernet := make(map[string]bool)
	hypernet := make(map[string]bool)

	for i, c := range strings.Split(s, "[") {
		target := supernet
		if i%2 != 0 {
			target = hypernet
		}

		for _, aba := range findAllABA(c) {
			target[aba] = true
		}
	}

	for aba := range supernet {
		bab := []rune{rune(aba[1]), rune(aba[0]), rune(aba[1])}
		if hypernet[string(bab)] {
			return true
		}
	}

	return false
}

func findAllABA(s string) []string {
	if len(s) < 3 {
		return nil
	}
	var aba []string
	for i := 0; i <= len(s)-3; i++ {
		if s[i] != s[i+1] && s[i] == s[i+2] {
			aba = append(aba, s[i:i+3])
		}
	}
	return aba
}

func (d *Day07) Part2() (int, error) {
	count := 0
	for _, l := range d.data {
		if supportsSSL(l) {
			count++
		}
	}
	return count, nil
}
