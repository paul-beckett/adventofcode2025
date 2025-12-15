package day10

import (
	"iter"
	"strconv"
	"strings"
)

type Day10 struct {
	data []string
}

func NewDay10(data []string) *Day10 {
	return &Day10{data: data}
}

func LookAndSay(number string) string {
	var output []string
	for i := 0; i < len(number); i++ {
		n := number[i]
		count := 1
		for j := i + 1; j < len(number); j++ {
			if number[j] != n {
				break
			}
			count++
			i++
		}
		output = append(output, strconv.Itoa(count))
		output = append(output, string(n))
	}
	return strings.Join(output, "")
}

func LookAndSaySeq(start string, limit int) iter.Seq[string] {
	return func(yield func(string) bool) {
		current := start
		for range limit {
			current = LookAndSay(current)
			if !yield(current) {
				break
			}
		}
	}
}

func (d *Day10) Part1() (int, error) {
	start := d.data[0]
	last := start
	for next := range LookAndSaySeq(start, 40) {
		last = next
	}
	return len(last), nil
}

func (d *Day10) Part2() (int, error) {
	start := d.data[0]
	last := start
	for next := range LookAndSaySeq(start, 50) {
		last = next
	}
	return len(last), nil
}
