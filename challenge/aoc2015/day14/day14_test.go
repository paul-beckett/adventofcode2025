package day14_test

import (
	"aoc2025/challenge/aoc2015/day14"
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
			data: `Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`,
			want: 1120,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//given
			day := day14.NewDay14(strings.Split(tc.data, "\n"))

			//when
			got, err := day.Part1WithTime(1000)

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
			data: `Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`,
			want: 689,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//given
			day := day14.NewDay14(strings.Split(tc.data, "\n"))

			//when
			got, err := day.Part2WithTime(1000)

			//then
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
