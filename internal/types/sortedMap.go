package types

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type SortedMap[K constraints.Ordered, V any] struct {
	items        map[K]V
	sortedKeys   *SortableSlice[K]
	needsRefresh *bool
}

func NewSortedMap[K constraints.Ordered, V any]() *SortedMap[K, V] {
	return &SortedMap[K, V]{
		items:        make(map[K]V),
		sortedKeys:   new(SortableSlice[K]),
		needsRefresh: new(bool),
	}
}

func (sortedMap SortedMap[K, V]) Add(key K, value V) {
	sortedMap.items[key] = value
	*sortedMap.needsRefresh = true
}

func (sortedMap SortedMap[K, V]) Get(key K) (V, bool) {
	value, ok := sortedMap.items[key]
	return value, ok
}

func (sortedMap SortedMap[K, V]) GetSortedKeys() []K {
	if *sortedMap.needsRefresh {
		sortedMap.sortedKeys.items = make([]K, 0, len(sortedMap.items))
		for k := range sortedMap.items {
			sortedMap.sortedKeys.items = append(sortedMap.sortedKeys.items, k)
		}

		sort.Stable(sortedMap.sortedKeys)

		*sortedMap.needsRefresh = false
	}

	return sortedMap.sortedKeys.items
}
