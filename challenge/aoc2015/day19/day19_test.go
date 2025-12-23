package day19_test

import (
	"aoc2025/challenge/aoc2015/day19"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	var testCases = map[string]struct {
		data string
		want int
	}{
		"example1": {
			data: `H => HO
H => OH
O => HH

HOH`,
			want: 4,
		},
		"example2": {
			data: `H => HO
H => OH
O => HH

HOHOHO`,
			want: 7,
		},
		"example3": {
			data: `H => OO

H20`,
			want: 1,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//given
			day := day19.NewDay19(strings.Split(tc.data, "\n"))

			//when
			got, err := day.Part1()

			//then
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPart2(t *testing.T) {
	var testCases = map[string]struct {
		data string
		want int
	}{
		"example1": {
			data: `e => H
e => O
H => HO
H => OH
O => HH

HOH`,
			want: 3,
		},
		"example2": {
			data: `e => H
e => O
H => HO
H => OH
O => HH

HOHOHO`,
			want: 6,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//given
			day := day19.NewDay19(strings.Split(tc.data, "\n"))

			//when
			got, err := day.Part2()

			//then
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
