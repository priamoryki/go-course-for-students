package filters

import "sort"

type Filter[T any] interface {
	Filter([]T) []T
}

type DefaultFilter[T any] struct {
	condition func(T) bool
}

func (f DefaultFilter[T]) Filter(arr []T) []T {
	result := make([]T, 0, len(arr))
	for _, elem := range arr {
		if f.condition(elem) {
			result = append(result, elem)
		}
	}
	return result
}

type itemsSorter[T any] struct {
	arr        []T
	comparator func(T, T) bool
}

func (s itemsSorter[T]) Len() int {
	return len(s.arr)
}

func (s itemsSorter[T]) Less(i, j int) bool {
	return s.comparator(s.arr[i], s.arr[j])
}

func (s itemsSorter[T]) Swap(i, j int) {
	s.arr[i], s.arr[j] = s.arr[j], s.arr[i]
}

type SortFilter[T any] struct {
	comparator func(T, T) bool
}

func (f SortFilter[T]) Filter(arr []T) []T {
	sorter := itemsSorter[T]{
		arr,
		f.comparator,
	}
	sort.Sort(sorter)
	return sorter.arr
}

type Filters[T any] []Filter[T]

func (f Filters[T]) Filter(arr []T) []T {
	for _, filter := range f {
		arr = filter.Filter(arr)
	}
	return arr
}
