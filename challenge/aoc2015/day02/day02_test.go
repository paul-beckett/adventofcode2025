package day02_test

import (
	"aoc2025/challenge/aoc2015/day02"
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
			data: `2x3x4
1x1x10`,
			want: 101,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//given
			day := day02.NewDay02(strings.Split(tc.data, "\n"))

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
			data: `2x3x4
1x1x10`,
			want: 48,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//given
			day := day02.NewDay02(strings.Split(tc.data, "\n"))

			//when
			got, err := day.Part2()

			//then
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
