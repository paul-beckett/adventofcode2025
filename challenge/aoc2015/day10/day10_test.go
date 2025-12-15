package day10_test

import (
	"aoc2025/challenge/aoc2015/day10"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	var testCases = map[string]struct {
		data string
		want string
	}{
		"example1": {
			data: `1`,
			want: `11`,
		},
		"example2": {
			data: `11`,
			want: `21`,
		},
		"example3": {
			data: `21`,
			want: `1211`,
		},
		"example4": {
			data: `1211`,
			want: `111221`,
		},
		"example5": {
			data: `111221`,
			want: `312211`,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//given

			//when
			got := day10.LookAndSay(tc.data)

			//then
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPart2(t *testing.T) {
	t.Skip("no part2 example")
}
