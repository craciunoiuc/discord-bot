package types

import "golang.org/x/exp/constraints"

type SortableSlice[T constraints.Ordered] struct {
	items []T
}

func (slice SortableSlice[T]) Len() int {
	return len(slice.items)
}

func (slice SortableSlice[T]) Less(i, j int) bool {
	return slice.items[i] < slice.items[j]
}

func (slice SortableSlice[T]) Swap(i, j int) {
	temp := slice.items[i]
	slice.items[i] = slice.items[j]
	slice.items[j] = temp
}
