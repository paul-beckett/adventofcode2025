package day16

import (
	"aoc2025/internal/must"
	"errors"
	"strings"
)

type Day16 struct {
	aunts []aunt
}

type aunt = map[string]int

func NewDay16(data []string) *Day16 {
	var aunts []aunt
	for _, line := range data {
		a := make(aunt)
		fields := strings.Fields(line)
		for i := 2; i < len(fields); i += 2 {
			thing, _ := strings.CutSuffix(fields[i], ":")
			amount, _ := strings.CutSuffix(fields[i+1], ",")
			a[thing] = must.ParseInt(amount)
		}
		aunts = append(aunts, a)
	}
	return &Day16{aunts: aunts}
}

func (d *Day16) Part1() (int, error) {
	target := map[string]func(int) bool{
		"children":    eq(3),
		"cats":        eq(7),
		"samoyeds":    eq(2),
		"pomeranians": eq(3),
		"akitas":      eq(0),
		"vizslas":     eq(0),
		"goldfish":    eq(5),
		"trees":       eq(3),
		"cars":        eq(2),
		"perfumes":    eq(1),
	}
	for i, a := range d.aunts {
		if matches(target, a) {
			return i + 1, nil
		}
	}
	return -1, errors.New("no match")
}

func (d *Day16) Part2() (int, error) {
	target := map[string]func(int) bool{
		"children":    eq(3),
		"cats":        gt(7),
		"samoyeds":    eq(2),
		"pomeranians": lt(3),
		"akitas":      eq(0),
		"vizslas":     eq(0),
		"goldfish":    lt(5),
		"trees":       gt(3),
		"cars":        eq(2),
		"perfumes":    eq(1),
	}
	for i, a := range d.aunts {
		if matches(target, a) {
			return i + 1, nil
		}
	}
	return -1, errors.New("no match")
}

func eq(n int) func(int) bool {
	return func(m int) bool {
		return n == m
	}
}

func gt(n int) func(int) bool {
	return func(m int) bool {
		return m > n
	}
}

func lt(n int) func(int) bool {
	return func(m int) bool {
		return m < n
	}
}

func matches(target map[string]func(int) bool, a aunt) bool {
	for k, f := range target {
		n, exists := a[k]
		if !exists {
			continue
		}
		if !f(n) {
			return false
		}
	}
	return true
}
