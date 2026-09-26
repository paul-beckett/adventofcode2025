package day04

import (
	"aoc2025/internal/must"
	"cmp"
	"fmt"
	"slices"
	"strings"
)

type Day04 struct {
	data []string
}

func NewDay04(data []string) *Day04 {
	return &Day04{data: data}
}

func (d *Day04) Part1() (int, error) {
	sum := 0
	for _, line := range d.data {
		r := parseRoom(line)
		if isReal(r) {
			sum += r.sectorID
		}
	}
	return sum, nil
}

func (d *Day04) Part2() (int, error) {
	for _, line := range d.data {
		r := parseRoom(line)
		if isReal(r) {
			decrypt := Part2Decrypt(line)
			if strings.Contains(decrypt, "north") {
				fmt.Println(decrypt)
				return r.sectorID, nil
			}
		}
	}
	return -1, fmt.Errorf("not found")
}

func Part2Decrypt(s string) string {
	i := strings.LastIndex(s, "-")
	a, _, _ := strings.Cut(s[i+1:], "[")
	id := must.ParseInt(a)
	var sentence []rune
	for _, c := range s[:i] {
		if c == '-' {
			sentence = append(sentence, ' ')
			continue
		}
		l := int(c - 'a')
		l = l + id
		l = l % 26
		sentence = append(sentence, int32(l)+'a')
	}

	return string(sentence)
}

func topN(counts map[rune]int) map[rune]bool {
	type pair struct {
		k rune
		v int
	}
	var pairs []pair
	for k, v := range counts {
		pairs = append(pairs, pair{k: k, v: v})
	}
	slices.SortFunc(pairs, func(a, b pair) int {
		if c := cmp.Compare(b.v, a.v); c != 0 {
			return c
		}
		return cmp.Compare(a.k, b.k)
	})
	top := make(map[rune]bool)
	for _, p := range pairs[:5] {
		top[p.k] = true
	}
	return top
}

type room struct {
	counts   map[rune]int
	sectorID int
	checksum []rune
}

func parseRoom(s string) room {
	f := strings.Split(s, "-")
	counts := make(map[rune]int)
	for _, l := range f[:len(f)-1] {
		for _, c := range l {
			counts[c]++
		}
	}

	sectorCheck := f[len(f)-1]
	sector := sectorCheck[:len(sectorCheck)-7]
	check := sectorCheck[len(sectorCheck)-6 : len(sectorCheck)-1]

	return room{
		counts:   counts,
		sectorID: must.ParseInt(sector),
		checksum: []rune(check),
	}
}

func isReal(r room) bool {
	top := topN(r.counts)
	for _, l := range r.checksum {
		if !top[l] {
			return false
		}
	}
	return true
}
