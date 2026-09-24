package day02

import (
	"aoc2025/internal/math"
	"strings"
)

type Day02 struct {
	data []string
}

func NewDay02(data []string) *Day02 {
	return &Day02{data: data}
}

var (
	directions = map[rune]math.Vector2{
		'U': {X: 0, Y: -1},
		'D': {X: 0, Y: 1},
		'L': {X: -1, Y: 0},
		'R': {X: 1, Y: 0},
	}
)

func (d *Day02) Part1() (string, error) {
	keypad := map[math.Vector2]string{
		{X: 0, Y: 0}: "1", {X: 1, Y: 0}: "2", {X: 2, Y: 0}: "3",
		{X: 0, Y: 1}: "4", {X: 1, Y: 1}: "5", {X: 2, Y: 1}: "6",
		{X: 0, Y: 2}: "7", {X: 1, Y: 2}: "8", {X: 2, Y: 2}: "9",
	}
	return d.solve(keypad, math.Vector2{X: 1, Y: 1}), nil
}

func (d *Day02) Part2() (string, error) {
	keypad := map[math.Vector2]string{
		{X: 2, Y: 0}: "1",
		{X: 1, Y: 1}: "2", {X: 2, Y: 1}: "3", {X: 3, Y: 1}: "4",
		{X: 0, Y: 2}: "5", {X: 1, Y: 2}: "6", {X: 2, Y: 2}: "7", {X: 3, Y: 2}: "8", {X: 4, Y: 2}: "9",
		{X: 1, Y: 3}: "A", {X: 2, Y: 3}: "B", {X: 3, Y: 3}: "C",
		{X: 2, Y: 4}: "D",
	}
	return d.solve(keypad, math.Vector2{X: 0, Y: 2}), nil
}

func (d *Day02) solve(keypad map[math.Vector2]string, start math.Vector2) string {
	var nums []string
	pos := start
	for _, line := range d.data {
		for _, r := range line {
			next := pos.Add(directions[r])
			if _, exists := keypad[next]; exists {
				pos = next
			}
		}
		nums = append(nums, keypad[pos])
	}

	return strings.Join(nums, "")
}
