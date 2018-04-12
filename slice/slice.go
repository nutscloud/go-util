package slice

import (
	"reflect"
	"sort"
)

// MatchInSlice returns true if the given value is in the Slice.
func MatchInSlice(sli []interface{}, elems interface{}) bool {
	for _, v := range sli {
		if reflect.DeepEqual(v, elems) {
			return true
		}
	}
	return false
}

// DeleteElement delete the first equal element from slice.
func DeleteElement(sli []interface{}, elems interface{}) []interface{} {
	for i, v := range sli {
		if reflect.DeepEqual(v, value) {
			return append(sli[:i], sli[i+1:]...)
		}
	}

	return sli
}

// IsEqual determine whether two slices are equal.
func IsEqual(s1, s2 []interface{}) bool {
	return reflect.DeepEqual(s1, s2)
}

// IsEqualIgnoreOrder determine whether two slices are equal without order.
func IsEqualWithoutOrder(s1, s2 []interface{}) bool {
	if len(s1) != len(s2) {
		return false
	}

	// The effect of removal order on results.
	sort.Strings(s1)
	sort.Strings(s2)

	return reflect.DeepEqual(s1, s2)
}
