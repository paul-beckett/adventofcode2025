package math

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllPermutations(t *testing.T) {
	//given
	permutations := [][]int{
		{1, 2, 3},
		{2, 1, 3},
		{3, 1, 2},
		{1, 3, 2},
		{2, 3, 1},
		{3, 2, 1},
	}

	//when
	seq := AllPermutations([]int{1, 2, 3})

	//then
	assert.Equal(t, permutations, slices.Collect(seq))
}
