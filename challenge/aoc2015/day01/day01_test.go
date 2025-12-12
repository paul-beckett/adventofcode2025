package day01_test

import (
	"aoc2025/challenge/aoc2015/day01"
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
			data: `(())`,
			want: 0,
		},
		"example2": {
			data: `()()`,
			want: 0,
		},
		"example3": {
			data: `(((`,
			want: 3,
		},
		"example4": {
			data: `(()(()(`,
			want: 3,
		},
		"example5": {
			data: `))(((((`,
			want: 3,
		},
		"example6": {
			data: `())`,
			want: -1,
		},
		"example7": {
			data: `))(`,
			want: -1,
		},
		"example8": {
			data: `)))`,
			want: -3,
		},
		"example9": {
			data: `)())())`,
			want: -3,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//given
			day := day01.NewDay01(strings.Split(tc.data, "\n"))

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
			data: `)`,
			want: 1,
		},
		"example2": {
			data: `()())`,
			want: 5,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//given
			day := day01.NewDay01(strings.Split(tc.data, "\n"))

			//when
			got, err := day.Part2()

			//then
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
