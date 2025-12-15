package day11_test

import (
	"aoc2025/challenge/aoc2015/day11"
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
			data: `abcdefgh`,
			want: `abcdffaa`,
		},
		"example2": {
			data: `ghijklmn`,
			want: `ghjaabcc`,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//given
			day := day11.NewDay11(strings.Split(tc.data, "\n"))

			//when
			got, err := day.Part1()

			//then
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPart2(t *testing.T) {
	t.Skip("no examples")
}
