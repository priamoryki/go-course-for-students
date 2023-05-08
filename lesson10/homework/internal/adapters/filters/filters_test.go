package filters

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type Test[T any] struct {
	In     []T
	Expect []T
}

func TestSortFilter_Filter(t *testing.T) {
	comparator := func(elem1 int, elem2 int) bool { return elem1 > elem2 }
	filter := SortFilter[int]{
		comparator: comparator,
	}

	tests := []Test[int]{
		{In: []int{}, Expect: []int{}},
		{In: []int{1, 2, 3}, Expect: []int{3, 2, 1}},
		{In: []int{123, 132, 123, 1231, 32123, 123, 13213}, Expect: []int{32123, 13213, 1231, 132, 123, 123, 123}},
	}

	for _, test := range tests {
		test := test
		t.Run("TestSortFilter_Filter", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.Expect, filter.Filter(test.In))
		})
	}
}

func TestFilters(t *testing.T) {
	condition := func(elem int) bool { return elem > 123 }
	comparator := func(elem1 int, elem2 int) bool { return elem1 > elem2 }
	filter := Filters[int]{
		DefaultFilter[int]{condition: condition},
		SortFilter[int]{comparator: comparator},
	}

	tests := []Test[int]{
		{In: []int{}, Expect: []int{}},
		{In: []int{1, 2, 3}, Expect: []int{}},
		{In: []int{123, 132, 123, 1231, 32123, 123, 13213}, Expect: []int{32123, 13213, 1231, 132}},
	}

	for _, test := range tests {
		test := test
		t.Run("TestFilters", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.Expect, filter.Filter(test.In))
		})
	}
}

func FuzzFilters_Filter(f *testing.F) {
	condition := func(s string) bool { return len(s) > 2 }
	filter := DefaultFilter[string]{
		condition: condition,
	}

	testcases := []string{
		"",
		"a a",
		"lol kek",
		"a aa bb bbb aaaaa ccccccccc",
	}

	for _, tc := range testcases {
		f.Add(tc)
	}

	f.Fuzz(func(t *testing.T, s string) {
		arr := strings.Split(s, " ")
		for _, s := range filter.Filter(arr) {
			assert.True(t, condition(s))
		}
	})
}
