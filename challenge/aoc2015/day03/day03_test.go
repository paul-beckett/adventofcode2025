package day03_test

import (
	"aoc2025/challenge/aoc2015/day03"
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
			data: `>`,
			want: 2,
		},
		"example2": {
			data: `^>v<`,
			want: 4,
		},
		"example3": {
			data: `^v^v^v^v^v`,
			want: 2,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//given
			day := day03.NewDay03(strings.Split(tc.data, "\n"))

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
			data: `^v`,
			want: 3,
		},
		"example2": {
			data: `^>v<`,
			want: 3,
		},
		"example3": {
			data: `^v^v^v^v^v`,
			want: 11,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//given
			day := day03.NewDay03(strings.Split(tc.data, "\n"))

			//when
			got, err := day.Part2()

			//then
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
