package day06_test

import (
	"aoc2025/challenge/aoc2025/day06"
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
			data: `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   + `,
			want: 4277556,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//given
			day := day06.NewDay06(strings.Split(tc.data, "\n"))

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
			data: `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `,
			want: 3263827,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//given
			day := day06.NewDay06(strings.Split(tc.data, "\n"))

			//when
			got, err := day.Part2()

			//then
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
