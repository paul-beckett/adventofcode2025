package day18_test

import (
	"aoc2025/challenge/aoc2015/day18"
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
			data: `.#.#.#
...##.
#....#
..#...
#.#..#
####..`,
			want: 4,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//given
			day := day18.NewDay18(strings.Split(tc.data, "\n"))

			//when
			got, err := day.Part1WithSteps(4)

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
			data: `##.#.#
...##.
#....#
..#...
#.#..#
####.#`,
			want: 17,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//given
			day := day18.NewDay18(strings.Split(tc.data, "\n"))

			//when
			got, err := day.Part2WithSteps(5)

			//then
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
