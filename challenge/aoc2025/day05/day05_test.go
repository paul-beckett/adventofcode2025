package day05_test

import (
	"aoc2025/challenge/aoc2025/day05"
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
			data: `3-5
10-14
16-20
12-18

1
5
8
11
17
32`,
			want: 3,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//given
			day := day05.NewDay05(strings.Split(tc.data, "\n"))

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
			data: `3-5
10-14
16-20
12-18

1
5
8
11
17
32`,
			want: 14,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//given
			day := day05.NewDay05(strings.Split(tc.data, "\n"))

			//when
			got, err := day.Part2()

			//then
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
