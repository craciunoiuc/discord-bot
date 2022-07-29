package types

import (
	"testing"
)

func error(t *testing.T, message string, args ...any) {
	t.Errorf(message, args...)
	t.FailNow()
}

func TestSortedMap(t *testing.T) {
	sortedMap := NewSortedMap[string, int]()

	sortedKeys := sortedMap.GetSortedKeys()

	if len(sortedKeys) != 0 {
		error(t, "Sorted keys should be empty at this step")
	}

	sortedMap.Add("z", 1)
	if len(sortedMap.items) != 1 {
		error(t, "Failed to add item")
	}

	if sortedMap.items["z"] != 1 {
		error(t, "Added value should be %d but is instead %d", 1, sortedMap.items["z"])
	}

	if *sortedMap.needsRefresh != true {
		error(t, "Sorted map is not marked for refresh")
	}

	sortedMap.Add("y", 2)
	sortedMap.Add("x", 3)

	value, found := sortedMap.Get("y")

	if !found {
		error(t, "Could not find value of key '%s'", "y")
	}

	if value != 2 {
		error(t, "Value found for key '%s' should be '%d' but is instead '%d'", "y", 2, value)
	}

	sortedKeys = sortedMap.GetSortedKeys()

	if len(sortedKeys) != 3 {
		error(t, "Sorted keys should be of length '%d' but is instead of length '%d'", 3, len(sortedKeys))
	}

	if sortedKeys[0] != "x" || sortedKeys[1] != "y" || sortedKeys[2] != "z" {
		error(t, "Keys are not sorted properly '%s'", sortedKeys)
	}

	if sortedMap.sortedKeys.items == nil {
		error(t, "Keys have been sorted but not persisted")
	}

	sortedMap.Add("a", 4)

	sortedKeys = sortedMap.GetSortedKeys()

	if len(sortedKeys) != 4 {
		error(t, "Sorted keys should be of length '%d' but is instead of length '%d'", 4, len(sortedKeys))
	}

	if sortedKeys[0] != "a" || sortedKeys[1] != "x" || sortedKeys[2] != "y" || sortedKeys[3] != "z" {
		error(t, "Keys are not sorted properly '%s'", sortedKeys)
	}
}
