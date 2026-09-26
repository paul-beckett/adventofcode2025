package day04_test

import (
	"aoc2025/challenge/aoc2016/day04"
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
			data: `aaaaa-bbb-z-y-x-123[abxyz]
a-b-c-d-e-f-g-h-987[abcde]
not-a-real-room-404[oarel]
totally-real-room-200[decoy]`,
			want: 1514,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//given
			day := day04.NewDay04(strings.Split(tc.data, "\n"))

			//when
			got, err := day.Part1()

			//then
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestPart2Decrypt(t *testing.T) {
	var testCases = map[string]struct {
		data string
		want string
	}{
		"example1": {
			data: "qzmt-zixmtkozy-ivhz-343",
			want: "very encrypted name",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			//given

			//when
			got := day04.Part2Decrypt(tc.data)

			//then
			assert.Equal(t, tc.want, got)
		})
	}
}
