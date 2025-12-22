package day23_test

import (
	"aoc2025/challenge/aoc2015/day23"
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
			data: `inc a
jio a, +2
tpl a
inc a`,
			want: 2,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//given
			day := day23.NewDay23(strings.Split(tc.data, "\n"))

			//when
			got, err := day.Part1WithRegister("a")

			//then
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPart2(t *testing.T) {
	t.Skip("no part2")
}
