package day01_test

import (
	"aoc2025/challenge/aoc2024/day01"
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
			data: `3   4
4   3
2   5
1   3
3   9
3   3`,
			want: 11,
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
			data: `3   4
4   3
2   5
1   3
3   9
3   3`,
			want: 31,
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
