package day12_test

import (
	"aoc2025/challenge/aoc2015/day12"
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
			data: `[1,2,3]`,
			want: 6,
		},
		"example2": {
			data: `{"a":2,"b":4}`,
			want: 6,
		},
		"example3": {
			data: `[[[3]]]`,
			want: 3,
		},
		"example4": {
			data: `{"a":{"b":4},"c":-1}`,
			want: 3,
		},
		"example5": {
			data: `{"a":[-1,1]}`,
			want: 0,
		},
		"example6": {
			data: `[-1,{"a":1}]`,
			want: 0,
		},
		"example7": {
			data: `[]`,
			want: 0,
		},
		"example8": {
			data: `{}`,
			want: 0,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//given
			day := day12.NewDay12(strings.Split(tc.data, "\n"))

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
			data: ``,
			want: -1,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//given
			day := day12.NewDay12(strings.Split(tc.data, "\n"))

			//when
			got, err := day.Part2()

			//then
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
