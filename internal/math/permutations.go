package math

import (
	"iter"
	"slices"
)

// AllPermutations is a non-recursive version of Heap's algorithm returning a Seq of all permutations
func AllPermutations[T any](elements []T) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		n := len(elements)
		c := make([]int, n)
		if !yield(slices.Clone(elements)) {
			return
		}
		a := slices.Clone(elements)
		i := 0
		for i < n {
			if c[i] < i {
				if i%2 == 0 {
					a[0], a[i] = a[i], a[0]
				} else {
					a[c[i]], a[i] = a[i], a[c[i]]
				}
				if !yield(slices.Clone(a)) {
					break
				}
				c[i]++
				i = 0
			} else {
				c[i] = 0
				i++
			}
		}
	}
}
