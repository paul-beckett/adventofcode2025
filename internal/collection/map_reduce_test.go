package collection_test

import (
	"aoc2025/internal/collection"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	t.Run("sum ints", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := collection.Sum(numbers)
		want := 6
		assert.Equal(t, want, got)
	})
	t.Run("sum int64", func(t *testing.T) {
		numbers := []int64{1, 2, 3}
		got := collection.Sum(numbers)
		want := int64(6)
		assert.Equal(t, want, got)
	})
	t.Run("sum float64", func(t *testing.T) {
		numbers := []float64{1, 2, 3}
		got := collection.Sum(numbers)
		want := 6.0
		assert.Equal(t, want, got)
	})
}

func TestSumFunc(t *testing.T) {
	stringLengthFunc := func(s string) int {
		return len(s)
	}
	t.Run("sum string length", func(t *testing.T) {
		words := []string{"sum", "of", "string", "lengths"}
		got := collection.SumFunc(words, stringLengthFunc)
		want := 18
		assert.Equal(t, want, got)
	})
}

func TestMap(t *testing.T) {
	t.Run("map to uppercase string", func(t *testing.T) {
		toUpper := func(upper string) string { return strings.ToUpper(upper) }
		words := []string{"abc", "def", "ghi"}
		got := collection.Map(words, toUpper)
		want := []string{"ABC", "DEF", "GHI"}
		assert.Equal(t, want, got)
	})
}

func TestReduce(t *testing.T) {
	t.Run("join words", func(t *testing.T) {
		join := func(acc string, word string) string { return acc + "," + word }
		words := []string{"abc", "def", "ghi"}
		got := collection.Reduce(words, join, "_")
		want := "_,abc,def,ghi"
		assert.Equal(t, want, got)
	})
}

func TestFilter(t *testing.T) {
	t.Run("filter even numbers", func(t *testing.T) {
		isEven := func(num int) bool { return num%2 == 0 }
		nums := []int{1, 2, 3, 4, 5}
		got := collection.Filter(nums, isEven)
		want := []int{2, 4}
		assert.Equal(t, want, got)
	})
}

func TestAll(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	t.Run("all values match predicate", func(t *testing.T) {
		allLessThan10 := func(num int) bool { return num < 10 }
		assert.True(t, collection.All(nums, allLessThan10))
	})
	t.Run("all values do not match predicate", func(t *testing.T) {
		allLessThan4 := func(num int) bool { return num < 4 }
		assert.False(t, collection.All(nums, allLessThan4))
	})
}

func TestChunkBy(t *testing.T) {
	var testCases = map[string]struct {
		data []string
		want [][]string
	}{

		"split in middle": {
			data: []string{"abc", "def", "", "ghi"},
			want: [][]string{{"abc", "def"}, {"ghi"}},
		},
		"multiple splits": {
			data: []string{"abc", "", "def", "", "ghi"},
			want: [][]string{{"abc"}, {"def"}, {"ghi"}},
		},
		"split at start": {
			data: []string{"", "abc", "def", "ghi"},
			want: [][]string{{"abc", "def", "ghi"}},
		},
		"split at end": {
			data: []string{"abc", "def", "ghi", ""},
			want: [][]string{{"abc", "def", "ghi"}},
		},
	}
	matchWhiteSpace := func(s string) bool { return strings.TrimSpace(s) == "" }
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			got := collection.ChunkBy(tc.data, matchWhiteSpace)
			assert.Equal(t, tc.want, got)
		})
	}
}
