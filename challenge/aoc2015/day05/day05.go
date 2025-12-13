package day05

import (
	"strings"
)

type Day05 struct {
	data []string
}

func NewDay05(data []string) *Day05 {
	return &Day05{data: data}
}

func (d *Day05) Part1() (int, error) {
	count := 0
	for _, line := range d.data {
		if threeVowels(line) && repeated(line) && !badWords(line) {
			count++
		}
	}
	return count, nil
}

func (d *Day05) Part2() (int, error) {
	count := 0
	for _, line := range d.data {
		if repeatedPair(line) && repeatedSplit(line) {
			count++
		}
	}
	return count, nil
}

func threeVowels(s string) bool {
	total := 0
	vowels := []string{"a", "e", "i", "o", "u"}
	for _, v := range vowels {
		total += strings.Count(s, v)
	}
	return total >= 3
}

func badWords(s string) bool {
	words := []string{"ab", "cd", "pq", "xy"}
	for _, w := range words {
		if strings.Contains(s, w) {
			return true
		}
	}
	return false
}

func repeated(s string) bool {
	for i := range len(s) - 1 {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func repeatedPair(s string) bool {
	for i := range len(s) - 1 {
		if strings.Contains(s[i+2:], s[i:i+2]) {
			return true
		}
	}
	return false
}

func repeatedSplit(s string) bool {
	for i := range len(s) - 2 {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}
