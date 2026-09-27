package day09_test

import (
	"aoc2025/challenge/aoc2016/day09"
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
			data: `ADVENT`,
			want: 6,
		},
		"example2": {
			data: `A(1x5)BC`,
			want: 7,
		},
		"example3": {
			data: `(3x3)XYZ`,
			want: 9,
		},
		"example4": {
			data: `A(2x2)BCD(2x2)EFG`,
			want: 11,
		},
		"example5": {
			data: `(6x1)(1x3)A`,
			want: 6,
		},
		"example6": {
			data: `X(8x2)(3x3)ABCY`,
			want: 18,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//given
			day := day09.NewDay09(strings.Split(tc.data, "\n"))

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
			data: `(3x3)XYZ`,
			want: 9,
		},
		"example2": {
			data: `X(8x2)(3x3)ABCY`,
			want: 20,
		},
		"example3": {
			data: `(27x12)(20x12)(13x14)(7x10)(1x12)A`,
			want: 241920,
		},
		"example4": {
			data: `(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN`,
			want: 445,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//given
			day := day09.NewDay09(strings.Split(tc.data, "\n"))

			//when
			got, err := day.Part2()

			//then
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
