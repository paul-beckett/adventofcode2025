package day17_test

import (
	"aoc2025/challenge/aoc2016/day17"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	var testCases = map[string]struct {
		data string
		want string
	}{
		"example1": {
			data: `ihgpwlah`,
			want: "DDRRRD",
		},
		"example2": {
			data: `kglvqrro`,
			want: "DDUDRLRRUDRD",
		},
		"example3": {
			data: `ulqzkmiv`,
			want: "DRURDRUDDLLDLUURRDULRLDUUDDDRR",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//given
			day := day17.NewDay17(strings.Split(tc.data, "\n"))

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
			day := day17.NewDay17(strings.Split(tc.data, "\n"))

			//when
			got, err := day.Part2()

			//then
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}
